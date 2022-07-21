package itemsrepository

import (
	"test_patterns/connection"
)

func NewRepositoryDatabase() ItemsRepository {
	return &itemsRepositoryDatabase{
		dbConnection: connection.New(),
	}
}

type itemsRepositoryDatabase struct {
	dbConnection *connection.Connection
}

func (i *itemsRepositoryDatabase) GetItems() (items []Output, err error) {
	rows, err := i.dbConnection.Query("select description, price from ex.items")
	if err != nil {
		return
	}

	var item Output
	for rows.Next() {
		if err = rows.Scan(&item.Description, &item.Price); err != nil {
			return
		}
		items = append(items, item)
	}

	return
}

type Output struct {
	Description string
	Price       string
}
