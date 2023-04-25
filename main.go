package main

import (
//	"github.com/arashmo/globalsync/db"
	"log"
  "github.com/arashmo/globalsync/datasets"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	
	

	router := gin.Default()

	router.GET("/datasets", datasets.GetDatasets)
	router.GET("/datasets/search",datasets.SearchDatasets)
	router.GET("/datasets/:id", datasets.GetDataset)
	router.POST("/datasets", datasets.CreateDataset)
	router.PUT("/datasets/:id", datasets.UpdateDataset)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

