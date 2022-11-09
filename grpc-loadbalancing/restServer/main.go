package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)
func main() {
	router := gin.Default()
	router.GET("/", response)

	port := "9090"
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(fmt.Sprintf("localhost:%s", port))
}

func response(c *gin.Context) {
	var podIp string
	if len(os.Getenv("POD_IP")) > 0 {
		podIp = os.Getenv("POD_IP")
	}
	c.String(http.StatusOK, "responding from a rest server with ip %s", podIp)
}
