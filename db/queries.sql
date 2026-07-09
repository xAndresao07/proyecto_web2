-- ===================== SOLICITANTES =====================

-- name: ListarSolicitantes :many
SELECT id, nombre, facultad, semestre, nivel_urgencia FROM solicitantes;

-- name: BuscarSolicitantePorID :one
SELECT id, nombre, facultad, semestre, nivel_urgencia FROM solicitantes
WHERE id = ?;

-- name: CrearSolicitante :one
INSERT INTO solicitantes (nombre, facultad, semestre, nivel_urgencia)
VALUES (?, ?, ?, ?)
RETURNING id, nombre, facultad, semestre, nivel_urgencia;

-- name: ActualizarSolicitante :one
UPDATE solicitantes
SET nombre = ?, facultad = ?, semestre = ?, nivel_urgencia = ?
WHERE id = ?
RETURNING id, nombre, facultad, semestre, nivel_urgencia;

-- name: BorrarSolicitante :execrows
DELETE FROM solicitantes WHERE id = ?;

-- ===================== DISPOSITIVOS =====================

-- name: ListarDispositivos :many
SELECT id, solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo FROM dispositivos;

-- name: BuscarDispositivoPorID :one
SELECT id, solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo FROM dispositivos
WHERE id = ?;

-- name: CrearDispositivo :one
INSERT INTO dispositivos (solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo;

-- name: ActualizarDispositivo :one
UPDATE dispositivos
SET solicitante_id = ?, marca = ?, modelo = ?, tipo_almacenamiento = ?, ram_gb = ?, sistema_operativo = ?
WHERE id = ?
RETURNING id, solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo;

-- name: BorrarDispositivo :execrows
DELETE FROM dispositivos WHERE id = ?;

-- ===================== TICKET_AYUDAS =====================

-- name: ListarTicketAyudas :many
SELECT id, solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket FROM ticket_ayudas;

-- name: BuscarTicketAyudaPorID :one
SELECT id, solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket FROM ticket_ayudas
WHERE id = ?;

-- name: CrearTicketAyuda :one
INSERT INTO ticket_ayudas (solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket)
VALUES (?, ?, ?, ?, ?)
RETURNING id, solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket;

-- name: ActualizarTicketAyuda :one
UPDATE ticket_ayudas
SET solicitante_id = ?, dispositivo_id = ?, descripcion_falla = ?, software_requerido = ?, estado_ticket = ?
WHERE id = ?
RETURNING id, solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket;

-- name: BorrarTicketAyuda :execrows
DELETE FROM ticket_ayudas WHERE id = ?;

-- ===================== USUARIOS =====================

-- name: CrearUsuario :one
INSERT INTO usuarios (email, password, rol)
VALUES (?, ?, ?)
RETURNING id, email, password, rol;

-- name: BuscarUsuarioPorEmail :one
SELECT id, email, password, rol FROM usuarios
WHERE email = ?;