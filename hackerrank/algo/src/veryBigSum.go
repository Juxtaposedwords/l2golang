package main
import "fmt"

func main() {
    a := loadSlice(1)
    b := loadSlice(a[0])
    e := sumSlice(b)
    fmt.Printf("%d",e)
}

func sumSlice(a []int)  int{
    var sum int
    for _, i := range(a) {
        sum += i
    }
    return sum
}
func loadSlice(x int) []int{
    var a []int
    for i:=0; i<x; i++{
        var b int
        fmt.Scanf("%d", &b)
        a = append(a, b)
    }
    return a
}