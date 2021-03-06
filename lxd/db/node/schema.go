package node

// DO NOT EDIT BY HAND
//
// This code was generated by the schema.DotGo function. If you need to
// modify the database schema, please add a new schema update to update.go
// and the run 'make update-schema'.
const freshSchema = `
CREATE TABLE config (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    key VARCHAR(255) NOT NULL,
    value TEXT,
    UNIQUE (key)
);
CREATE TABLE patches (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    applied_at DATETIME NOT NULL,
    UNIQUE (name)
);
CREATE TABLE raft_nodes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    address TEXT NOT NULL,
    UNIQUE (address)
);

INSERT INTO schema (version, updated_at) VALUES (38, strftime("%s"))
`
