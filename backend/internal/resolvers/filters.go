package resolvers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"

	"codeberg.org/tomkoid/stravule/backend/db"
	"codeberg.org/tomkoid/stravule/backend/internal/database"
	"codeberg.org/tomkoid/stravule/backend/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type FilterReq struct {
	Category string `json:"category"`
	Value    string `json:"value"`
}

type Filter struct {
	Category  string `json:"category"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	Weight    int    `json:"weight"`
}

// func (f FilterReq) ToFilter() Filter {
// 	flt := Filter{
// 		Category:  f.Category,
// 		Value:     f.Value,
// 		CreatedAt: filter.CreatedAt.Time.String(),
// 		Weight:    int(filter.Weight),
// 	}
// 	return flt
// }

type Filters struct {
	Include []Filter `json:"include"`
	Exclude []Filter `json:"exclude"`
}

func GetFilters(userHash *string) (*Filters, error) {
	dbFilters, err := database.DB.ListFilters(context.Background(), *userHash)

	if err != nil {
		return nil, err
	}

	var filters Filters = Filters{Include: make([]Filter, 0), Exclude: make([]Filter, 0)}

	for _, filter := range dbFilters {
		flt := Filter{
			Category:  filter.Category.String,
			Value:     filter.FilterText,
			CreatedAt: filter.CreatedAt.Time.String(),
			Weight:    int(filter.Weight),
		}
		if filter.Category.String == "include" {
			filters.Include = append(filters.Include, flt)
		} else if filter.Category.String == "exclude" {
			filters.Exclude = append(filters.Exclude, flt)
		}
	}

	filters.Include = sortFilters(filters.Include)
	filters.Exclude = sortFilters(filters.Exclude)

	log.Println(filters.Include)

	return &filters, nil
}

func sortFilters(filterType []Filter) []Filter {
	layout := "2006-01-02 15:04:05 -0700 MST"
	sort.Slice(filterType, func(i, j int) bool {
		filter1time, err := time.Parse(layout, utils.NormalizeDateString(filterType[i].CreatedAt))
		if err != nil {
			log.Println(err)
		}

		filter2time, err := time.Parse(layout, utils.NormalizeDateString(filterType[j].CreatedAt))
		if err != nil {
			log.Println(err)
		}

		return filter1time.Before(filter2time)
	})

	return filterType
}

func AddFilter(userHash *string, filter Filter) (*Filters, error) {
	if !isFilterValid(&filter.Value) {
		return nil, errors.New("invalid filter value")
	}

	finalFilters, err := GetFilters(userHash)

	if err != nil {
		return nil, err
	}

	// if filter already exists, do nothing
	for _, f := range finalFilters.Include {
		if f.Value == filter.Value {
			return nil, errors.New("Filter už existuje v kategorii chci.")
		}
	}
	for _, f := range finalFilters.Exclude {
		if f.Value == filter.Value {
			return nil, errors.New("Filter už existuje v kategorii nechci.")
		}
	}

	if filter.Category == "include" {
		database.DB.AddFilter(context.Background(), db.AddFilterParams{
			FilterText: filter.Value,
			Category: pgtype.Text{
				String: "include",
				Valid:  true,
			},
			Weight:   int32(filter.Weight),
			UserHash: *userHash,
		})
		finalFilters.Include = append(finalFilters.Include, filter)
	} else if filter.Category == "exclude" {
		database.DB.AddFilter(context.Background(), db.AddFilterParams{
			FilterText: filter.Value,
			Category: pgtype.Text{
				String: "exclude",
				Valid:  true,
			},
			Weight:   int32(filter.Weight),
			UserHash: *userHash,
		})
		finalFilters.Exclude = append(finalFilters.Exclude, filter)
	} else {
		return nil, errors.New("invalid filter category")
	}

	return finalFilters, nil
}

func RemoveFilter(userHash *string, filter Filter) (*Filters, error) {
	if !isFilterValid(&filter.Value) {
		return nil, errors.New("invalid filter value")
	}

	removed := false
	finalFilters, err := GetFilters(userHash)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(finalFilters.Include); i++ {
		if finalFilters.Include[i].Value == filter.Value {
			database.DB.DeleteFilter(context.Background(), db.DeleteFilterParams{
				FilterText: filter.Value,
				Category: pgtype.Text{
					String: "include",
					Valid:  true,
				},
				UserHash: *userHash,
			})
			finalFilters.Include = append(finalFilters.Include[:i], finalFilters.Include[i+1:]...)
			removed = true
		}
	}
	for i := 0; i < len(finalFilters.Exclude); i++ {
		if finalFilters.Exclude[i].Value == filter.Value {
			database.DB.DeleteFilter(context.Background(), db.DeleteFilterParams{
				FilterText: filter.Value,
				Category: pgtype.Text{
					String: "exclude",
					Valid:  true,
				},
				UserHash: *userHash,
			})
			finalFilters.Exclude = append(finalFilters.Exclude[:i], finalFilters.Exclude[i+1:]...)
			removed = true
		}
	}

	if !removed {
		return nil, errors.New("Filter nikde nenalezen.")
	}

	return finalFilters, nil
}

func UpdateFilterWeight(userHash *string, filter Filter) error {
	if !isFilterValid(&filter.Value) {
		return errors.New("invalid filter value")
	}

	filters, err := database.DB.ListFilters(context.Background(), *userHash)

	if err != nil {
		return errors.New(fmt.Sprintf("filters not found: %v", err))
	}

	for _, f := range filters {
		if f.FilterText == filter.Value {
			database.DB.UpdateFilterWeight(context.Background(), db.UpdateFilterWeightParams{
				FilterText: f.FilterText,
				Category: pgtype.Text{
					String: f.Category.String,
					Valid:  true,
				},
				Weight:   int32(filter.Weight),
				UserHash: *userHash,
			})
			return nil
		}
	}

	return errors.New("Filter nenalezen.")
}

func isFilterValid(filter *string) bool {
	if strings.TrimSpace(*filter) == "" {
		return false
	}

	// Regular expression for alphabetic characters, including Czech diacritics
	var re = regexp.MustCompile(`^[a-zA-ZáčďéěíňóřšťůúýžÁČĎÉĚÍŇÓŘŠŤŮÚÝŽ ]+$`)

	return re.MatchString(*filter)
}
