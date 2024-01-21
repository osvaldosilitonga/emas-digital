package repository

import (
	"check-harga-service/dto"
	"context"
	"database/sql"
)

type HargaRepository struct {
	DB *sql.DB
}

func NewHargaRepository(db *sql.DB) *HargaRepository {
	return &HargaRepository{
		DB: db,
	}
}

func (h *HargaRepository) FindOne(ctx context.Context) (dto.Data, error) {
	data := dto.Data{}

	query := `SELECT topup, buyback
	FROM prices
	ORDER BY created_at DESC
	LIMIT 1`

	row, err := h.DB.QueryContext(ctx, query)
	if err != nil {
		return data, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&data.HargaTopup, &data.HargaBuyback); err != nil {
			return data, err
		}
	}

	return data, nil
}
