package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"

	"github.cm/gin-gonic/gin"
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

func SearchHandler(c *gin.Context) {
	// Load JSON file
	jsonFile, err := ioutil.ReadFile("servers.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading JSON file"})
		return
	}

	// Parse JSON data
	var serverAssets ServerAssets
	err = json.Unmarshal(jsonFile, &serverAssets)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON data"})
		return
	}

	// Parse query parameters
	q := c.Request.URL.Query()
	searchFields := make(map[string]string)
	for key, value := range q {
		searchFields[strings.ToLower(key)] = strings.ToLower(value[0])
	}

	// Search for matching server assets
	var results []ServerAsset
	for _, serverAsset := range serverAssets.ServerAssets {
		matched := true
		for key, value := range searchFields {
			switch key {
			case "name":
				if !strings.Contains(strings.ToLower(serverAsset.Name), value) {
					matched = false
					break
				}
			case "id":
				if fmt.Sprintf("%d", serverAsset.ID) != value {
					matched = false
					break
				}
			case "freespace":
				freespace, err := strconv.Atoi(value)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid freespace value"})
					return
				}
				if serverAsset.FreeSpace < freespace {
					matched = false
					break
				}
			case "usedspace":
				usedspace, err := strconv.Atoi(value)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid usedspace value"})
					return
				}
				if serverAsset.UsedSpace < usedspace {
					matched = false
					break
				}
			case "location":
				if !strings.Contains(strings.ToLower(serverAsset.Location), value) {
					matched = false
					break
				}
			case "belonginggroup":
				if !strings.Contains(strings.ToLower(serverAsset.BelongingGroup), value) {
					matched = false
					break
				}

			case "dataset":
				found := false
				for _, v := range serverAsset.Dataset {
					if strings.Contains(strings.ToLower(v), value) {
						found = true
						break
					}
				}
				if !found {
					matched = false
					break
				}
			}
		}
	}
}

		
