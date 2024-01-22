package repository

import (
	"check-saldo-service/entity"
	"context"
	"database/sql"
)

type WalletRepository struct {
	DB *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (h *WalletRepository) FindByNorek(ctx context.Context, norek string) (entity.Wallets, error) {
	data := entity.Wallets{}

	query := `SELECT id, customer_id, balance
	FROM wallets
	WHERE id = $1`

	row, err := h.DB.QueryContext(ctx, query, norek)
	if err != nil {
		return data, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&data.ID, &data.CustomerID, &data.Balance); err != nil {
			return data, err
		}
	}

	return data, nil
}
