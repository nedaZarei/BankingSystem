CREATE TABLE account (
    id SERIAL,
    username VARCHAR(20) UNIQUE,
    accountNumber VARCHAR(16) UNIQUE,
    password VARCHAR(30) NOT NULL,
    firstName VARCHAR(20) NOT NULL,
    lastName VARCHAR(20) NOT NULL,
    dateOfBirth DATE NOT NULL,
    nationalID VARCHAR(10) NOT NULL,
    email VARCHAR(50) NOT NULL,
    phoneNumber VARCHAR(15) NOT NULL,
    accountType VARCHAR(10) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY(id) 
); 