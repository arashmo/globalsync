package store

import (
	"database/sql"
)

type Store interface {
	// Datacenters
	InsertDatacenter(name, location, comment string) error
	// Add more methods for datacenters as needed

	// Racks
	InsertRack(number, aisleNumber int, location string, datacenterID int) error
	// Add more methods for racks as needed

	// Owner Groups
	InsertOwnerGroup(name string) error
	// Add more methods for owner groups as needed

	// Datasets
	InsertDataset(name string, size int, version, status string) error
	// Add more methods for datasets as needed

	// Servers
	InsertServer(hostname, ipAddress string, cpuPercent, usedGpus, freeDisk, rackID, datacenterID int) error
	// Add more methods for servers as needed

	// Server Owner Groups
	InsertServerOwnerGroup(serverID, ownerGroupID int) error
	// Add more methods for server owner groups as needed

	// Server Datasets
	InsertServerDataset(serverID, datasetID, attachedStorageID int, folderName string) error
	// Add more methods for server datasets as needed

	// Attached Storage
	InsertAttachedStorage(serverID int, location string) error
	// Add more methods for attached storage as needed
}

type DBStore struct {
	DB *sql.DB
}

// Implement the remaining Store interface methods for DBStore

func (store *DBStore) InsertRack(number, aisleNumber int, location string, datacenterID int) error {
	query := `INSERT INTO racks (number, aisle_number, location, datacenter_id) VALUES (?, ?, ?, ?)`
	_, err := store.DB.Exec(query, number, aisleNumber, location, datacenterID)
	return err
}

func (store *DBStore) InsertOwnerGroup(name string) error {
	query := `INSERT INTO owner_groups (name) VALUES (?)`
	_, err := store.DB.Exec(query, name)
	return err
}

func (store *DBStore) InsertDataset(name string, size int, version, status string) error {
	query := `INSERT INTO datasets (name, size, version, status) VALUES (?, ?, ?, ?)`
	_, err := store.DB.Exec(query, name, size, version, status)
	return err
}

func (store *DBStore) InsertServer(hostname, ipAddress string, cpuPercent, usedGpus, freeDisk, rackID, datacenterID int) error {
	query := `INSERT INTO servers (hostname, ip_address, cpupercent, usedgpus, freedisk, rack_id, datacenter_id) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := store.DB.Exec(query, hostname, ipAddress, cpuPercent, usedGpus, freeDisk, rackID, datacenterID)
	return err
}

func (store *DBStore) InsertServerOwnerGroup(serverID, ownerGroupID int) error {
	query := `INSERT INTO server_owner_groups (server_id, owner_group_id) VALUES (?, ?)`
	_, err := store.DB.Exec(query, serverID, ownerGroupID)
	return err
}

func (store *DBStore) InsertServerDataset(serverID, datasetID, attachedStorageID int, folderName string) error {
	query := `INSERT INTO server_datasets (server_id, dataset_id, attached_storage_id, folder_name) VALUES (?, ?, ?, ?)`
	_, err := store.DB.Exec(query, serverID, datasetID, attachedStorageID, folderName)
	return err
}

func (store *DBStore) InsertAttachedStorage(serverID int, location string) error {
	query := `INSERT INTO attached_storage (server_id, location) VALUES (?, ?)`
	_, err := store.DB.Exec(query, serverID, location)
	return err
}

// Update functions for DBStore

func (store *DBStore) UpdateDatacenter(id int, name, location, comment string) error {
	query := `UPDATE datacenters SET name = ?, location = ?, comment = ? WHERE id = ?`
	_, err := store.DB.Exec(query, name, location, comment, id)
	return err
}

func (store *DBStore) UpdateRack(id, number, aisleNumber int, location string, datacenterID int) error {
	query := `UPDATE racks SET number = ?, aisle_number = ?, location = ?, datacenter_id = ? WHERE id = ?`
	_, err := store.DB.Exec(query, number, aisleNumber, location, datacenterID, id)
	return err
}

func (store *DBStore) UpdateOwnerGroup(id int, name string) error {
	query := `UPDATE owner_groups SET name = ? WHERE id = ?`
	_, err := store.DB.Exec(query, name, id)
	return err
}

func (store *DBStore) UpdateDataset(id int, name string, size int, version, status string) error {
	query := `UPDATE datasets SET name = ?, size = ?, version = ?, status = ? WHERE id = ?`
	_, err := store.DB.Exec(query, name, size, version, status, id)
	return err
}

func (store *DBStore) UpdateServer(id int, hostname, ipAddress string, cpuPercent, usedGpus, freeDisk, rackID, datacenterID int) error {
	query := `UPDATE servers SET hostname = ?, ip_address = ?, cpupercent = ?, usedgpus = ?, freedisk = ?, rack_id = ?, datacenter_id = ? WHERE id = ?`
	_, err := store.DB.Exec(query, hostname, ipAddress, cpuPercent, usedGpus, freeDisk, rackID, datacenterID, id)
	return err
}

func (store *DBStore) UpdateServerOwnerGroup(id, serverID, ownerGroupID int) error {
	query := `UPDATE server_owner_groups SET server_id = ?, owner_group_id = ? WHERE id = ?`
	_, err := store.DB.Exec(query, serverID, ownerGroupID, id)
	return err
}

func (store *DBStore) UpdateServerDataset(id, serverID, datasetID, attachedStorageID int, folderName string) error {
	query := `UPDATE server_datasets SET server_id = ?, dataset_id = ?, attached_storage_id = ?, folder_name = ? WHERE id = ?`
	_, err := store.DB.Exec(query, serverID, datasetID, attachedStorageID, folderName, id)
	return err
}

func (store *DBStore) UpdateAttachedStorage(id, serverID int, location string) error {
	query := `UPDATE attached_storage SET server_id = ?, location = ? WHERE id = ?`
	_, err := store.DB.Exec(query, serverID, location, id)
	return err
}

// Delete functions for DBStore

func (store *DBStore) DeleteDatacenter(id int) error {
	query := `DELETE FROM datacenters WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteRack(id int) error {
	query := `DELETE FROM racks WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteOwnerGroup(id int) error {
	query := `DELETE FROM owner_groups WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteDataset(id int) error {
	query := `DELETE FROM datasets WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteServer(id int) error {
	query := `DELETE FROM servers WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteServerOwnerGroup(id int) error {
	query := `DELETE FROM server_owner_groups WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteServerDataset(id int) error {
	query := `DELETE FROM server_datasets WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

func (store *DBStore) DeleteAttachedStorage(id int) error {
	query := `DELETE FROM attached_storage WHERE id = ?`
	_, err := store.DB.Exec(query, id)
	return err
}

