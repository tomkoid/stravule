package api

import (
	"fmt"
	"strings"
	"time"

	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
)

// returns a list of orders, used filters and possibly an error
func PickOrders(sid string, canteen string, userHash string) ([][]order, error) {
	res, err := Orders(sid, canteen)
	if err != nil {
		return nil, err
	}

	filters, err := resolvers.GetFilters(&userHash)

	if err != nil {
		return nil, err
	}

	// score and compare orders by filters within each order table
	for i := range res {
		for j := range res[i] {
			order := &res[i][j]

			if strings.HasSuffix(order.Omezeni, "B") {
				continue
			}

			timeKonec, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", order.CasKonec))
			if err != nil {
				return nil, err
			}

			if !timeKonec.After(time.Now()) {
				continue
			}

			for _, filter := range filters.Include {
				if strings.Contains(strings.ToLower(order.Nazev), strings.ToLower(filter)) {
					order.Score += 1
					order.selectedByIncludeFilter = true
				}
			}

			for _, filter := range filters.Exclude {
				if strings.Contains(strings.ToLower(order.Nazev), strings.ToLower(filter)) {
					order.Score -= 1
					order.selectedByExcludeFilter = true
				}
			}

			order.Nazev = fmt.Sprintf("%s (S: %d)", order.Nazev, order.Score)
		}
	}

	for i := range res {
		if len(res[i]) > 1 {
			highestScoreIdx := 0

			// rank
			for j := 1; j < len(res[i]); j++ {
				if res[i][j].Score > res[i][highestScoreIdx].Score {
					highestScoreIdx = j
				}
			}

			// set amount to 0 for all other orders
			for j := 0; j < len(res[i]); j++ {
				if j != highestScoreIdx {
					res[i][j].Pocet = 0
				}
			}

			// order the order if there is a positive score and it is the highest score in the orderTable
			if res[i][highestScoreIdx].Score >= 0 {
				res[i][highestScoreIdx].Pocet = 1
			} else {
				res[i][highestScoreIdx].Pocet = 0
			}
		}
	}

	return res, nil
}
