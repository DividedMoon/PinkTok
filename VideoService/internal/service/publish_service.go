package service

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"
	"video_service/biz"
	"video_service/internal/constants"
	"video_service/internal/dal/db"
	"video_service/internal/middleware/ffmpeg"
	"video_service/internal/middleware/minio"
	"video_service/internal/utils"
)

func (s *VideoService) PublishAction(req *biz.PublishReq) error {
	userId := req.UserId
	title := req.Title

	nowTime := time.Now()
	filename := utils.NewFileName(userId, nowTime.Unix())
	//req.Data.Filename = filename + path.Ext(req.Data.Filename)
	mimeType := http.DetectContentType(req.Data)
	extensions, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		hlog.Error("get extensions err:" + err.Error())
		return constants.NewErrNo(constants.GetExtensionErrCode, constants.GetExtensionErrMsg)
	}

	fileType := strings.TrimPrefix(extensions[0], ".")
	uploadFilename := filename + "." + fileType
	hlog.Infof("Uploading File:" + uploadFilename)
	videoBuffer := bytes.NewBuffer(req.Data)

	uploadInfo, err := minio.PutToBucketByBuf(s.ctx, constants.MinioVideoBucketName, uploadFilename, videoBuffer)
	hlog.Infof("video upload size:" + strconv.FormatInt(uploadInfo.Size, 10))

	PlayURL := constants.MinioVideoBucketName + "/" + uploadFilename
	buf, err := ffmpeg.GetSnapshot(utils.URLconvert(s.ctx, PlayURL))
	uploadInfo, err = minio.PutToBucketByBuf(s.ctx, constants.MinioImgBucketName, filename+".png", buf)
	hlog.Infof("image upload size:" + strconv.FormatInt(uploadInfo.Size, 10))
	if err != nil {
		hlog.Error("err:" + err.Error())
		return constants.NewErrNo(constants.UploadImgErrCode, constants.UploadImgErrMsg)
	}

	_, err = db.CreateVideo(&db.VideoDBInfo{
		AuthorID:    userId,
		PlayURL:     PlayURL,
		CoverURL:    constants.MinioImgBucketName + "/" + filename + ".png",
		PublishTime: nowTime,
		Title:       title,
	})
	if err != nil {
		hlog.Error("CreateVideo Err:" + err.Error())
		return constants.NewErrNo(constants.DBErrCode, constants.DBErrMsg)
	}
	return nil
}

func (s *VideoService) GetPublishList(req *biz.GetPublishListReq) (resp *biz.GetPublishListResp, err error) {
	resp = &biz.GetPublishListResp{}
	currentUserId := req.UserId

	dbVideos, err := db.GetVideoByUserID(currentUserId)
	hlog.Infof("dbVideos:%+v", dbVideos)
	if err != nil {
		return nil, constants.NewErrNo(constants.DBErrCode, constants.DBErrMsg)
	}
	var videos []*biz.VideoInfo

	f := NewVideoService(s.ctx)
	err = f.CopyVideos(&videos, &dbVideos, currentUserId)
	hlog.Infof("videos:%+v", videos)
	if err != nil {
		return nil, constants.NewErrNo(constants.VideoCopyErrCode, constants.VideoCopyErrMsg)
	}
	for _, item := range videos {
		video := &biz.VideoInfo{
			Id: item.Id,
			Author: &biz.UserInfo{
				Id:              item.Author.Id,
				Name:            item.Author.Name,
				FollowCount:     item.Author.FollowCount,
				FollowerCount:   item.Author.FollowerCount,
				Avatar:          item.Author.Avatar,
				BackgroundImage: item.Author.BackgroundImage,
				Signature:       item.Author.Signature,
				TotalFavorited:  item.Author.TotalFavorited,
				WorkCount:       item.Author.WorkCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		resp.VideoList = append(resp.VideoList, video)
	}
	return resp, nil
}
