CREATE TABLE bookservices(
  user_account VARCHAR(50) NOT NULL,
  car_license VARCHAR(20) NOT NULL,
  hours DOUBLE NOT NULL,
  park_id VARCHAR(255)
  time_stamp TIMESTAMP NOT NULL,

  PRIMARY KEY(user_account, car_license)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;