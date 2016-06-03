CREATE TABLE carofuser (
    id INT NOT NULL AUTO_INCREMENT,
    car_license VARCHAR(20),
    user_account VARCHAR(50) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (car_license) REFERENCES users(car_license) ON DELETE CASCADE
)  ENGINE=MYISAM character set utf8 collate utf8_unicode_ci