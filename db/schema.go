package db

const Schema = `
CREATE TABLE IF NOT EXISTS goly (
	id SERIAL PRIMARY KEY,
	redirect_url VARCHAR(255) NOT NULL,
	goly VARCHAR(255) NOT NULL UNIQUE,
	random BOOLEAN default False NOT NULL,
	count INT default 0 NOT NULL);`
