package dbo

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
)

var DB *sql.DB

func Connect(connection string) (*sql.DB, error) {
    splitInfo := strings.Split(connection, "/")
    dbName := splitInfo[len(splitInfo)-1]
    splitInfo[len(splitInfo)-1] = "?charset=utf8&parseTime=True&loc=Local" // remove DB name
    db, err := sql.Open("mysql", strings.Join(splitInfo, "/")) // Open DB without DB name
    if err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %v", err)
    }

    _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
    if err != nil {
        return nil, fmt.Errorf("failed to create database: %v", err)
    }

    _, err = db.Exec("USE " + dbName)
    if err != nil {
        return nil, fmt.Errorf("failed to switch to database: %v", err)
    }

    // Perform any additional database configuration or setup if needed
    return db, nil
}

func Close(db *sql.DB) error {
    if db != nil {
        return db.Close()
    }
    return nil
}

func InitTables() {
    // Define table creation queries
    createDatacentersQuery := `
        CREATE TABLE IF NOT EXISTS datacenters (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            location VARCHAR(255),
            comment VARCHAR(255)
        )
    `
    createRacksQuery := `
        CREATE TABLE IF NOT EXISTS racks (
            id INT AUTO_INCREMENT PRIMARY KEY,
            number INT,
            aisle_number INT,
            location VARCHAR(255),
            datacenter_id INT,
            FOREIGN KEY (datacenter_id) REFERENCES datacenters(id)
        )
    `
    createOwnerGroupsQuery := `
        CREATE TABLE IF NOT EXISTS owner_groups (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255)
        )
    `
    createDatasetsQuery := `
        CREATE TABLE IF NOT EXISTS datasets (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            size INT,
            version VARCHAR(255),
            status VARCHAR(255),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        )
    `
    createServersQuery := `
        CREATE TABLE IF NOT EXISTS servers (
            id INT AUTO_INCREMENT PRIMARY KEY,
            hostname VARCHAR(255),
            ip_address VARCHAR(255),
            cpupercent INT,
            usedgpus INT,
            freedisk INT,
            rack_id INT,
            datacenter_id INT,
            FOREIGN KEY (rack_id) REFERENCES racks(id),
            FOREIGN KEY (datacenter_id) REFERENCES datacenters(id)
        )
    `
    createServerOwnerGroupsQuery := `
        CREATE TABLE IF NOT EXISTS server_owner_groups (
            id INT AUTO_INCREMENT PRIMARY KEY,
            server_id INT,
            owner_group_id INT,
            FOREIGN KEY (server_id) REFERENCES servers(id),
            FOREIGN KEY (owner_group_id) REFERENCES owner_groups(id)
        )
    `
    createServerDatasetsQuery := `
        CREATE TABLE IF NOT EXISTS server_datasets (
            id INT AUTO_INCREMENT PRIMARY KEY,
            server_id INT,
            dataset_id INT,
            attached_storage_id INT,
            folder_name VARCHAR(255),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (server_id) REFERENCES servers(id),
            FOREIGN KEY (dataset_id) REFERENCES datasets(id)
        )
    `
    createAttachedStorageQuery := `
        CREATE TABLE IF NOT EXISTS attached_storage (
            id INT AUTO_INCREMENT PRIMARY KEY,
            server_id INTEGER,
            location VARCHAR(255) NOT NULL,
            UNIQUE (server_id, location),
            FOREIGN KEY (server_id) REFERENCES servers(id)
        )
    `

    // Execute each query
    _, err := DB.Exec(createDatacentersQuery)
    if err != nil {
        log.Fatalf("Error creating datacenters table: %v", err)
    }
    
    _, err = DB.Exec(createRacksQuery)
    if err != nil {
        log.Fatalf("Error creating racks table: %v", err)
    }
    
    _, err = DB.Exec(createOwnerGroupsQuery)
    if err != nil {
        log.Fatalf("Error creating owner_groups table: %v", err)
    }
    
    _, err = DB.Exec(createDatasetsQuery)
    if err != nil {
        log.Fatalf("Error creating datasets table: %v", err)
    }
    
    _, err = DB.Exec(createServersQuery)
    if err != nil {
        log.Fatalf("Error creating servers table: %v", err)
    } 
	_, err = DB.Exec(createServerDatasetsQuery)
    if err != nil {
        log.Fatalf("Error creating servers table: %v", err)
    }
    
    _, err = DB.Exec(createServerOwnerGroupsQuery)
    if err != nil {
        log.Fatalf("Error creating server_owner_groups table: %v", err)
    }  
	 _, err = DB.Exec(createAttachedStorageQuery)
    if err != nil {
        log.Fatalf("Error creating server_owner_groups table: %v", err)
    }
    
}
