# HackerRank Submission
Note here I had to tinker with the scan and blindly steal a read function.
I really like the read function though, it's short and simple. In a real application,
I wouldn't have omitted the error returns(it was really to get myself to drop them).
```
func read(n int) ([]int) {
  in := make([]int, n)
  for i := range in {
    fmt.Scan(&in[i])
  }
  return in
}
func main() {
    var queries, length int
    var inputSlice []int
    fmt.Scanf("%d", &queries)
    for i:=0; i < queries; i++{
       fmt.Scanf("%d", &length)
        inputSlice = read(length)
       fmt.Println(IsAbsSortable(inputSlice))
    }

}
```