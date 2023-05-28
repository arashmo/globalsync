package main

import (
	"github.com/arashmo/globalsync/dbo"
	"github.com/arashmo/globalsync/store"
	"github.com/arashmo/globalsync/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dsn := "root@tcp(localhost)/globalsync"
	db, err := dbo.Connect(dsn)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbo.Close(db)
	dbo.DB = db
	dbo.InitTables() 
	dbStore := &store.DBStore{DB: db}
	h := handler.NewHandler(dbStore)
	
	r := gin.Default()
	r.POST("/datasets", h.CreateDataset)
	
	r.Run()
}
