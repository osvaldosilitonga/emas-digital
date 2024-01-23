package handlers

import (
	"buyback-storage/dto"
	"buyback-storage/repository"
	"context"
	"fmt"
	"strconv"
	"time"
)

type BuybackHandler struct {
	TransactionRepo *repository.TransactionRepository
}

func NewBuybackHandler(tr *repository.TransactionRepository) *BuybackHandler {
	return &BuybackHandler{
		TransactionRepo: tr,
	}
}

func (th *BuybackHandler) Buyback(message *dto.Buyback) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	saldo, err := th.TransactionRepo.CheckSaldo(ctx, message.NoRek)
	if err != nil {
		return err
	}

	saldo -= message.Gram
	s := fmt.Sprintf("%.3f", saldo)
	f, _ := strconv.ParseFloat(s, 32)
	message.Saldo = float32(f)

	if err := th.TransactionRepo.BuybackSaldo(ctx, message.NoRek, message.Saldo); err != nil {
		return err
	}

	if err := th.TransactionRepo.SaveTransaction(ctx, message); err != nil {
		return err
	}

	return nil
}
