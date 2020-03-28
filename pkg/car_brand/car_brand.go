package car_brand

// type List map[string]string
//
// func (l List) Add(id string, brand string) List {
// 	if _, ok := l[id]; ok {
// 		return l
// 	}
// 	l[id] = brand
// 	return l
// }

type Brand struct {
	ID string
	Name string
}

type Chan chan Brand
