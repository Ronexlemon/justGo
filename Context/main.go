package main

import (
	"fmt"
	"gocontext/contexts"
)

func main() {
	welcome := "Yollow here we go again Context"
	fmt.Println(welcome)
	cont := "Context are a way of adding cancelation or timeout"
	fmt.Println(cont)
	// contexts.Contexts()
	// contexts.CallDoWork()
	// contexts.PassWithValue()
	// contexts.WithTimeOut()
	contexts.WithCancel()
}
