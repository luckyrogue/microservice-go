package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ashishsingh4u/bookmicroservice/config"
	"github.com/ashishsingh4u/bookmicroservice/controllers"
	docs "github.com/ashishsingh4u/bookmicroservice/docs"
	"github.com/ashishsingh4u/bookmicroservice/models"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	pprof.Register(router)

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(ctx *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}))

	// This middleware is already used when router is created
	// with router := gin.Default(). Use gin.New() to start router with no middleware attached
	// and then use specific middleware like logger or Recovery
	// router.Use(gin.Recovery())
	var conf config.Configuration
	if err := config.GetConfig(&conf); err != nil {
		panic(fmt.Sprintf("Couldn't read the configuration file. Error: %s", err.Error()))
	}

	machineIP := fmt.Sprintf("%s:%s", conf.SERVER_IP, conf.PORT)
	log.Printf("Server will be starting on %s\n", machineIP)

	router.SetTrustedProxies([]string{conf.SERVER_IP})

	router.GET("/", func(ctx *gin.Context) {
		// If the client is 192.168.86.22, use the X-Forwarded-For
		// header to deduce the original client IP from the trust-
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
	})

	models.ConnectDatabase()

	// Grouping
	v1 := router.Group("/v1")
	{
		v1.GET("/books", controllers.GetBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}

	// Swagger related declarations
	// This adds some extra memory overhead (20 MB)
	docs.SwaggerInfo.BasePath = "/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(machineIP)
}
