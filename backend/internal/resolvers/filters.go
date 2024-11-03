package resolvers

func GetFilters() ([]string, []string) {
	var filterInclude []string = []string{"bramborové knedlíky", "vrabec"}
	var filterExclude []string = []string{"kuskus", "brambory", "bramborová kaše", "okurka", "játra", "rýže"}

	return filterInclude, filterExclude
}
