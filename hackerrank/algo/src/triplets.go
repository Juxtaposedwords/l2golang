package triplet
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    a := loadSlice(3)
    b := loadSlice(3)
    c,d := compareScores(a,b)
    fmt.Printf("%d %d", c,d )
}
func compareScores(a, b []int) (c,d int){
    for i, _ := range(a){
        switch{ 
            case a[i] > b[i]:
                c += 1
            case a[i] < b[i]:
                d += 1a
        }        
    }
    return c, d
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