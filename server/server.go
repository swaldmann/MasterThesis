package main

import (
	"image/color"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startServer() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.GET("/queryGraph/:graphType/relations/:numberOfRelations", func(c *gin.Context) {
			graphType := c.Param("graphType")

			numberOfRelations, err := strconv.ParseUint(c.Param("numberOfRelations"), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if graphType == "moerkotte" {
				numberOfRelations = 5
			}
			QG := GetQueryGraph(graphType, uint(numberOfRelations))
			c.JSON(http.StatusOK, gin.H{
				"queryGraph": QG,
			})
		})

		api.GET("/algorithm/:type/relations/:numberOfRelations/graphType/:graphType", func(c *gin.Context) {
			algorithmType := c.Param("type")
			graphType := c.Param("graphType")

			numberOfRelations, err := strconv.ParseUint(c.Param("numberOfRelations"), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if graphType == "moerkotte" {
				numberOfRelations = 5
			}
			QG := GetQueryGraph(graphType, uint(numberOfRelations))

			// First, define the initial state.
			// Second, generate all the steps from the algorithm implementation.

			switch algorithmType {
			case "dpccp":
				//obeservedRelations := []string{"S", "X", "N", "emit/S"}
				configuration := &Configuration{}

				Costfunctions := []costfunctionT{Cnlj, Chj, Csmj}
				JTC := JoinTreeCreator{false, false, Costfunctions}
				routines := visualizeDPccp(QG, JTC)

				c.JSON(http.StatusOK, gin.H{
					"routines":      routines,
					"configuration": configuration,
					"queryGraph":    QG,
				})
			case "adaptiveRadixTree":
				color := color.RGBA{85, 165, 34, 1}
				nodeColor := NodeColor{NodeIndex: 0, Color: color}
				graphState := &GraphState{NodeColors: []NodeColor{nodeColor}}
				counter := &AlgorithmCounter{Name: "LohmannCounter", Value: 0}

				c.JSON(http.StatusOK, gin.H{
					"counters":   counter,
					"graphState": graphState,
					"queryGraph": QG,
				})
			}
		})
	}
	// Listen and serve on 0.0.0.0:8080
	router.Run()
}
