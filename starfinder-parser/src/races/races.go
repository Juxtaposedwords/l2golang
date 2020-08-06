package races

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Race struct {
	Name        string
	Description string
	Source      string
	Type        string
	Mods        []*Mod
	Size        string
	SubType     string
	HP          int
}
type RaceTuple struct {
	URL  string
	Name string
}
type Mod struct {
	Stat  string
	Value int
}

func ListRaces(baseURL string, input io.Reader) ([]*RaceTuple, error) {
	f, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	raceRE := regexp.MustCompile(`href="Races\.aspx\?ItemName=.*?<\/a`)
	var output []*RaceTuple
	for _, entry := range raceRE.FindAllString(string(f), -1) {
		url := strings.TrimSuffix(regexp.MustCompile(`\?.*"`).FindString(entry), `"`)

		output = append(output, &RaceTuple{
			Name: strings.TrimPrefix(strings.TrimSuffix(regexp.MustCompile(">.*<").FindString(entry), "<"), ">"),
			URL:  fmt.Sprintf("%s%s", baseURL, url),
		})

	}
	return output, nil
}
func ParseRace(input io.Reader) (*Race, error) {
	bytesContent, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	content := string(bytesContent)

	mods, err := raceModifiers(content)
	if err != nil {
		return nil, err
	}
	output := &Race{
		Mods:   mods,
		Source: raceSource(content),
	}
	if err := sizeTypes(content, output); err != nil {
		return nil, err
	}
	rawHP := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(regexp.MustCompile(`Hit Points.*?<h`).FindString(content), "Hit Points</b>"), "<h"))
	hp, err := strconv.Atoi(rawHP)
	if err != nil {
		return nil, err
	}
	output.HP = hp
	output.Description = raceDesc(content)
	return output, nil
}

func raceDesc(input string) string {
	desc := regexp.MustCompile(`<b>Source<\/b>.*?(<h1|<\/td)`).FindString(input)
	desc = strings.Replace(desc, `target=_blank`, "", -1)
	desc = strings.Replace(desc, `class="title"`, "", -1)
	desc = regexp.MustCompile(`^.*?<br />`).ReplaceAllString(desc, "")
	desc = strings.TrimSuffix(strings.TrimSuffix(fmt.Sprintf("<p>%s", desc), "<h1"), "</td")
	return desc
}
func raceSource(input string) string {
	return strings.TrimSuffix(strings.TrimPrefix(regexp.MustCompile(`<i>.*<\/i>`).FindString(regexp.MustCompile(`<b>Source<\/b>.*?</i>`).FindString(input)), "<i>"), "</i>")
}

func raceModifiers(input string) ([]*Mod, error) {
	line := strings.TrimSuffix(regexp.MustCompile(`^.*?>`).ReplaceAllString(regexp.MustCompile(`Ability Modifiers.*?<br`).FindString(input), ""), "<br")
	var output []*Mod
	for _, entry := range strings.Split(line, ",") {
		parts := strings.Split(strings.TrimSpace(entry), " ")
		//	log.Infof(" output: %#v parts: %#v line: %q base: %q", output, parts, line, input)
		strVals := map[string]int{
			"+2": 2,
			"-2": -2,
			"+1": 1,
			"-1": -1,
			"-4": -4,
		}
		output = append(output, &Mod{
			Stat:  strings.ToLower(parts[1]),
			Value: strVals[parts[0]],
		})
	}
	return output, nil
}
