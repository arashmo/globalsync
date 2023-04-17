package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"

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

	router.GET("/search",SearchHandler )
	router.Run(":8080")

}
	