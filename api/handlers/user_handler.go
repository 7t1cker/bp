// handlers/user_handler.go

package handlers

import (
	"encoding/json"
	"net/http"

	"yourproject/models" // замените "yourproject" на путь к вашему пакету моделей

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var newUser models.User
        err := json.NewDecoder(r.Body).Decode(&newUser)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Здесь можно добавить валидацию нового пользователя, если необходимо

        // Создание нового пользователя в базе данных с использованием Gorm
        result := db.Create(&newUser)
        if result.Error != nil {
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
            return
        }

        // Отправка ответа об успешном выполнении операции
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("User created successfully"))
    }
}
