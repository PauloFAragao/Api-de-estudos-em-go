Create DataBase If Not Exists devbook;

Use devbook;

Drop Table If Exists publicacoes;
Drop Table If Exists seguidores;
Drop Table If Exists usuarios;


Create Table usuarios(
    id Int Auto_increment Primary Key,
    nome Varchar(50) Not Null,
    nick Varchar(50) Not Null Unique,
    email Varchar(50) Not Null Unique,
    senha Varchar(100) Not Null,
    criadoEm Timestamp Default current_timestamp()
) ENGINE=INNODB;

Create Table seguidores(
    usuario_id Int Not Null, 
    Foreign Key (usuario_id)
    References usuarios(id)
    On Delete Cascade,

    seguidor_id Int Not Null,
    Foreign Key (seguidor_id)
    References usuarios(id)
    On Delete Cascade,

    Primary key (usuario_id, seguidor_id)
) ENGINE=INNODB;

Create Table publicacoes(
    id Int Auto_increment Primary Key,
    titulo Varchar(50) Not Null,
    conteudo Varchar(300) Not Null,

    autor_id Int Not Null,
    Foreign Key (autor_id)
    References usuarios(id)
    On Delete Cascade,

    curtidas Int Default 0,
    criadaEm Timestamp Default current_timestamp
) ENGINE=INNODB;