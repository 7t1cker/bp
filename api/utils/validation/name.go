package validation

import (
	"errors"
	"regexp"

	"github.com/7t1cker/bp/api/models"
)

func ValidateName(name string) bool {

    regex := regexp.MustCompile(`^[a-zA-Zа-яА-Я\s]+$`)
    return regex.MatchString(name)
}

func ValidateUser(newUser models.User) error {
 
    if !ValidateName(newUser.FirstName) {
        return errors.New("Invalid FirstName. Only letters and spaces are allowed")
    }
    if !ValidateName(newUser.LastName) {
        return errors.New("Invalid LastName. Only letters and spaces are allowed")
    }
    if newUser.MiddleName != "" {
        if !ValidateName(newUser.MiddleName) {
            return errors.New("Invalid MiddleName. Only letters and spaces are allowed")
        }
    }

    return nil
}
