CREATE DATABASE accounts;

\c accounts

CREATE TABLE users (
    userID UUID PRIMARY KEY,
    username VARCHAR NOT NULL, 
    password VARCHAR NOT NULL,
    createdAt TIMESTAMP DEFAULT NOW() 
);

CREATE INDEX UsernameInex ON users USING hash (username);
