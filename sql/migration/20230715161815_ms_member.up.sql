-- Table ms_member
CREATE TABLE ms_member (
  id VARCHAR(100) PRIMARY KEY,
  id_customer VARCHAR(50) REFERENCES ms_customer (id),
  type VARCHAR(50) NOT NULL,
  expire TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NULL,
  updated_by VARCHAR(50) NULL
);