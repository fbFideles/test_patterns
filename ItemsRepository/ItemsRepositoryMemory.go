package itemsrepository

func NewRepositoryMemory() ItemsRepository {
	return &itemsRepositoryMemory{
		[]Output{
			{
				Price:       "30.99",
				Description: "Chainsaw Man",
			},
		},
	}
}

type itemsRepositoryMemory struct {
	Items []Output
}

func (i *itemsRepositoryMemory) GetItems() (items []Output, err error) {
	return i.Items, nil
}
