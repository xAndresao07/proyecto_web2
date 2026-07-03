-- ===================== CITAS =====================

-- name: ListarCitas :many
SELECT id, solicitante_id, tecnico_id, estado, hora_acordada, punto_encuentro FROM cita;

-- name: BuscarCitaPorID :one
SELECT id, solicitante_id, tecnico_id, estado, hora_acordada, punto_encuentro FROM cita
WHERE id = ?;

-- name: CrearCita :one
INSERT INTO cita (solicitante_id, tecnico_id, estado, hora_acordada, punto_encuentro)
VALUES (?, ?, ?, ?, ?)
RETURNING id, solicitante_id, tecnico_id, estado, hora_acordada, punto_encuentro;

-- name: ActualizarCita :one
UPDATE cita
SET solicitante_id = ?, tecnico_id = ?, estado = ?, hora_acordada = ?, punto_encuentro = ?
WHERE id = ?
RETURNING id, solicitante_id, tecnico_id, estado, hora_acordada, punto_encuentro;

-- name: BorrarCita :execrows
DELETE FROM cita WHERE id = ?;

-- ===================== PUNTOS DE ENCUENTRO =====================

-- name: ListarPuntosEncuentro :many
SELECT id, nombre_lugar, facultad_perteneciente, disponible_para_soporte FROM punto_encuentros;

-- name: BuscarPuntoEncuentroPorID :one
SELECT id, nombre_lugar, facultad_perteneciente, disponible_para_soporte FROM punto_encuentros
WHERE id = ?;

-- name: CrearPuntoEncuentro :one
INSERT INTO punto_encuentros (nombre_lugar, facultad_perteneciente, disponible_para_soporte)
VALUES (?, ?, ?)
RETURNING id, nombre_lugar, facultad_perteneciente, disponible_para_soporte;

-- name: ActualizarPuntoEncuentro :one
UPDATE punto_encuentros
SET nombre_lugar = ?, facultad_perteneciente = ?, disponible_para_soporte = ?
WHERE id = ?
RETURNING id, nombre_lugar, facultad_perteneciente, disponible_para_soporte;

-- name: BorrarPuntoEncuentro :execrows
DELETE FROM punto_encuentros WHERE id = ?;

-- ===================== SOPORTES =====================

-- name: ListarSoportes :many
SELECT id, cita_id, dispositivo_id, solucion, piezas_cambiadas FROM soportes;

-- name: BuscarSoportePorID :one
SELECT id, cita_id, dispositivo_id, solucion, piezas_cambiadas FROM soportes
WHERE id = ?;

-- name: CrearSoporte :one
INSERT INTO soportes (cita_id, dispositivo_id, solucion, piezas_cambiadas)
VALUES (?, ?, ?, ?)
RETURNING id, cita_id, dispositivo_id, solucion, piezas_cambiadas;

-- name: ActualizarSoporte :one
UPDATE soportes
SET cita_id = ?, dispositivo_id = ?, solucion = ?, piezas_cambiadas = ?
WHERE id = ?
RETURNING id, cita_id, dispositivo_id, solucion, piezas_cambiadas;

-- name: BorrarSoporte :execrows
DELETE FROM soportes WHERE id = ?;
