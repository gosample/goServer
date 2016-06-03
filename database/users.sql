# Create
CREATE TABLE users (
	user_account VARCHAR(50) NOT NULL,
    user_pwd VARCHAR(50) NOT NULL,
    money DOUBLE,
    token VARCHAR(100),

    PRIMARY KEY (user_account)
) ENGINE=MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;
