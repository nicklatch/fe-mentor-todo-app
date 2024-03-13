CREATE TYPE status AS ENUM (
    'Active',
    'Completed'
);

CREATE TABLE todos (
    id serial primary key,
    task varchar(100) not null ,
    status status not null ,
    created_at timestamp default now(),
    updated_at timestamp default null
);