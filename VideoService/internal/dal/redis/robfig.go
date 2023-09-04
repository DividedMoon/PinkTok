package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"video_service/internal/utils"
)

func InitChangedVideoSet4Robfig() error {
	//1. 查询redis中是否已经存在了该集合
	hlog.Infof("InitChangedVideoSet4Robfig")
	exist, err := querySetExist(rdRobfig, utils.GetChangedVideoKey())
	if err != nil {
		hlog.Error("querySetExist error", err.Error())
		return err
	}
	//2. 如果不存在则创建
	hlog.Infof("InitChangedVideoSet4Robfig exist: %v", exist)
	if !exist {
		err = initSet(rdRobfig, utils.GetChangedVideoKey(), []int64{})
		hlog.Infof("InitChangedVideoSet4Robfig initSet err: %v", err)
		if err != nil {
			hlog.Error("createSet error", err.Error())
			return err
		}
	}
	return nil
}

func GetChangedVideoIds() ([]int64, error) {
	//1. 查询redis中是否已经存在了该集合 如果不存在就报错
	exist, err := querySetExist(rdRobfig, utils.GetChangedVideoKey())
	if !exist && err == nil {
		hlog.Infof("GetChangedVideoIds not exist")
		return nil, nil
	}
	if exist && err != nil {
		hlog.Error("querySetExist error", err.Error())
		return nil, err
	}

	//2. 以Pop的方式获取集合中的所有元素
	ids, err := getAllFromSetAndClear(rdRobfig, utils.GetChangedVideoKey())
	if err != nil {
		hlog.Error("getAllFromSetAndClear error", err.Error())
		return nil, err
	}
	return ids, nil
}

// TODO 还没有将该方法加入业务实现中
func AddChangedVideo(videoId int64) error {
	err := addIntoSet(rdRobfig, utils.GetChangedVideoKey(), videoId)
	if err != nil {
		hlog.Error("addIntoSet error", err.Error())
		return err
	}
	return nil
}
