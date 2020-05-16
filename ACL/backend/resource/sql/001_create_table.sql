USE ACL;


# admin, read/write or read only
CREATE TABLE IF NOT EXISTS userRole(
	id 	INT		PRIMARY KEY,
	role 		VARCHAR(20)		UNIQUE KEY
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO userRole(id, role) VALUES(1, "admin");
INSERT INTO userRole(id, role) VALUES(2, "write");
INSERT INTO userRole(id, role) VALUES(3, "read");

CREATE TABLE IF NOT EXISTS users (
	id 	INT		AUTO_INCREMENT 		PRIMARY KEY,
	first_name	CHAR(25) 	NOT NULL,
	last_name	CHAR(25) 	NOT NULL,
	username	CHAR(25)	NOT NULL,
	password 	VARBINARY(128)	NOT NULL,
	creation_date	DATETIME	DEFAULT CURRENT_TIMESTAMP
)ENGINE = INNODB CHARACTER SET=utf8;

# groups table
CREATE TABLE IF NOT EXISTS groups1 (
	id 	INT 	AUTO_INCREMENT PRIMARY KEY,
	group_name	CHAR(25),
	group_info	CHAR(50),
	group_creation_date		DATETIME 	DEFAULT CURRENT_TIMESTAMP 
)ENGINE = INNODB CHARACTER SET=utf8;

# group-user table 
CREATE TABLE IF NOT EXISTS userGroup (
	id 	INT 	AUTO_INCREMENT UNIQUE KEY,
	user_id 	INT,
	group_id 	INT,
	PRIMARY KEY (user_id, group_id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (group_id) REFERENCES groups1(id)
)ENGINE = INNODB CHARACTER SET=utf8;

# file and folder is considered as single entity
CREATE TABLE IF NOT EXISTS files (
	id		INT 	AUTO_INCREMENT 	PRIMARY KEY,
	file_name	CHAR(20),
	file_type 	CHAR(20),
	parent_id 	INT
)ENGINE = INNODB CHARACTER SET=utf8;


# user file permission/role
CREATE TABLE IF NOT EXISTS userFileRole (
	id 	INT 	AUTO_INCREMENT UNIQUE KEY,
	user_id 	INT,
	file_id 	INT,
	role_id 	INT,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (file_id) REFERENCES files(id),
	FOREIGN KEY (role_id) REFERENCES userRole(id)
)ENGINE = INNODB CHARACTER SET=utf8;


# group file permission/role
CREATE TABLE IF NOT EXISTS groupFileRole (
	id 	INT 	AUTO_INCREMENT UNIQUE KEY,
	group_id 	INT,
	file_id 	INT,
	role_id 	INT,
	FOREIGN KEY (group_id) REFERENCES groups1(id),
	FOREIGN KEY (file_id) REFERENCES files(id),
	FOREIGN KEY (role_id) REFERENCES userRole(id)
);

