package resolvers

import "errors"

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

func GetFilters(sid *string, canteen *string) Filters {
	return filtersGlob
}

func AddFilter(sid string, canteen string, filter Filter) (*Filters, error) {
	// if filter already exists, do nothing
	for _, f := range filtersGlob.Include {
		if f == filter.Value {
			return nil, errors.New("filter already exists in includeFilters, please remove it first")
		}
	}
	for _, f := range filtersGlob.Exclude {
		if f == filter.Value {
			return nil, errors.New("filter already exists in excludeFilters, please remove it first")
		}
	}

	if filter.Category == "include" {
		filtersGlob.Include = append(filtersGlob.Include, filter.Value)
	} else if filter.Category == "exclude" {
		filtersGlob.Exclude = append(filtersGlob.Exclude, filter.Value)
	} else {
		return nil, errors.New("invalid filter category")
	}

	finalFilters := GetFilters(&sid, &canteen)
	return &finalFilters, nil
}

func RemoveFilter(sid string, canteen string, filter Filter) (*Filters, error) {
	removed := false

	if filter.Category == "include" {
		for i := 0; i < len(filtersGlob.Include); i++ {
			if filtersGlob.Include[i] == filter.Value {
				filtersGlob.Include = append(filtersGlob.Include[:i], filtersGlob.Include[i+1:]...)
			}
		}
	} else if filter.Category == "exclude" {
		for i := 0; i < len(filtersGlob.Exclude); i++ {
			if filtersGlob.Exclude[i] == filter.Value {
				filtersGlob.Exclude = append(filtersGlob.Exclude[:i], filtersGlob.Exclude[i+1:]...)
			}
		}
	} else {
		return nil, errors.New("invalid filter category")
	}

	if !removed {
		return nil, errors.New("filter not found")
	}

	finalFilters := GetFilters(&sid, &canteen)
	return &finalFilters, nil
}
