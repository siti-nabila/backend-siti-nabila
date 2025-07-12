package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (u *userRepository) Register(request models.RegisterRequest) (result domain.User, err error) {
	query := `
	WITH u AS (
			INSERT INTO users(user_email, user_password)
			VALUES ($1, $2)
			ON CONFLICT (user_email) DO NOTHING
			RETURNING user_id
		)
	INSERT INTO pivot_users_roles (user_id, role_id)
	SELECT u.user_id, $3 FROM u RETURNING user_id;
	`
	tx, err := u.db.Begin()
	if err != nil {
		log.Error(err)
		return result, err
	}
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Error(err)
		return result, err
	}

	err = stmt.QueryRow(request.Email, request.Password, request.RoleId).Scan(
		&result.UserId,
	)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		log.Error(err)
		return result, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return result, err
	}

	return result, err
}

func (u *userRepository) Login(request models.LoginReqeust) (result domain.User, err error) {

	query := `
		SELECT user_id, user_password FROM users WHERE user_email = $1 
	`

	stmt, err := u.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	err = stmt.QueryRow(request.Email).Scan(
		&result.UserId,
		&result.UserPassword,
	)

	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	return result, err
}
