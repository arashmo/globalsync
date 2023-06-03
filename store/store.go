package store


import (
	"database/sql"
)

type Store interface {
	// Upserts
	UpsertDatacenter(id int, name, location, comment string) error
	UpsertRack(id, number, aisleNumber int, location string, datacenterID int) error
	UpsertOwnerGroup(id int, name string) error
	UpsertDataset(id int, name string, size int, version, status string) error
	UpsertServer(id int, hostname, ipAddress string, cpuPercent, usedGpus, freeDisk, rackID, datacenterID int) error
	UpsertServerOwnerGroup(id, serverID, ownerGroupID int) error
	UpsertServerDataset(id, serverID, datasetID, attachedStorageID int, folderName string) error
	UpsertAttachedStorage(id, serverID int, location string) error
}

type DBStore struct {
	DB *sql.DB
}

func (store *DBStore) UpsertDatacenter(id int, name, location, comment string) error {
	query := `INSERT INTO datacenters (id, name, location, comment) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE name = VALUES(name), location = VALUES(location), comment = VALUES(comment)`
	_, err := store.DB.Exec(query, id, name, location, comment)
	return err
}

func (store *DBStore) UpsertRack(id, number, aisleNumber int, location string, datacenterID int) error {
	query := `INSERT INTO racks (id, number, aisle_number, location, datacenter_id) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE number = VALUES(number), aisle_number = VALUES(aisle_number), location = VALUES(location), datacenter_id = VALUES(datacenter_id)`
	_, err := store.DB.Exec(query, id, number, aisleNumber, location, datacenterID)
	return err
}

func (store *DBStore) UpsertOwnerGroup(id int, name string) error {
	query := `INSERT INTO owner_groups (id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = VALUES(name)`
	_, err := store.DB.Exec(query, id, name)
	return err
}

func (store *DBStore) UpsertDataset(id int, name string, size int, version, status string) error {
	query := `INSERT INTO datasets (id, name, size, version, status) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE name = VALUES(name), size = VALUES(size), version = VALUES(version), status = VALUES(status)`
	_, err := store.DB.Exec(query, id, name, size, version, status)
	return err
}

func (store *DBStore) UpsertServer(id int, hostname, ipAddress string, cpuPercent, usedGpus, freeDisk, rackID, datacenterID int) error {
	query := `INSERT INTO servers (id, hostname, ip_address, cpupercent, usedgpus, freedisk, rack_id, datacenter_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE hostname = VALUES(hostname), ip_address = VALUES(ip_address), cpupercent = VALUES(cpupercent), usedgpus = VALUES(usedgpus), freedisk = VALUES(freedisk), rack_id = VALUES(rack_id), datacenter_id = VALUES(datacenter_id)`
	_, err := store.DB.Exec(query, id, hostname, ipAddress, cpuPercent, usedGpus, freeDisk, rackID, datacenterID)
	return err
}

func (store *DBStore) UpsertServerOwnerGroup(id, serverID, ownerGroupID int) error {
	query := `INSERT INTO server_owner_groups (id, server_id, owner_group_id) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE server_id = VALUES(server_id), owner_group_id = VALUES(owner_group_id)`
	_, err := store.DB.Exec(query, id, serverID, ownerGroupID)
	return err
}

func (store *DBStore) UpsertServerDataset(id, serverID, datasetID, attachedStorageID int, folderName string) error {
	query := `INSERT INTO server_datasets (id, server_id, dataset_id, attached_storage_id, folder_name) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE server_id = VALUES(server_id), dataset_id = VALUES(dataset_id), attached_storage_id = VALUES(attached_storage_id), folder_name = VALUES(folder_name)`
	_, err := store.DB.Exec(query, id, serverID, datasetID, attachedStorageID, folderName)
	return err
}

func (store *DBStore) UpsertAttachedStorage(id, serverID int, location string) error {
	query := `INSERT INTO attached_storage (id, server_id, location) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE server_id = VALUES(server_id), location = VALUES(location)`
	_, err := store.DB.Exec(query, id, serverID, location)
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

