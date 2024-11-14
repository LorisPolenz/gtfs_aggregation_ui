package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	initdb := fetch_db()

	db1 = db_meta{
		instance:  "db1",
		db:        initdb,
		lastFetch: time.Unix(1, 1),
	}
	db2 = db_meta{
		instance:  "db2",
		db:        initdb,
		lastFetch: time.Unix(0, 0),
	}

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to specific domains for better security
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// load html
	r.GET("/", load_index)

	r.GET("/route/:route_short_name", get_route_by_short_name)
	r.POST("/route/search", get_route_by_search)

	r.Run() // listen and serve on 0.0.0.0:8080
}
