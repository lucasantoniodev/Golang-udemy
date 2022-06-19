CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id) -- Declarando que é uma chave extrangeira
    REFERENCES usuarios(id) -- Referenciando a tabela e a coluna que vão ser a chave primária
    ON DELETE CASCADE, -- Deletando todas as linhas caso o ID seja deletado na TABELA "usuarios"
    
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id) -- A chave primária da tabela vai ser uma junção de duas colunas "1-3, 1-2..."
) ENGINE=INNODB;

