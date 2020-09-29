package main

import (
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
			return origin == "https://swaldmann.github.io"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.GET("/algorithm/:type/relations/:numberOfRelations/graphType/:graphType", func(c *gin.Context) {
			algorithmType := c.Param("type")
			graphType := c.Param("graphType")

			numberOfRelations, err := strconv.ParseUint(c.Param("numberOfRelations"), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if graphType == "moerkotte" {
				numberOfRelations = 5 // This is a specific example and not auto-generated
			}
			QG := GetQueryGraph(graphType, uint(numberOfRelations))

			switch algorithmType {
			case "dpccp":
				configuration := &Configuration{}

				Costfunctions := []costfunctionT{Cnlj, Chj, Csmj}
				JTC := JoinTreeCreator{false, false, Costfunctions}
				routines := visualize(DPccp, QG, JTC)

				c.JSON(http.StatusOK, gin.H{
					"routines":      routines,
					"configuration": configuration,
					"queryGraph":    QG,
				})
			case "dpsize":
				configuration := &Configuration{}

				Costfunctions := []costfunctionT{Cnlj, Chj, Csmj}
				JTC := JoinTreeCreator{false, false, Costfunctions}
				routines := visualize(DPsize, QG, JTC)

				c.JSON(http.StatusOK, gin.H{
					"routines":      routines,
					"configuration": configuration,
					"queryGraph":    QG,
				})
			}
		})
	}
	// Listen and serve on 0.0.0.0:8080
	router.Run()
}
