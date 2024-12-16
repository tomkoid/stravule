package api

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"codeberg.org/tomkoid/stravule/backend/internal/resolvers"
	"codeberg.org/tomkoid/stravule/backend/internal/utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

func normalizeString(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, s)
	if err != nil {
		return "", err
	}

	return strings.ToLower(result), nil
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

	orderDayExceptions, err := resolvers.ListOrderDayExceptions(&userHash)
	if err != nil {
		return nil, nil, err
	}

	noOrderDays, err := resolvers.ListNoOrderDays(&userHash)
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

			orderNazev, err := normalizeString(order.Nazev)
			if err != nil {
				// log it and use the not normalized val
				log.Println(err)
				orderNazev = order.Nazev
			}

			for _, filter := range filters.Include {
				filterNazev, err := normalizeString(filter.Value)
				if err != nil {
					log.Println(err)
					filterNazev = filter.Value
				}

				if strings.Contains(orderNazev, filterNazev) {
					order.Score += filter.Weight
					order.selectedByIncludeFilter = true
				}
			}

			for _, filter := range filters.Exclude {
				filterNazev, err := normalizeString(filter.Value)
				if err != nil {
					log.Println(err)
					filterNazev = filter.Value
				}

				if strings.Contains(orderNazev, filterNazev) {
					order.Score -= filter.Weight
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
				orderingTimeCheck, err := CheckOrderingTime(resPicked[i][highestScoreIdx].CasKonec)
				if err != nil {
					return nil, nil, err
				}

				// if ordering time is expired, don't modify
				if !orderingTimeCheck {
					resPicked[i][j].Pocet = res[i][j].Pocet
					continue
				}

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

				// don't apply if order can't be picked
				if strings.HasSuffix(resPicked[i][highestScoreIdx].Omezeni, "B") {
					continue
				}

				resPicked[i][highestScoreIdx].Pocet = 1
			} else {
				resPicked[i][highestScoreIdx].Pocet = 0
			}

			// don't order on days we don't want to
			for j := 0; j < len(resPicked[i]); j++ {
				parsedDate, err := utils.ParseDate(resPicked[i][j].Datum)
				if err != nil {
					// print to console but don't interrupt the request
					log.Println(err)
				} else {
					for _, exceptionDay := range orderDayExceptions {
						if parsedDate.Weekday() == time.Weekday(exceptionDay) {
							for k := 0; k < len(resPicked[i]); k++ {
								if resPicked[i][k].Datum == resPicked[i][j].Datum {
									resPicked[i][k].Pocet = 0
								}
							}
						}
					}
					for _, noOrderDay := range noOrderDays {
						layout := "2006-01-02"
						noOrderDayParsed, err := time.Parse(layout, noOrderDay)
						if err != nil {
							// print to console but don't interrupt the request
							log.Println(err)
							continue
						}
						noOrderDayParsed = time.Date(noOrderDayParsed.Year(), noOrderDayParsed.Month(), noOrderDayParsed.Day(), 0, 0, 0, 0, time.Local)

						fmt.Println(parsedDate, noOrderDayParsed)
						if parsedDate == noOrderDayParsed {
							for k := 0; k < len(resPicked[i]); k++ {
								if resPicked[i][k].Datum == resPicked[i][j].Datum {
									resPicked[i][k].Pocet = 0
								}
							}
						}
					}
				}
			}

		}

	}

	return resPicked, res, nil
}
