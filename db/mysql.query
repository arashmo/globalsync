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
  status VARCHAR(255)
);

CREATE TABLE servers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  hostname VARCHAR(255),
  ip_address VARCHAR(255),
  cpupercent INT,
  usedgpus INT,
  owner_group_id INT,
  freedisk INT,
  rack_id INT,
  datacenter_id INT,
  FOREIGN KEY (datacenter_id) REFERENCES datacenters(id),
  FOREIGN KEY (rack_id) REFERENCES racks(id),
  FOREIGN KEY (owner_group_id) REFERENCES owner_groups(id)
);

CREATE TABLE server_datasets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  location_on_disk VARCHAR(255),
  server_id INT,
  dataset_id INT,
  FOREIGN KEY (server_id) REFERENCES servers(id),
  FOREIGN KEY (dataset_id) REFERENCES datasets(id)
);
-- psudu data_entry 

-- datacenters table
INSERT INTO datacenters (name, location, comment) VALUES
('Datacenter 1', 'New York', 'Primary datacenter'),
('Datacenter 2', 'Los Angeles', 'Secondary datacenter');

-- racks table
INSERT INTO racks (number, aisle_number, location, datacenter_id) VALUES
(1, 1, 'North-East corner', 1),
(2, 1, 'North-West corner', 1),
(3, 1, 'South-East corner', 2),
(4, 1, 'South-West corner', 2);

-- owner_groups table
INSERT INTO owner_groups (name) VALUES
('Group A'),
('Group B'),
('Group C');

-- datasets table
INSERT INTO datasets (name, version) VALUES
('Dataset 1', 'v1.0'),
('Dataset 2', 'v2.0'),
('Dataset 3', 'v1.5');

-- servers table
INSERT INTO servers (hostname, ip_address, cpupercent, usedgpus, freedisk, owner_group_id, rack_id, datacenter_id) VALUES
('server1', '192.168.1.1', 20, 2, 500, 1, 1, 1),
('server2', '192.168.1.2', 10, 0, 1000, 1, 1, 1),
('server3', '192.168.2.1', 30, 4, 200, 2, 3, 2),
('server4', '192.168.2.2', 5, 1, 750, 2, 3, 2);

-- server_datasets table
INSERT INTO server_datasets (location_on_disk, server_id, dataset_id) VALUES
('/mnt/data1', 1, 1),
('/mnt/data2', 1, 2),
('/mnt/data3', 2, 1),
('/mnt/data4', 3, 2),
('/mnt/data5', 4, 3);
