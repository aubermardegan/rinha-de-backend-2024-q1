CREATE UNLOGGED TABLE cliente (
    id INT PRIMARY KEY,
    limite INT,
    saldo INT
);

CREATE UNLOGGED TABLE transacao (
    id SERIAL PRIMARY KEY,
    clienteId INT REFERENCES cliente(id),
    valor INT,
    tipo CHAR,
    descricao VARCHAR(200),
    realizadaEm timestamp 
);

CREATE INDEX idx_cliente_id_includes ON cliente (id) include (limite, saldo);
CREATE INDEX idx_transacao_cliente_id_id_realizada_em ON transacao (clienteId, realizadaEm DESC);

