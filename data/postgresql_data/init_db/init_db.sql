create table IF NOT EXISTS Radio.stations (
    id int primary key unique not null,
    slug varchar(255),
    name varchar,
    site varchar(255),
    email varchar(255),
    region varchar(255),
    city varchar(255),
    address varchar(255),
    facebook varchar(255),
    twitter varchar(255),
    ok varchar(255),
    vk varchar(255),
    wiki varchar(255),
    Genre varchar(255),
    PhoneNumber varchar(255),
    Stream text
);
