CREATE TABLE carofuser (
    id INT NOT NULL AUTO_INCREMENT,
    car_license VARCHAR(20) NOT NULL,
    user_account VARCHAR(50) NOT NULL,

    PRIMARY KEY (id)
)  ENGINE=MYISAM character set utf8 collate utf8_unicode_ci