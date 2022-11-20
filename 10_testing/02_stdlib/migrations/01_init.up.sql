create table posts(
    id serial primary key,
    title varchar(256) not null,
    text text not null,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
);
