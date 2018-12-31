package ordering

import (
	"field/pkg"
	"log"
)

type Service struct {
	db OrderRepository
}

type OrderRepository interface {
	Save(o *material.Order) error
	Find(id material.OrderID) (*material.Order, error)
	FindAllFromProject(id material.ProjectID) ([]*material.Order, error)
}

func NewOrderingService(db OrderRepository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateNewOrder(o material.OrderID, p material.ProjectID) (*material.Order, error) {
	order, err := material.NewOrder(o, p)
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.Save(order)
	if err != nil {
		log.Fatal(err)
	}

	order.AddItemToOrder(material.LineItem{
		ProductID: "1",
		Name:      "Test",
		UOM:       0,
		Quantity:  7,
		PO:        "",
	})

	order.AddItemToOrder(material.LineItem{
		ProductID: "2",
		Name:      "Test",
		UOM:       0,
		Quantity:  2,
		PO:        "",
	})

	order.AddItemToOrder(material.LineItem{
		ProductID: "3",
		Name:      "Test",
		UOM:       0,
		Quantity:  13,
		PO:        "",
	})

	//order.RemoveItemFromOrder("1")
	order.UpdateItemQuantity("3", 9)

	return order, nil
}

func (s *Service) FindOrder(id material.OrderID) (*material.Order, error) {
	findAll, err := s.db.Find(id)
	if err != nil {
		log.Fatal(err)
	}
	return findAll, nil
}

func (s *Service) FindAllProjectOrders(id material.ProjectID) ([]*material.Order, error) {
	findAll, err := s.db.FindAllFromProject(id)
	if err != nil {
		log.Fatal(err)
	}
	return findAll, nil
}
