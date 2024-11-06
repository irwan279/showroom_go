-- Table ms_customer
CREATE TABLE ms_customer (
  id VARCHAR(100) PRIMARY KEY,
  id_user VARCHAR(50) REFERENCES ms_user (id) NOT NULL,
  full_name VARCHAR(100) NOT NULL, 
  NIK VARCHAR(16) NOT NULL,
  noPhone VARCHAR(15) NOT NULL,
  email VARCHAR(100) NOT NULL,
  address VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NULL,
  updated_by VARCHAR(50) NULL,
  UNIQUE(email),
  UNIQUE(NIK)
);