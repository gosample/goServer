# Create
CREATE TABLE parks(
  park_id VARCHAR(255) NOT NULL,
  park_name VARCHAR(255) NOT NULL,
  park_pwd VARCHAR(225) NOT NULL,
  storey_num INT NOT NULL,
  empty_num INT NOT NULL,
  longitude DOUBLE NOT NULL,
  latitude DOUBLE NOT NULL,
  park_ip VARCHAR(20) NOT NULL,
  money DOUBLE,
  unit_price DOUBLE NOT NULL,
  book_uint_price DOUBLE NOT NULL,
  park_pt POINT,

  PRIMARY KEY(park_id)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;
