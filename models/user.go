package models

import "time"

type User struct {
	UsuarioID        int       `json:"usuario_id"`
	Nombre           string    `json:"nombre"`
	App              string    `json:"app"`
	Apm              string    `json:"apm"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Telefono         string    `json:"telefono"`
	Direccion        string    `json:"direccion"`
	FechaNacimiento  time.Time `json:"fecha_nacimiento"`
	Rol              string    `json:"rol"`
	FechaRegistro    time.Time `json:"fecha_registro"`
	CedulaProfesional string   `json:"cedula_profesional"`
	Especialidad     string    `json:"especialidad"`
	TipoSangre       string    `json:"tipo_sangre"`
	NumeroSeguro     string    `json:"numero_seguro"`
	EstatusUs        string    `json:"estatus_us"`
}


type UserInput struct {
	Nombre            string `json:"nombre"`
	App               string `json:"app"`
	Apm               string `json:"apm"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Telefono          string `json:"telefono"`
	Direccion         string `json:"direccion"`
	FechaNacimiento   string `json:"fecha_nacimiento"` // string para validación manual
	Rol               string `json:"rol"`
	CedulaProfesional string `json:"cedula_profesional"`
	Especialidad      string `json:"especialidad"`
	TipoSangre        string `json:"tipo_sangre"`
	NumeroSeguro      string `json:"numero_seguro"`
	EstatusUs         string `json:"estatus_us"`
}
