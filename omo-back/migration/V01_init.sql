create or replace function set_updated()
    returns trigger as
$$
begin
    new.updated = now();
    return new;
end;
$$ language plpgsql;

create table if not exists users
(
    id            uuid         not null primary key,
    created       timestamp    not null default now(),
    updated       timestamp    not null default now(),
    username      varchar(64)  not null unique,
    password_hash varchar(512) not null,
    email         varchar(128) not null unique,
    avatar        varchar(64),
    last_login    timestamp    not null default now()
);

create index if not exists idx_username on users (username);