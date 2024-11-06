-- Table User
CREATE TABLE ms_user (
  id VARCHAR(100) PRIMARY KEY,
  username VARCHAR(50),
  password VARCHAR(100),
  role VARCHAR(8),
  is_active BOOLEAN
);

-- password = qwertyuiop
INSERT INTO ms_user (id, username, password, role, is_active) VALUES ('U001', 'Admin 1', '$2a$10$1J1xeBkXMxobbrAirqKv6Ow2kH5YyZBPnfb0bNZyBMxeoXMB8WviS', 'Admin', true);