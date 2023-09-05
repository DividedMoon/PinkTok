package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v7"
	interactBiz "interact_service/biz"
	client "video_service/biz"
	internalClient "video_service/internal/client"
	"video_service/internal/constants"
	"video_service/internal/dal/db"
	rd "video_service/internal/dal/redis"
)

func (s *VideoService) FavoriteAction(videoId, userId int64, actionType int32) error {
	//1为点赞，0为取消点赞
	//1. 调用interactService去修改点赞表 注意要等待返回值
	updateFavoriteRecordReq := &interactBiz.AddFavoriteRecordReq{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: int64(actionType),
	}
	updateFavoriteRecordResp, err := internalClient.InteractServiceClient.AddFavoriteRecord(s.ctx, updateFavoriteRecordReq)
	if err != nil || updateFavoriteRecordResp.StatusCode != 0 {
		hlog.Error("AddFavoriteRecord error", err.Error())
		return constants.NewErrNo(constants.RPCCallErrCode, constants.RPCCallErrMsg)
	}

	// 2. 修改video表中的点赞数，有两种，对于热门视频直接修改redis中的值，对于冷门视频可以直接修改数据库
	if actionType == 1 { // 如果操作类型为1则为点赞
		err = AddVideoFavoriteCount(videoId, 1)
	} else if actionType == 0 { // 如果操作类型为0 则为取消赞操作
		err = AddVideoFavoriteCount(videoId, -1)
	} else { // 如果不是这两种类型 则返回错误
		hlog.Error("FavoriteAction error", "actionType is not 1 or 0")
		return constants.NewErrNo(constants.ParameterErrCode, constants.ParameterErrMsg)
	}

	// 3. 如果过程中出现错误则调用interactService去取消修改点赞表
	if err != nil {
		hlog.Error("AddVideoFavoriteCount error", err.Error())
		if actionType == 1 {
			updateFavoriteRecordReq.ActionType = 2
		} else if actionType == 2 {
			updateFavoriteRecordReq.ActionType = 1
		}
		cancelUpdateResp, err := internalClient.InteractServiceClient.AddFavoriteRecord(s.ctx, updateFavoriteRecordReq)
		if err != nil || cancelUpdateResp.StatusCode != 0 {
			hlog.Error("AddFavoriteRecord error", err.Error())
			return constants.NewErrNo(constants.RPCCallErrCode, constants.RPCCallErrMsg+"while cancel update")
		}
		return constants.NewErrNo(constants.RedisErrCode, constants.RedisErrMsg)
	}
	// 3. 返回结果
	return nil
}

func (s *VideoService) GetFavoriteVideoList(userId int64) ([]*client.VideoInfo, error) {

	// 1. 调用 interactService 获取用户点赞的视频列表 返回值是videoIds 视频ID列表
	getFavoriteVideoIdsReq := &interactBiz.FavoriteVideoReq{
		UserId: userId,
	}
	getFavoriteVideoIdsResp, err := internalClient.InteractServiceClient.QueryUserFavoriteVideoIds(s.ctx, getFavoriteVideoIdsReq)
	if err != nil || getFavoriteVideoIdsResp.StatusCode != 0 {
		hlog.Error("QueryUserFavoriteVideoIds error", err.Error())
		return nil, constants.NewErrNo(constants.RPCCallErrCode, constants.RPCCallErrMsg)
	}
	favoriteVideoIds := getFavoriteVideoIdsResp.VideoList

	// 2. 调用copyVideo方法获取视频信息并返回
	videoDBInfos, err := db.GetVideoDBInfoByIDs(favoriteVideoIds)
	if err != nil {
		hlog.Error("GetVideoDBInfoByIDs error", err.Error())
		return nil, constants.NewErrNo(constants.DBErrCode, constants.DBErrMsg)
	}
	videos := make([]*client.VideoInfo, 0, len(favoriteVideoIds))
	err = s.CopyVideos(&videos, &videoDBInfos, userId)
	if err != nil {
		hlog.Error("CopyVideos error", err.Error())
		return nil, constants.NewErrNo(constants.FunctionErrCode, constants.FunctionErrMsg)
	}
	return videos, nil
}

func AddVideoFavoriteCount(videoId, increment int64) error {
	// 1.先查redis里有没有 如果有的话直接加一或者减一返回
	video, err := rd.GetVideoHash(videoId)
	if err != nil && err != redis.Nil {
		hlog.Error("GetVideoFavoriteCount error", err.Error())
		return err
	}
	if video != nil { // video不为空 直接更新redis后返回
		err = rd.SetVideoField(videoId, "FavoriteCount", video.FavoriteCount+increment)
		if err != nil {
			hlog.Error("SetVideoField error", err.Error())
			return err
		}
		// redis里的video被修改，加入修改集合
		err = rd.AddChangedVideo(videoId)
		if err != nil {
			hlog.Error("AddChangedVideo error", err.Error())
			return err
		}
		return nil
	}
	// 2.如果没有的话，查数据库
	video, err = db.GetVideoByID(videoId)
	if err != nil {
		hlog.Error("GetVideoFavoriteCount error", err.Error())
		return err
	}
	// 2.1如果数据库中的点赞数大于1000 则认为是热门视频，将点赞数 + 1，写入redis
	if video.FavoriteCount >= 1000 {
		video.FavoriteCount += increment
		err = buildVideoInfoCache(video)
		if err != nil {
			hlog.Error("buildVideoInfoCache error", err.Error())
			return err
		}
		return nil
	}
	// 2.2更新数据库
	err = db.UpdateVideoFavoriteCount(videoId, video.FavoriteCount+increment)
	if err != nil {
		hlog.Error("UpdateVideoFavoriteCount error", err.Error())
		return err
	}
	return nil
}
