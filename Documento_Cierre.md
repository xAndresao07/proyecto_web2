# Documento de Cierre — EmparejaTech

**Proyecto:** API de Optimización Técnica para Equipos Personales en Universidad (ULEAM)
**Equipo:** Mario Soriano (Intervenciones) · Anthony Solórzano (Técnicos) · Jandry Cedeño (Solicitantes)
**Componente:** C2 — Proyecto Semestral · Aplicaciones Web II (TDI-601) · Hito 3

## Qué construimos

Una API REST en Go que resuelve un problema real del campus: cruzar los horarios de
estudiantes con equipos defectuosos (Solicitantes) con los de compañeros que saben repararlos
(Técnicos), generando una Intervención (cita) con punto de encuentro y cierre auditado
(Soporte). El proyecto se dividió en tres módulos de dominio con arquitectura en capas idéntica
(handler → service → repository), autenticación JWT compartida, persistencia con GORM y
despliegue con Docker Compose (API + PostgreSQL + seeders).

## Qué aprendimos

- A mantener **tres módulos independientes con la misma convención de capas** para que la
  integración final no exigiera reescribir lógica de negocio, solo cablear rutas y compartir
  el módulo de `Usuario`/Auth.
- El valor de las **interfaces de repository** (`TecnicoRepository`, etc.): permitieron cambiar
  de backend en memoria → GORM → sqlc sin tocar el service, y facilitaron los mocks con
  `testify` para los tests unitarios.
- A usar el **patrón de opciones funcionales** (`Options`) para configurar `AuthService` y el
  servidor HTTP (timeouts, puerto) en vez de valores globales hardcodeados — mover el secreto
  JWT y el puerto a `config.Cargar()` fue un cambio pequeño con impacto grande en testabilidad.
- Coordinar **Git en equipo real**: feature branches por módulo, un `go.mod` con el mismo
  nombre de paquete (`proyecto`) en los tres módulos para que la fusión no rompiera imports, y
  Pull Requests una vez protegida la rama `main`.
- Diferencias prácticas entre **SQLite en desarrollo y PostgreSQL en Docker** (DSN, healthcheck
  del contenedor `db`, `depends_on: condition: service_healthy`).

## Qué haríamos distinto

- Unificar desde la **semana 1** el nombre de las rutas de auth (`/auth/registrar` vs
  `/auth/register` vs `/registro` terminaron distintas entre módulos) para evitar el trabajo de
  normalización al integrar.
- Definir el **contrato de roles** (`admin`, `solicitante`, `tecnico`) antes de escribir los
  handlers, para que el middleware de autorización por rol no se agregara al final.
- Escribir la colección de Postman **en paralelo** al desarrollo de cada endpoint, no al cierre
  del hito — habría detectado antes las inconsistencias de nombres de campos entre módulos.
- Configurar el pipeline de CI/CD (build → vet → test) desde un hito anterior, en vez de
  dejarlo para la semana 14, para llegar al Hito 3 con el gate G1 ya validado con margen.

## Próximos pasos del producto

1. Terminar de fusionar los tres `main.go` en un único router `/api/v1` (hoy `proyecto/`
   contiene la base de Intervenciones; falta montar las rutas de Técnicos y Solicitantes).
2. Middleware de **autorización por rol** (ej. solo `tecnico` puede actualizar su propio
   horario; solo `admin` puede borrar un `PuntoEncuentro`).
3. Motor de matching real en `intervenciones`: cruzar `HorarioTecnico` disponible contra la
   urgencia y el horario libre del `Solicitante` para sugerir citas automáticamente.
4. Notificaciones (correo o push) cuando una `Cita` cambia de estado.
5. Frontend mínimo que consuma la API para demo pública, más allá del `frontend/index.html`
   de prueba.
