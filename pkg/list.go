package field

import "log"

type MaterialList []Item

func (m MaterialList) UpdateQuantityRequested(uuid ProductUUID, quantity uint) {
	i, item := m.findItem(uuid)

	if item.ProductUUID == "" {
		log.Println(ErrItemNotFound)
		return
	}

	m[i] = m[i].updateQuantityRequested(quantity)
}

func (m MaterialList) ReceiveItem(uuid ProductUUID, quantity uint) {
	i, item := m.findItem(uuid)

	if item.ProductUUID == "" {
		log.Println(ErrItemNotFound)
		return
	}

	m[i] = m[i].receiveItem(quantity)
}

func (m MaterialList) findItem(uuid ProductUUID) (int, Item) {
	for i := range m {
		if m[i].ProductUUID == uuid {
			return i, m[i]
		}
	}
	return -1, Item{}
}

func (m MaterialList) receivedAll() bool {
	for i := range m {
		if m[i].ItemStatus != Filled {
			return false
		}
	}
	return true
}

//func (m MaterialList) AdjustQuantityRequested(id ProductUUID, qr QuantityRequested) {
//	if qr <= 0 {
//		log.Println(ErrQuantityZero)
//		return
//	}
//
//	i := m.FindItem(id)
//	if i < 0 {
//		log.Println(ErrItemNotFound)
//		return
//	}
//
//	m[i].QuantityRequested = qr
//}

//func (m MaterialList) FindItem(uuid ProductUUID) Item {
//	for i := range m {
//		if m[i].ProductUUID == uuid {
//			m[i].Index = i
//			return m[i]
//		}
//	}
//	return Item{}
//}
