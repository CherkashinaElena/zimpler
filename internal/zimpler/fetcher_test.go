package zimpler_test

import (
	"github.com/stretchr/testify/require"
	"testing"
	"zimpler/internal/zimpler"
)

func TestFileCustomerFetcher(t *testing.T) {
	f := zimpler.NewFileCustomerFetcher()

	t.Run("should fetch customers from HTML file", func(t *testing.T) {
		// given

		// when
		customers, err := f.Fetch("./test/valid.html")

		// then
		require.NoError(t, err)
		require.Len(t, customers, 17)
	})

	t.Run("customers should be empty if no file", func(t *testing.T) {
		// given

		// when
		_, err := f.Fetch("./test/no_file.html")

		// then
		require.Error(t, err)
	})

	t.Run("customers should be empty if no table", func(t *testing.T) {
		// given

		// when
		customers, err := f.Fetch("./test/no_table.html")

		// then
		require.NoError(t, err)
		require.Empty(t, customers)
	})

	t.Run("customers should be empty if no customers", func(t *testing.T) {
		// given

		// when
		customers, err := f.Fetch("./test/no_customers.html")

		// then
		require.NoError(t, err)
		require.Empty(t, customers)
	})

	t.Run("should skip customers with invalid data", func(t *testing.T) {
		// given

		// when
		customers, err := f.Fetch("./test/no_customers.html")

		// then
		require.NoError(t, err)
		require.Empty(t, customers)
	})

	t.Run("should fetch customers from HTML file", func(t *testing.T) {
		// given

		// when
		customers, err := f.Fetch("./test/invalid_data.html")

		// then
		require.NoError(t, err)
		require.Len(t, customers, 2)
	})
}

func TestHttpCustomerFetcher(t *testing.T) {
	f := zimpler.NewHttpCustomerFetcher()

	t.Run("should fetch customers from web page", func(t *testing.T) {
		// given

		// when
		customers, err := f.Fetch("https://candystore.zimpler.net/")

		// then
		require.NoError(t, err)
		require.Len(t, customers, 17)
	})

	t.Run("customers should be empty if not found", func(t *testing.T) {
		// given

		// when
		_, err := f.Fetch("https://candystore.invalid.net/")

		// then
		require.Error(t, err)
	})
}
