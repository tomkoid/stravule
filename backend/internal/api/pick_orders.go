package api

import (
	"strings"

	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
)

// returns a list of orders, used filters and possibly an error
func PickOrders(sid string, canteen string) ([][]order, error) {
	res, err := Orders(sid, canteen)
	if err != nil {
		return nil, err
	}

	filterInclude, filterExclude := resolvers.GetFilters()

	// compare orders by filters within each order table
	for i := range res {
		somethingAlreadyRemoved := false
		for j := range res[i] {
			order := &res[i][j]

			// if the previous thing was removed, set the amount of this to 1
			if somethingAlreadyRemoved {
				if order.Pocet == 0 {
					order.Pocet = 1
					somethingAlreadyRemoved = false // reset
				}
			}

			for _, filter := range filterInclude {
				if strings.Contains(order.Nazev, filter) {
					if j >= 1 && res[i][j-1].selectedByIncludeFilter {
						order.Pocet = 0
					} else {
						order.Pocet = 1
					}

					if j >= 1 && res[i][j-1].Pocet == 1 && !res[i][j-1].selectedByIncludeFilter {
						res[i][j-1].Pocet = 0
					}

					order.selectedByIncludeFilter = true
				}
			}

			for _, filter := range filterExclude {
				if strings.Contains(order.Nazev, filter) {
					order.Pocet = 0
					somethingAlreadyRemoved = true
				}
			}
		}
	}

	return res, nil
}
