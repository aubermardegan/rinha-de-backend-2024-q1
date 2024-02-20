CREATE TABLE cliente (
    id SERIAL PRIMARY KEY,
    limite INT,
    saldoInicial INT
);

CREATE TABLE transacao (
    id SERIAL PRIMARY KEY,
    clienteId INT REFERENCES cliente(id),
    valor INT,
    tipo CHAR,
    descricao VARCHAR(200),
    realizadaEm timestamp 
);