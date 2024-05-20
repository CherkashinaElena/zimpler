package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
	"zimpler/internal/zimpler"
)

const url = "https://candystore.zimpler.net/"

func TestE2e(t *testing.T) {
	// given
	f := zimpler.NewHttpCustomerFetcher()

	//when
	customers, err := f.Fetch(url)
	require.NoError(t, err)

	stats := zimpler.CollectCustomerStats(customers)

	// then
	expected, err := fetchExpected()
	require.NoError(t, err)

	actual, err := json.Marshal(stats)
	require.NoError(t, err)

	fmt.Println(string(actual))

	require.JSONEq(t, string(expected), string(actual))
}

func fetchExpected() ([]byte, error) {
	file, err := os.Open("./expected.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return io.ReadAll(file)
}
