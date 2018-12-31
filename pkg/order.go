package material

import (
	"errors"
	"log"
	"time"
)

var (
	ErrOrderNotFound    = errors.New("order not found")
	ErrQuantityZero     = errors.New("quantity must be greater than 0")
	ErrMustHaveItems    = errors.New("order must have at least 1 item")
	ErrOrderAlreadySent = errors.New("order already sent")
)

type OrderID string
type ProjectID string

type Order struct {
	OrderID   OrderID
	ProjectID ProjectID
	version   int
	Statuses  []status
	LineItems []LineItem
}

func NewOrder(o OrderID, p ProjectID) (*Order, error) {
	return &Order{
		OrderID:   o,
		ProjectID: p,
		version:   0,
		Statuses: []status{
			{
				Date: time.Now(),
				Type: New,
			},
		},
		LineItems: nil,
	}, nil
}

func (o *Order) AddItemToOrder(id ProductID, name string, uom UOM, quantity int) error {
	if err := checkQuantityNotZero(quantity); err != nil {
		return err
	}
	o.LineItems = append(o.LineItems, LineItem{
		ProductID: id,
		Name:      name,
		UOM:       uom,
		Quantity:  quantity,
		status:    Waiting,
		PO:        "",
	})
	return nil
}

func (o *Order) RemoveItemFromOrder(id ProductID) error {
	for i := range o.LineItems {
		if o.LineItems[i].ProductID == id {
			copy(o.LineItems[i:], o.LineItems[i+1:])
			o.LineItems[len(o.LineItems)-1] = LineItem{}
			o.LineItems = o.LineItems[:len(o.LineItems)-1]
			break
		}
	}
	return nil
}

func (o *Order) UpdateItemQuantity(id ProductID, q int) error {
	if q <= 0 {
		return ErrQuantityZero
	}

	for i, l := range o.LineItems {
		if l.ProductID == id {
			o.LineItems[i].Quantity = q
		}
	}
	return nil
}

func (o *Order) SendOrder() error {
	if err := o.checkOrderHasItems(); err != nil {
		log.Fatal(err)
		return err
	}

	if o.Statuses[len(o.Statuses)-1].Type != New {
		return ErrOrderAlreadySent
	}
	o.Statuses = append(o.Statuses, status{
		Date: time.Time{},
		Type: Sent,
	})
	return nil
}

func (o *Order) checkOrderHasItems() error {
	if len(o.LineItems) <= 0 {
		return ErrMustHaveItems
	}
	return nil
}

func checkQuantityNotZero(quantity int) error {
	if quantity <= 0 {
		return ErrQuantityZero
	}
	return nil
}

type status struct {
	Date time.Time
	Type OrderStatus
}

type OrderStatus int

const (
	New OrderStatus = iota
	Sent
	Received
	Processing
	OnRoute
	Complete
)

func (s OrderStatus) String() string {
	switch s {
	case New:
		return "New"
	case Sent:
		return "Sent"
	case Received:
		return "Received"
	case Processing:
		return "Processing"
	case OnRoute:
		return "OnRoute"
	case Complete:
		return "Complete"
	}
	return ""
}
