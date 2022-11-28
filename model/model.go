package model

type Customer struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type Customers []Customer

// remove item on database
func (database Customers) Remove(index int) Customers {
	return append(database[:index], database[index+1:]...)
}
