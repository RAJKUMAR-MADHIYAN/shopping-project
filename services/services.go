package services

import (
	"errors"
	"log"
	"shops/models"
	"shops/repository"
)

type ItemService interface {
	CreateItem(input models.CreateItemInput) (*models.Item, error)
	FindItem(id int) (*models.Item, bool, error)
	UpdateItem(id int, input models.UpdateItemInput) (*models.Item, error)
	ListItems() ([]models.Item, error)
	DeleteItem(id int) error
}
type itemService struct {
	r *repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) *ItemService {
	if r == nil {
		log.Fatal("Failed to initialize item service, repository is nil")
		return nil
	}
	var s ItemService = &itemService{r: &r}
	return &s
}

func (s *itemService) GetItemRepository() repository.ItemRepository {
	if s.r == nil {
		log.Fatal("Failed to get item repository, it is nil")
		return nil
	}

	return *s.r
}
func (s *itemService) ListItems() ([]models.Item, error) {
	items, err := s.GetItemRepository().ListItems()
	return items, err
}

func (s *itemService) FindItem(id int) (*models.Item, bool, error) {
	r := s.GetItemRepository()
	item, found, err := r.FindItem(id)
	return item, found, err
}
func (s *itemService) CreateItem(input models.CreateItemInput) (*models.Item, error) {

	r := s.GetItemRepository()

	sold := false
	data := &models.Item{
		Name:  &input.Name,
		Price: &input.Price,
		Sold:  &sold,
	}
	item, err := r.CreateItem(data)
	return item, err
}
func (s *itemService) UpdateItem(id int, input models.UpdateItemInput) (*models.Item, error) {
	// Assumed input is validated on upper (handlers) layer
	r := s.GetItemRepository()

	item, found, err := r.FindItem(id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("item not found")
	}

	// if a field is missing in the update input, keep the old value for this field
	name := input.Name
	if name == nil {
		name = item.Name
	}
	price := input.Price
	if price == nil {
		price = item.Price
	}
	sold := input.Sold
	if sold == nil {
		sold = item.Sold
	}

	data := &models.Item{
		Name:  name,
		Price: price,
		Sold:  sold,
	}
	item, err = r.UpdateItem(id, data)
	return item, err
}
func (s *itemService) DeleteItem(id int) error {
	r := s.GetItemRepository()
	err := r.DeleteItem(id)
	return err
}
