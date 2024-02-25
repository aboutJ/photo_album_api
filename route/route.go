package route

import "github.com/gin-gonic/gin"

func InitRoute() {
	r := gin.Default()

	InitUserRoute(r)

	_ = r.Run(":8081")
}
