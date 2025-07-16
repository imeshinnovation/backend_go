package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"` // El "-" omite el campo en JSON
}

// HashPassword encripta la contraseña del usuario
func (u *User) HashPassword() error {
	// Verificar que la contraseña no esté vacía
	if len(u.Password) == 0 {
		return errors.New("password cannot be empty")
	}

	// Verificar longitud mínima (opcional pero recomendado)
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Convertir a slice de bytes
	passwordBytes := []byte(u.Password)

	// Generar el hash con un coste adecuado
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 14)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifica si la contraseña proporcionada coincide con la almacenada
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
