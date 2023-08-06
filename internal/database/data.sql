DROP TABLE IF EXISTS public.countries CASCADE;
DROP TABLE IF EXISTS public.cities CASCADE;
DROP TABLE IF EXISTS public.airlines CASCADE;
DROP TABLE IF EXISTS public.passengers CASCADE;
DROP TABLE IF EXISTS public.directions CASCADE;
DROP TABLE IF EXISTS public.flights CASCADE;
DROP TABLE IF EXISTS public.bookings CASCADE;

CREATE TABLE public.countries (
    ID SERIAL PRIMARY KEY,
    "name" varchar(255) NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE TABLE public.cities (
    ID SERIAL PRIMARY KEY,
    "name" varchar(255) NULL,
    country_id int8 NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE TABLE public.airlines (
    ID SERIAL PRIMARY KEY,
    "name" varchar(255) NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE TABLE public.passengers (
    ID SERIAL PRIMARY KEY,
    "name" varchar(255) NULL,
    passport varchar(50) NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE TABLE public.directions (
    ID SERIAL PRIMARY KEY,
    origin_city_id int8 NULL,
    destination_city_id int8 NULL,
    airline_id int8 NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE TABLE public.flights (
    ID SERIAL PRIMARY KEY,
    flight_number varchar(255) NULL,
    direction_id int8 NULL,
    departure_time timestamptz NULL,
    arrival_time timestamptz NULL,
    seats_number int8 NULL,
    price float8 NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

CREATE TABLE public.bookings (
    ID SERIAL PRIMARY KEY,
    booking_number text NULL,
    flight_id int8 NULL,
    passenger_id int8 NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

--countries
insert into public.countries (name, created_at) values ('Kazakhstan', current_timestamp);
insert into public.countries (name, created_at) values ('Turkey', current_timestamp);
insert into public.countries (name, created_at) values ('USA', current_timestamp);
insert into public.countries (name, created_at) values ('Australia', current_timestamp);
insert into public.countries (name, created_at) values ('China', current_timestamp);

--cities
insert into public.cities (name, country_id, created_at) Select 'Almaty', id, current_timestamp from public.countries where name = 'Kazakhstan';
insert into public.cities (name, country_id, created_at) Select 'Astana', id, current_timestamp from public.countries where name = 'Kazakhstan';
insert into public.cities (name, country_id, created_at) Select 'Istanbul', id, current_timestamp from public.countries where name = 'Turkey';
insert into public.cities (name, country_id, created_at) Select 'Antalya', id, current_timestamp from public.countries where name = 'Turkey';
insert into public.cities (name, country_id, created_at) Select 'Ankara', id, current_timestamp from public.countries where name = 'Turkey';
insert into public.cities (name, country_id, created_at) Select 'New York', id, current_timestamp from public.countries where name = 'USA';
insert into public.cities (name, country_id, created_at) Select 'Los Angeles', id, current_timestamp from public.countries where name = 'USA';
insert into public.cities (name, country_id, created_at) Select 'Sydney', id, current_timestamp from public.countries where name = 'Australia';
insert into public.cities (name, country_id, created_at) Select 'Canberra', id, current_timestamp from public.countries where name = 'Australia';
insert into public.cities (name, country_id, created_at) Select 'Beijing', id, current_timestamp from public.countries where name = 'China';
insert into public.cities (name, country_id, created_at) Select 'Shanghai', id, current_timestamp from public.countries where name = 'China';
insert into public.cities (name, country_id, created_at) Select 'Suzhou', id, current_timestamp from public.countries where name = 'China';

--airlines
insert into public.airlines (name, created_at) values ('Air Astana', current_timestamp);
insert into public.airlines (name, created_at) values ('Turkish Airlines', current_timestamp);
insert into public.airlines (name, created_at) values ('American Airlines', current_timestamp);
insert into public.airlines (name, created_at) values ('China Airlines', current_timestamp);
insert into public.airlines (name, created_at) values ('Scat', current_timestamp);

--passengers
insert into public.passengers (name, passport, created_at) values ('ZAURE MAZHIKOVA', 71812828, current_timestamp);
insert into public.passengers (name, passport, created_at) values ('ANARA SAGATBEKOVA', 24848747, current_timestamp);
insert into public.passengers (name, passport, created_at) values ('ASSEL JEXEMBEKOVA', 09238498, current_timestamp);
insert into public.passengers (name, passport, created_at) values ('KAMIZHAN ISSANOVA', 09834841, current_timestamp);
insert into public.passengers (name, passport, created_at) values ('MARYA KIM', 12387890, current_timestamp);
insert into public.passengers (name, passport, created_at) values ('ZAMZAGUL SADIMOVA', 42897992, current_timestamp);

--directions
insert into public.directions
    (origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Almaty') as origin,
     (Select id from public.cities where name = 'Istanbul') as destination,
     (Select id from public.airlines where name = 'Air Astana') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Almaty') as origin,
     (Select id from public.cities where name = 'Astana') as destination,
     (Select id from public.airlines where name = 'Air Astana') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Almaty') as origin,
     (Select id from public.cities where name = 'Ankara') as destination,
     (Select id from public.airlines where name = 'Air Astana') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Almaty') as origin,
     (Select id from public.cities where name = 'Antalya') as destination,
     (Select id from public.airlines where name = 'Air Astana') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Antalya') as origin,
     (Select id from public.cities where name = 'Ankara') as destination,
     (Select id from public.airlines where name = 'Turkish Airlines') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Almaty') as origin,
     (Select id from public.cities where name = 'Ankara') as destination,
     (Select id from public.airlines where name = 'Turkish Airlines') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Almaty') as origin,
     (Select id from public.cities where name = 'Istanbul') as destination,
     (Select id from public.airlines where name = 'Turkish Airlines') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Istanbul') as origin,
     (Select id from public.cities where name = 'Astana') as destination,
     (Select id from public.airlines where name = 'Turkish Airlines') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Ankara') as origin,
     (Select id from public.cities where name = 'Antalya') as destination,
     (Select id from public.airlines where name = 'Turkish Airlines') as airline;
insert into public.directions
(origin_city_id, destination_city_id, airline_id, created_at) Select origin.id, destination.id, airline.id, current_timestamp
from (Select id from public.cities where name = 'Antalya') as origin,
     (Select id from public.cities where name = 'Istanbul') as destination,
     (Select id from public.airlines where name = 'Turkish Airlines') as airline;


--flights
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A782', 1, date('2023-09-24 10:00:00'), date('2023-09-24 19:00:00'), 110, 110000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A797', 2, date('2023-09-25 10:00:00'), date('2023-09-25 19:00:00'), 100, 185000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A745', 3, date('2023-09-30 10:00:00'), date('2023-09-30 19:00:00'), 70, 145000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A709', 4, date('2023-09-10 10:00:00'), date('2023-09-10 19:00:00'), 150, 140000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A134', 5, date('2023-09-09 10:00:00'), date('2023-09-09 19:00:00'), 120, 130000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A567', 6, date('2023-09-15 10:00:00'), date('2023-09-15 19:00:00'), 110, 200000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A354', 7, date('2023-09-18 10:00:00'), date('2023-09-18 19:00:00'), 90, 180000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A745', 8, date('2023-09-23 10:00:00'), date('2023-09-23 19:00:00'), 180, 150000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A798', 9, date('2023-09-30 21:00:00'), date('2023-09-30 23:00:00'), 180, 50000, current_timestamp);
insert into public.flights (flight_number, direction_id, departure_time, arrival_time, seats_number, price, created_at)
values ('A545', 10, date('2023-10-01 03:00:00'), date('2023-10-30 05:00:00'), 180, 30000, current_timestamp);


--bookings
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('AT8745744', 1, 1, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('AT8734564', 1, 2, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('OR834Y893', 1, 3, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('TR893U432', 1, 4, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('PQ9834934', 1, 5, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('EU4938423', 1, 6, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('UE8389393', 1, 7, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('QO239184U', 1, 8, current_timestamp);

insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('IE9830830', 2, 1, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('QP3938983', 2, 2, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('BR9390393', 2, 3, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('OX8389383', 2, 4, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('AP2389389', 2, 5, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('KS8938938', 2, 6, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('LA8938E8E', 2, 7, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('DH2390390', 2, 8, current_timestamp);

insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('WO4095489', 3, 1, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('EO8934834', 3, 2, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('OD8484843', 3, 3, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('VR8938923', 3, 4, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('CM3993934', 3, 5, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('XO4893489', 3, 6, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('KE8938943', 3, 7, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('PZ83E8948', 3, 8, current_timestamp);

insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('QI3843849', 4, 1, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('OC9393934', 4, 2, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('SJ3892389', 4, 3, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('AK8934834', 4, 4, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('LD8934E83', 4, 5, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('OW8923483', 4, 6, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('DJ3894834', 4, 7, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('SH90494I4', 4, 8, current_timestamp);

insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('DU348934U', 5, 1, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('JE8348344', 5, 2, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('OC8348483', 5, 3, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('EU9023490', 5, 4, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('HF8348344', 5, 5, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('SJ8238032', 5, 6, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('KE9384844', 5, 7, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('LV8438423', 5, 8, current_timestamp);

insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('VF8348934', 6, 1, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('BG8343824', 6, 2, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('MH9304903', 6, 3, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('XS8293823', 6, 4, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('ME8934823', 6, 5, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('KC83E84RE', 6, 6, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('ID8934E84', 6, 7, current_timestamp);
insert into public.bookings (booking_number, flight_id, passenger_id, created_at) values ('JE2893813', 6, 8, current_timestamp);
