package connection_test

import (
	"test_patterns/connection"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItems(t *testing.T) {
	t.Run("Deve retornar dados do banco de dados", func(t *testing.T) {
		var connection = connection.New()
		_, err := connection.Query("select * from ex.items")
		assert.Equal(t, nil, err)
	})
}
