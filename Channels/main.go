package main

import (
	"fmt"
	"gochannels/channel02"
	"gochannels/channels"
)


func main(){
	welcome := "Yollow welcome to golang channels"
	fmt.Println(welcome)
	channels.Channels()
	channels.BufferedChannels()
	channels.ChannelsSimultenous()
	channels.ChannelWithLoad()
	channel02.Channel02()
	channel02.BuffereChan()
}