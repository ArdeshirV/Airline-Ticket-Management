delete from tickets ;
delete from payments ;
delete from order_items ;
delete from orders ;
delete from passengers ;
delete from users where username not in ('admin');
delete from roles where name not in ('admin','user');
delete from flights ;
delete from airplanes ;
delete from airports ;
delete from airlines ;
delete from cities ;