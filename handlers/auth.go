package handlers

import (
	"context"
	"backend/config"
	"backend/models" 
	"backend/utils"
	"time" 
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// POST /users
func Register(c *fiber.Ctx) error {
	var u models.UserInput

	// Parsear JSON
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).SendString("Error al parsear entrada: " + err.Error())
	}

	// Parsear fecha (string a time.Time)
	var fechaNacimiento time.Time
	if u.FechaNacimiento != "" {
		var err error
		fechaNacimiento, err = time.Parse("02/01/2006", u.FechaNacimiento)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Fecha inválida, formato dd/mm/yyyy"})
		}
	}

	// Hashear password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al encriptar contraseña"})
	}

	// Variables para Scan
	var usuarioID int
	var fechaRegistro time.Time

	// Ejecutar query
	err = config.Conn.QueryRow(context.Background(), `
		INSERT INTO usuarios (
			nombre, app, apm, email, password, telefono, direccion,
			fecha_nacimiento, rol, cedula_profesional, especialidad,
			tipo_sangre, numero_seguro, estatus_us
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
		RETURNING usuario_id, fecha_registro
	`,
		u.Nombre, u.App, u.Apm, u.Email, string(hashedPassword),
		u.Telefono, u.Direccion, fechaNacimiento, u.Rol,
		u.CedulaProfesional, u.Especialidad, u.TipoSangre,
		u.NumeroSeguro, u.EstatusUs,
	).Scan(&usuarioID, &fechaRegistro)

	if err != nil {
		return c.Status(500).SendString("Error al insertar usuario: " + err.Error())
	}

	// Retornar resultado
	return c.Status(201).JSON(fiber.Map{
		"usuario_id":     usuarioID,
		"fecha_registro": fechaRegistro,
		"mensaje":        "Usuario creado correctamente",
	})
}


// LOGIN
func Login(c *fiber.Ctx) error {
	// Parsear body con email y password
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al parsear datos"})
	}

	var id int
	var nombre string
	var hashedPassword string	

	err := config.Conn.QueryRow(context.Background(), `
		SELECT usuario_id, nombre, password FROM usuarios WHERE email=$1`, req.Email).Scan(&id, &nombre, &hashedPassword)
		
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email incorrecto"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Contraseña incorrecta"})
	}


	token, err := utils.CrearToken(id, nombre) 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo generar el token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Usuario logueado",
		"token": token,
	})
}
