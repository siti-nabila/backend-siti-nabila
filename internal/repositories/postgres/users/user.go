package users

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
)

func (u *userRepository) GetUserByUserId(userId int) (result domain.User, err error) {
	query := `
	SELECT 
		u.user_id, 
		u.user_email, 
		u.user_password, 
		r.role_id, 
		r.role_name
	FROM pivot_users_roles ur
	JOIN users u
		ON ur.user_id = u.user_id
	JOIN roles r
		ON ur.role_id = r.role_id
	WHERE u.user_id = $1
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	err = stmt.QueryRow(userId).Scan(
		&result.UserId,
		&result.UserEmail,
		&result.UserPassword,
		&result.RoleId,
		&result.RoleName,
	)

	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	return result, err

}
