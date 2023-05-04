package  sshsync 
import (
	"github.com/arashmo/globalsync/db"
		"fmt"
		"log"
	//	"strings"
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		
	)

	func SearchDstData(c *gin.Context)  {
		db.Connect("root@tcp(localhost:3306)/globalsync")
		defer db.Close()
	
		searchTerm := c.Query("search")
	
		rows, err := db.DB.Query(`
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
	func SearchSrcData(c *gin.Context)  {
		db.Connect("root@tcp(localhost:3306)/globalsync")
		defer db.Close()
	
		searchTerm := c.Query("search")
	
		rows, err := db.DB.Query(`
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
	