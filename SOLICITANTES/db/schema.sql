CREATE TABLE solicitantes (
    id             INTEGER PRIMARY KEY,
    usuario_id     INTEGER NOT NULL,
    matricula      TEXT    NOT NULL,
    nombre         TEXT    NOT NULL,
    facultad       TEXT    NOT NULL,
    semestre       INTEGER NOT NULL,
    nivel_urgencia TEXT    NOT NULL
);

CREATE TABLE dispositivos (
    id                  INTEGER PRIMARY KEY,
    solicitante_id      INTEGER NOT NULL,
    marca               TEXT    NOT NULL,
    modelo              TEXT    NOT NULL,
    tipo_almacenamiento TEXT    NOT NULL,
    ram_gb              INTEGER NOT NULL,
    sistema_operativo   TEXT    NOT NULL
);

CREATE TABLE ticket_ayudas (
    id                 INTEGER PRIMARY KEY,
    solicitante_id     INTEGER NOT NULL,
    dispositivo_id     INTEGER NOT NULL,
    descripcion_falla  TEXT    NOT NULL,
    software_requerido TEXT    NOT NULL,
    estado_ticket      TEXT    NOT NULL
);
