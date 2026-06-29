
CREATE TABLE cita (
    id              INTEGER PRIMARY KEY,
    solicitante_id  TEXT    NOT NULL,
    tecnico_id      TEXT    NOT NULL,
    estado          TEXT    NOT NULL,
    hora_acordada   TEXT    NOT NULL,
    punto_encuentro TEXT    NOT NULL
);

CREATE TABLE punto_encuentros (
    id                      INTEGER PRIMARY KEY,
    nombre_lugar            TEXT    NOT NULL,
    facultad_perteneciente  TEXT    NOT NULL,
    disponible_para_soporte BOOLEAN NOT NULL
);

CREATE TABLE soportes (
    id               INTEGER PRIMARY KEY ,
    cita_id          INTEGER NOT NULL,
    dispositivo_id   INTEGER NOT NULL,
    solucion         TEXT    NOT NULL,
    piezas_cambiadas TEXT    NOT NULL
);