create type genders as enum
    (
        'male',
        'female'
        );

create table users
(
    user_id  bigint primary key generated always as identity,
    login    text not null unique,
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
    bread_unit         bigint  not null,
    height             bigint  not null
);

create table products
(
    product_id   bigint primary key generated always as identity,
    product_name text   not null unique,
    bread_units        float not null
);

create table note_header
(
    note_id     bigint primary key generated always as identity,
    create_date date       not null,
    sugar_level real
);

create table note_user
(
    note_id bigint unique not null references note_header (note_id),
    user_id bigint not null references users (user_id)
);

create table note_text (
    note_id bigint not null references note_user (note_id),
    text varchar
);

INSERT INTO products (product_name, bread_units) VALUES
                                                    ('Белый хлеб', 5),
                                                    ('Черный хлеб', 4),
                                                    ('Вермишель', 2),
                                                    ('Гречневая крупа', 2),
                                                    ('Кукуруза', 1),
                                                    ('Кукуруза консервированная', 1.67),
                                                    ('Кукурузные хлопья', 6.67),
                                                    ('Попкорн', 6.67),
                                                    ('Картофельное пюре', 1.33),
                                                    ('Овсянка', 2),
                                                    ('Перловка', 2),
                                                    ('Пшенная каша', 2),
                                                    ('Рис', 2),
                                                    ('Жареный картофель', 2.5),
                                                    ('Чипсы', 4),
                                                    ('Абрикос', 0.91),
                                                    ('Айва', 0.71),
                                                    ('Ананас', 0.71),
                                                    ('Арбуз', 0.37),
                                                    ('Апельсин', 0.67),
                                                    ('Банан', 1.43),
                                                    ('Брусника', 0.71),
                                                    ('Виноград', 1.43),
                                                    ('Вишня', 1.11),
                                                    ('Гранат', 0.59),
                                                    ('Грейпфрут', 0.59),
                                                    ('Груша', 1.11),
                                                    ('Дыня', 1),
                                                    ('Ежевика', 0.71),
                                                    ('Инжир', 1.25),
                                                    ('Киви', 0.91),
                                                    ('Клубника', 0.63),
                                                    ('Крыжовник', 0.83),
                                                    ('Малина', 0.63),
                                                    ('Манго', 0.91),
                                                    ('Мандарин', 0.67),
                                                    ('Персик', 0.83),
                                                    ('Слива', 1.11),
                                                    ('Смородина', 0.83),
                                                    ('Финик', 6.67),
                                                    ('Хурма', 1.43),
                                                    ('Черешня', 1.11),
                                                    ('Черника', 1.11),
                                                    ('Яблоко', 1.11),
                                                    ('Сухофрукты', 5),
                                                    ('Морковь', 0.5),
                                                    ('Свекла', 0.67),
                                                    ('Арахис', 1),
                                                    ('Бобы', 5),
                                                    ('Горошек зеленый', 1),
                                                    ('Фасоль', 2),
                                                    ('Сахар', 10),
                                                    ('Шоколад', 5),
                                                    ('Мёд', 8.33);
