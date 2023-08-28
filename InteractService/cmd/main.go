package main

import (
	"interact_service/biz/handler"
	biz "interact_service/biz/interactservice"
	"log"
)

func main() {
	svr := biz.NewServer(new(handler.InteractServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
