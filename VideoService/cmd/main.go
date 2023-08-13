package main

import (
	"log"
	"video_service/biz/handler"
	biz "video_service/biz/videoservice"
)

func main() {
	svr := biz.NewServer(new(handler.VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
