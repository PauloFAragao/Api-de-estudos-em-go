Create DataBase If Not Exists devbook;

Use devbook;

Drop Table If Exists usuarios;

Create Table usuarios(
    id Int Auto_increment Primary Key,
    nome Varchar(50) Not Null,
    nick Varchar(50) Not Null Unique,
    email Varchar(50) Not Null Unique,
    senha Varchar(20) Not Null Unique,
    criadoEm Timestamp Default current_timestamp()
) ENGINE=INNODB;

