package getitems_test

import (
	getitems "test_patterns/GetItems"
	itemsrepository "test_patterns/ItemsRepository"
	"testing"

	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetItems(t *testing.T) {
	t.Run("Should get items", func(t *testing.T) {
		itemsRepository := itemsrepository.NewRepositoryDatabase()
		getItems := getitems.New(itemsRepository)
		items, err := getItems.Execute()
		require.Nil(t, err)

		assert.Equal(t, items[0].Description, "Chainsaw Man")
		assert.Equal(t, items[0].Price, 30.99)
	})

	t.Run("Should get items from a fake repository", func(t *testing.T) {
		itemsRepository := itemsrepository.NewRepositoryMemory()
		getItems := getitems.New(itemsRepository)
		items, err := getItems.Execute()
		require.Nil(t, err)

		assert.Equal(t, items[0].Description, "Chainsaw Man")
		assert.Equal(t, items[0].Price, 30.99)
	})

	t.Run("Should get items from a stub repository", func(t *testing.T) {
		itemsRepository := itemsrepository.NewRepositoryMemory()

		getItemsFunc := itemsRepository.GetItems
		stubs := gostub.StubFunc(&getItemsFunc, []itemsrepository.Output{{Price: "30.99", Description: "Chainsaw Man"}}, nil)
		defer stubs.Reset()

		getItems := getitems.New(itemsRepository)
		items, err := getItems.Execute()
		require.Nil(t, err)

		assert.Equal(t, items[0].Description, "Chainsaw Man")
		assert.Equal(t, items[0].Price, 30.99)
	})

	t.Run("Should get items from a spy repository", func(t *testing.T) {
		// wich is also a spy
		itemsRepository := itemsrepository.NewRepositoryMock()
		itemsRepository.On("GetItems").
			Return([]itemsrepository.Output{{Price: "30.99", Description: "Chainsaw Man"}}, nil)

		getItems := getitems.New(itemsRepository)
		items, err := getItems.Execute()
		require.Nil(t, err)

		itemsRepository.AssertNumberOfCalls(t, "GetItems", 1)

		assert.Equal(t, items[0].Description, "Chainsaw Man")
		assert.Equal(t, items[0].Price, 30.99)
	})

	t.Run("Should get items from a mock repository", func(t *testing.T) {
		itemsRepository := itemsrepository.NewRepositoryMock()
		itemsRepository.On("GetItems").
			Return([]itemsrepository.Output{{Price: "30.99", Description: "Chainsaw Man"}}, nil)

		getItems := getitems.New(itemsRepository)
		items, err := getItems.Execute()
		require.Nil(t, err)

		assert.Equal(t, items[0].Description, "Chainsaw Man")
		assert.Equal(t, items[0].Price, 30.99)
	})
}
