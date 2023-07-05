insert into cities(id,name,created_at) 
	values
	(1001,'Tehran',NOW()), 
	(1002,'Shiraz',NOW()),
	(1003,'Tabriz',NOW()),
	(1004,'Isfahan',NOW()),
	(1005,'Bandar Abbas',NOW()),
	(1006,'Mashhad',NOW());
---------------------------------
insert into airports (id,name,city_id,terminal,created_at) 
	values
	(501,'Mehrabad International Airport',1001,6,NOW()),
	(502,'Shahid Dastgheib International',1002,1,NOW()),
	(503,'Tabriz International Airport',1003,1,NOW()),
	(504,'Shahid Beheshti International',1004,1,NOW()),
	(505,'Bandar Abbas International Airport',1005,1,NOW()),
	(506,'Shahid Hashemi Nejad Airport',1006,1,NOW());
-----------------------------------------------
insert into airlines (id,name,logo,created_at) 
	values 
	(701,'caspian','CPN',NOW()),
	(702,'IRANAIR','IRA',NOW()),
	(703,'ASEMAN','IRC',NOW()),
	(704,'ATAAIR','TBZ',NOW()),
	(705,'ZAGROS','IZG',NOW()),
	(706,'QESHM AIR','QSM',NOW());
--------------------------------------------
insert into airplanes (id,name,airline_id,capacity,created_at) 
	values 
	(2001,'McDonnell Douglas MD-80',701,130,NOW()),
	(2002,'Boeing 747',702,550,NOW()),
	(2003,'ATR 72',703,81,NOW()),
	(2004,'Airbus A320 family',704,180,NOW()),
	(2005,'Tupolev Tu-334',705,102,NOW()),
	(2006,'Fokker 100',706,117,NOW());
-------------------------------------------
insert into flights (id,flight_no,departure_id,destination_id,
					departure_time,arrival_time,airplane_id,
					flight_class,price,remaining_capacity,created_at) 
	values
	(301,'CPN6909',501,502,NOW()::date + interval '4 hours',NOW()::date + interval '6 hours',2001,2,12000000,130,NOW()),
	(302,'IZG4076',502,506,NOW()::date + interval '48 hours',NOW()::date + interval '52 hours',2005,1,14000000,102,NOW()),
	(303,'TBZ5662',505,503,NOW()::date - interval '5 hours',NOW()::date + interval '1 hours',2004,0,17000000,180,NOW()),
	(304,'IRA320',501,504,NOW()::date + interval '14 hours',NOW()::date + interval '15 hours',2002,2,9000000,550,NOW()),
	(305,'IRC645',504,506,NOW()::date + interval '10 hours',NOW()::date + interval '14 hours'+ interval '20 minute',2003,1,10000000,130,NOW()),
	(306,'QSM1290',506,505,NOW()::date + interval '2 hours'+ interval '30 minute',NOW()::date + interval '6 hours',2006,2,16000000,117,NOW());
	----------------------------------------------------------------
insert into roles  (id,name,description,created_at) 
   values
   (3,'Customer','Ticket buyer',NOW());
------------------------------------------
insert into users  (id,username,password,email,phone,role_id,created_at) 
   values
   (10,'RezaAhmadi','$2a$10$E1iADe9slYWiRtSu24h3uyNG1L/CSrjN9N7D6abGo9QxjFoinEjC','Reza.Ahmadi@godragon.com','09111111111',3,NOW());
 ------------------------------------------------------------------
insert into passengers  (id,first_name,last_name,national_code,gender,birth_date,user_id,created_at) 
	values 
	(401,'Reza','Ahmadi','9876543210',0,'1999-06-16',10,NOW()),
	(402,'Ehsan','Rezvani','1234567890',0,'2001-06-16',10,NOW()); 
------------------------------------------------------------
insert into orders (id,order_num,amount,flight_id,status,created_at) 
	values 
	(6,'654987321A',17000000,303,2,NOW()); 
-----------------------------------------------------------
insert into order_items  (id,passenger_id,order_id,created_at) 
	values 
	(10,401,6,NOW());
----------------------------------------------------------
insert into payments  (id,pay_amount,pay_time,payment_serial,order_id,created_at) 
	values 
	(1,17000000,now()::date - interval '10 hours','98765434321',6,NOW());
-----------------------------------------------------------
insert into tickets  (id,flight_id,passenger_id,payment_id,payment_status,user_id,refund,created_at) 
	values 
	(1,303,401,1,'Paid',10,false,NOW());
