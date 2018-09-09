package palindrome

import (
	"testing"
)

var LongestPalindrome = longestPalindrome
var GetPalindrome = getPalindrome

func TestLongestPalindrome(t *testing.T) {

	tt := []struct {
		have string
		want string
	}{
		{"dabad", "dabad"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"bb", "bb"},
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

func TestGetPalindrome(t *testing.T) {

	type h struct {
		text  string
		index int
	}
	tt := []struct {
		have *h
		want string
	}{
		{&h{"asaddaa", 0}, "a"},
		{&h{"adaddadd", 1}, "ada"},
		{&h{"adaddadd", 3}, "daddad"},
		{&h{"adad", 3}, "d"},
		{&h{"dddddd", 1}, "dddd"},
		{&h{"a", 0}, "a"},
		{&h{"bb", 0}, "bb"},
	}

	for _, v := range tt {
		got := getPalindrome(v.have.text, v.have.index)
		if got != v.want {
			t.Errorf("getPalindrome:  have:%#v want:%s got: %s", v.have, v.want, got)
		}
	}
}
