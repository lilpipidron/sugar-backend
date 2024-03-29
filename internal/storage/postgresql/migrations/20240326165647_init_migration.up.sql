create type note_types as enum
    (
        'regular'
            'sugar'
        );

create type genders as enum
    (
        'male'
            'female'
        );

create table users
(
    user_id  bigint primary key generated always as identity,
    login    text not null,
    password text not null
);

create table user_info
(
    user_id            bigint  not null references users (user_id),
    name               text    not null,
    birthday           date    not null,
    gender             genders not null,
    weight             bigint  not null,
    carbohydrate_ratio real    not null,
    bread_unit         bigint  not null
);

create table products
(
    product_id   bigint primary key generated always as identity,
    product_name text   not null,
    carbs        bigint not null
);


create table note_user
(
    note_id bigint not null references note_header (note_id),
    user_id bigint not null references users (user_id)
);

create table note_header
(
    note_id     bigint primary key generated always as identity,
    note_type   note_types not null,
    create_date date       not null,
    sugar_level real
);

create table note_detail
(
    note_id        bigint not null references note_header (note_id),
    product_id     bigint not null references product (product_id),
    product_amount bigint not null
);
