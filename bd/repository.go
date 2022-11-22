package bd

import "CRMBackend/Udacity/model"

var Database = model.Customers{
	model.Customer{
		ID:        0,
		Name:      "Thiago Garcia",
		Role:      "Backend Developer CEO",
		Email:     "fisica.thiagogarcia@gmail.com",
		Phone:     "987654321",
		Contacted: false,
	},
	model.Customer{
		ID:        1,
		Name:      "Matt Smith",
		Role:      "Frontend Developer",
		Email:     "matt.smith@gmail.com",
		Phone:     "997654321",
		Contacted: false,
	},
	model.Customer{
		ID:        2,
		Name:      "Marta Waynne",
		Role:      "QA",
		Email:     "marta.waynne@gmail.com",
		Phone:     "990654321",
		Contacted: false,
	},
	model.Customer{
		ID:        3,
		Name:      "Leanne Graham",
		Role:      "Mobile Developer",
		Email:     "leanne.graham@gmail.com",
		Phone:     "987654321",
		Contacted: false,
	},
	model.Customer{
		ID:        4,
		Name:      "Glenna Reichert",
		Role:      "Project Manager",
		Email:     "glenna.reichert@gmail.com",
		Phone:     "987654321",
		Contacted: false,
	},
}
