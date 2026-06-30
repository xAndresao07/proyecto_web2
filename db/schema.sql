CREATE TABLE tecnicos (
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