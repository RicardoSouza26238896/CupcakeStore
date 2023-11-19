package services

import (
	"github.com/bitebait/cupcakestore/models"
	"github.com/bitebait/cupcakestore/repositories"
)

type ShoppingCartService interface {
	FindByUserId(id uint) (models.ShoppingCart, error)
}

type shoppingCartService struct {
	shoppingCartRepository repositories.ShoppingCartRepository
}

func NewShoppingCartService(shoppingCartRepository repositories.ShoppingCartRepository) ShoppingCartService {
	return &shoppingCartService{
		shoppingCartRepository: shoppingCartRepository,
	}
}

func (s shoppingCartService) FindByUserId(id uint) (models.ShoppingCart, error) {
	return s.shoppingCartRepository.FindByUserId(id)
}
