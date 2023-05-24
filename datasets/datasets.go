package datasets

import (
"github.com/arashmo/globalsync/dbo"
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	
)

type Dataset struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Size   *int    `json:"size"`
	Status string `json:"status"`
	Version string `json:"version"`
}
type ServerDataset struct {
    ID            int    `json:"id"`
    location_on_disk string `json:"location_on_disk"`
	Name          string `json: "name"`
	Hostname      string `json: "hostname"`
	IP_address    string `json: "ip_address"`
	Folder_name   string  `json: "folder_name"`
}

func GetDatasets(c *gin.Context) {

	rows, err := dbo.DB.Query("SELECT id, name, size, COALESCE(status, ''), version FROM datasets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var datasets []Dataset

	for rows.Next() {
		var dataset Dataset
		var size sql.NullInt64
		err = rows.Scan(&dataset.ID, &dataset.Name, &size, &dataset.Status, &dataset.Version)

		if err != nil {
			log.Fatal(err)
		}

		if size.Valid {
			dataset.Size = new(int)
			*dataset.Size = int(size.Int64)
		}
		datasets = append(datasets, dataset)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": datasets})
}
func SearchDatasets(c *gin.Context) {

	searchTerm := c.Query("search")

	rows, err := dbo.DB.Query(`
	SELECT d.name, s.hostname, s.ip_address, ass.location, sd.folder_name FROM datasets d 
	JOIN server_datasets sd ON sd.dataset_id = d.id 
	JOIN servers s ON s.id = sd.id 
	JOIN attached_storage ass ON ass.id = s.id WHERE d.name LIKE ?
    `, "%"+searchTerm+"%")
	if err != nil {
		log.Fatal(err)
	}

	datasets := []ServerDataset{}
	for rows.Next() {
		var dataset ServerDataset
		if err := rows.Scan(&dataset.Name, &dataset.Hostname, &dataset.IP_address, &dataset.location_on_disk, &dataset.Folder_name); err != nil {
			log.Fatal(err)
		}
		datasets = append(datasets, dataset)
	}

	c.JSON(http.StatusOK, gin.H{"data": datasets})
}


func GetDataset(c *gin.Context) {
	id := c.Param("id")
	var dataset Dataset
	err := dbo.DB.QueryRow("SELECT id, name, size, status FROM datasets WHERE id = ?", id).Scan(&dataset.ID, &dataset.Name, &dataset.Size, &dataset.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("dataset %s not found", id)})
			return
		}
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": dataset})
}

func CreateDataset(c *gin.Context) {
	var dataset Dataset
	if err := c.ShouldBindJSON(&dataset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := dbo.DB.Exec("INSERT INTO datasets (name, size, status) VALUES (?, ?, ?)", dataset.Name, dataset.Size, dataset.Status)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	dataset.ID = int(id)

	c.JSON(http.StatusCreated, gin.H{"data": dataset})
}

func UpdateDataset(c *gin.Context) {
	id := c.Param("id")

	var dataset Dataset
	if err := c.ShouldBindJSON(&dataset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := dbo.DB.Exec("UPDATE datasets SET name = ?, size = ?, status = ? WHERE id = ?", dataset.Name, dataset.Size, dataset.Status, id)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": dataset})
}
