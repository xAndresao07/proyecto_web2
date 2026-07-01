-- name: ListarTecnicos :many
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