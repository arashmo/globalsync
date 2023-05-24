
DROP TABLE IF EXISTS server_datasets;
DROP TABLE IF EXISTS attached_storage;
DROP TABLE IF EXISTS servers;
DROP TABLE IF EXISTS datasets;
DROP TABLE IF EXISTS owner_groups;
DROP TABLE IF EXISTS racks;
DROP TABLE IF EXISTS datacenters;

CREATE TABLE datacenters (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  location VARCHAR(255),
  comment VARCHAR(255)
);

CREATE TABLE racks (
  id INT AUTO_INCREMENT PRIMARY KEY,
  number INT,
  aisle_number INT,
  location VARCHAR(255),
  datacenter_id INT,
  FOREIGN KEY (datacenter_id) REFERENCES datacenters(id)
);

CREATE TABLE owner_groups (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255)
);

CREATE TABLE datasets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  size INT,
  version VARCHAR(255),
  status VARCHAR(255),
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE servers (
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
);

CREATE TABLE server_owner_groups (
  server_id INT,
  owner_group_id INT,
  PRIMARY KEY (server_id, owner_group_id),
  FOREIGN KEY (server_id) REFERENCES servers(id),
  FOREIGN KEY (owner_group_id) REFERENCES owner_groups(id)
);

CREATE TABLE attached_storage (
  id INT AUTO_INCREMENT PRIMARY KEY,
  server_id INT,
  location VARCHAR(255) NOT NULL,
  FOREIGN KEY (server_id) REFERENCES servers(id),
  UNIQUE KEY (server_id, location)
);

CREATE TABLE server_datasets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    server_id INT NOT NULL,
    dataset_id INT NOT NULL,
    attached_storage_id INT NOT NULL,
    folder_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (server_id) REFERENCES servers(id),
    FOREIGN KEY (dataset_id) REFERENCES datasets(id),
    FOREIGN KEY (attached_storage_id) REFERENCES attached_storage(id)
);
