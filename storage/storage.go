package storage

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hacdan/receipt-processor-challenge/types"
)

type Storage struct {
	store *map[string]types.Receipt
}

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewStorage() Storage {
	mStore := make(map[string]types.Receipt)
	return Storage{
		store: &mStore,
	}
}

func (s *Storage) AddReceipt(receipt types.Receipt) types.ReceiptId {
	id := uuid.NewString()
	(*s.store)[id] = receipt

	return types.ReceiptId{
		Id: id,
	}
}

func (s *Storage) GetReceipt(id string) (*types.Receipt, error) {
	receipt, ok := (*s.store)[id]
	if !ok {
		return nil, errors.New("Receipt not found")
	}
	return &receipt, nil
}

// - One point for every alphanumeric character in the retailer name.
// - 50 points if the total is a round dollar amount with no cents.
// - 25 points if the total is a multiple of `0.25`.
// - 5 points for every two items on the receipt.
// - If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
// - 6 points if the day in the purchase day is odd.
// - 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func (s *Storage) GetPoints(id string) (int, error) {
	receipt, err := s.GetReceipt(id)
	if err != nil {
		return -1, err
	}
	points := 0

	points += countAlphanumeric(receipt.Retailer)

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Trunc(total) {
		points += 50
	}

	totalItems := len(receipt.Items)
	if totalItems > 2 {
		points += totalItems / 2 * 5
	}

	purchasedDateTime, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%v %v", receipt.PurchaseDate, receipt.PurchaseTime))
	if err != nil {
		return points, err
	}

	if purchasedDateTime.Day()%2 != 0 {
		points += 6
	}

	if purchasedDateTime.Hour() >= 14 && purchasedDateTime.Hour() <= 16 {
		points += 10
	}

	return points, nil
}

func countAlphanumeric(s string) int {
	count := 0

	for _, v := range s {
		if strings.ContainsRune(alphanumeric, v) {
			count++
		}
	}
	return count
}
