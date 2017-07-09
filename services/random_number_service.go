package services

import "math/rand"

// RandomNumberService generates a random number using a seed
type RandomNumberService interface {
	GenerateRandomNumber(int) int
}

// DefaultRandomNumberService is the default implementation of RandomNumberService
type DefaultRandomNumberService struct {
}

// NewDefaultRandomNumberService creates a new DefaultRandomNumberService
func NewDefaultRandomNumberService() *DefaultRandomNumberService {
	return &DefaultRandomNumberService{}
}

// GenerateRandomNumber generates a random number
func (rns *DefaultRandomNumberService) GenerateRandomNumber(seed int) int {
	rand.Seed(int64(seed))
	return rand.Int()
}
