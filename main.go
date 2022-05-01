// @title		main
// @description	开启伺服器
// @auth		ryl				2022/4/7		11:30

package main

import (
	"dianasdog/communication"
	"fmt"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "dianasdog/docs"
)

// @title 特型卡片搜索系统
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host https://Backend-DianasDog.app.secoder.net
// @BasePath /
// schemes http
func main() {
	router := communication.SetupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":80")

	if err != nil {
		fmt.Println(err)
	}
}
