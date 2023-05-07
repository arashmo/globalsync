package main

import (
"log"
"github.com/arashmo/globalsync/datasets"
"github.com/arashmo/globalsync/sshsync"
"github.com/arashmo/globalsync/servers"
//"github.com/arashmo/globalsync/db"
"net/http"
"github.com/gin-gonic/gin"
_ "github.com/go-sql-driver/mysql"

)
func main() {
	

	
	router := gin.Default()

	router.Static("/public", "./frontend/public")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./frontend/public/index.html", nil)
	  })

	router.GET("/datasets", datasets.GetDatasets)
	router.GET("/datasets/search",datasets.SearchDatasets)
	router.GET("/datasets/:id", datasets.GetDataset)
	router.POST("/datasets", datasets.CreateDataset)
	router.PUT("/datasets/:id", datasets.UpdateDataset)
	router.GET("/searchdst/", servers.SearchDstData)


	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func sshsynci(){
	opts := &sshsync.Options{
        SourceDir:      "/home/araddsh/1",
        DestinationDir: "localhost:/home/arash/kir",
        Username:       "arash",
        Password:       "klfjhpi4sswo44riswwor??",
        Host:           "localhost:22",
    }

    err := sshsync.SyncFiles(opts)
    if err != nil {
        panic(err)
    }
}
