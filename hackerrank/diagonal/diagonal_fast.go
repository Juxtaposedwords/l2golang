package main
import "fmt"

func main(){
    a := readInt() 
    var b,c int
    for i:=0; i < a ; i++ {
        for j:=0; j < a ; j++ {
            x := readInt()
            switch {
                case (j == a-i-1 ) && (j == i):
                    b += x
                    c += x
                case j == i:
                    b += x
                case j == a-i-1:
                    c += x
                default:
                    _ = x
            }
        }
    }
    fmt.Printf("%d", Abs(b-c))
}
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func readInt() int{
    var b int
    fmt.Scanf("%d", &b)
    return b
}