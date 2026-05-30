package main

import (
	"encoding/json"
	"fmt"
)

// начало решения

// Genre описывает жанр фильма
type Genre string

func (g *Genre) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*g = Genre(temp.Name)
	return nil
}

/*
// (solution)
func (g *Genre) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        return nil
    }
    var obj map[string]string
    if err := json.Unmarshal(data, &obj); err != nil {
        return err
    }
    if val, ok := obj["name"]; ok {
        *g = Genre(val)
    }
    return nil
}
*/

// Movie описывает фильм
type Movie struct {
	Title  string  `json:"name"`
	Year   int     `json:"released_at"`
	Genres []Genre `json:"tags"`
}

// конец решения

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
