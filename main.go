// @title		main
// @description	开启伺服器
// @auth		ryl				2022/4/6		23:30

package main

import (
	"dianasdog/communication"
	"dianasdog/setup"
)

func main() {
	router := communication.SetupRouter()
	router.Run(":8080")

	setup.SetByCategory()
}
