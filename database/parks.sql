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
  park_address VARCHAR(255),

  PRIMARY KEY(park_id)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;

# INSERT INTO `parking`.parks VALUES ("9","浙江医药","qwerty",100,100,121.560883,29.821188,"10.80.136.52",5,5,5,ST_GEOMFROMTEXT('POINT(121.560883 29.821188)'),"http://upload.admin5.com/2015/0728/1438073690416.jpg")
