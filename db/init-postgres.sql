-- Schema adaptado para PostgreSQL (quitamos AUTOINCREMENT que es de SQLite)

CREATE TABLE solicitantes (
    id             SERIAL PRIMARY KEY,
    nombre         TEXT    NOT NULL,
    facultad       TEXT    NOT NULL,
    semestre       INTEGER NOT NULL,
    nivel_urgencia TEXT    NOT NULL
);

CREATE TABLE dispositivos (
    id                  SERIAL PRIMARY KEY,
    solicitante_id      INTEGER NOT NULL,
    marca               TEXT    NOT NULL,
    modelo              TEXT    NOT NULL,
    tipo_almacenamiento TEXT    NOT NULL,
    ram_gb              INTEGER NOT NULL,
    sistema_operativo   TEXT    NOT NULL
);

CREATE TABLE ticket_ayudas (
    id                 SERIAL PRIMARY KEY,
    solicitante_id     INTEGER NOT NULL,
    dispositivo_id     INTEGER NOT NULL,
    descripcion_falla  TEXT    NOT NULL,
    software_requerido TEXT    NOT NULL,
    estado_ticket      TEXT    NOT NULL
);

CREATE TABLE usuarios (
    id       SERIAL PRIMARY KEY,
    email    TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    rol      TEXT NOT NULL
);

CREATE TABLE tecnicos (
    id         SERIAL PRIMARY KEY,
    nombre     TEXT NOT NULL,
    reputacion REAL NOT NULL
);

CREATE TABLE servicio_ofrecidos (
    id                SERIAL PRIMARY KEY,
    tecnico_id        INTEGER NOT NULL,
    nombre_servicio   TEXT NOT NULL,
    nivel_experiencia TEXT NOT NULL,
    tiempo_estimado   TEXT NOT NULL,
    FOREIGN KEY(tecnico_id) REFERENCES tecnicos(id) ON DELETE CASCADE
);

CREATE TABLE horario_tecnicos (
    id                    SERIAL PRIMARY KEY,
    tecnico_id            INTEGER NOT NULL,
    dia_semana            TEXT NOT NULL,
    hora_inicio           TEXT NOT NULL,
    hora_fin              TEXT NOT NULL,
    estado_disponibilidad TEXT NOT NULL,
    FOREIGN KEY(tecnico_id) REFERENCES tecnicos(id) ON DELETE CASCADE
);

CREATE TABLE cita (
    id              SERIAL PRIMARY KEY,
    solicitante_id  TEXT    NOT NULL,
    tecnico_id      TEXT    NOT NULL,
    estado          TEXT    NOT NULL,
    hora_acordada   TEXT    NOT NULL,
    punto_encuentro TEXT    NOT NULL
);

CREATE TABLE punto_encuentros (
    id                      SERIAL PRIMARY KEY,
    nombre_lugar            TEXT    NOT NULL,
    facultad_perteneciente  TEXT    NOT NULL,
    disponible_para_soporte BOOLEAN NOT NULL
);

CREATE TABLE soportes (
    id               SERIAL PRIMARY KEY,
    cita_id          INTEGER NOT NULL,
    dispositivo_id   INTEGER NOT NULL,
    solucion         TEXT    NOT NULL,
    piezas_cambiadas TEXT    NOT NULL
);

-- SEEDERS REQUERIDOS POR LA RUBRICA
INSERT INTO usuarios (email, password, rol) VALUES ('admin@uleam.edu.ec', '123456', 'admin');
INSERT INTO solicitantes (nombre, facultad, semestre, nivel_urgencia) VALUES ('Juan Perez', 'FACCI', 5, 'alta');
INSERT INTO tecnicos (nombre, reputacion) VALUES ('Carlos Técnico', 4.8);
INSERT INTO punto_encuentros (nombre_lugar, facultad_perteneciente, disponible_para_soporte) VALUES ('Laboratorio 1', 'FACCI', true);
