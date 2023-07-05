CREATE SCHEMA IF NOT EXISTS "public";

CREATE SEQUENCE "public".airlines_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".airplanes_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".airports_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".cities_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".flights_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".orders_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".passengers_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".payments_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".roles_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".tickets_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".users_id_seq START WITH 1 INCREMENT BY 1;

CREATE SEQUENCE "public".order_items_id_seq START WITH 1 INCREMENT BY 1;


CREATE  TABLE "public".airlines ( 
	id                   bigint DEFAULT nextval('airlines_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar(255)  NOT NULL  ,
	logo                 varchar(255)    ,
	CONSTRAINT airlines_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_airlines_deleted_at ON "public".airlines  ( deleted_at );

CREATE  TABLE "public".airplanes ( 
	id                   bigint DEFAULT nextval('airplanes_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar(255)  NOT NULL  ,
	airline_id           bigint    ,
	capacity             bigint    ,
	CONSTRAINT airplanes_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_airplanes_deleted_at ON "public".airplanes  ( deleted_at );

CREATE  TABLE "public".cities ( 
	id                   bigint DEFAULT nextval('cities_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar(255)  NOT NULL  ,
	CONSTRAINT cities_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_cities_deleted_at ON "public".cities  ( deleted_at );

CREATE INDEX idx_cities_name ON "public".cities  ( name );

CREATE  TABLE "public".roles ( 
	id                   bigint DEFAULT nextval('roles_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar(255)  NOT NULL  ,
	description          varchar(255)    ,
	CONSTRAINT roles_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_roles_deleted_at ON "public".roles  ( deleted_at );

CREATE  TABLE "public".users ( 
	id                   bigint DEFAULT nextval('users_id_seq'::regclass) NOT NULL  ,
	username             varchar(255)  NOT NULL  ,
	"password"           varchar(255)  NOT NULL  ,
	email                varchar(255)    ,
	phone                varchar(255)    ,
	created_at           timestamptz    ,
	role_id              bigint    ,
	is_login_required    boolean DEFAULT false   ,
	CONSTRAINT users_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_users_username ON "public".users  ( username );

CREATE INDEX idx_users_phone ON "public".users  ( phone );

CREATE INDEX idx_users_email ON "public".users  ( email );

CREATE  TABLE "public".airports ( 
	id                   bigint DEFAULT nextval('airports_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar(255)  NOT NULL  ,
	city_id              bigint    ,
	terminal             varchar(255)    ,
	CONSTRAINT airports_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_airports_deleted_at ON "public".airports  ( deleted_at );

CREATE  TABLE "public".flights ( 
	id                   bigint DEFAULT nextval('flights_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	flight_no            varchar(255)  NOT NULL  ,
	departure_id         bigint    ,
	destination_id       bigint    ,
	departure_time       timestamptz  NOT NULL  ,
	arrival_time         timestamptz    ,
	airplane_id          bigint    ,
	flight_class         varchar(255)    ,
	price                bigint  NOT NULL  ,
	remaining_capacity   bigint  NOT NULL  ,
	cancel_condition     text    ,
	CONSTRAINT flights_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_flights_deleted_at ON "public".flights  ( deleted_at );

CREATE INDEX idx_flights_flight_no ON "public".flights  ( flight_no );

CREATE INDEX idx_flights_price ON "public".flights  ( price );

CREATE INDEX idx_flights_departure_time ON "public".flights  ( departure_time );

CREATE  TABLE "public".orders ( 
	id                   bigint  DEFAULT nextval('orders_id_seq'::regclass) NOT NULL,
	order_num            varchar(255)  NOT NULL  ,
	amount               bigint  NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	flight_id            bigint  NOT NULL  ,
	user_id	 			 bigint  NOT NULL  ,
	status               bigint  NOT NULL  ,
	CONSTRAINT pk_orders PRIMARY KEY ( id )
 );

CREATE INDEX idx_orders_deleted_at ON "public".orders  ( deleted_at );

CREATE INDEX idx_orders_order_num ON "public".orders USING hash ( order_num );

CREATE INDEX idx_orders_status ON "public".orders USING hash ( status );

CREATE  TABLE "public".passengers ( 
	id                   bigint DEFAULT nextval('passengers_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	first_name           varchar(255)  NOT NULL  ,
	last_name            varchar(255)  NOT NULL  ,
	national_code        varchar(255)  NOT NULL  ,
	email                varchar(255)    ,
	gender               bigint  NOT NULL  ,
	phone                varchar(255)  ,
	birth_date           timestamptz  NOT NULL  ,
	address              text    ,
	user_id              bigint    ,
	CONSTRAINT passengers_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_passengers_deleted_at ON "public".passengers  ( deleted_at );

CREATE  TABLE "public".payments ( 
	id                   bigint DEFAULT nextval('payments_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	pay_amount           bigint    ,
	pay_time             timestamptz    ,
	payment_serial       varchar(255)    ,
	order_id             integer    ,
	CONSTRAINT payments_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_payments_deleted_at ON "public".payments  ( deleted_at );

CREATE  TABLE "public".tickets ( 
	id                   bigint DEFAULT nextval('tickets_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	flight_id            bigint    ,
	passenger_id         bigint    ,
	payment_id           bigint    ,
	payment_status		 varchar(255),
	user_id              bigint    ,
	refund               boolean    ,
	CONSTRAINT tickets_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_tickets_deleted_at ON "public".tickets  ( deleted_at );

CREATE  TABLE "public".order_items ( 
	id                   bigint   DEFAULT nextval('order_items_id_seq'::regclass) NOT NULL, 
	passenger_id         bigint  NOT NULL  ,
	order_id             bigint  NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	CONSTRAINT pk_order_items PRIMARY KEY ( id )
 );

CREATE INDEX idx_order_items_deleted_at ON "public".order_items  ( deleted_at );

ALTER TABLE "public".airplanes ADD CONSTRAINT fk_airplanes_airline FOREIGN KEY ( airline_id ) REFERENCES "public".airlines( id );

ALTER TABLE "public".airports ADD CONSTRAINT fk_airports_city FOREIGN KEY ( city_id ) REFERENCES "public".cities( id );

ALTER TABLE "public".flights ADD CONSTRAINT fk_flights_airplane FOREIGN KEY ( airplane_id ) REFERENCES "public".airplanes( id );

ALTER TABLE "public".flights ADD CONSTRAINT fk_flights_departure FOREIGN KEY ( departure_id ) REFERENCES "public".airports( id );

ALTER TABLE "public".flights ADD CONSTRAINT fk_flights_destination FOREIGN KEY ( destination_id ) REFERENCES "public".airports( id );

ALTER TABLE "public".order_items ADD CONSTRAINT fk_order_items_order_id FOREIGN KEY ( order_id ) REFERENCES "public".orders( id );

ALTER TABLE "public".order_items ADD CONSTRAINT fk_order_items_passengers FOREIGN KEY ( passenger_id ) REFERENCES "public".passengers( id );

ALTER TABLE "public".orders ADD CONSTRAINT fk_orders_flight_id FOREIGN KEY ( flight_id ) REFERENCES "public".flights( id );

ALTER TABLE "public".orders ADD CONSTRAINT fk_orders_user_id FOREIGN KEY ( user_id ) REFERENCES "public".users( id );

ALTER TABLE "public".passengers ADD CONSTRAINT fk_users_passengers FOREIGN KEY ( user_id ) REFERENCES "public".users( id );

ALTER TABLE "public".payments ADD CONSTRAINT fk_payments_orders FOREIGN KEY ( order_id ) REFERENCES "public".orders( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_tickets_flight FOREIGN KEY ( flight_id ) REFERENCES "public".flights( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_passengers_tickets FOREIGN KEY ( passenger_id ) REFERENCES "public".passengers( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_tickets_payment FOREIGN KEY ( payment_id ) REFERENCES "public".payments( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_tickets_user FOREIGN KEY ( user_id ) REFERENCES "public".users( id );

ALTER TABLE "public".users ADD CONSTRAINT fk_users_role FOREIGN KEY ( role_id ) REFERENCES "public".roles( id );
