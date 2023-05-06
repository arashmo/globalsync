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
    id SERIAL PRIMARY KEY,
    server_id INT NOT NULL REFERENCES servers(id),
    dataset_id INT NOT NULL REFERENCES datasets(id),
    attached_storage_id INT NOT NULL REFERENCES attached_storage(id),
    folder_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE attached_storage (
  id SERIAL PRIMARY KEY,
  server_id INTEGER REFERENCES servers(id),
  location VARCHAR(255) NOT NULL,
  UNIQUE (server_id, location)
);



--datacenters table
INSERT INTO datacenters (name, location, comment) VALUES
('Datacenter 1', 'New York', 'Primary datacenter'),
('Datacenter 2', 'Los Angeles', 'Secondary datacenter');

-- racks table
INSERT INTO racks (number, aisle_number, location, datacenter_id) VALUES
(1, 1, 'North-East corner', (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1)),
(2, 1, 'North-West corner', (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1)),
(3, 1, 'South-East corner', (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1)),
(4, 1, 'South-West corner', (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1));

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
('server1', '192.168.1.1', 20, 2, 500, (SELECT id FROM owner_groups ORDER BY RAND() LIMIT 1), (SELECT id FROM racks ORDER BY RAND() LIMIT 1), (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1)),
('server2', '192.168.1.2', 10, 0, 1000, (SELECT id FROM owner_groups ORDER BY RAND() LIMIT 1), (SELECT id FROM racks ORDER BY RAND() LIMIT 1), (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1)),
('server3', '192.168.2.1', 30, 4, 200, (SELECT id FROM owner_groups ORDER BY RAND() LIMIT 1), (SELECT id FROM racks ORDER BY RAND() LIMIT 1), (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1)),
('server4', '192.168.2.2', 5, 1, 750, (SELECT id FROM owner_groups ORDER BY RAND() LIMIT 1), (SELECT id FROM racks ORDER BY RAND() LIMIT 1), (SELECT id FROM datacenters ORDER BY RAND() LIMIT 1));

-- server_datasets table
INSERT INTO server_datasets (dataset_id, server_id, attached_storage_id, folder_name) VALUES
    ((SELECT id FROM datasets ORDER BY RAND() LIMIT 1), (SELECT id FROM servers ORDER BY RAND() LIMIT 1), (SELECT id FROM attached_storage ORDER BY RAND() LIMIT 1),'data_folder_1'),
    ((SELECT id FROM datasets ORDER BY RAND() LIMIT 1), (SELECT id FROM servers ORDER BY RAND() LIMIT 1), (SELECT id FROM attached_storage ORDER BY RAND() LIMIT 1),'data_folder_2'),
    ((SELECT id FROM datasets ORDER BY RAND() LIMIT 1), (SELECT id FROM servers ORDER BY RAND() LIMIT 1), (SELECT id FROM attached_storage ORDER BY RAND() LIMIT 1),'data_folder_3');
 


-- attached_storage data_entry 
INSERT INTO attached_storage (server_id, location) VALUES
  ((SELECT id FROM servers ORDER BY RAND() LIMIT 1), '/path/to/attached/storage/for/server1'),
  ((SELECT id FROM servers ORDER BY RAND() LIMIT 1), '/path/to/attached/storage/for/server2'),
  ((SELECT id FROM servers ORDER BY RAND() LIMIT 1), '/path/to/attached/storage/for/server3');

----

