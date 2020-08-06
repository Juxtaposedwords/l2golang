package parser

import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"testing"
)

func readerfy(input string, t *testing.T) io.ReadCloser {
	f, err := os.Open(input)
	if err != nil {
		t.Fatal(err)
	}
	return f

}

func TestSpellURLS(t *testing.T) {
	tests := []struct {
		desc           string
		url            string
		readCloser     io.ReadCloser
		wantResp       []*Spell
		wantStatusCode codes.Code
	}{
		{
			desc:       "single Entry",
			url:        "http://foo.com",
			readCloser: readerfy("/workspace/starscraper/src/parser/testdata/zone.html", t),
			wantResp: []*Spell{
				{
					Name:      "Zone of Truth",
					ShortDesc: "Creatures within range can’t lie.",
					URL:       "http://foo.com/SpellDisplay.aspx?ItemName=Zone%20of%20Truth&Family=None",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Parallel()
		tc := tc // to make sure TC scope is kept
		t.Run(tc.desc, func(t *testing.T) {
			defer tc.readCloser.Close()
			got, err := SpellURLs(tc.url, tc.readCloser)
			if gotErr, want := status.Code(err), tc.wantStatusCode; gotErr != want {
				t.Fatalf("SpellURLS() unexpected error. want: %s got: %s - %s", want, gotErr, err.Error())
			}
			if tc.wantStatusCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.wantResp, got); diff != "" {
				t.Errorf("SpellURLS() unexpected Response (-want +got):\n%s\n got: %#v want: %#v", diff, got, tc.wantResp)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		desc           string
		url            string
		readCloser     io.ReadCloser
		wantResp       *Spell
		wantStatusCode codes.Code
	}{
		{
			desc:       "single Entry",
			url:        "http://foo.com",
			readCloser: readerfy("/workspace/starscraper/src/parser/testdata/spell-description.html", t),
			wantResp: &Spell{
				Level:       2,
				Mystic:      true,
				Witchwarper: true,
				School:      "evo",
			},
		},
	}
	t.Parallel()
	for _, tc := range tests {
		tc := tc // to make sure TC scope is kept
		t.Run(tc.desc, func(t *testing.T) {
			defer tc.readCloser.Close()
			got := &Spell{}
			err := Update(got, tc.readCloser)
			if gotErr, want := status.Code(err), tc.wantStatusCode; gotErr != want {
				t.Fatalf("Update() unexpected error. want: %s got: %s - %s", want, gotErr, err.Error())
			}
			if tc.wantStatusCode != codes.OK {
				return
			}
			if diff := cmp.Diff(tc.wantResp, got); diff != "" {
				t.Errorf("Update() unexpected Response (-want +got):\n%s\n got: %#v want: %#v", diff, got, tc.wantResp)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		desc  string
		line  string
		spell *Spell
		want  string
	}{
		{
			desc: "single Entry",
			line: `{"_id":"0c25gVm7RKKg1nIy","name":"Contact Other Plane","permission":{"default":0,"umHu8DFnCrMQVZ37":3},"type":"spell","data":{"description":{"value":"<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">You send your mind to another plane of existence (an Elemental Plane or some plane further removed) in order to receive advice and information from powers there. The powers reply in a language you understand, but they resent such contact and give only brief answers to your questions. All questions are answered with “Yes,” “No,” “Maybe,” “Never,” “Irrelevant,” or some other one-word answer.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">You must concentrate on maintaining the spell in order to ask questions at the rate of one per round. A question is answered by the power during the same round. You can ask one question for every 2 caster levels. On rare occasions, this divination may be blocked by an act of certain deities or forces.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">Contacting a minor planar power is relatively safe but may not result in useful answers. For each question you ask, the GM secretly rolls 1d20.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>1–2</strong>: The power gives you no answer, the spell ends, and you must attempt a DC 7 Intelligence check. On a failed check, your Intelligence and Charisma scores each fall to 8 for a week and you are unable to cast spells for that period.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>3–5</strong>: You receive a random answer to the question.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>6–10</strong>: You receive an incorrect answer to the question. Based on the nature and needs of the creature contacted, this may be a lie designed to harm you.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>11–15</strong>: You receive no answer to the question.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>16 or More</strong>: You receive a truthful and useful one-word answer. If the question can’t be truthfully answered in this way, no answer is received.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">Contact with minds further removed from your home plane increases the probability that you will incur a decrease in Intelligence and Charisma due to your brain being overwhelmed by the power’s sheer strangeness and force, but it also increases the chance of the power knowing the answer and answering correctly. You can add any value from +1 to +5 to the d20 roll to represent contacting increasingly powerful planar beings. However, on a roll of 1 or 2, the result is still no answer, the spell ends, and you must attempt an Intelligence check to avoid losing Intelligence and Charisma. The DC of this Intelligence check is increased by the same amount added to the d20 check to contact a planar creature.</span></p>","chat":"","unidentified":""},"source":"Core Rulebook","type":"","activation":{"type":"min","cost":10,"condition":""},"duration":{"value":"concentration","units":""},"target":{"value":"","type":""},"area":{"value":15,"units":"ft","shape":"","effect":""},"range":{"value":null,"units":"personal","additional":"","per":""},"uses":{"value":0,"max":0,"per":""},"ability":"","actionType":"","attackBonus":0,"chatFlavor":"","critical":{"parts":[],"effect":""},"damage":{"parts":[]},"formula":"","save":{"type":"will","dc":null,"descriptor":"negate"},"descriptors":[],"level":1,"school":"div","sr":true,"dismissible":false,"concentration":false,"materials":{"value":"","consumed":false,"cost":0,"supply":0},"preparation":{"mode":"","prepared":false},"allowedClasses":{"myst":true,"tech":true,"wysh":true},"abilityMods":{"parts":[]}},"flags":{},"img":"systems/sfrpg/icons/spells/contact_other_plane.png"}`,
			spell: &Spell{
				Tags:        []string{"SFS Legal"},
				Level:       2,
				Mystic:      true,
				Witchwarper: true,
				School:      "evo",
			},
			want: `{"_id":"0c25gVm7RKKg1nIy","name":"Contact Other Plane","permission":{"default":0,"umHu8DFnCrMQVZ37":3},"type":"spell","data":{"description":{"value":"<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">You send your mind to another plane of existence (an Elemental Plane or some plane further removed) in order to receive advice and information from powers there. The powers reply in a language you understand, but they resent such contact and give only brief answers to your questions. All questions are answered with “Yes,” “No,” “Maybe,” “Never,” “Irrelevant,” or some other one-word answer.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">You must concentrate on maintaining the spell in order to ask questions at the rate of one per round. A question is answered by the power during the same round. You can ask one question for every 2 caster levels. On rare occasions, this divination may be blocked by an act of certain deities or forces.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">Contacting a minor planar power is relatively safe but may not result in useful answers. For each question you ask, the GM secretly rolls 1d20.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>1–2</strong>: The power gives you no answer, the spell ends, and you must attempt a DC 7 Intelligence check. On a failed check, your Intelligence and Charisma scores each fall to 8 for a week and you are unable to cast spells for that period.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>3–5</strong>: You receive a random answer to the question.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>6–10</strong>: You receive an incorrect answer to the question. Based on the nature and needs of the creature contacted, this may be a lie designed to harm you.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>11–15</strong>: You receive no answer to the question.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\"><strong>16 or More</strong>: You receive a truthful and useful one-word answer. If the question can’t be truthfully answered in this way, no answer is received.<br></span></p>\n<p><span id=\"ctl00_MainContent_DataListTalentsAll_ctl00_LabelName\">Contact with minds further removed from your home plane increases the probability that you will incur a decrease in Intelligence and Charisma due to your brain being overwhelmed by the power’s sheer strangeness and force, but it also increases the chance of the power knowing the answer and answering correctly. You can add any value from +1 to +5 to the d20 roll to represent contacting increasingly powerful planar beings. However, on a roll of 1 or 2, the result is still no answer, the spell ends, and you must attempt an Intelligence check to avoid losing Intelligence and Charisma. The DC of this Intelligence check is increased by the same amount added to the d20 check to contact a planar creature.</span></p>","chat":"","unidentified":""},"source":"Core Rulebook","type":"","activation":{"type":"min","cost":10,"condition":""},"duration":{"value":"concentration","units":""},"target":{"value":"","type":""},"area":{"value":15,"units":"ft","shape":"","effect":""},"range":{"value":null,"units":"personal","additional":"","per":""},"uses":{"value":0,"max":0,"per":""},"ability":"","actionType":"","attackBonus":0,"chatFlavor":"","critical":{"parts":[],"effect":""},"damage":{"parts":[]},"formula":"","save":{"type":"will","dc":null,"descriptor":"negate"},"descriptors":[],"level":2,"school":"evo","sr":true,"dismissible":false,"concentration":false,"materials":{"value":"","consumed":false,"cost":0,"supply":0},"preparation":{"mode":"","prepared":false},"allowedClasses":{"myst":true,"tech":false,"wysh":true},"abilityMods":{"parts":[]}},"flags":{},"img":"systems/sfrpg/icons/spells/contact_other_plane.png"}`,
		},
	}
	for _, tc := range tests {
		t.Parallel()
		tc := tc // to make sure TC scope is kept
		t.Run(tc.desc, func(t *testing.T) {
			got := Merge(tc.spell, tc.line)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Merge() unexpected Response (-want +got):\n%s\n got: %#v want: %#v", diff, got, tc.want)
			}
		})
	}
}
