import "fmt"
var mappy = map[string]int{
"M":1000,
"CM": 900,
"D": 500,
"CD": 400,
"C": 100,
"XC": 90,
"L": 50,
"XL": 40,
"X": 10,
"IX": 9,
"V": 5,
"IV": 4,
"I": 1,}
	
// Ther are a few solutions to this problem, i've opted for the simplest
//   1.  Most Performant: o:1 
//          With 3000+ test, the most performant solution is to make a map of 
//              every possible roman numeral
//   2.  Simplest(this version): o:N
//       Create a mapof all the string mappings, with the edge cases. Then 
//          iterate over the slice, looking ahead and compare 2 rune before 1 rune
//          note: this most closely resembles the process most people use for 
//          this problem
//      
//    3. Truest to problem:
//      Turn the string into the unicode representation of each actual roman 
//        numeral and then use a map with those
//
// I see these as mapping to 1: Engineer, 2: developer, 3: academic
func romanToInt(s string) int {
    o, err := roman(s)
    if err != nil {
        fmt.Println(err)
        return 0
    }
    return o
    
}

func roman(s string)(o int, err error){
	r := []rune(s)
	for i := 0; i < len(r); i++{
		double, tc := mappy[string(r[i:i+2])]
		single, sc := mappy[string(r[i])]
		switch{
			case tc:
				o+=double
				i++
			case sc:
				o+=single
			default:
				return o, fmt.Errorf("An invalid character was entered")
		}
	}
	return o, nil
}