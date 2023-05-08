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

)
