// @title		main
// @description	开启伺服器
// @auth		ryl				2022/4/6		23:30

package main

import (
	"dianasdog/communication"
	"dianasdog/setup"
	"fmt"
)

func main() {
	router := communication.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
	setup.SetByCategory()
}
