package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	rounded := time.Duration(d).Round(time.Minute)

	hours := int(rounded / time.Hour)
	minutes := int((rounded % time.Hour) / time.Minute)

	var result string

	if minutes == 0 {
		result = fmt.Sprintf("%dh", hours)
	} else if hours == 0 {
		result = fmt.Sprintf("%dm", minutes)
	} else {
		result = fmt.Sprintf("%dh%dm", hours, minutes)
	}

	return json.Marshal(result)
}

/* (solution)
func (d Duration) MarshalJSON() ([]byte, error) {
    b := new(bytes.Buffer)
    b.WriteByte('"')
    dur := int(time.Duration(d).Minutes())
    hours := dur / 60
    if hours != 0 {
        fmt.Fprintf(b, "%dh", hours)
    }
    mins := dur % 60
    if mins != 0 {
        fmt.Fprintf(b, "%dm", mins)
    }
    b.WriteByte('"')
    return b.Bytes(), nil
}
*/

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	result := "\""
	i := 0
	for ; i < int(r); i++ {
		result += "★"
	}
	for ; i < 5; i++ {
		result += "☆"
	}
	result += "\""
	return []byte(result), nil
}

/* (solution)
func (r Rating) MarshalJSON() ([]byte, error) {
    b := new(bytes.Buffer)
    b.WriteByte('"')
    for i := 0; i < int(r); i++ {
        b.WriteRune('★')
    }
    for i := 0; i < 5-int(r); i++ {
        b.WriteRune('☆')
    }
    b.WriteByte('"')
    return b.Bytes(), nil
}
*/

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	var err error
	var result []byte

	if indent > 0 {
		spaces := strings.Repeat(" ", indent)
		result, err = json.MarshalIndent(movies, "", spaces)
	} else {
		result, err = json.Marshal(movies)
	}
	if err != nil {
		fmt.Printf("Error while marshalling: %v", err)
		return "", err
	}
	return string(result), nil
}

// конец решения
func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
