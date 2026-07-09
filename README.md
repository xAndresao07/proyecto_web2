# EmparejaTech — API de Optimización Técnica para Equipos Personales (ULEAM)

API REST que resuelve un problema cotidiano del campus: estudiantes con equipos lentos o
defectuosos (**Solicitantes**) no logran coincidir en horario con compañeros que sí saben
resolver el problema (**Técnicos**). El sistema cruza horarios, ubicación y tipo de habilidad
requerida para generar una **Intervención** (cita técnica) y darle seguimiento hasta su cierre.

> Proyecto Semestral — Aplicaciones Web II (TDI-601) · Período 2026-1 · ULEAM
> Docente: Ing. John Cevallos Macías, Mg.

## Tabla de contenido
- [Equipo y módulos](#equipo-y-módulos)
- [Stack tecnológico](#stack-tecnológico)
- [Arquitectura](#arquitectura)
- [Cómo correrlo](#cómo-correrlo)
- [Variables de entorno](#variables-de-entorno)
- [Autenticación](#autenticación)
- [Endpoints por módulo](#endpoints-por-módulo)
- [Colección de Postman](#colección-de-postman)
- [Testing](#testing)
- [Estructura de carpetas](#estructura-de-carpetas)

## Equipo y módulos

El proyecto está organizado en **micro-módulos de dominio**: cada integrante es dueño de su
capa completa (handler → service → repository) y la responsabilidad grupal cubre lo transversal.

| Integrante | Módulo | Responsabilidad |
|---|---|---|
| **Anthony Solórzano** | `tecnicos` | Perfiles técnicos, catálogo de habilidades/servicios ofrecidos, horarios de disponibilidad y reputación. |
| **Jandry Cedeño** | `solicitantes` | Perfiles de estudiantes, dispositivos con fallas y tickets de ayuda. |
| **Mario Soriano** | `intervenciones` | Citas (matching solicitante–técnico), puntos de encuentro y soporte/cierre del servicio. |
| **Grupal** | transversal | Autenticación JWT, middlewares, Docker y CI/CD. |

## Stack tecnológico

- **Lenguaje:** Go 1.22+
- **Router:** [Chi](https://github.com/go-chi/chi)
- **ORM / persistencia:** GORM (backend alterno vía `sqlc`), SQLite en desarrollo / PostgreSQL en Docker
- **Auth:** JWT (`golang-jwt`) + `bcrypt` para hash de contraseñas
- **Testing:** `testify` (mocks y asserts) + Postman para pruebas de integración manual
- **Contenedores:** Docker multi-stage + Docker Compose

## Arquitectura

Arquitectura en capas, replicada de forma idéntica en los tres módulos, con inyección de
dependencias desde `main.go`:

```
Request HTTP
   │
   ▼
Handler        → recibe/valida payload, traduce a llamadas de servicio, arma la respuesta HTTP
   │
   ▼
Service        → reglas de negocio (ej. matching de horarios, validación de estado de ticket)
   │
   ▼
Repository     → acceso a datos vía GORM (interfaces, así el service no conoce SQL/GORM directo)
   │
   ▼
Base de datos  → SQLite (local) / PostgreSQL (Docker)
```

Los tres módulos comparten el mismo patrón y la misma tabla `Usuario` (autenticación), y se
integran bajo un único router `/api/v1` con middleware de auth compartido. Ver el diagrama
completo en `diagrama_arquitectura.svg`.

## Cómo correrlo

```bash
git clone https://github.com/xAndresao07/proyecto_web2.git
cd proyecto_web2
docker-compose up --build
```

Esto levanta:
- **api** — el backend Go en `http://localhost:8080`
- **db** — PostgreSQL 16 con *seeders* iniciales

La API queda disponible en `http://localhost:8080/api/v1`. No se requieren pasos manuales
adicionales: las migraciones (`AutoMigrate` de GORM) corren al iniciar el contenedor.

### Correr un módulo de forma individual (desarrollo)

Cada módulo puede levantarse por separado con SQLite local mientras se desarrolla:

```bash
cd TECNICOS   # o SOLICITANTES / CITAS
go run ./cmd/api
```

## Variables de entorno

| Variable | Descripción | Default |
|---|---|---|
| `PUERTO` | Puerto HTTP del servidor | `:8080` |
| `DB_DRIVER` | Motor de base de datos: `sqlite` o `postgres` | `sqlite` |
| `DB_DSN` | Cadena de conexión (solo si `DB_DRIVER=postgres`) | — |
| `RUTA_DB` | Archivo SQLite local | `proyecto.db` |
| `STORAGE` | Backend de acceso a datos: `gorm` o `sqlc` | `gorm` |
| `JWT_SECRETO` | Clave para firmar/verificar los tokens | *(cambiar en producción)* |
| `JWT_DURACION` | Vigencia del token | `24h` |

## Autenticación

El acceso es vía **JWT Bearer Token**. Los endpoints de auth son públicos; todo lo demás exige
`Authorization: Bearer <token>`.

```
POST /api/v1/auth/registrar   { "email", "password", "rol" }   → 201
POST /api/v1/auth/login       { "email", "password" }          → 200 { "token": "..." }
```

Roles soportados: `admin`, `solicitante`, `tecnico`. El rol viaja en el registro del usuario y
condiciona qué puede crear/consultar cada cuenta.

## Endpoints por módulo

### 🔧 Técnicos — *responsable: Anthony Solórzano*

| Método | Endpoint | Descripción |
|---|---|---|
| GET | `/api/v1/tecnicos` | Lista todos los técnicos |
| POST | `/api/v1/tecnicos` | Crea un técnico (con servicios y horarios) |
| GET | `/api/v1/tecnicos/{id}` | Obtiene un técnico por ID |
| PUT | `/api/v1/tecnicos/{id}` | Actualiza un técnico |
| DELETE | `/api/v1/tecnicos/{id}` | Elimina un técnico |

**Entidades:** `Tecnico` (1:N `ServicioOfrecido`, 1:N `HorarioTecnico`).

### 🎓 Solicitantes — *responsable: Jandry Cedeño*

| Método | Endpoint | Descripción |
|---|---|---|
| GET/POST | `/api/v1/solicitantes` | Lista / crea solicitantes |
| GET/PUT/DELETE | `/api/v1/solicitantes/{id}` | Detalle / actualiza / elimina |
| GET/POST | `/api/v1/dispositivos` | Lista / registra dispositivos con falla |
| GET/PUT/DELETE | `/api/v1/dispositivos/{id}` | Detalle / actualiza / elimina |
| GET/POST | `/api/v1/tickets` | Lista / crea tickets de ayuda |
| GET/PUT/DELETE | `/api/v1/tickets/{id}` | Detalle / actualiza / elimina |

**Entidades:** `Solicitante` (1:1 `Usuario`) → 1:N `Dispositivo` → 1:N `TicketAyuda`.

### 🤝 Intervenciones — *responsable: Mario Soriano*

| Método | Endpoint | Descripción |
|---|---|---|
| GET/POST | `/api/v1/citas` | Lista / crea citas (matching solicitante–técnico) |
| GET/PUT/DELETE | `/api/v1/citas/{id}` | Detalle / actualiza estado / elimina |
| GET/POST | `/api/v1/puntos-encuentro` | Lista / crea puntos de encuentro en el campus |
| GET/PUT/DELETE | `/api/v1/puntos-encuentro/{id}` | Detalle / actualiza / elimina |
| GET/POST | `/api/v1/soportes` | Lista / crea el cierre/auditoría de una cita |
| GET/PUT/DELETE | `/api/v1/soportes/{id}` | Detalle / actualiza / elimina |

**Entidades:** `Cita` (referencia a `Solicitante` y `Tecnico` por ID) → `Soporte` (cierra la
`Cita` y referencia el `Dispositivo` intervenido). `PuntoEncuentro` es independiente,
consultado por `Cita`.

## Colección de Postman

La colección `postman_collection.json` incluye las tres carpetas (Auth, Técnicos, Solicitantes,
Intervenciones) con variables `{{base_url}}` (default `http://localhost:8080/api/v1`) y
`{{token}}`, que se autocompleta al correr el request de *Login*.

## Testing

Cada módulo trae tests unitarios de service (con mocks de repository vía `testify`) y de
storage, cubriendo casos de éxito y *error paths* (recurso no encontrado, input inválido).

```bash
cd TECNICOS   # o SOLICITANTES / CITAS
go test ./... -v -cover
```

## Estructura de carpetas

```
proyecto_web2/
├── TECNICOS/        # módulo de Anthony (fuente original, ver rama feature/tecnicos)
├── SOLICITANTES/     # módulo de Jandry (fuente original, ver rama feature/solicitantes)
├── CITAS/            # módulo de Mario (fuente original, ver rama feature/intervenciones)
├── proyecto/         # integración: config, Docker, docker-compose
│   ├── cmd/api/main.go
│   ├── internal/{config,handlers,httpserver,middleware,models,service,storage}/
│   ├── Dockerfile
│   └── docker_compose.yml
└── README.md
```

Cada carpeta de módulo (`TECNICOS/`, `SOLICITANTES/`, `CITAS/`) replica esta misma estructura
interna (`cmd/api`, `internal/{handlers,middleware,models,service,storage}`), lo que permitió
fusionarlas en `proyecto/` sin reescribir la lógica de negocio de cada integrante.
