# Backend - Sistema de Citas y Reportes Hospitalarios

Este proyecto es el backend de un sistema hospitalario que permite gestionar usuarios, citas médicas y reportes clínicos. Está desarrollado en **Go** utilizando el framework [Fiber](https://gofiber.io/) y se conecta a una base de datos **PostgreSQL** mediante [Supabase](https://supabase.com/), usando una cadena de conexión segura (Connection String).

## 🩺 Características

- Registro y login de usuarios con autenticación JWT
- Hashing de contraseñas con bcrypt
- CRUD completo para usuarios (GET, POST, PUT, DELETE)
- Conexión a Supabase (PostgreSQL) usando pgx y cadena de conexión
- Middleware para rutas protegidas
- Manejo seguro de variables de entorno
- Enfoque en privacidad, seguridad y cumplimiento de normativas médicas

## 🛠 Tecnologías utilizadas

- Go v1.24.x
- Fiber v2.x
- Supabase (PostgreSQL)
- pgx v5 (para conexión PostgreSQL)
- bcrypt (para cifrado de contraseñas)
- JWT (github.com/golang-jwt/jwt/v5)
- godotenv (para manejo de variables de entorno)

## 🚀 Instalación

1. Clona el repositorio:


git clone https://github.com/TuUsuario/backend-hospital.git
cd backend-hospital


2. Instala las dependencias:


go mod tidy


3. Crea un archivo .env en la raíz del proyecto con tus credenciales Supabase:


user=postgres.kzarztpehygcitcjozpp
password=TU_PASSWORD
host=aws-0-us-east-2.pooler.supabase.com
port=6543
dbname=postgres


4. Corre la base de datos Supabase desde tu panel, y asegúrate de tener las tablas necesarias creadas (usuarios, etc.)


5. Ejecuta el servidor:


go run main.go

## 🔐 Autenticación
Para acceder a rutas protegidas, primero debes iniciar sesión y obtener un token JWT. Luego, inclúyelo en el encabezado de tus peticiones:


Authorization: Bearer <tu_token>


## 📡 Endpoints principales

| Método | Ruta               | Descripción                              |
|:------:|:------------------:|:----------------------------------------:|
| POST   | `/usuarios/register`    | Registrar un nuevo usuario               |
| POST   | `/usuarios/login`       | Iniciar sesión (retorna JWT)             |
| GET    | `/usuarios/users`       | Obtener todos los usuarios               |
| GET    | `/usuarios/users/:id`   | Obtener un usuario por ID                |
| PUT    | `/usuarios/users/:id`   | Actualizar los datos de un usuario       |
| DELETE | `/usuarios/users/:id`   | Eliminar un usuario                      |


⚠️ Próximamente: módulos para citas médicas y reportes clínicos.


## 📁 Estructura del proyecto

/backend
│
├── config/           # Conexión a Supabase (db.go)
├── handlers/         # Lógica de endpoints (auth.go, users.go, etc.)
├── middleware/       # Autenticación y validaciones
├── models/           # Estructuras de datos (User, etc.)
├── routes/           # Definición de rutas API
├── utils/            # Funciones JWT y herramientas auxiliares
├── main.go           # Archivo principal
├── .env              # Variables de entorno (NO subir a Git)
└── .gitignore


## 👩‍⚕️ Autor
María Dolores Salinas Jiménez