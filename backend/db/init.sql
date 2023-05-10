create table if not exists contract (
    name    text primary key,
    address text
);

create table if not exists proof (
    address             text primary key,
    last_commited_date  text,
    last_commited_value bigint,
    date                text not null,
    v                   smallint not null,
    r                   text not null,
    s                   text not null,
    hash                text not null,
    value               bigint not null

);

create table if not exists content (
    id                   serial primary key,
    author               text not null,
    header               text not null,
    content              text not null
);

create table if not exists click (
    content_id integer not null references content(id),
    date       text not null,
    address    text not null,
    commited   boolean not null default false,
    primary key(content_id, date, address)
)
