# Create
CREATE TABLE parkingspace(
  space_id VARCHAR(32) NOT NULL,
  park_id VARCHAR(255) NOT NULL,
  state INT NOT NULL,
  floor INT NOT NULL,
  user_account VARCHAR(50) NOT NULL,
  assign_time TIMESTAMP,

  PRIMARY KEY(space_id, park_id)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;

# Insert
INSERT
INTO
  `parkingspace`(
    `space_id`,
    `park_id`,
    `state`,
    `floor`,
    `assign_time`
  )
VALUES('1', '1', '0', '1', CURRENT_TIMESTAMP);
