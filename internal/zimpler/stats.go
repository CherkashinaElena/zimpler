package zimpler

import (
	"sort"
)

// CustomerStats describes customer candy statistics
type CustomerStats struct {
	Name      string `json:"name"`
	Total     int    `json:"totalSnacks"`
	Favourite string `json:"favouriteSnack"`
}

// CollectCustomerStats returns customer candy statistic sorted by amount of eaten candies.
func CollectCustomerStats(customers []*Customer) []*CustomerStats {
	stats := make(map[string]candyStats)
	for _, customer := range customers {
		if candyStats, ok := stats[customer.Name]; ok {
			if _, ok := candyStats[customer.Candy]; ok {
				candyStats[customer.Candy] += customer.Eaten
			} else {
				candyStats[customer.Candy] = customer.Eaten
			}
		} else {
			stats[customer.Name] = map[string]int{
				customer.Candy: customer.Eaten,
			}
		}
	}

	var customerStats []*CustomerStats
	for name, candyStats := range stats {
		customerStats = append(customerStats, &CustomerStats{
			Name:      name,
			Total:     candyStats.total(),
			Favourite: candyStats.favourite(),
		})
	}

	sort.Slice(customerStats, func(i, j int) bool {
		return customerStats[i].Total > customerStats[j].Total
	})

	return customerStats
}

type candyStats map[string]int

func (s candyStats) total() int {
	var sum int
	for _, eaten := range s {
		sum += eaten
	}

	return sum
}

func (s candyStats) favourite() string {
	var favourite string

	var count int
	for candy, eaten := range s {
		if eaten > count {
			count = eaten
			favourite = candy
		}
	}

	return favourite
}
