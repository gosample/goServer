# Create
CREATE TABLE parks(
  park_id VARCHAR(255) NOT NULL,
  park_name VARCHAR(255) NOT NULL,
  park_name VARCHAR(225) NOT NULL,
  park_num_of_storey INT NOT NULL,
  num_of_empty_parking INT NOT NULL,
  longitude DOUBLE NOT NULL,
  latitude DOUBLE NOT NULL,
  park_ip VARCHAR(20) NOT NULL,
  money DOUBLE,
  unit_price DOUBLE NOT NULL,
  book_uint_price DOUBLE NOT NULL,

  PRIMARY KEY(park_id)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;
