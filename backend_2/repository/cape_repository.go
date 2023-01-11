package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"quantumclient.org/backend/v2/models"
)

type CapeRepositoryInterface interface {
	FindCapeByUsername(string) (*models.Cape, error)
	FindCapeByUuid(uuid.UUID) (*models.Cape, error)
	FindCapesByOwnerUuid(uuid.UUID) ([]*models.Cape, error)
	FindAllCapes() ([]*models.Cape, error)
	UpdateCape(models.Cape) error
	DeleteCape(uuid.UUID) error
}

// CapeRepository is the repository for Cape
type CapeRepository struct {
	db *sql.DB
}

// NewCapeRepository returns a new instance of CapeRepository
func NewCapeRepository(db *sql.DB) *CapeRepository {
	return &CapeRepository{db}
}

func (c CapeRepository) FindCapeByUsername(s string) (*models.Cape, error) {
	row := c.db.QueryRow("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` WHERE player_username = ?", s)
	return scanRowForCape(row)
}

func (c CapeRepository) FindCapeByUuid(uuid uuid.UUID) (*models.Cape, error) {
	row := c.db.QueryRow("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` WHERE player_uuid = ?", uuid)
	return scanRowForCape(row)
}

func (c CapeRepository) FindCapesByOwnerUuid(uuid uuid.UUID) ([]*models.Cape, error) {
	rows, err := c.db.Query("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` WHERE account_uuid = ?", uuid)
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	defer rows.Close()

	var capes []*models.Cape
	for rows.Next() {
		cape, er := scanRowsForCape(rows)
		if er != nil {
			slog.Error(er)
			return nil, er
		}
		capes = append(capes, cape)
	}
	return capes, nil
}

func (c CapeRepository) FindAllCapes() ([]*models.Cape, error) {
	rows, err := c.db.Query("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys`")
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	defer rows.Close()

	var capes []*models.Cape
	for rows.Next() {
		cape, er := scanRowsForCape(rows)
		if er != nil {
			slog.Error(er)
			return nil, er
		}
		capes = append(capes, cape)
	}
	return capes, nil
}

func (c CapeRepository) UpdateCape(cape models.Cape) error {
	_, err := c.db.Exec("UPDATE capes INNER JOIN verified_mc_accounts vma on capes.`key` = vma.`keys` SET enabled = ?, url = ?, type = ? WHERE vma.player_uuid = ?", cape.Enabled, cape.Url, cape.Type, cape.Uuid)
	if err != nil {
		slog.Error(err)
		return err
	}
	return nil
}

func (c CapeRepository) DeleteCape(id uuid.UUID) error {
	_, err := c.db.Exec("DELETE capes FROM capes INNER JOIN verified_mc_accounts vma on capes.`key` = vma.`keys` WHERE vma.player_uuid = ?", id)
	if err != nil {
		slog.Error(err)
		return err
	}
	return nil
}

func scanRowForCape(row *sql.Row) (*models.Cape, error) {
	var cape models.Cape
	err := row.Scan(&cape.Uuid, &cape.Username, &cape.Type, &cape.Url, &cape.Enabled, &cape.OwnerUuid)
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	return &cape, nil
}

func scanRowsForCape(row *sql.Rows) (*models.Cape, error) {
	var cape models.Cape
	err := row.Scan(&cape.Uuid, &cape.Username, &cape.Type, &cape.Url, &cape.Enabled, &cape.OwnerUuid)
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	return &cape, nil
}
