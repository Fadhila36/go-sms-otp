package cmd

import (
	"github.com/Fadhila36/go-sms-otp/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// initialize config
	app := api.Config{Router: router}

	// routes
	app.Routes()

	router.Run(":8000")
}
