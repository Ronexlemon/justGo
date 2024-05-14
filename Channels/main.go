package main

import (
	"fmt"
	"gochannels/channel02"
	"gochannels/channel03"
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
	channel02.RangeAndClose()
	channel03.WithoutRoutineAndChannel()
	channel03.RoutinesChannel()
}