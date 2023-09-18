drop table if exists todos;
drop table if exists owners;

create table owners (
    id varchar(50) primary key,
    name varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null default now()
);
create table todos (
    id varchar(50) primary key,
    title varchar(255) not null,
    completed boolean default false,
    created_at timestamp not null,
    updated_at timestamp not null default now(),
    owner_id varchar(50) not null references owners(id) on delete cascade
);

INSERT INTO owners (id, name, created_at) VALUES ('1', 'Alice', now());
INSERT INTO owners (id, name, created_at) VALUES ('2', 'Bob', now());
INSERT INTO owners (id, name, created_at) VALUES ('3', 'Cathy', now());
INSERT INTO owners (id, name, created_at) VALUES ('4', 'Danny', now());
INSERT INTO owners (id, name, created_at) VALUES ('5', 'Eli', now());

INSERT INTO todos (id, title, created_at, owner_id) VALUES ('1', 'Buy milk', now(), '1');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('2', 'Buy eggs', now(), '1');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('3', 'Buy bread', now(), '2');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('4', 'Buy butter', now(), '2');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('5', 'Finish homework', now(), '3');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('6', 'Do chores', now(), '3');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('7', 'Read Documentation', now(), '4');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('8', 'Change oil', now(), '5');