package main

import (
	"fmt"
	"gochannels/channels"
)


func main(){
	welcome := "Yollow welcome to golang channels"
	fmt.Println(welcome)
	channels.Channels()
	channels.BufferedChannels()
	channels.ChannelsSimultenous()
	channels.ChannelWithLoad()
}