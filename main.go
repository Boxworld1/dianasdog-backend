// @title		main
// @description	开启伺服器
// @auth		ryl				2022/4/7		11:30

package main

import (
	"dianasdog/communication"
	"dianasdog/setup"
	"fmt"
)

func main() {
	setup.SetByCategory()
	router := communication.SetupRouter()
	err := router.Run(":80")
	if err != nil {
		fmt.Println(err)
	}
}
