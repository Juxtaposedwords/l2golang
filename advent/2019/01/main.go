package main
import (
    "bufio"
    "fmt"
    "io"
    "strconv"
    "strings"
)
func main() {
}

func readFuel( )([]int, error) {
	f, err := os.Open("/tmp/dat")
	return readInts(bufio.NewReader(f))

}
func readInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}