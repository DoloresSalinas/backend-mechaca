package handlers

import (
	"context"
	"backend/config"
	"backend/models" 
	"time" 
	"github.com/gofiber/fiber/v2"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GET /users
func GetUsers(c *fiber.Ctx) error {
	rows, err := config.Conn.Query(context.Background(), `
		SELECT usuario_id, nombre, app, apm, email, password, telefono, direccion,
		       fecha_nacimiento, rol, fecha_registro, cedula_profesional,
		       especialidad, tipo_sangre, numero_seguro, estatus_us
		FROM usuarios`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.UsuarioID, &u.Nombre, &u.App, &u.Apm, &u.Email, &u.Password, &u.Telefono,
			&u.Direccion, &u.FechaNacimiento, &u.Rol, &u.FechaRegistro, &u.CedulaProfesional,
			&u.Especialidad, &u.TipoSangre, &u.NumeroSeguro, &u.EstatusUs,
		)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		users = append(users, u)
	}
	return c.JSON(users)
}

// GET /users/:id
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var u models.User
	err := config.Conn.QueryRow(context.Background(), `
		SELECT usuario_id, nombre, app, apm, email, password, telefono, direccion,
		       fecha_nacimiento, rol, fecha_registro, cedula_profesional,
		       especialidad, tipo_sangre, numero_seguro, estatus_us
		FROM usuarios WHERE usuario_id=$1`, id).
		Scan(
			&u.UsuarioID, &u.Nombre, &u.App, &u.Apm, &u.Email, &u.Password, &u.Telefono,
			&u.Direccion, &u.FechaNacimiento, &u.Rol, &u.FechaRegistro, &u.CedulaProfesional,
			&u.Especialidad, &u.TipoSangre, &u.NumeroSeguro, &u.EstatusUs,
		)
	if err != nil {
		return c.Status(404).SendString("Usuario no encontrado")
	}
	return c.JSON(u)
}


// PUT /users/:id
func UpdateUser(c *fiber.Ctx) error {
	inputID := c.Params("id")
	var u models.UserInput
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).SendString("Error al parsear entrada")
	}

	query := "UPDATE usuarios SET "
	args := []interface{}{}
	argIndex := 1

	if u.Nombre != "" {
		query += fmt.Sprintf("nombre=$%d, ", argIndex)
		args = append(args, u.Nombre)
		argIndex++
	}
	if u.App != "" {
		query += fmt.Sprintf("app=$%d, ", argIndex)
		args = append(args, u.App)
		argIndex++
	}
	if u.Apm != "" {
		query += fmt.Sprintf("apm=$%d, ", argIndex)
		args = append(args, u.Apm)
		argIndex++
	}
	if u.Email != "" {
		query += fmt.Sprintf("email=$%d, ", argIndex)
		args = append(args, u.Email)
		argIndex++
	}
	if u.Password != "" {
		// Hashear la contraseña
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se pudo encriptar la contraseña",
			})
		}
		query += fmt.Sprintf("password=$%d, ", argIndex)
		args = append(args, string(hash))
		argIndex++
	}
	if u.Telefono != "" {
		query += fmt.Sprintf("telefono=$%d, ", argIndex)
		args = append(args, u.Telefono)
		argIndex++
	}
	if u.Direccion != "" {
		query += fmt.Sprintf("direccion=$%d, ", argIndex)
		args = append(args, u.Direccion)
		argIndex++
	}
	if u.FechaNacimiento != "" {
		nac, err := time.Parse("01/01/1980", u.FechaNacimiento)
		if err != nil {
			return c.Status(400).SendString("Formato de fecha_nacimiento inválido. Usa dd/mm/yyyy")
		}
		query += fmt.Sprintf("fecha_nacimiento=$%d, ", argIndex)
		args = append(args, nac)
		argIndex++
	}
	if u.Rol != "" {
		query += fmt.Sprintf("rol=$%d, ", argIndex)
		args = append(args, u.Rol)
		argIndex++
	}
	if u.CedulaProfesional != "" {
		query += fmt.Sprintf("cedula_profesional=$%d, ", argIndex)
		args = append(args, u.CedulaProfesional)
		argIndex++
	}
	if u.Especialidad != "" {
		query += fmt.Sprintf("especialidad=$%d, ", argIndex)
		args = append(args, u.Especialidad)
		argIndex++
	}
	if u.TipoSangre != "" {
		query += fmt.Sprintf("tipo_sangre=$%d, ", argIndex)
		args = append(args, u.TipoSangre)
		argIndex++
	}
	if u.NumeroSeguro != "" {
		query += fmt.Sprintf("numero_seguro=$%d, ", argIndex)
		args = append(args, u.NumeroSeguro)
		argIndex++
	}
	if u.EstatusUs != "" {
		query += fmt.Sprintf("estatus_us=$%d, ", argIndex)
		args = append(args, u.EstatusUs)
		argIndex++
	}

	// Si no hay campos para actualizar
	if len(args) == 0 {
		return c.Status(400).SendString("No se enviaron campos para actualizar")
	}

	// Eliminar la última coma
	query = query[:len(query)-2]
	// WHERE
	query += fmt.Sprintf(" WHERE usuario_id=$%d", argIndex)
	args = append(args, inputID)

	// Ejecutar la consulta
	_, err := config.Conn.Exec(context.Background(), query, args...)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Usuario actualizado correctamente")
}

// DELETE /users/:id
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := config.Conn.Exec(context.Background(), "DELETE FROM usuarios WHERE usuario_id=$1", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Usuario eliminado correctamente")
}
