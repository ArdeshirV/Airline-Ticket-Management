CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE "public".airlines ( 
	id                   bigint DEFAULT nextval('airlines_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar  NOT NULL  ,
	logo                 varchar    ,
	CONSTRAINT airlines_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_airlines_deleted_at ON "public".airlines  ( deleted_at );

CREATE  TABLE "public".airplanes ( 
	id                   bigint DEFAULT nextval('airplanes_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar  NOT NULL  ,
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
	name                 varchar  NOT NULL  ,
	CONSTRAINT cities_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_cities_deleted_at ON "public".cities  ( deleted_at );

CREATE INDEX idx_cities_name ON "public".cities USING btree ( name );

CREATE  TABLE "public".payments ( 
	id                   bigint DEFAULT nextval('payments_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	pay_amount           bigint    ,
	pay_time             timestamptz    ,
	payment_serial       varchar    ,
	CONSTRAINT payments_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_payments_deleted_at ON "public".payments  ( deleted_at );

CREATE  TABLE "public".roles ( 
	id                   bigint DEFAULT nextval('roles_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar  NOT NULL  ,
	description          text    ,
	CONSTRAINT roles_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_roles_deleted_at ON "public".roles  ( deleted_at );

CREATE  TABLE "public".users ( 
	id                   bigint DEFAULT nextval('users_id_seq'::regclass) NOT NULL  ,
	username             varchar  NOT NULL  ,
	"password"           varchar  NOT NULL  ,
	email                varchar    ,
	phone                varchar    ,
	created_at           timestamptz    ,
	role_id              bigint    ,
	is_login_required    boolean DEFAULT false   ,
	CONSTRAINT users_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_users_username ON "public".users USING hash ( username );

CREATE INDEX idx_users_phone ON "public".users USING btree ( phone );

CREATE INDEX idx_users_email ON "public".users USING hash ( email );

CREATE  TABLE "public".airports ( 
	id                   bigint DEFAULT nextval('airports_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	name                 varchar  NOT NULL  ,
	city_id              bigint    ,
	terminal             varchar    ,
	CONSTRAINT airports_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_airports_deleted_at ON "public".airports  ( deleted_at );

CREATE  TABLE "public".flights ( 
	id                   bigint DEFAULT nextval('flights_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	flight_no            varchar  NOT NULL  ,
	departure_id         bigint    ,
	destination_id       bigint    ,
	departure_time       timestamptz  NOT NULL  ,
	arrival_time         timestamptz    ,
	airplane_id          bigint    ,
	flight_class         bigint    ,
	price                bigint  NOT NULL  ,
	remaining_capacity   bigint  NOT NULL  ,
	cancel_condition     varchar    ,
	CONSTRAINT flights_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_flights_deleted_at ON "public".flights  ( deleted_at );

CREATE INDEX idx_flights_flight_no ON "public".flights USING btree ( flight_no );

CREATE INDEX idx_flights_price ON "public".flights USING btree ( price );

CREATE INDEX idx_flights_departure_time ON "public".flights USING btree ( departure_time );

CREATE  TABLE "public".passengers ( 
	id                   bigint DEFAULT nextval('passengers_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	first_name           varchar  NOT NULL  ,
	last_name            varchar  NOT NULL  ,
	national_code        varchar  NOT NULL  ,
	email                varchar    ,
	gender               bigint  NOT NULL  ,
	phone                varchar  NOT NULL  ,
	birth_date           timestamptz  NOT NULL  ,
	address              text    ,
	user_id              bigint    ,
	CONSTRAINT passengers_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_passengers_deleted_at ON "public".passengers  ( deleted_at );

CREATE  TABLE "public".tickets ( 
	id                   bigint DEFAULT nextval('tickets_id_seq'::regclass) NOT NULL  ,
	created_at           timestamptz    ,
	updated_at           timestamptz    ,
	deleted_at           timestamptz    ,
	flight_id            bigint    ,
	passenger_id         bigint    ,
	payment_id           bigint    ,
	user_id              bigint    ,
	payment_status       varchar    ,
	refund               boolean    ,
	CONSTRAINT tickets_pkey PRIMARY KEY ( id )
 );

CREATE INDEX idx_tickets_deleted_at ON "public".tickets  ( deleted_at );

ALTER TABLE "public".airplanes ADD CONSTRAINT fk_airplanes_airline FOREIGN KEY ( airline_id ) REFERENCES "public".airlines( id );

ALTER TABLE "public".airports ADD CONSTRAINT fk_airports_city FOREIGN KEY ( city_id ) REFERENCES "public".cities( id );

ALTER TABLE "public".flights ADD CONSTRAINT fk_flights_airplane FOREIGN KEY ( airplane_id ) REFERENCES "public".airplanes( id );

ALTER TABLE "public".flights ADD CONSTRAINT fk_flights_departure FOREIGN KEY ( departure_id ) REFERENCES "public".airports( id );

ALTER TABLE "public".flights ADD CONSTRAINT fk_flights_destination FOREIGN KEY ( destination_id ) REFERENCES "public".airports( id );

ALTER TABLE "public".passengers ADD CONSTRAINT fk_users_passengers FOREIGN KEY ( user_id ) REFERENCES "public".users( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_tickets_flight FOREIGN KEY ( flight_id ) REFERENCES "public".flights( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_passengers_tickets FOREIGN KEY ( passenger_id ) REFERENCES "public".passengers( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_tickets_payment FOREIGN KEY ( payment_id ) REFERENCES "public".payments( id );

ALTER TABLE "public".tickets ADD CONSTRAINT fk_tickets_user FOREIGN KEY ( user_id ) REFERENCES "public".users( id );

ALTER TABLE "public".users ADD CONSTRAINT fk_users_role FOREIGN KEY ( role_id ) REFERENCES "public".roles( id );
