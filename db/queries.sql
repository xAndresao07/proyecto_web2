-- ===================== SOLICITANTES =====================

-- name: ListarSolicitantes :many
SELECT id, nombre, facultad, semestre, nivel_urgencia FROM solicitantes;

-- name: BuscarSolicitantePorID :one
SELECT id, nombre, facultad, semestre, nivel_urgencia FROM solicitantes
WHERE id = ?;

-- name: CrearSolicitante :one
INSERT INTO solicitantes (id, nombre, facultad, semestre, nivel_urgencia)
VALUES (?, ?, ?, ?, ?)
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
INSERT INTO dispositivos (id, solicitante_id, marca, modelo, tipo_almacenamiento, ram_gb, sistema_operativo)
VALUES (?, ?, ?, ?, ?, ?, ?)
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
INSERT INTO ticket_ayudas (id, solicitante_id, dispositivo_id, descripcion_falla, software_requerido, estado_ticket)
VALUES (?, ?, ?, ?, ?, ?)
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
INSERT INTO usuarios (id, email, password, rol)
VALUES (?, ?, ?, ?)
RETURNING id, email, password, rol;

-- name: BuscarUsuarioPorEmail :one
SELECT id, email, password, rol FROM usuarios
WHERE email = ?;-- name: ListarTecnicos :many
SELECT id, nombre, reputacion FROM tecnicos;

-- name: BuscarTecnicoPorID :one
SELECT id, nombre, reputacion FROM tecnicos WHERE id = ?;

-- name: CrearTecnico :one
INSERT INTO tecnicos (nombre, reputacion) VALUES (?, ?) RETURNING id, nombre, reputacion;

-- name: ActualizarTecnico :one
UPDATE tecnicos SET nombre = ?, reputacion = ? WHERE id = ? RETURNING id, nombre, reputacion;

-- name: BorrarTecnico :execrows
DELETE FROM tecnicos WHERE id = ?;

-- name: ListarServiciosPorTecnico :many
SELECT id, tecnico_id, nombre_servicio, nivel_experiencia, tiempo_estimado FROM servicio_ofrecidos WHERE tecnico_id = ?;

-- name: CrearServicio :one
INSERT INTO servicio_ofrecidos (tecnico_id, nombre_servicio, nivel_experiencia, tiempo_estimado) VALUES (?, ?, ?, ?) RETURNING id, tecnico_id, nombre_servicio, nivel_experiencia, tiempo_estimado;

-- name: BorrarServiciosPorTecnico :exec
DELETE FROM servicio_ofrecidos WHERE tecnico_id = ?;

-- name: ListarHorariosPorTecnico :many
SELECT id, tecnico_id, dia_semana, hora_inicio, hora_fin, estado_disponibilidad FROM horario_tecnicos WHERE tecnico_id = ?;

-- name: CrearHorario :one
INSERT INTO horario_tecnicos (tecnico_id, dia_semana, hora_inicio, hora_fin, estado_disponibilidad) VALUES (?, ?, ?, ?, ?) RETURNING id, tecnico_id, dia_semana, hora_inicio, hora_fin, estado_disponibilidad;

-- name: BorrarHorariosPorTecnico :exec
DELETE FROM horario_tecnicos WHERE tecnico_id = ?;
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
