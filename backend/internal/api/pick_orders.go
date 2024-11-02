package api

func PickOrders(sid string, canteen string) ([][]order, error) {
	res, err := Orders(sid, canteen)
	if err != nil {
		return nil, err
	}

	for i := range res {
		for j := range res[i] {
			res[i][j].Pocet = 1
		}
	}

	return res, nil
}
