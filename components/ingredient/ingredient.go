package ingredient

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/mitchellh/mapstructure"
)

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

func (ing *Ingredient) Remove(value string) {
	// TODO User must be able to remove a value from an ingredient
}

// NewFromReader returns an instance to a new Ingredient from the
// specified name of the ingredient from a Reader.
//
// Will return nil if none is found or if there is an error along the process.
func NewFromReader(rw io.Reader, name string) (*Ingredient, error) {
	// Reader into bytes[]
	data, err := ioutil.ReadAll(rw)

	if err != nil {
		return nil, err
	}

	// Create an empty interface to unmarshal JSON into
	var j interface{}

	err = json.Unmarshal(data, &j)

	if err != nil {
		return nil, err
	}

	// Use type assertion of string interface map to step through JSON
	mp, ok := j.(map[string]interface{})

	if !ok {
		return nil, err
	}

	// Get the ingredients as an interface array
	ings, ok := mp["ingredients"].([]interface{})

	if !ok {
		return nil, err
	}

	// Create ingredient variable
	var ing Ingredient

	// Loop through the ingredients to find the one with the corresponding name
	for _, v := range ings {
		vv, _ := v.(map[string]interface{})

		if vv["name"] == name {
			mapstructure.Decode(v, &ing)
			return &ing, err
		}
	}

	// None found, return nil
	return nil, err
}
