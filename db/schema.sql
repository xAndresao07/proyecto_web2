CREATE TABLE solicitantes (
    id             INTEGER PRIMARY KEY,
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

CREATE TABLE usuarios (
    id       INTEGER PRIMARY KEY,
    email    TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    rol      TEXT NOT NULL
);CREATE TABLE tecnicos (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre     TEXT NOT NULL,
    reputacion REAL NOT NULL
);

CREATE TABLE servicio_ofrecidos (
    id                INTEGER PRIMARY KEY AUTOINCREMENT,
    tecnico_id        INTEGER NOT NULL,
    nombre_servicio   TEXT NOT NULL,
    nivel_experiencia TEXT NOT NULL,
    tiempo_estimado   TEXT NOT NULL,
    FOREIGN KEY(tecnico_id) REFERENCES tecnicos(id) ON DELETE CASCADE
);

CREATE TABLE horario_tecnicos (
    id                    INTEGER PRIMARY KEY AUTOINCREMENT,
    tecnico_id            INTEGER NOT NULL,
    dia_semana            TEXT NOT NULL,
    hora_inicio           TEXT NOT NULL,
    hora_fin              TEXT NOT NULL,
    estado_disponibilidad TEXT NOT NULL,
    FOREIGN KEY(tecnico_id) REFERENCES tecnicos(id) ON DELETE CASCADE
);

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
