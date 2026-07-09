# API de Optimización Técnica para Equipos Personales en Universidad (ULEAM)

## Descripción del Proyecto
Esta API REST busca resolver un problema real y cotidiano en el campus universitario: estudiantes que tienen equipos lentos o desactualizados y no pueden correr las herramientas necesarias para sus clases, pero que no logran coincidir en sus horas libres ("huecos") con compañeros técnicos que tienen el conocimiento y las herramientas para ayudarles.

El sistema cruza dinámicamente los horarios académicos de estudiantes solicitantes y técnicos dentro de la universidad, permitiendo el emparejamiento exacto según la compatibilidad de tiempo, punto de encuentro y tipo de habilidad técnica requerida.

## Equipo de Desarrollo y Módulos de Dominio
El proyecto está construido bajo una arquitectura de micro-módulos de dominio, donde cada integrante es dueño absoluto de su lógica de negocio:

* **Jandry Cedeño:** Módulo de Solicitantes (Gestión de perfiles de estudiantes, especificaciones de hardware defectuoso y registro de horarios libres).
* **Anthony Solórzano:** Módulo de Técnicos (Gestión de perfiles técnicos, inventario de habilidades como instalación de SO o disco duro, y reputación).
* **Mario Soriano:** Módulo de Intervenciones (Máquina de estados del servicio y motor de emparejamiento/matching multicriterio temporal y espacial).
* **Responsabilidad Grupal:** Autenticación JWT, Middlewares, Docker y CI/CD.

## Endpoints Documentados (Por Responsable)

### Módulo de Solicitantes (Responsable: Jandry Cedeño)
- `GET /api/v1/solicitantes`: Lista todos los solicitantes.
- `POST /api/v1/solicitantes`: Crea un nuevo solicitante.
- `GET /api/v1/solicitantes/{id}`: Obtiene un solicitante por su ID.
- `PUT /api/v1/solicitantes/{id}`: Actualiza los datos de un solicitante.
- `DELETE /api/v1/solicitantes/{id}`: Elimina un solicitante del sistema.
*(También incluye endpoints similares para `/dispositivos` y `/tickets`)*

### Módulo de Técnicos (Responsable: Anthony Solórzano)
- `GET /api/v1/tecnicos`: Lista todos los técnicos registrados, junto con sus servicios y horarios.
- `POST /api/v1/tecnicos`: Registra un nuevo técnico.
- `GET /api/v1/tecnicos/{id}`: Obtiene el perfil de un técnico.
- `PUT /api/v1/tecnicos/{id}`: Actualiza los datos de un técnico.
- `DELETE /api/v1/tecnicos/{id}`: Elimina a un técnico y sus dependencias.

### Módulo de Intervenciones (Responsable: Mario Soriano)
- `GET /api/v1/citas`: Lista todas las citas programadas.
- `POST /api/v1/citas`: Agenda una nueva cita (Intervención).
- `GET /api/v1/citas/{id}`: Detalles de una cita.
- `PUT /api/v1/citas/{id}`: Actualiza el estado de la cita (Ej. Pendiente -> Finalizada).
- `DELETE /api/v1/citas/{id}`: Cancela una cita.
*(También incluye endpoints similares para `/puntos-encuentro` y `/soportes`)*

## Stack Tecnológico Obligatorio
El proyecto hace uso de las siguientes herramientas estándar para garantizar consistencia y buenas prácticas:

* **Lenguaje:** Go (Golang) 1.22+
* **Enrutador:** Chi Router
* **ORM y Base de Datos:** `sqlc` y `gorm` con PostgreSQL (vía Docker) y SQLite (local).
* **Seguridad:** `golang-jwt` para autenticación.
* **Testing:** Pruebas unitarias nativas en Go (`testing`, `httptest`) con cobertura mayor al 50%. CI configurado con GitHub Actions.

## Cómo ejecutar el proyecto con Docker (Local)
El proyecto está configurado para levantar toda la infraestructura (API y PostgreSQL con seeders) usando `docker-compose`.

1. Clona el repositorio.
2. Ejecuta el siguiente comando en la raíz del proyecto:
   ```bash
   docker-compose up --build -d
   ```
3. La API estará disponible en `http://localhost:8088` y la base de datos PostgreSQL estará corriendo en el puerto `5432`.

## Estructura Inicial del Proyecto
La arquitectura está pensada en capas:

- `/cmd/api/`: Punto de entrada principal y configuración del servidor web con inyección de dependencias.
- `/internal/models/`: Estructuras de datos (Structs) y entidades del negocio.
- `/internal/handlers/`: Controladores que reciben peticiones HTTP y devuelven respuestas.
- `/internal/service/`: Lógica de negocio dura y validaciones de reglas.
- `/internal/storage/`: Operaciones directas contra la base de datos e interfaces de repositorios.