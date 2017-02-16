package kitchen

// Holds buffers of ingredients
type Storage struct {
	Ingredients map[Item] chan bool

	Burger chan bool
}

func NewStorage() Storage {
	s := Storage{}
	s.Ingredients = make(map[Item] chan bool)

	for _, item := range Items {
		s.Ingredients[item] = make(chan bool, 100)
	}

	s.Burger = make(chan bool, 100)

	return s
}

func (s Storage) Get(item Item) chan bool {
	return s.Ingredients[item]
}
