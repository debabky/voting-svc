-- +migrate Up

create table votings(
    id          uuid                     primary key default gen_random_uuid(),
    name        text                        not null,
    description text                        not null,
    created_at  timestamp without time zone not null default now()
);

create table voting_options(
    name         text    not null,
    voting_id    uuid    not null,
    description  text,
    primary key (name, voting_id)
);

create table votes(
    voting_id     uuid not null,
    voting_option text not null,
    nullifier     text not null,
    primary key (voting_id, voting_option, nullifier)
);

-- +migrate Down

drop table votings;
drop table voting_options;
drop table votes;