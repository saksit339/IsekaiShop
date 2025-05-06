package server

import (
	_itemShopController "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/controller"
	_itemShopRepository "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/repository"
	_itemShopService "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, &s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewitemShopControllerImpl(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
