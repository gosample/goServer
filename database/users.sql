# Create
CREATE TABLE users (
	UserAccount VARCHAR(50) NOT NULL,
    UserPwd VARCHAR(50) NOT NULL,
    PhoneNum VARCHAR(50) NOT NULL,
    Money DOUBLE,
    Token VARCHAR(100),

    PRIMARY KEY (UserAccount)
) ENGINE=MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;
