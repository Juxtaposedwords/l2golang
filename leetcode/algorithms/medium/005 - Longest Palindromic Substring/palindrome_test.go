package palindrome

import (
	"testing"
)

var LongestPalindrome = longestPalindrome
var IsPalindromic = isPalindromic

func TestLongestPalindrome(t *testing.T) {

	tt := []struct {
		have string
		want string
	}{
		{"dabad", "dabad"},
		{"cbbd", "bb"},
		{"dabadddddddddddddddddddddddddddddddddddd", "dddddddddddddddddddddddddddddddddddd"},
		//		{"jrjnbctoqgzimtoklkxcknwmhiztomaofwwzjnhrijwkgmwwuazcowskjhitejnvtblqyepxispasrgvgzqlvrmvhxusiqqzzibcyhpnruhrgbzsmlsuacwptmzxuewnjzmwxbdzqyvsjzxiecsnkdibudtvthzlizralpaowsbakzconeuwwpsqynaxqmgngzpovauxsqgypinywwtmekzhhlzaeatbzryreuttgwfqmmpeywtvpssznkwhzuqewuqtfuflttjcxrhwexvtxjihunpywerkktbvlsyomkxuwrqqmbmzjbfytdddnkasmdyukawrzrnhdmaefzltddipcrhuchvdcoegamlfifzistnplqabtazunlelslicrkuuhosoyduhootlwsbtxautewkvnvlbtixkmxhngidxecehslqjpcdrtlqswmyghmwlttjecvbueswsixoxmymcepbmuwtzanmvujmalyghzkvtoxynyusbpzpolaplsgrunpfgdbbtvtkahqmmlbxzcfznvhxsiytlsxmmtqiudyjlnbkzvtbqdsknsrknsykqzucevgmmcoanilsyyklpbxqosoquolvytefhvozwtwcrmbnyijbammlzrgalrymyfpysbqpjwzirsfknnyseiujadovngogvptphuyzkrwgjqwdhtvgxnmxuheofplizpxijfytfabx", "qosoq"},
	}

	for _, v := range tt {
		got := longestPalindrome(v.have)
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}

}
func TestIsPalindromic(t *testing.T) {

	tt := []struct {
		have string
		want bool
	}{
		{"aba", true},
		{"dbddd", false},
	}

	for _, v := range tt {
		got := isPalindromic(v.have)
		if got != v.want {
			t.Errorf("IsPalindromic: %#v want:%b", v, got)
		}
	}
}
