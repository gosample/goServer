# Create
CREATE TABLE service(
  user_account VARCHAR(50) NOT NULL,
  car_license VARCHAR(20) NOT NULL,
  park_id VARCHAR(255) NOT NULL,
  space_id VARCHAR(255) NOT NULL,
  start_time TIMESTAMP NOT NULL,
  exit_time TIMESTAMP NOT NULL,
  total_money DOUBLE NOT NULL,

  PRIMARY KEY(user_account, car_license)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;

