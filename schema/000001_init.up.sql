CREATE TABLE users
(
    id            serial       not null unique,
    login         varchar(255) not null unique,
    email         varchar(320) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE notes
(
    id           serial                                      not null unique,
    title        varchar(255)                                not null,
    body         varchar(255)                                not null,
    date_created timestamp                                   not null,
    user_id      int references users (id) on delete cascade not null
);

