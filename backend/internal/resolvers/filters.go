package resolvers

import (
	"context"
	"errors"
	"regexp"

	"codeberg.org/tomkoid/stravule/backend/db"
	"codeberg.org/tomkoid/stravule/backend/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Filter struct {
	Category string `json:"category"`
	Value    string `json:"value"`
}

type Filters struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

var filtersGlob = Filters{Include: []string{
	"bramborové knedlíky",
	"vrabec",
}, Exclude: []string{
	"kuskus",
	"brambory",
	"bramborová kaže",
	"okurka",
	"játra",
	"rýže",
},
}

func GetFilters(userHash *string) (*Filters, error) {
	dbFilters, err := database.DB.ListFilters(context.Background(), *userHash)

	if err != nil {
		return nil, err
	}

	var filters Filters = Filters{Include: []string{}, Exclude: []string{}}

	for _, filter := range dbFilters {
		if filter.Category.String == "include" {
			filters.Include = append(filters.Include, filter.FilterText)
		} else if filter.Category.String == "exclude" {
			filters.Exclude = append(filters.Exclude, filter.FilterText)
		}
	}

	return &filters, nil
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
		if f == filter.Value {
			return nil, errors.New("filter already exists in includeFilters, please remove it first")
		}
	}
	for _, f := range finalFilters.Exclude {
		if f == filter.Value {
			return nil, errors.New("filter already exists in excludeFilters, please remove it first")
		}
	}

	if filter.Category == "include" {
		database.DB.AddFilter(context.Background(), db.AddFilterParams{
			FilterText: filter.Value,
			Category: pgtype.Text{
				String: "include",
				Valid:  true,
			},
			UserHash: *userHash,
		})
		finalFilters.Include = append(finalFilters.Include, filter.Value)
	} else if filter.Category == "exclude" {
		database.DB.AddFilter(context.Background(), db.AddFilterParams{
			FilterText: filter.Value,
			Category: pgtype.Text{
				String: "exclude",
				Valid:  true,
			},
			UserHash: *userHash,
		})
		finalFilters.Exclude = append(finalFilters.Exclude, filter.Value)
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
		if finalFilters.Include[i] == filter.Value {
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
		if finalFilters.Exclude[i] == filter.Value {
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
		return nil, errors.New("filter not found anywhere")
	}

	return finalFilters, nil
}

func isFilterValid(filter *string) bool {
	// Regular expression for alphabetic characters, including Czech diacritics
	var re = regexp.MustCompile(`^[a-zA-ZáčďéěíňóřšťůúýžÁČĎÉĚÍŇÓŘŠŤŮÚÝŽ ]+$`)

	return re.MatchString(*filter)
}
