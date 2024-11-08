package resolvers

type Filters struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

func GetFilters(sid string, canteen string) Filters {
	var filterInclude []string = []string{"bramborové knedlíky", "vrabec"}
	var filterExclude []string = []string{"kuskus", "brambory", "bramborová kaše", "okurka", "játra", "rýže"}

	return Filters{Include: filterInclude, Exclude: filterExclude}
}
