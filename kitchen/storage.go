package kitchen

// Holds buffers of ingredients
type Storage struct {
	Node

	Ingredients map[Item] chan bool

	Meals map[Item] chan bool
}

func NewStorage() *Storage {
	s := Storage{}
	s.Node = Node{NodeTypeStorage, "storage"}

	s.Ingredients = make(map[Item] chan bool)

	for _, item := range Items {
		s.Ingredients[item] = make(chan bool, 100)
	}

	s.Meals = make(map[Item] chan bool)
	s.Meals[Burger] = make(chan bool)

	return &s
}

func (s Storage) GetMeal(item Item) chan bool {
	return s.Meals[item]
}

func (s Storage) GetIngredient(item Item) chan bool {
	return s.Ingredients[item]
}
