package models

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
)

// User is the model of a user object
type User struct {
    UUID     uuid.UUID `json:"uuid,omitempty"`
    UserName string    `json:"user_name"`
    Password string    `json:"password"`
    ConfirmPassword string    `json:"confirm_password,omitempty"`
}

// GenerateUUID genereates a new UUID for a user.
func (user *User) GenerateUUID() {
    user.UUID = uuid.New()
}

// HashPassword hashes the given password and sets the hash to the users password.
func (user *User) HashPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if (err != nil) {
        return err
    }

    user.Password = string(bytes)

    return nil
}

// CheckPassword compares the provided password and the users password hash to make sure
// they're the same.
func (user *User) CheckPassword(providedPassword string) error {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
    if (err != nil) {
        return err
    }

    return nil
}
