insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$.B1T4JjMTdicbMO1bxAeO.PQ4ZpQmnGhioEjGJc1E660xRGMJ6FFm"), -- usuario 1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$.B1T4JjMTdicbMO1bxAeO.PQ4ZpQmnGhioEjGJc1E660xRGMJ6FFm"), -- usuário 2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$.B1T4JjMTdicbMO1bxAeO.PQ4ZpQmnGhioEjGJc1E660xRGMJ6FFm"); -- usuário 3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);