package diag
import "fmt"
func diag() {
    a := diagDif()
    fmt.Printf("%d", a)
}
func diagDif() int {
    a := readInt()
    b := loadMatrix(a)
    c := sumSlice(primaryDiag(b))
    d := sumSlice(secondaryDiag(b))
    e := Abs(c - d)
    return e
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
    a := len(x[0]) -1
    b := []int{}
    for i:=0; i <= a; i++ {
        b = append(b, x[i][i])
    }
    return b
}
func secondaryDiag(x [][]int) []int{
    a := len(x[0])
    b := a
    c := []int{}
    for i:=0; i < b ; i++ {
        a -= 1
        c = append(c, x[i][a])
    }
    return c
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
        a = append(a, readInt())
    }
    return a
}

func readInt() int{
    var b int
    fmt.Scanf("%d", &b)
    return b
}