
create table deviceMaster
(
	id int not null AUTO_INCREMENT PRIMARY key,
	device varchar(100),
	unit varchar(2)
)

create table deviceValue
(
	id int not null AUTO_INCREMENT PRIMARY key,
	device_id int,
	value decimal(5,4),
	datetime datetime
)