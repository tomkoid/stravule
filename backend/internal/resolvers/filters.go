package resolvers

func GetFilters() ([]string, []string) {
	var filterInclude []string = []string{"rýže", "bramborové knedlíky", "vrabec"}
	var filterExclude []string = []string{"kuskus", "brambory", "bramborová kaše", "okurka", "játra"}

	return filterInclude, filterExclude
}
