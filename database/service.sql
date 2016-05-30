# Create
CREATE TABLE service(
  UserAccount VARCHAR(50) NOT NULL,
  CarLicense VARCHAR(20) NOT NULL,
  ParkID VARCHAR(255) NOT NULL,
  SpaceID VARCHAR(255) NOT NULL,
  StartTime TIMESTAMP NOT NULL,
  ExitTime TIMESTAMP NOT NULL,
  TotalMoney DOUBLE NOT NULL,

  PRIMARY KEY(UserAccount, CarLicense)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;

# Insert
INSERT
INTO
  `parks`(
    `ParkID`,
    `ParkName`,
    `ParkPwd`,
    `ParkNumOfStorey`,
    `NumOfEmptyParking`,
    `Longitude`,
    `Latitude`,
    `ParkIP`,
    `Money`,
    `UnitPrice`,
    `BookUintPrice`
  )
VALUES(
  '1',
  'a',
  'a',
  '100',
  '20',
  '120.1',
  '30.1',
  '20.199.1.2',
  NULL,
  '2',
  '1'
)
