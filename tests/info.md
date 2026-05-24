## Filter tests by time/name

```
func Sum(nums ...int) int {    // func to test
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func TestSumFew(t *testing.T) {   //short test
    if Sum(1, 2, 3, 4, 5) != 15 {
        t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
    }
}


func TestSumN(t *testing.T) {  // long test
    n := 1_000_000_000
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    got := Sum(nums...)
    want := n * (n + 1) / 2
    if got != want {
        t.Errorf("Expected sum[i=1..n](i) == n*(n+1)/2")
    }
}
```
### v1: filter short tests (testing.Short(), t.Skip(reason) and -short in cmd)

```
func TestSumN(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    // сам тест
}
```

`$ go test -v -short`




//--------------------------------------
### v2: filter test by name mask:
```
$ go test -v -run N
``` 
(only SumN)

```
$ go test -v -run Few
``` 
(only SumFew)

//-------------------------------------------------------

## Test Coverage:
`$ go test -cover`
(%)

Построчный: создать файл cover.prof
`$ go test -coverprofile="cover.prof"`
(создать отчет)
`go tool cover -html="cover.prof"`
(посмотреть, по каким строчкам прошли тесты)
