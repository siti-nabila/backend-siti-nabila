package settings

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
)

type settingRepository struct {
	db *sql.DB
}

func NewSettingPostgresRepository(db *sql.DB) domain.SettingRepository {
	return &settingRepository{
		db: db,
	}
}

func (s *settingRepository) GetSettingByKey(settingKey string) (result domain.Setting, err error) {
	query := `
		SELECT 
			setting_key,
			setting_value
		FROM settings
		WHERE setting_key = $1
	`
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	err = stmt.QueryRow(settingKey).Scan(
		&result.SettingKey,
		&result.SettingValue,
	)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, err
}
