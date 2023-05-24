package main

import (
	"log"

	"github.com/arashmo/globalsync/dbo"
	"github.com/gin-gonic/gin"
		"github.com/arashmo/globalsync/servers"
	//	"github.com/arashmo/globalsync/sshsync"
	//	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main(){
	db, err := dbo.Connect("user:password@/globalsync")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	router := gin.Default()
	router.GET("/dst",servers.Show_dst_data_location)
	router.Run()

}

// func sshsynci(){
// 	opts := &sshsync.Options{
//         SourceDir:      "/home/araddsh/1",
//         DestinationDir: "localhost:/home/arash/kir",
//         Username:       "arash",
//         Password:       "klfjhpi4sswo44riswwor??",
//         Host:           "localhost:22",
//     }


	
//     err := sshsync.SyncFiles(opts)
//     if err != nil {
//         panic(err)
//     }
// }
