package main

import (
	"fmt"
	"github.com/gomystery/easyws"
)

func main() {

	ws:=easyws.NewEasyWs("127.0.0.1",9001)
	fmt.Println(ws)



}