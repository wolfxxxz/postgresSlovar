CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null unique,
    password varchar(255) not null,
    email varchar(255) not null
);

CREATE TABLE words 
(
	id serial not null unique,
	english varchar(255) not null,
	russian varchar(255) not null,
	theme varchar(255) not null,
	right_answer integer
);

CREATE TABLE words_learn 
(
	id integer not null unique,
	english varchar(255) not null,
	russian varchar(255) not null,
	theme varchar(255) not null
);
 