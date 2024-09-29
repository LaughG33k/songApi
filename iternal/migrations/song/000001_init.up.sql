


create table if not exists songs (
    Id serial primary key,
    Song text not null,
    Groups text not null,
    Song_Text text not null,
    Link text not null,
    Realese_Date varchar(30) not null
);

create index groupIndex on songs using hash (Groups);
create index songIndex on songs using hash (Song);