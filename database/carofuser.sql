CREATE TABLE carofuser (
    car_license VARCHAR(20) NOT NULL,
    user_account VARCHAR(50) NOT NULL,

    PRIMARY KEY (car_license)
)  ENGINE=MYISAM character set utf8 collate utf8_unicode_ci