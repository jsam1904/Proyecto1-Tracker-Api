# Proyecto1-Tracker-Api

REST API en Go + PostgreSQL para el Bundesliga Tracker.

> 🔗 **Frontend:** https://github.com/jsam1904/Proyecto1-Bundesliga-Client

## Stack
- Go 1.21
- chi v5 (router)
- PostgreSQL

## ¿Qué es CORS?
CORS es una política del navegador que bloquea peticiones a un origen distinto al del cliente. Como el cliente y la API corren en puertos diferentes, el servidor debe permitirlo explícitamente. Esta API permite todos los orígenes (`*`) durante desarrollo.

## Setup
```bash
cp .env.example .env
# Edita .env con tus credenciales de PostgreSQL
chmod +x setup.sh && ./setup.sh
go run main.go
```

Si estás en Windows, ejecuta esos comandos desde Git Bash o WSL, o crea la base de datos y corre la migración manualmente.

Servidor en `http://localhost:8080` por defecto, o en el puerto que definas con `PORT`.

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/teams` | Listar equipos |
| GET | `/teams?q=bayern` | Buscar por nombre |
| GET | `/teams?page=1&limit=5` | Paginación (por defecto `page=1`, `limit=10`) |
| GET | `/teams?sort=founded&order=asc` | Ordenamiento (`name`, `city`, `founded`, `stadium`) |
| GET | `/teams/{id}` | Obtener un equipo |
| POST | `/teams` | Crear equipo |
| PUT | `/teams/{id}` | Editar equipo |
| DELETE | `/teams/{id}` | Eliminar equipo |

### Filtros disponibles
- `q`: busca por nombre con coincidencia parcial.
- `page` y `limit`: controlan la paginación.
- `sort`: acepta `name`, `city`, `founded` o `stadium`.
- `order`: acepta `asc` o `desc`.

### Payload de equipo
- `name`, `city`, `stadium` y `founded` son obligatorios al crear.
- `image_url` es opcional.
- En actualización, todos los campos son opcionales.

## Códigos HTTP
| Código | Cuándo |
|--------|--------|
| 200 | GET / PUT exitoso |
| 201 | POST exitoso |
| 204 | DELETE exitoso |
| 400 | Input inválido |
| 404 | No encontrado |
| 500 | Error del servidor |

## Challenges
- [x] Códigos HTTP correctos
- [x] Validación server-side con JSON descriptivo
- [x] Búsqueda `?q=`
- [x] Paginación `?page=&limit=`
- [x] Ordenamiento `?sort=&order=`

## Reflexión
Go fue bastante limpio para una API REST. La librería estándar maneja HTTP bien y chi agrega routing sin mucho overhead. PostgreSQL con `database/sql` requiere más boilerplate que un ORM, pero da control total sobre las queries. Lo volvería a usar para APIs donde el rendimiento importa.