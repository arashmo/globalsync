package main

import (

	sshsyncer "github.com/arashmo/globalsync/sshsyncer"
	"github.com/gin-gonic/gin"
)

type ServerAsset struct {
	Name           string `json:"name"`
	ID             int    `json:"id"`
	FreeSpace      int    `json:"freespace"`
	UsedSpace      int    `json:"usedspace"`
	Location       string `json:"location"`
	BelongingGroup string `json:"belonging_group"`
	Dataset        map[string]string `json:"dataset"`
}

type ServerAssets struct {
	ServerAssets []ServerAsset `json:"server_assets"`
}

func main() {
	router := gin.Default()

	router.GET("/search", )
	router.Run(":8080")

}
func sa(){
sshsyncer.copyFiles
}

	