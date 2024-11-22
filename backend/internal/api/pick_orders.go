package api

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
)

func CheckOrderingTime(orderingEndTime string) (bool, error) {
	timeKonec, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", orderingEndTime))
	if err != nil {
		return false, err
	}

	if !timeKonec.After(time.Now()) {
		return false, nil
	}

	return true, nil
}

func deepCopyOrders(src [][]order) [][]order {
	dest := make([][]order, len(src))
	for i := range src {
		dest[i] = make([]order, len(src[i]))
		copy(dest[i], src[i])
	}
	return dest
}

// returns a list of picked orders, originally used orders and used filters and possibly an error
func PickOrders(sid string, canteen string, userHash string) ([][]order, [][]order, error) {
	res, err := Orders(sid, canteen)
	if err != nil {
		return nil, nil, err
	}

	var resPicked [][]order = deepCopyOrders(res)

	filters, err := resolvers.GetFilters(&userHash)

	if err != nil {
		return nil, nil, err
	}

	// score and compare orders by filters within each order table
	for i := range resPicked {
		for j := range resPicked[i] {
			order := &resPicked[i][j]

			if strings.HasSuffix(order.Omezeni, "B") {
				continue
			}

			orderingTimeCheck, err := CheckOrderingTime(order.CasKonec)
			if err != nil {
				return nil, nil, err
			}

			if !orderingTimeCheck {
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
					order.Score -= 2
					order.selectedByExcludeFilter = true
				}
			}

			order.Nazev = fmt.Sprintf("%s (S: %d)", order.Nazev, order.Score)
		}
	}

	for i := range resPicked {
		if len(resPicked[i]) > 1 {
			highestScoreIdx := 0

			// rank
			for j := 1; j < len(resPicked[i]); j++ {
				orderingTimeCheck, err := CheckOrderingTime(resPicked[i][j].CasKonec)
				if err != nil {
					return nil, nil, err
				}

				if !orderingTimeCheck {
					continue
				}

				if resPicked[i][j].Score > resPicked[i][highestScoreIdx].Score {
					highestScoreIdx = j
				}
			}

			// set amount to 0 for all other orders
			for j := 0; j < len(resPicked[i]); j++ {
				if j != highestScoreIdx {
					resPicked[i][j].Pocet = 0
				}
			}

			// order the order if there is a positive score and it is the highest score in the orderTable
			if resPicked[i][highestScoreIdx].Score >= 0 {
				orderingTimeCheck, err := CheckOrderingTime(resPicked[i][highestScoreIdx].CasKonec)
				if err != nil {
					return nil, nil, err
				}

				if !orderingTimeCheck {
					continue
				}

				resPicked[i][highestScoreIdx].Pocet = 1
			} else {
				resPicked[i][highestScoreIdx].Pocet = 0
			}
		}
	}

	fmt.Println(reflect.DeepEqual(res, resPicked))
	return resPicked, res, nil
}
