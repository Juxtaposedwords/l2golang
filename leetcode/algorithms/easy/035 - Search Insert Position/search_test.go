package search

import (
	"testing"
)

var SearchInsert = searchInsert

func TestSearchInsert(t *testing.T) {
	type h struct {
		slice  []int
		target int
	}
	tt := []struct {
		have *h
		want int
	}{
		{&h{[]int{1, 3, 5, 6}, 5}, 2},
		{&h{[]int{1, 3, 5, 6}, 2}, 1},
		{&h{[]int{1, 3, 5, 6}, 7}, 4},
		{&h{[]int{1, 3, 5, 6}, 0}, 0},
		{&h{[]int{1, 3}, 1}, 0},
		{&h{[]int{1, 3, 5}, 4}, 2},
		//		{"jrjnbctoqgzimtoklkxcknwmhiztomaofwwzjnhrijwkgmwwuazcowskjhitejnvtblqyepxispasrgvgzqlvrmvhxusiqqzzibcyhpnruhrgbzsmlsuacwptmzxuewnjzmwxbdzqyvsjzxiecsnkdibudtvthzlizralpaowsbakzconeuwwpsqynaxqmgngzpovauxsqgypinywwtmekzhhlzaeatbzryreuttgwfqmmpeywtvpssznkwhzuqewuqtfuflttjcxrhwexvtxjihunpywerkktbvlsyomkxuwrqqmbmzjbfytdddnkasmdyukawrzrnhdmaefzltddipcrhuchvdcoegamlfifzistnplqabtazunlelslicrkuuhosoyduhootlwsbtxautewkvnvlbtixkmxhngidxecehslqjpcdrtlqswmyghmwlttjecvbueswsixoxmymcepbmuwtzanmvujmalyghzkvtoxynyusbpzpolaplsgrunpfgdbbtvtkahqmmlbxzcfznvhxsiytlsxmmtqiudyjlnbkzvtbqdsknsrknsykqzucevgmmcoanilsyyklpbxqosoquolvytefhvozwtwcrmbnyijbammlzrgalrymyfpysbqpjwzirsfknnyseiujadovngogvptphuyzkrwgjqwdhtvgxnmxuheofplizpxijfytfabx", "qosoq"},
	}

	for _, v := range tt {
		got := searchInsert(v.have.slice, v.have.target)
		if got != v.want {
			t.Errorf("SearchInsert failed with\n  have: %#v\n  want: %d\n  got:  %#v\n", v.have, v.want, got)
		}
	}

}
