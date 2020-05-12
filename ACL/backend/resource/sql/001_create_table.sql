USE ACL;

// admin, read/write or read only
CREATE TABLE IF NOT EXISTS userRole(
	role_id 	INT		PRIMARY KEY,
	role 		VARCHAR(20)		UNIQUE KEY
);

// user tabel
CREATE TABLE IF NOT EXISTS user (
	user_id 	INT		AUTO_INCREMENT 		PRIMARY KEY,
	first_name	CHAR(25) 	NOT NULL,
	last_name	CHAR(25) 	NOT NULL,
	password 	VARBINARY(128)	NOT NULL,
	creation_date	DATETIME	DEFAULT CURRENT_TIMESTAMP
);

// group table
CREATE TABLE IF NOT EXISTS group (
	group_id 	INT 	AUTO_INCREMENT PRIMARY KEY,
	group_name	CHAR(25),
	group_info	CHAR(50),
	group_creation_date		DATETIME 	DEFAULT CURRENT_TIMESTAMP 
);

// group-user table 
CREATE TABLE IF NOT EXISTS userGroup (
	user_id 	INT,
	group_id 	INT,
	PRIMARY KEY (user_id, group_id),
	FOREIGN KEY (user_id) REFERENCES user(user_id),
	FOREIGN KEY (group_id) REFERENCES group(group_id)
);

// file and folder is considered as single entity
CREATE TABLE IF NOT EXISTS file (
	file_id		INT 	AUTO_INCREMENT 	PRIMARY KEY,
	file_name	CHAR(20),
	file_type 	CHAR(20),
	parent_id 	INT,
	FOREIGN KEY (parent_id) REFERENCES file(file_id)
);

// user file permission/role
CREATE TABLE IF NOT EXISTS userFileRole (
	user_id 	INT,
	file_id 	INT,
	role_id 	INT,
	FOREIGN KEY (user_id) REFERENCES user(user_id),
	FOREIGN KEY (file_id) REFERENCES file(file_id),
	FOREIGN KEY (role_id) REFERENCES userRole(role_id)
);

// group file permission/role
CREATE TABLE IF NOT EXISTS groupFileRole (
	group_id 	INT,
	file_id 	INT,
	role_id 	INT,
	FOREIGN KEY (group_id) REFERENCES group(group_id),
	FOREIGN KEY (file_id) REFERENCES file(file_id),
	FOREIGN KEY (role_id) REFERENCES userRole(role_id)
);



