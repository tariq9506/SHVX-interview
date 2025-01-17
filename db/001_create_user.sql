create table shvx_user (
    id serial primary key,
    phone_number varchar (15) unique,
    email varchar(50) unique NOT NULL,
    password varchar(15) NOT NULL,
    user_name varchar(50)
);