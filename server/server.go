package main

import (
	"image/color"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// AlgorithmCounter Defines a counter in an algorithm.
type AlgorithmCounter struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Algorithm Defines an algorithm.
type Algorithm struct {
	Name        string             `json:"name"`
	Counters    []AlgorithmCounter `json:"counters"`
	GraphStates []GraphState       `json:"graphStates"`
}

// GraphState Defines the state of a graph.
type GraphState struct {
	NodeColors []NodeColor `json:"nodeColors"`
}

// NodeColor Defines the color of a graph node.
type NodeColor struct {
	NodeIndex int        `json:"nodeIndex"`
	Color     color.RGBA `json:"color"`
}

func startServer() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/graph/:type", func(c *gin.Context) {
			graphType := c.Param("type")
			c.JSON(200, gin.H{
				"nodePoints": "" + graphType,
			})
		})

		api.GET("/algorithm/:type/", func(c *gin.Context) {
			algorithmType := c.Param("type")

			// First, define the initial state.
			// Second, generate all the diffs from the algorithm implementation.

			switch algorithmType {
			case "dpccp":
				color := color.RGBA{85, 165, 34, 1}
				nodeColor := NodeColor{NodeIndex: 0, Color: color}
				graphStates := &GraphState{NodeColors: []NodeColor{nodeColor}}
				counter := &AlgorithmCounter{Name: "LohmannCounter", Value: 0}

				changes := []interface{}{}

				c.JSON(http.StatusOK, gin.H{
					"begin": gin.H{"counter": counter, "graphStates": graphStates},
					"diffs": changes,
				})
			case "adaptiveRadixTree":
				color := color.RGBA{85, 165, 34, 1}
				nodeColor := NodeColor{NodeIndex: 0, Color: color}
				graphStates := &GraphState{NodeColors: []NodeColor{nodeColor}}
				counter := &AlgorithmCounter{Name: "LohmannCounter", Value: 0}
				c.JSON(200, gin.H{
					"counters":    counter,
					"graphStates": graphStates,
				})
			}
		})
	}
	// Listen and serve on 0.0.0.0:8080
	router.Run()
}
