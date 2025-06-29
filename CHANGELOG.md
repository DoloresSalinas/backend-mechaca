# Changelog

Todas las modificaciones importantes en este proyecto se documentan en este archivo, siguiendo el formato semántico.

## [Unreleased]

- Incorporación de módulo para gestión de citas médicas.
- Generación automática de reportes clínicos en PDF.
- Implementación de roles (admin, médico, paciente).
- Control de acceso por permisos específicos.
- Registro de logs y auditoría de cambios.

## [1.0.0] - 2025-06-28

### Añadido
- Proyecto inicial del backend del Sistema de Citas y Reportes Hospitalarios.
- Conexión a base de datos Supabase (PostgreSQL) usando `pgx` y variables de entorno.
- Estructura modular: `handlers`, `routes`, `models`, `config`, `utils`.
- Endpoints REST para usuarios: registro, login, obtener todos, obtener por ID, actualizar y eliminar.
- Hashing de contraseñas con `bcrypt`.
- Autenticación segura con JWT (válidos por 10 minutos).
- Middleware para protección de rutas mediante token.
- Validación de datos de entrada para usuarios.
- Manejo de errores y respuestas estandarizadas.
- Archivo `.env` para configuración sensible y `.gitignore` para evitar su inclusión en Git.
- Documentación inicial (`README.md`).
- Avisos de privacidad contemplados en el diseño y flujo.

### Corregido 
- Manejo de errores en conexión con Supabase.
- Ajuste en validación de campos vacíos para actualización parcial.

### Mejorado
- Modularización del código para facilitar mantenimiento. 
- Refuerzo en el control de errores HTTP y mensajes amigables.
