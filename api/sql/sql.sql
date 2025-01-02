Create DataBase If Not Exists devbook;

Use devbook;

Drop Table If Exists usuarios;

Create Table usuarios(
    id Int Auto_increment Primary Key,
    nome Varchar(50) Not Null,
    nick Varchar(50) Not Null Unique,
    email Varchar(50) Not Null Unique,
    senha Varchar(100) Not Null,
    criadoEm Timestamp Default current_timestamp()
) ENGINE=INNODB;

