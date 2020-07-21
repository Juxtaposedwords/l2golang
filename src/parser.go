// Package parser takes stdios and attempts to collect certain information from those blocks.
package parser

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

type Spell struct {
	URL          string
	Tags         []string
	ShortDesc    string
	Name         string
	Level        int
	School       string
	Technomancer bool
	Mystic       bool
	Witchwarper  bool
}

// SpellURLs takes the body of a website and atttempts to return all spell entries found within that site.
func SpellURLs(baseURL string, input io.Reader) ([]*Spell, error) {
	log.Init("dontlookatme", true, false, ioutil.Discard)

	resp, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s", err)
	}
	var spells []*Spell
	spansRE := regexp.MustCompile(`<span.*?<\/span?>`)
	idRE := regexp.MustCompile(`href=".*?"`)
	shortRE := regexp.MustCompile(`<\/a><\/b>: .*?</span`)
	for _, entry := range spansRE.FindAllString(string(resp), -1) {
		if !strings.Contains(entry, "SpellDisplay") || ! strings.Contains(entry,"Family=None") {
			continue
		}

		spellQuery := strings.TrimSuffix(strings.TrimPrefix(idRE.FindString(entry), `href="`), `"`)
		shortDesc := strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(shortRE.FindString(entry), `</a></b>: `), "</span"), "<br />")
		spellName := strings.Split(strings.TrimPrefix(spellQuery, "SpellDisplay.aspx?ItemName="), "&")[0]

		url := fmt.Sprintf("%s/%s", baseURL, strings.Replace(spellQuery, " ", "%20", -1))
		var tags []string
		spellEntry := &Spell{
			URL:       url,
			ShortDesc: shortDesc,
			Name:      spellName,
			Tags:      tags,
		}
		spells = append(spells, spellEntry)
	}
	return spells, nil
}

// Update fparses a given buffer and provides any details provided by the buffer to the provided spell struct.
func Update(spell *Spell, input io.Reader) error {
	if spell == nil {
		return status.Error(codes.FailedPrecondition, "must have a spell to operate upon")
	}
	resp, err := ioutil.ReadAll(input)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "%s", err)
	}
	level, err := getLevel(string(resp))
	if err != nil {
		return status.Errorf(codes.InvalidArgument,"failed to update %s with %s ",spell,err)
	}
	spell.Level = level
	mystic, technomancer, witchwarper := classes(string(resp))
	spell.Mystic = mystic
	spell.Technomancer = technomancer
	spell.Witchwarper = witchwarper
	spell.School = getSchool(string(resp))
//	log.Infof("%#v",spell)
	return nil
}

// Merge takes a spell and updates a line entry from the stafinder packdb
func Merge(spell *Spell, line string) string {
	//"allowedClasses":{"myst":true,"tech":true,"wysh":true}
	log.Infof("%#v", spell)
	classStanza := regexp.MustCompile(`"allowedClasses":.*?}`).FindString(line)
	newClassStanza := fmt.Sprintf(`"allowedClasses":{"myst":%t,"tech":%t,"wysh":%t}`,spell.Mystic,spell.Technomancer,spell.Witchwarper)
	newLine := strings.Replace(line, classStanza,newClassStanza, 1)
	if spell.Level != 0{
		levelStanza := regexp.MustCompile(`"level":[0-9]`).FindString(line)
		newLevelStanza := fmt.Sprintf(`"level":%d`,spell.Level)
		newLine = strings.Replace(newLine, levelStanza,newLevelStanza,1)
	}
	schoolStanza := regexp.MustCompile(`"school":.*?,`).FindString(line)
	newSchoolStanza := fmt.Sprintf(`"school":"%s",`, spell.School)
	newLine = strings.Replace(newLine, schoolStanza,newSchoolStanza,1)

	return newLine
}
func getSchool(input string) string {
	schoolRE := regexp.MustCompile("<b>School</b> .*?<br />")
	shortMap  := map[string]string{
		"abjuration":    "abj",
		"conjuration":   "con",
		"divination":    "div",
		"enchantment":   "enc",
		"evocation":     "evo",
		"illusion":      "ill",
		"necromancy":    "nec",
		"transmutation": "trs",
	}
	fullName := strings.Split(strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(schoolRE.FindString(input), "<b>School</b> "), "<br />")), " ")[0]
	val, ok :=  shortMap[fullName]
	if ! ok {
		log.Errorf("%s\n", fullName)
		return fullName
	}
	return val
}

func getLevel(input string) (int, error) {
	levelRE := regexp.MustCompile("(Witchwarper|Mystic|Technomancer) [0-9]")
	first := levelRE.FindString(input)
	parts := strings.Split(first, " ")
	if len(parts) != 2 {
		return 0, status.Errorf(codes.InvalidArgument, "%#v doesn't have 2 parts", parts)
	}
	resp, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	return resp, nil
}
func classes(input string) (bool, bool, bool) {
	return regexp.MustCompile("Mystic [0-9]").MatchString(input), regexp.MustCompile("Technomancer [0-9]").MatchString(input), regexp.MustCompile("Witchwarper [0-9]").MatchString(input)
}
