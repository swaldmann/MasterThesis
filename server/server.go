package main

import (
	"fmt"
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
		/*api.GET("/graph/:type", func(c *gin.Context) {
			graphType := c.Param("type")
			c.JSON(200, gin.H{
				"nodePoints": "" + graphType,
			})
		})*/

		api.GET("/algorithm/:type/relations/:numberOfRelations/graphType/:graphType", func(c *gin.Context) {
			algorithmType := c.Param("type")
			graphType := c.Param("graphType")

			numberOfRelations, err := strconv.ParseUint(c.Param("numberOfRelations"), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			QG := GetQueryGraph(graphType, uint(numberOfRelations))

			/*QG := GetQueryGraph("clique", 5)*/
			fmt.Println(QG.N[0])

			// First, define the initial state.
			// Second, generate all the diffs from the algorithm implementation.

			switch algorithmType {
			case "dpccp":
				color := color.RGBA{85, 165, 34, 1}
				nodeColor := NodeColor{NodeIndex: 0, Color: color}
				graphState := &GraphState{NodeColors: []NodeColor{nodeColor}}
				counter := &AlgorithmCounter{Name: "LohmannCounter", Value: 0}

				observedVariables := []string{"S", "X", "N", "emit/S"}
				configuration := &Configuration{ObserverdVariables: observedVariables}

				changes := visualizeDPccp(QG)

				c.JSON(http.StatusOK, gin.H{
					"begin": gin.H{
						"counter":    counter,
						"graphState": graphState},
					"diffs":         changes,
					"configuration": configuration,
				})
			case "adaptiveRadixTree":
				color := color.RGBA{85, 165, 34, 1}
				nodeColor := NodeColor{NodeIndex: 0, Color: color}
				graphState := &GraphState{NodeColors: []NodeColor{nodeColor}}
				counter := &AlgorithmCounter{Name: "LohmannCounter", Value: 0}

				c.JSON(http.StatusOK, gin.H{
					"counters":   counter,
					"graphState": graphState,
				})
			}
		})
	}
	// Listen and serve on 0.0.0.0:8080
	router.Run()
}
