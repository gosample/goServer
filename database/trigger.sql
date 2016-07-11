
CREATE TRIGGER `add_empty` AFTER DELETE ON `bookservices`
FOR EACH ROW update `parks` set empty_num = empty_num + 1 where park_id = old.park_id;


CREATE TRIGGER `remove_empty` AFTER INSERT ON `bookservices`
FOR EACH ROW update `parks` set empty_num = empty_num -1 where park_id = new.park_id;


CREATE TRIGGER `add_service` AFTER DELETE ON `bookservices`
FOR EACH ROW insert into `service` values (old.`user_account`,old.`car_license`,old.`hours`,old.`park_id`,old.`time_stamp`);