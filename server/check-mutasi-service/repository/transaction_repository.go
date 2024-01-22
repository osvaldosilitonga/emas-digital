package repository

import (
	"check-mutasi-service/entity"
	"context"
	"database/sql"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (h *TransactionRepository) FindByNorek(ctx context.Context, norek string, start, end int64) ([]entity.Transaction, error) {
	data := []entity.Transaction{}

	query := `SELECT t.date, t.type, t.gram, p.topup, p.buyback, t.balance
	FROM transactions AS t
	JOIN prices AS p ON t.price_id = p.id
	WHERE t.wallet_id = $1 
	AND t.date BETWEEN $2 AND $3
	ORDER BY t.date DESC`

	row, err := h.DB.QueryContext(ctx, query, norek, start, end)
	if err != nil {
		return data, err
	}
	defer row.Close()

	for row.Next() {
		d := entity.Transaction{}
		if err := row.Scan(&d.Date, &d.Type, &d.Gram, &d.HargaTopup, &d.HargaBuyback, &d.Balance); err != nil {
			return data, err
		}

		data = append(data, d)
	}

	return data, nil
}
