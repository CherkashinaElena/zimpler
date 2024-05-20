package zimpler_test

import (
	"github.com/stretchr/testify/require"
	"testing"
	"zimpler/internal/zimpler"
)

func TestCustomerStats(t *testing.T) {
	tests := []struct {
		name   string
		input  []*zimpler.Customer
		output []*zimpler.CustomerStats
	}{
		{
			name: "should collect customer stats with large data",
			input: []*zimpler.Customer{
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 200,
				},
				{
					Name:  "Jonas",
					Candy: "Kexchoklad",
					Eaten: 100,
				},
				{
					Name:  "Aadya",
					Candy: "Nötchoklad",
					Eaten: 2,
				},
				{

					Name:  "Jonas",
					Candy: "Nötchoklad",
					Eaten: 3,
				},
				{
					Name:  "Jane",
					Candy: "Nötchoklad",
					Eaten: 17,
				},
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},

				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 700,
				},
				{
					Name:  "Jane",
					Candy: "Nötchoklad",
					Eaten: 4,
				},
				{
					Name:  "Aadya",
					Candy: "Center",
					Eaten: 7,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 900,
				},
				{
					Name:  "Jane",
					Candy: "Nötchoklad",
					Eaten: 1,
				},
				{
					Name:  "Jonas",
					Candy: "Kexchoklad",
					Eaten: 12,
				},
				{
					Name:  "Jonas",
					Candy: "Plopp",
					Eaten: 40,
				},
				{
					Name:  "Jonas",
					Candy: "Center",
					Eaten: 27,
				},
				{
					Name:  "Aadya",
					Candy: "Center",
					Eaten: 2,
				},
				{
					Name:  "Annika",
					Candy: "Center",
					Eaten: 8,
				},
			},
			output: []*zimpler.CustomerStats{
				{
					Name:      "Jonas",
					Total:     1982,
					Favourite: "Geisha",
				},
				{
					Name:      "Annika",
					Total:     208,
					Favourite: "Geisha",
				},
				{
					Name:      "Jane",
					Total:     22,
					Favourite: "Nötchoklad",
				},
				{
					Name:      "Aadya",
					Total:     11,
					Favourite: "Center",
				},
			},
		},
		{
			name: "should collect customer stats with few data",
			input: []*zimpler.Customer{
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 200,
				},
				{
					Name:  "Aadya",
					Candy: "Nötchoklad",
					Eaten: 2,
				},
			},
			output: []*zimpler.CustomerStats{
				{
					Name:      "Jonas",
					Total:     200,
					Favourite: "Geisha",
				},
				{
					Name:      "Annika",
					Total:     100,
					Favourite: "Geisha",
				},
				{
					Name:      "Aadya",
					Total:     2,
					Favourite: "Nötchoklad",
				},
			},
		},
		{
			name: "should collect customer stats with one data",
			input: []*zimpler.Customer{
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},
			},
			output: []*zimpler.CustomerStats{
				{
					Name:      "Annika",
					Total:     100,
					Favourite: "Geisha",
				},
			},
		},
		{
			name: "should collect customer stats with same amount of eaten",
			input: []*zimpler.Customer{
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 100,
				},
				{
					Name:  "Aadya",
					Candy: "Geisha",
					Eaten: 100,
				},
			},
			output: []*zimpler.CustomerStats{
				{
					Name:      "Annika",
					Total:     100,
					Favourite: "Geisha",
				},
				{
					Name:      "Jonas",
					Total:     100,
					Favourite: "Geisha",
				},
				{
					Name:      "Aadya",
					Total:     100,
					Favourite: "Geisha",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// given

			// when
			stats := zimpler.CollectCustomerStats(test.input)

			// then
			require.Equal(t, test.output, stats)
		})
	}
}
