create table users (
    macadd varchar unique,
    username varchar unique,
    nickname varchar
);

create table messagerelation (
    mID integer unique,
    username varchar
);

create table messages (
    mID integer unique,
    mText varchar
);

insert into users values ('0', 'dbCreator', 'Hal');
insert into messagerelation values (0, 'dbCreator');
insert into messages values (0, 'db Created');