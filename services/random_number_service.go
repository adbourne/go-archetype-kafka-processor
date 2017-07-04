package services

import "math/rand"

// A service providing a random number
type RandomNumberService interface {
	GenerateRandomNumber(int) int
}

// Default implementation of RandomNumberService
type DefaultRandomNumberService struct {
}

// Generates a random number
func (rns *DefaultRandomNumberService) GenerateRandomNumber(seed int) int {
	rand.Seed(int64(seed))
	return rand.Int()
}
