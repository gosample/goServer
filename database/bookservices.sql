CREATE TABLE bookservices(
  UserAccount VARCHAR(50) NOT NULL,
  CarLicense VARCHAR(20) NOT NULL,
  Hours DOUBLE NOT NULL,
  StartTime TIMESTAMP NOT NULL,

  PRIMARY KEY(UserAccount, CarLicense)
) ENGINE = MYISAM CHARACTER SET utf8 COLLATE utf8_unicode_ci;