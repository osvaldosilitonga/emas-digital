package repository

import (
	"context"
	"database/sql"
	"input-harga-storage/dto"
)

type HargaRepository struct {
	DB *sql.DB
}

func NewHargaRepository(db *sql.DB) *HargaRepository {
	return &HargaRepository{
		DB: db,
	}
}

func (hr *HargaRepository) Save(ctx context.Context, harga *dto.Harga) error {
	query := `INSERT INTO prices (id, admin_id, topup, buyback) VALUES ($1, $2, $3, $4)`

	_, err := hr.DB.ExecContext(ctx, query, harga.ID, harga.AdminID, harga.HargaTopup, harga.HargaBuyback)
	if err != nil {
		return err
	}

	return nil
}
