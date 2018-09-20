package ingredient

type Ingredient struct {
	Name      string
	Type      string
	Size      int
	Precision int
	Values    []string
}

func (ing *Ingredient) Add(value string) {
	for _, i := range ing.Values {
		if i == value {
			return
		}
	}

	ing.Values = append(ing.Values, value)
}

func(ing *Ingredient) Remove(value string) {
	ing.
}