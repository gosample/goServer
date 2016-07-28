
CREATE TRIGGER `add_empty` AFTER DELETE ON `bookservices`
FOR EACH ROW update `parks` set empty_num = empty_num + 1 where park_id = old.park_id;


CREATE TRIGGER `remove_empty` AFTER INSERT ON `bookservices`
FOR EACH ROW update `parks` set empty_num = empty_num -1 where park_id = new.park_id;


CREATE TRIGGER `add_service` AFTER DELETE ON `bookservices`
FOR EACH ROW insert into `service` values (old.`user_account`,old.`car_license`,old.`hours`,old.`park_id`,old.`book_time`);



CREATE TRIGGER `mark_state` AFTER INSERT ON `bookservices`
FOR EACH ROW update `parkingspace` set state = 1 where park_id = new.park_id and space_id = new.space_id;

CREATE TRIGGER `free_state` AFTER DELETE ON `bookservices`
FOR EACH ROW update `parkingspace` set state = 0 where park_id = old.park_id and space_id = old.space_id;