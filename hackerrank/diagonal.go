package main
import "fmt"
func main() {
    a := loadSlice(1)
    b := loadMatrix(a[0])
    c := sumSlice(primaryDiag(b))
    d := sumSlice(secondaryDiag(b))
    e := Abs(Abs(c) - Abs(d))
    fmt.Printf("%d", e)
}


func sumSlice(a []int)  int{
    var sum int
    for _, i := range(a) {
        sum += i
    }
    return sum
}
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func primaryDiag(x [][]int) []int{
    a := len(x[0])
    b := []int{}
    for i:=0; i < a; i++ {
        b = append(b, x[i][i])
    }
    return b
}
func secondaryDiag(x [][]int) []int{
    a := len(x[0]) - 1
    b := []int{}
    for i:=0; i < 3  ; i++ {
        b = append(b, x[i][a])
        a -= 1
    }
    return b
}
func loadMatrix(x int)[][]int{
    a := [][]int{}
    for i:=0; i<x; i ++{
        b := loadSlice(x)
        a = append(a,b)
    }
    return a
}
func loadSlice(x int) []int{
    a := []int{}
    for i:=0; i<x; i++{
        var b int
        fmt.Scanf("%d", &b)
        a = append(a, b)
    }
    return a
}