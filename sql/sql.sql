CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=InnoDB;


CREATE TABLE seguidores (
    usuario_id int not null,
    foreign key (usuario_id) references usuarios(id) on delete cascade,
    seguidor_id int not null,
    foreign key (seguidor_id) references usuarios(id) on delete cascade,

    primary key (usuario_id, seguidor_id)
) ENGINE=InnoDB;