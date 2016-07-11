
DROP PROCEDURE IF EXISTS `del_bookservers` ;
 

DELIMITER $$
CREATE PROCEDURE del_bookservers()
BEGIN

  DELETE FROM bookservices WHERE UNIX_TIMESTAMP(time_stamp) <= UNIX_TIMESTAMP() - hours*3600;
END ;

 
CREATE EVENT myevent
    ON SCHEDULE EVERY 60 SECOND
    DO
      CALL del_bookservers();
       
