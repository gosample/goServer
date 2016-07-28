# Create
CREATE TABLE service(
  service_id INT NOT NULL AUTO_INCREMENT,
  user_account VARCHAR(50) NOT NULL,
  car_license VARCHAR(20) NOT NULL,
  hours DOUBLE NOT NULL,
  park_id VARCHAR(255) NOT NULL,
  start_time DATATIME NOT NULL,
  exit_time DATATIME NOT NULL,
  total_money DOUBLE NOT NULL,

  PRIMARY KEY(service_id)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;

