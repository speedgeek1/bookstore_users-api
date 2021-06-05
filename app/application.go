package app

import "github.com/gin-gonic/gin"

const PORT = ":8080"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrl()
	router.Run(PORT)
}
