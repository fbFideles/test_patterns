package itemsrepository

type ItemsRepository interface {
	GetItems() (items []Output, err error)
}
