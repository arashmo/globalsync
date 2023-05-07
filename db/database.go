package db

import (
	"database/sql"
	"log"
"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dsn string) {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}



func DBhadler(dsn string) (*sql.DB ,error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Create database
	if _, err := DB.Exec("CREATE DATABASE IF NOT EXISTS globalsync"); err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	// Switch to the created database
	if _, err := DB.Exec("USE globalsync"); err != nil {
		return nil, fmt.Errorf("failed to switch to database: %v", err)
	}

	// Create tables
	if _, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS datacenters (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		location VARCHAR(255),
		comment VARCHAR(255)
	  );
	  
	  CREATE TABLE IF NOT EXISTS  racks (
		id INT AUTO_INCREMENT PRIMARY KEY,
		number INT,
		aisle_number INT,
		location VARCHAR(255),
		datacenter_id INT,
		FOREIGN KEY (datacenter_id) REFERENCES datacenters(id)
	  );
	  
	  CREATE TABLE IF NOT EXISTS  owner_groups (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255)
	  );
		CREATE TABLE IF NOT EXISTS datasets (
			id INT(11) NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		);

		CREATE TABLE IF NOT EXISTS servers (
			id INT(11) NOT NULL AUTO_INCREMENT,
			hostname VARCHAR(255) NOT NULL,
			ip_address VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		);

		CREATE TABLE IF NOT EXISTS attached_storage (
			id INT(11) NOT NULL AUTO_INCREMENT,
			server_id INT(11) NOT NULL,
			location VARCHAR(255) NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS server_datasets (
			server_id INT(11) NOT NULL,
			dataset_id INT(11) NOT NULL,
			folder_name VARCHAR(255) NOT NULL,
			PRIMARY KEY (server_id, dataset_id),
			FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE,
			FOREIGN KEY (dataset_id) REFERENCES datasets(id) ON DELETE CASCADE
		);
	`); err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}
	return db, nill
}

