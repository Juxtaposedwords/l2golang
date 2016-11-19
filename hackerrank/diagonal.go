package main
import "fmt"
func main() {
    a := loadSlice(1)
    b := loadMatrix(a[0])
    c := primaryDiag(b)
    d := secondaryDiag(b)
    e := absSlices(c,d)
    fmt.Printf("%z",d)
    fmt.Printf("%z",e)
}

func absSlices(a,b []int)  int{
    var abs int
    for i, _ := range(a) {
        abs += Abs(a[i])+ Abs(b[i])
    }
    return abs
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
    a := len(x[0])
    b := []int{}
    for i:=a; i > 0 ; i-- {
        a -= 1
        b = append(b, x[i][a])
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