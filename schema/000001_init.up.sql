CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    is_done     boolean      not null default false,
    created_at  timestamp             default CURRENT_TIMESTAMP,
    updated_at  timestamp             default CURRENT_TIMESTAMP
);

CREATE TABLE todo_items
(
    id         serial       not null unique,
    text       varchar(255) not null,
    is_done    boolean      not null default false,
    created_at timestamp             default CURRENT_TIMESTAMP,
    updated_at timestamp             default CURRENT_TIMESTAMP
);

CREATE TABLE user_todo_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_list_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);