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
  location VARCHAR(255) NOT NULL
);



--datacenters table
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
('server1', '192.168.1.1', 20, 2, 500, 1, 5, 1),
('server2', '192.168.1.2', 10, 0, 1000, 1, 6, 1),
('server3', '192.168.2.1', 30, 4, 200, 2, 7, 2),
('server4', '192.168.2.2', 5, 1, 750, 2, 8, 2);

-- server_datasets table
INSERT INTO server_datasets (dataset_id, server_id, attached_storage_id, folder_name) VALUES
    (1, 1, 1, 'data_folder_1'),
    (2, 1, 2, 'data_folder_2'),
    (3, 2, 3, 'data_folder_3'),
    (4, 2, 1, 'data_folder_4'),
    (5, 3, 2, 'data_folder_5');


-- psudu data_entry 
INSERT INTO attached_storage (server_id, location) VALUES
  (10, '/path/to/attached/storage/for/server1'),
  (11, '/path/to/attached/storage/for/server2'),
  (12, '/path/to/attached/storage/for/server3');

----


SELECT d.name, s.hostname, s.ip_address, ass.location, sd.folder_name
FROM datasets d
JOIN server_datasets sd ON sd.dataset_id = d.id
JOIN servers s ON s.id = sd.server_id
JOIN attached_storage ass ON ass.server_id = s.id
WHERE d.name LIKE 'Dataset 1';
