package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"jjaa.me/server"
	"jjaa.me/util"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if util.InitConfig() == false {
		print("no config")
		return
	}
	fmt.Println(util.AllConfig)
	if len(os.Args) > 2 {
		util.AllConfig.Http.Host = os.Args[2]
		server.Serve(os.Args[1])
	} else {
		server.Serve(util.AllConfig.Http.Port)
	}
}
