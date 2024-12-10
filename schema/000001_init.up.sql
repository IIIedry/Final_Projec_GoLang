create table products (
    product_id serial primary key,
    name text,
    description text,
    price numeric(10, 2),
    count integer,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table product_changes (
    product_change_id serial primary key,
    product_id integer not null,
    change text,
    updated_at timestamp with time zone default now()
);
