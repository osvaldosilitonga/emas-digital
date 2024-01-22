package repository

import (
	"buyback-storage/dto"
	"context"
	"database/sql"
	"log"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (tr *TransactionRepository) SaveTransaction(ctx context.Context, data *dto.Buyback) error {
	query := `INSERT INTO transactions (wallet_id, type, gram, price_id, balance) VALUES ($1, $2, $3, $4, $5)`

	_, err := tr.DB.ExecContext(ctx, query, data.NoRek, "buyback", data.Gram, data.PriceID, data.Saldo)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TransactionRepository) BuybackSaldo(ctx context.Context, id string, saldo float32) error {
	query := `UPDATE wallets SET balance = $1 WHERE id = $2`

	_, err := tr.DB.ExecContext(ctx, query, saldo, id)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TransactionRepository) CheckSaldo(ctx context.Context, norek string) (float32, error) {
	query := `SELECT balance FROM wallets WHERE id = $1`

	row, err := tr.DB.QueryContext(ctx, query, norek)
	if err != nil {
		return 0, err
	}

	var saldo float32
	for row.Next() {
		if err := row.Scan(&saldo); err != nil {
			log.Println(err)
		}
	}

	return saldo, nil
}
