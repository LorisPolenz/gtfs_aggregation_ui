package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Query struct {
	StopName       string `json:"stop_name"`
	RouteShortname string `json:"route_short_name"`
	TripHeadsign   string `json:"trip_headsign"`
	DelayThreshold int    `json:"delay_threshold"`
}

func get_route_by_short_name(c *gin.Context) {
	route_short_name := c.Param("route_short_name")
	entities := fetch_by_routeShortName(route_short_name)

	c.JSON(200, gin.H{
		"entities": entities,
	})
}

func get_route_by_search(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindBodyWith(&query, binding.JSON)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	entities := fetch_by_query(query)

	c.JSON(200, gin.H{
		"entities": entities,
	})

}
