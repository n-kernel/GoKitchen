package kitchen

// Holds buffers of ingredients
type storage struct {
	Ingredients map[item] chan bool

	Burger chan bool
}

func NewStorage() storage {
	s := storage{}
	s.Ingredients = make(map[item] chan bool)

	for _, item := range Items {
		s.Ingredients[item] = make(chan bool, 100)
	}

	s.Burger = make(chan bool, 100)

	return s
}

func (s storage) Get(item item) chan bool {
	return s.Ingredients[item]
}
