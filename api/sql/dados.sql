Insert Into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$SOS/yU0TWFn3k.oYatGoxOvYuWQadQKsydn.JvY8N6Bv1n1hoVzci"), #senha: 123456
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$SOS/yU0TWFn3k.oYatGoxOvYuWQadQKsydn.JvY8N6Bv1n1hoVzci"), #senha: 123456
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$SOS/yU0TWFn3k.oYatGoxOvYuWQadQKsydn.JvY8N6Bv1n1hoVzci"); #senha: 123456

Insert Into seguidores (usuario_id, seguidor_id)
values
(1, 2), #usuário 2 segue o usuário 1
(3, 1), #usuário 1 segue o usuário 3
(1, 3), #usuário 3 segue o usuário 1
