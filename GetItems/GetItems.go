package getitems

import (
	"strconv"
	itemsrepository "test_patterns/ItemsRepository"
)

func New(repo itemsrepository.ItemsRepository) *getItems {
	return &getItems{
		itemsRepository: repo,
	}
}

type getItems struct {
	itemsRepository itemsrepository.ItemsRepository
}

type Item struct {
	Price       float64
	Description string
}

func (g *getItems) Execute() (output []Item, err error) {
	items, err := g.itemsRepository.GetItems()
	if err != nil {
		return
	}
	return g.convertItems(items)
}

func (g *getItems) convertItems(repoItems []itemsrepository.Output) (items []Item, err error) {
	items = make([]Item, len(repoItems))
	for i := range repoItems {
		price, err := strconv.ParseFloat(repoItems[i].Price, 64)
		if err != nil {
			return nil, err
		}
		items[i].Price = price
		items[i].Description = repoItems[i].Description
	}
	return
}
