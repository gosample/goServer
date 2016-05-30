CREATE TABLE carofuser (
    ID INT NOT NULL AUTO_INCREMENT,
    CarLicense VARCHAR(20),
    UserAccount VARCHAR(50) NOT NULL,

    PRIMARY KEY (ID),
    FOREIGN KEY (CarLicense) REFERENCES users(CarLicense) ON DELETE CASCADE
)  ENGINE=MYISAM character set utf8 collate utf8_unicode_ci