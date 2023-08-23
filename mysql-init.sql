create database if not exists rinha;

use rinha;

create table if not exists people (
	id text,
	name varchar(100),
	nickname varchar(32),
	birthdate varchar(10),
	unique(nickname)
);

create table if not exists person_stack (
	id text,
	stack text
);
