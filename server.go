package main

import (
	"fmt"
	"log"
)

type RtcToken struct {
	UID         uint32 `json:uid`
	ChannelName string `json:channel_name`
	Role        uint32 `json:role`
}

func main() {
	var token RtcToken
	log.Println(token)
	fmt.Println("Hello Agora!")
}
