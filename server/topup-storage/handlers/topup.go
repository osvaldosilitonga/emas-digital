package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"topup-storage/dto"
	"topup-storage/repository"
)

type TopupHandler struct {
	TransactionRepo *repository.TransactionRepository
}

func NewTopupHandler(tr *repository.TransactionRepository) *TopupHandler {
	return &TopupHandler{
		TransactionRepo: tr,
	}
}

func (th *TopupHandler) Topup(message *dto.Topup) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	saldo, err := th.TransactionRepo.CheckSaldo(ctx, message.NoRek)
	if err != nil {
		return err
	}

	saldo += message.Gram
	s := fmt.Sprintf("%.3f", saldo)
	f, _ := strconv.ParseFloat(s, 32)
	message.Saldo = float32(f)

	if err := th.TransactionRepo.TopupSaldo(ctx, message.NoRek, message.Saldo); err != nil {
		return err
	}

	if err := th.TransactionRepo.SaveTransaction(ctx, message); err != nil {
		return err
	}

	return nil
}
