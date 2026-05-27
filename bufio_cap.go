package main

// v1: my

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	result := []string{}
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		result = append(result, string(runes))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(result, " "))
}

//-------------------------------------------------------------------
// // v2: strings.Title (deprecated):
// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     r := bufio.NewReader(os.Stdin)
//     s, _ := r.ReadString('\n')
//     fmt.Print(strings.Title(strings.ToLower(s)))
// }

//-------------------------------------------------------------------

// // v3: cases.Title
// package main

// import (
//     "fmt"

//     "golang.org/x/text/cases"
//     "golang.org/x/text/language"
// )

// func main() {
//     s := "Привет, ромашки"
//     titled := cases.Title(language.Russian).String(s)
//     fmt.Print(titled)
//     // Привет, Ромашки
// }
