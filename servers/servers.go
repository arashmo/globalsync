package  servers 
import (
	"github.com/arashmo/globalsync/dbo"
		"fmt"
		"log"
	//	"strings"
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		
	)

	func Show_dst_data_location(c *gin.Context)  {
	

		searchTerm := c.Query("search")
	
		rows, err := dbo.DB.Query(`
			SELECT servers.ip_address, attached_storage.location
			FROM servers
			JOIN attached_storage ON servers.id = attached_storage.server_id
			WHERE servers.hostname = ?
		`, searchTerm)
	
		if err != nil {
			log.Fatal(err)
		}
	
		defer rows.Close()
	
		var results []string
	
		for rows.Next() {
			var ip, location string
			err := rows.Scan(&ip, &location)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, fmt.Sprintf("IP: %s, Location: %s", ip, location))
		}
	
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	fmt.Println(results)
		//return strings.Join(results, "\n")
	}