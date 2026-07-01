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

## Stack Tecnológico Obligatorio
El proyecto hace uso de las siguientes herramientas estándar para garantizar consistencia y buenas prácticas:

* **Lenguaje:** Go (Golang) 1.22+
* **Enrutador:** Chi Router
* **ORM y Base de Datos:** GORM con SQLite (aún indefinido)
* **Seguridad:** golang-jwt para autenticación
* **Testing:** Postman (para pruebas unitarias y mocks en hitos posteriores)

## Estructura Inicial del Proyecto
La arquitectura está pensada en capas(con posibilidad a cambios):

- `/cmd/api/`: Punto de entrada principal y configuración del servidor web.
- `/internal/models/`: Estructuras de datos (Structs) y entidades del negocio.
- `/internal/handlers/`: Controladores que reciben peticiones HTTP y devuelven respuestas.
- `/internal/services/`: Lógica de negocio dura y validaciones de reglas.
- `/internal/repositories/`: Operaciones directas contra la base de datos a través de GORM.