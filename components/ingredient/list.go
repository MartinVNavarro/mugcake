package ingredient

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// List contains a list of Ingredients
type List struct {
	Ingredients []Ingredient
}

// NewList returns new empty Ingredients List
func NewList() *List {
	return &List{}
}

// ListFromReader will load a create an IngredientList from a reader
// that reads ingredients from a JSON source
func ListFromReader(rw io.Reader) (*List, error) {
	data, err := ioutil.ReadAll(rw)

	if err != nil {
		return nil, err
	}

	var il List

	err = json.Unmarshal(data, &il)

	if err != nil {
		return nil, err
	}

	return &il, err
}

// WriteToFile writes an ingredient list to a file
func (il *List) WriteToFile(w io.Writer) error {
	var data []byte

	data, err := json.Marshal(il)

	if err != nil {
		return err
	}

	_, err = w.Write(data)

	if err != nil {
		return err
	}

	return nil
}

// Get an ingredient from an ingredient list by name
func (il *List) Get(name string) *Ingredient {
	if i := il.IndexOf(name); i >= 0 {
		return &il.Ingredients[i]
	}

	return nil
}

// Add an ingredient to the list if it does not exist
// No action if it currently exists in Ingredients
func (il *List) Add(ing Ingredient) {
	// return if it exists already
	if i := il.IndexOf(ing.Name); i >= 0 {
		return
	}

	il.Ingredients = append(il.Ingredients, ing)
}

// Remove an ingredient from a List by name
// No action if it does not exist
func (il *List) Remove(name string) {
	if i := il.IndexOf(name); i >= 0 {
		il.Ingredients = append(il.Ingredients[:i], il.Ingredients[i+1:]...)
	}
}

// IndexOf returns the index of an ingredient in a list
// or returns -1 if not found
func (il *List) IndexOf(name string) int {
	for i, v := range il.Ingredients {
		if v.Name == name {
			return i
		}
	}

	return -1
}
