package races

import (
	"fmt"
	log "github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readerfy(input string, t *testing.T) io.ReadCloser {
	f, err := os.Open(input)
	if err != nil {
		t.Fatal(err)
	}
	return f

}
func TestListRaces(t *testing.T) {
	log.Init("dontlookatme", true, false, ioutil.Discard)

	tests := []struct {
		desc           string
		url            string
		readCloser     io.ReadCloser
		wantResp       []*RaceTuple
		wantStatusCode codes.Code
	}{
		{
			desc:       "single Entry",
			url:        "http://foo.com",
			readCloser: readerfy("/workspace/starscraper/src/parser/testdata/races.html", t),
			wantResp: []*RaceTuple{{URL: "http://foo.com/?ItemName=Android", Name: "Android"},
				{URL: "http://foo.com/?ItemName=Human", Name: "Human"},
				{URL: "http://foo.com/?ItemName=Kasatha", Name: "Kasatha"},
				{URL: "http://foo.com/?ItemName=Lashunta", Name: "Lashunta"},
				{URL: "http://foo.com/?ItemName=Shirren", Name: "Shirren"},
				{URL: "http://foo.com/?ItemName=Vesk", Name: "Vesk"},
				{URL: "http://foo.com/?ItemName=Ysoki", Name: "Ysoki"},
				{URL: "http://foo.com/?ItemName=Aasimar", Name: "Aasimar"},
				{URL: "http://foo.com/?ItemName=Amrantah", Name: "Amrantah"},
				{URL: "http://foo.com/?ItemName=Anassanoi", Name: "Anassanoi"},
				{URL: "http://foo.com/?ItemName=Astrazoan", Name: "Astrazoan"},
				{URL: "http://foo.com/?ItemName=Bantrid", Name: "Bantrid"},
				{
					URL:  "http://foo.com/?ItemName=Barathu (Early Stage)",
					Name: "Barathu (Early Stage)",
				},
				{URL: "http://foo.com/?ItemName=Bear (Uplifted)", Name: "Bear (Uplifted)"},
				{URL: "http://foo.com/?ItemName=Bolida", Name: "Bolida"},
				{URL: "http://foo.com/?ItemName=Borai", Name: "Borai"},
				{URL: "http://foo.com/?ItemName=Brakim", Name: "Brakim"},
				{URL: "http://foo.com/?ItemName=Brenneri", Name: "Brenneri"},
				{URL: "http://foo.com/?ItemName=Cephalume", Name: "Cephalume"},
				{URL: "http://foo.com/?ItemName=Contemplative", Name: "Contemplative"},
				{URL: "http://foo.com/?ItemName=Copaxi", Name: "Copaxi"},
				{URL: "http://foo.com/?ItemName=Damai", Name: "Damai"},
				{URL: "http://foo.com/?ItemName=Dirindi", Name: "Dirindi"},
				{URL: "http://foo.com/?ItemName=Draelik", Name: "Draelik"},
				{URL: "http://foo.com/?ItemName=Dragonkin", Name: "Dragonkin"},
				{URL: "http://foo.com/?ItemName=Dromada", Name: "Dromada"},
				{URL: "http://foo.com/?ItemName=Drow", Name: "Drow"},
				{URL: "http://foo.com/?ItemName=Dwarf", Name: "Dwarf"},
				{URL: "http://foo.com/?ItemName=Elebrian", Name: "Elebrian"},
				{URL: "http://foo.com/?ItemName=Elf", Name: "Elf"},
				{URL: "http://foo.com/?ItemName=Embri", Name: "Embri"},
				{URL: "http://foo.com/?ItemName=Endiffian", Name: "Endiffian"},
				{URL: "http://foo.com/?ItemName=Espraksa", Name: "Espraksa"},
				{URL: "http://foo.com/?ItemName=Ferran", Name: "Ferran"},
				{URL: "http://foo.com/?ItemName=Formian", Name: "Formian"},
				{URL: "http://foo.com/?ItemName=Ghibrani", Name: "Ghibrani"},
				{URL: "http://foo.com/?ItemName=Ghoran", Name: "Ghoran"},
				{URL: "http://foo.com/?ItemName=Gnome", Name: "Gnome"},
				{URL: "http://foo.com/?ItemName=Goblin (Space)", Name: "Goblin (Space)"},
				{URL: "http://foo.com/?ItemName=Gosclaw", Name: "Gosclaw"},
				{URL: "http://foo.com/?ItemName=Gray", Name: "Gray"},
				{URL: "http://foo.com/?ItemName=Haan", Name: "Haan"},
				{URL: "http://foo.com/?ItemName=Half-Elf", Name: "Half-Elf"},
				{URL: "http://foo.com/?ItemName=Halfling", Name: "Halfling"},
				{URL: "http://foo.com/?ItemName=Half-Orc", Name: "Half-Orc"},
				{URL: "http://foo.com/?ItemName=Hanakan", Name: "Hanakan"},
				{URL: "http://foo.com/?ItemName=Hobgoblin", Name: "Hobgoblin"},
				{URL: "http://foo.com/?ItemName=Hortus", Name: "Hortus"},
				{URL: "http://foo.com/?ItemName=Ifrit", Name: "Ifrit"},
				{URL: "http://foo.com/?ItemName=Ijtikri", Name: "Ijtikri"},
				{URL: "http://foo.com/?ItemName=Ikeshti", Name: "Ikeshti"},
				{URL: "http://foo.com/?ItemName=Ilthisarian", Name: "Ilthisarian"},
				{URL: "http://foo.com/?ItemName=Imago", Name: "Imago"},
				{URL: "http://foo.com/?ItemName=Instar", Name: "Instar"},
				{URL: "http://foo.com/?ItemName=Izalguun", Name: "Izalguun"},
				{URL: "http://foo.com/?ItemName=Kalo", Name: "Kalo"},
				{URL: "http://foo.com/?ItemName=Kanabo", Name: "Kanabo"},
				{URL: "http://foo.com/?ItemName=Kayal", Name: "Kayal"},
				{URL: "http://foo.com/?ItemName=Khizar", Name: "Khizar"},
				{URL: "http://foo.com/?ItemName=Kiirinta", Name: "Kiirinta"},
				{URL: "http://foo.com/?ItemName=Kish", Name: "Kish"},
				{URL: "http://foo.com/?ItemName=Maraquoi", Name: "Maraquoi"},
				{URL: "http://foo.com/?ItemName=Megalonyxa", Name: "Megalonyxa"},
				{URL: "http://foo.com/?ItemName=Morlamaw", Name: "Morlamaw"},
				{URL: "http://foo.com/?ItemName=Neskinti", Name: "Neskinti"},
				{URL: "http://foo.com/?ItemName=Nuar", Name: "Nuar"},
				{URL: "http://foo.com/?ItemName=Orc", Name: "Orc"},
				{URL: "http://foo.com/?ItemName=Oread", Name: "Oread"},
				{URL: "http://foo.com/?ItemName=Osharu", Name: "Osharu"},
				{URL: "http://foo.com/?ItemName=Pahtra", Name: "Pahtra"},
				{URL: "http://foo.com/?ItemName=Phentomite", Name: "Phentomite"},
				{URL: "http://foo.com/?ItemName=Quorlu", Name: "Quorlu"},
				{URL: "http://foo.com/?ItemName=Ramiyel", Name: "Ramiyel"},
				{URL: "http://foo.com/?ItemName=Raxilite", Name: "Raxilite"},
				{URL: "http://foo.com/?ItemName=Reptoid", Name: "Reptoid"},
				{URL: "http://foo.com/?ItemName=Ryphorian", Name: "Ryphorian"},
				{URL: "http://foo.com/?ItemName=Sarcesian", Name: "Sarcesian"},
				{URL: "http://foo.com/?ItemName=Sazaron", Name: "Sazaron"},
				{URL: "http://foo.com/?ItemName=Screedreep", Name: "Screedreep"},
				{URL: "http://foo.com/?ItemName=Scyphozoan", Name: "Scyphozoan"},
				{URL: "http://foo.com/?ItemName=Selamid", Name: "Selamid"},
				{URL: "http://foo.com/?ItemName=Seprevoi", Name: "Seprevoi"},
				{URL: "http://foo.com/?ItemName=Shakatla", Name: "Shakatla"},
				{URL: "http://foo.com/?ItemName=Shatori", Name: "Shatori"},
				{URL: "http://foo.com/?ItemName=Shimreen", Name: "Shimreen"},
				{URL: "http://foo.com/?ItemName=Shobhad", Name: "Shobhad"},
				{URL: "http://foo.com/?ItemName=Skittermander", Name: "Skittermander"},
				{URL: "http://foo.com/?ItemName=Spathinae", Name: "Spathinae"},
				{URL: "http://foo.com/?ItemName=SRO", Name: "SRO"},
				{URL: "http://foo.com/?ItemName=Stellifera", Name: "Stellifera"},
				{URL: "http://foo.com/?ItemName=Strix", Name: "Strix"},
				{URL: "http://foo.com/?ItemName=Suli", Name: "Suli"},
				{URL: "http://foo.com/?ItemName=Svartalfar", Name: "Svartalfar"},
				{URL: "http://foo.com/?ItemName=Sylph", Name: "Sylph"},
				{URL: "http://foo.com/?ItemName=Telia", Name: "Telia"},
				{URL: "http://foo.com/?ItemName=Tiefling", Name: "Tiefling"},
				{URL: "http://foo.com/?ItemName=Trinir", Name: "Trinir"},
				{URL: "http://foo.com/?ItemName=Trox", Name: "Trox"},
				{URL: "http://foo.com/?ItemName=Undine", Name: "Undine"},
				{URL: "http://foo.com/?ItemName=Urog", Name: "Urog"},
				{URL: "http://foo.com/?ItemName=Varculak", Name: "Varculak"},
				{URL: "http://foo.com/?ItemName=Verthani", Name: "Verthani"},
				{URL: "http://foo.com/?ItemName=Vilderaro", Name: "Vilderaro"},
				{URL: "http://foo.com/?ItemName=Vlaka", Name: "Vlaka"},
				{URL: "http://foo.com/?ItemName=Witchwyrd", Name: "Witchwyrd"},
				{URL: "http://foo.com/?ItemName=Woioko", Name: "Woioko"},
				{URL: "http://foo.com/?ItemName=Wrikreechee", Name: "Wrikreechee"},
			},
		},
	}
	for _, tc := range tests {
		t.Parallel()
		tc := tc // to make sure TC scope is kept
		t.Run(tc.desc, func(t *testing.T) {
			defer tc.readCloser.Close()
			got, err := ListRaces(tc.url, tc.readCloser)
			if gotErr, want := status.Code(err), tc.wantStatusCode; gotErr != want {
				t.Fatalf("ListRaces() unexpected error. want: %s got: %s - %s", want, gotErr, err.Error())
			}
			if tc.wantStatusCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.wantResp, got); diff != "" {
				t.Errorf("ListRaces() unexpected Response (-want +got):\n%s\n got: %#v want: %#v", diff, got, tc.wantResp)
			}
		})
	}
}
func TestParseRace(t *testing.T) {

	tests := []struct {
		desc           string
		readCloser     io.ReadCloser
		wantResp       *Race
		wantStatusCode codes.Code
	}{
		{
			desc:       "single Entry",
			readCloser: readerfy("/workspace/starscraper/src/parser/testdata/race_Astrazoan.html", t),
			wantResp: &Race{
				Source: "Pact Worlds pg. 210",
				Mods: []*Mod{
					{"dex", 2},
					{"cha", 2},
					{"con", -2},
				},
				Size:    "Medium",
				Type:    "aberration",
				SubType: "shapechanger",
				HP:      4,
			},
		},
		{
			desc:       "dwarf (Existing entry)",
			readCloser: readerfy("/workspace/starscraper/src/parser/testdata/race_dwarf.html", t),
			wantResp: &Race{
				Source: "Starfinder Core Rulebook pg. 506",
				Mods: []*Mod{
					{"con", 2},
					{"wis", 2},
					{"cha", -2},
				},
				Size:    "Medium",
				Type:    "humanoid",
				SubType: "dwarf",
				HP:      6,
			},
		},
	}
	for _, tc := range tests {
		//t.Parallel()
		tc := tc // to make sure TC scope is kept
		t.Run(tc.desc, func(t *testing.T) {
			log.Init("dontlookatme", true, false, ioutil.Discard)
			defer tc.readCloser.Close()
			got, err := ParseRace(tc.readCloser)
			if gotErr, want := status.Code(err), tc.wantStatusCode; gotErr != want {
				t.Fatalf("ListRaces() unexpected error. want: %s got: %s - %s", want, gotErr, err.Error())
			}
			if tc.wantStatusCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.wantResp, got); diff != "" {
				t.Errorf("ListRaces() unexpected Response (-want +got):\n%s\n got: %#v want: %#v", diff, got, tc.wantResp)
			}
		})
	}
}
