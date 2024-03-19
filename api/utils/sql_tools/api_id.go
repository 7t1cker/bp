package sql_tools

import (
	"database/sql"
	"net/http"
)

// токены
func GetUserRequest( db *sql.DB, token string) (int, int) {
	sqlStatement := `SELECT id FROM users WHERE access_token = $1 AND status IN ('работает', 'перерыв')`
	row := db.QueryRow(sqlStatement, token)
	var userID int
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusUnauthorized, 0
		}
		return http.StatusInternalServerError, 0
	}

	return http.StatusOK, userID
}
