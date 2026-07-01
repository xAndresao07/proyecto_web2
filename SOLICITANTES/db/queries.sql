-- ===================== SOLICITANTES =====================

-- name: ListarSolicitantes :many
SELECT id, usuario_id, matricula, nombre, facultad, semestre, nivel_urgencia FROM solicitantes;

-- name: BuscarSolicitantePorID :one
SELECT id, usuario_id, matricula, nombre, facultad, semestre, nivel_urgencia FROM solicitantes
WHERE id = ?;

-- name: CrearSolicitante :one
INSERT INTO solicitantes (usuario_id, matricula, nombre, facultad, semestre, nivel_urgencia)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, usuario_id, matricula, nombre, facultad, semestre, nivel_urgencia;

-- name: ActualizarSolicitante :one
UPDATE solicitantes
SET matricula = ?, nombre = ?, facultad = ?, semestre = ?, nivel_urgencia = ?
WHERE id = ?
RETURNING id, usuario_id, matricula, nombre, facultad, semestre, nivel_urgencia;

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
SET marca = ?, modelo = ?, tipo_almacenamiento = ?, ram_gb = ?, sistema_operativo = ?
WHERE id = ?
RETURNING id, solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo;

-- name: BorrarDispositivo :execrows
DELETE FROM dispositivos WHERE id = ?;

-- ===================== TICKET AYUDAS =====================

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
SET dispositivo_id = ?, descripcion_falla = ?, software_requerido = ?, estado_ticket = ?
WHERE id = ?
RETURNING id, solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket;

-- name: BorrarTicketAyuda :execrows
DELETE FROM ticket_ayudas WHERE id = ?;
