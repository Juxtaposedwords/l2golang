package diag
import (
    "fmt"   
    "io"
//    "strings"
) 
func diagValue(r io.Reader)(int,error) {
    var err error
    a, err := readInt(r) 
    if err != nil {
        return 0, err
    }
    var b,c int
    for i:=0; i < a ; i++ {
        for j:=0; j < a ; j++ {
            x, err := readInt(r)
            if err != nil {
                return 0, err
            }
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
    return Abs(b-c), nil
}
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func readInt(r io.Reader) (int, error){
    var i int
    _, err := fmt.Fscan(r,&i)
    return i,err
}	