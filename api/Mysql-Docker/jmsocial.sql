
USE jmsocial

create table if not exists users(
    id int auto_increment primary key, 
    name varchar(50)  not null, 
    nick varchar(20) not null unique,
    email varchar(50) not null unique,
    passwd varchar(120) not null,
    createdin timestamp default current_timestamp()
    ) ENGINE=INNODB;


create table if not exists followers(
    userid int not null,
    FOREIGN KEY(userid)
    REFERENCES users(id)
    ON DELETE CASCADE,

    followerid int not null,
    FOREIGN KEY(followerid)
    REFERENCES users(id)
    ON DELETE CASCADE,
    primary key(userid, followerid)

)ENGINE=INNODB;

create table if not exists publications(
    id int auto_increment primary key, 
    title varchar(100) not null, 
    content varchar(1000) not null,
    creatorid int  not null, 
    FOREIGN KEY (creatorid)
    REFERENCES users(id)
    ON DELETE CASCADE,
    likes int default 0,
    createdin timestamp default current_timestamp
  
)ENGINE=INNODB;


insert into users (name, nick, email, passwd)
values
("user01", "user01", "user01@mail.com", "$2a$10$aIX9XSdknNb56dTL9G9Wz.FIw0HFf/ZqiShyoG4a5L/ORD2p5SEoq"),
("user02", "user02", "user02@mail.com", "$2a$10$aIX9XSdknNb56dTL9G9Wz.FIw0HFf/ZqiShyoG4a5L/ORD2p5SEoq"),
("user03", "user03", "user03@mail.com", "$2a$10$aIX9XSdknNb56dTL9G9Wz.FIw0HFf/ZqiShyoG4a5L/ORD2p5SEoq");
insert into followers(userid, followerid)
values
(1,2),
(3,1),
(1,3);

