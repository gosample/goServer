# Create
CREATE TABLE parkingspace(
  SpaceID VARCHAR(32) NOT NULL,
  ParkID VARCHAR(255) NOT NULL,
  State INT NOT NULL,
  Floor INT NOT NULL,
  UserAccount VARCHAR(50) NOT NULL,
  AssignTime TIMESTAMP,

  PRIMARY KEY(SpaceID, ParkID)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;

# Insert
INSERT
INTO
  `parkingspace`(
    `SpaceID`,
    `ParkID`,
    `State`,
    `Floor`,
    `UserAccount`,
    `AssignTime`
  )
VALUES('1', '1', '0', '1', 'Linus', CURRENT_TIMESTAMP);
