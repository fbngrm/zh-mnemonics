package pinyin

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable()

	tests := []struct {
		pinyin   string
		expected Word
	}{
		{
			"qu",
			Word{
				Initial:        "qtip (qu-)",
				Final:          "space station (-Ø)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"zong",
			Word{
				Initial:        "steve gomeZ (zu-)",
				Final:          "charles house (-eng)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"sui",
			Word{
				Initial: "saul goodman (su-)",
				Final:   "subway station (-ei)",
			},
		},
		{
			"a",
			Word{
				Initial: "ghost (Ø-)",
				Final:   "anna’s flat (-a)",
			},
		},
		{
			"ai",
			Word{
				Initial: "ghost (Ø-)",
				Final:   "iceland hut (-ai)",
			},
		},
		{
			"an",
			Word{
				Initial: "ghost (Ø-)",
				Final:   "hanoi flat (-an)",
			},
		},
		{
			"ang",
			Word{
				Initial: "ghost (Ø-)",
				Final:   "angkor flat (-ang)",
			},
		},
		{
			"ao",
			Word{
				Initial: "ghost (Ø-)",
				Final:   "autobahn stop (-ao)",
			},
		},
		{
			"ca",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "anna’s flat (-a)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"cai",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "iceland hut (-ai)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"can",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "hanoi flat (-an)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"cang",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "angkor flat (-ang)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"cao",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "autobahn stop (-ao)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"ce",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "1st flat berlin (-e)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"cen",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "endelich’s flat (-en)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"ceng",
			Word{
				Initial:        "cosmonaut (c-)",
				Final:          "charles house (-eng)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"cha",
			Word{
				Initial:        "(jesus) christ (ch-)",
				Final:          "anna’s flat (-a)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"chai",
			Word{
				Initial:        "(jesus) christ (ch-)",
				Final:          "iceland hut (-ai)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"chan",
			Word{
				Initial:        "(jesus) christ (ch-)",
				Final:          "hanoi flat (-an)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"chang",
			Word{
				Initial:        "(jesus) christ (ch-)",
				Final:          "angkor flat (-ang)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"xuan",
			Word{
				Initial:        "xhibit (xu-)",
				Final:          "hanoi flat (-an)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"xue",
			Word{
				Initial:        "xhibit (xu-)",
				Final:          "1st flat berlin (-e)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"xun",
			Word{
				Initial:        "xhibit (xu-)",
				Final:          "endelich’s flat (-en)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"yong",
			Word{
				Initial: "yung hurn (yu-)",
				Final:   "charles house (-eng)",
			},
		},
		{
			"yu",
			Word{
				Initial: "yung hurn (yu-)",
				Final:   "space station (-Ø)",
			},
		},
		{
			"yuan",
			Word{
				Initial: "yung hurn (yu-)",
				Final:   "hanoi flat (-an)",
			},
		},
		{
			"zhuo",
			Word{
				Initial:        "zafiro añejo (zhu-)",
				Final:          "osaka flat (-o)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"zong",
			Word{
				Initial:        "steve gomeZ (zu-)",
				Final:          "charles house (-eng)",
				Pronounciation: "tongue behind top teeth",
			},
		},
		{
			"wang",
			Word{
				Initial: "walter white (w-)",
				Final:   "angkor flat (-ang)",
			},
		},
		{
			"wei",
			Word{
				Initial: "walter white (w-)",
				Final:   "subway station (-ei)",
			},
		},
		{
			"shui",
			Word{
				Initial:        "skyler white (shu-)",
				Final:          "subway station (-ei)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"shun",
			Word{
				Initial:        "skyler white (shu-)",
				Final:          "endelich’s flat (-en)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"chuai",
			Word{
				Initial:        "chuck mcgill (chu-)",
				Final:          "iceland hut (-ai)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"chuan",
			Word{
				Initial:        "chuck mcgill (chu-)",
				Final:          "hanoi flat (-an)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"chuang",
			Word{
				Initial:        "chuck mcgill (chu-)",
				Final:          "angkor flat (-ang)",
				Pronounciation: "tongue on roof of the mouth",
			},
		},
		{
			"qie",
			Word{
				Initial:        "qualle (qi-)",
				Final:          "1st flat berlin (-e)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"qin",
			Word{
				Initial:        "qualle (qi-)",
				Final:          "endelich’s flat (-en)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"qing",
			Word{
				Initial:        "qualle (qi-)",
				Final:          "charles house (-eng)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
		{
			"qiu",
			Word{
				Initial:        "qualle (qi-)",
				Final:          "sis’ house (-ou)",
				Pronounciation: "tongue behind bottom teeth",
			},
		},
	}

	for _, tt := range tests {
		actual, err := table.Find(tt.pinyin)
		if err != nil {
			t.Errorf("%s: %v", tt.pinyin, err)
			continue
		}
		if actual.Initial != tt.expected.Initial {
			t.Errorf("expect initial: %s but got: %s", tt.expected.Initial, actual.Initial)
			continue
		}
		if actual.Final != tt.expected.Final {
			t.Errorf("expect final: %s but got: %s", tt.expected.Final, actual.Final)
			continue
		}
		if actual.Pronounciation != tt.expected.Pronounciation {
			t.Errorf("%s expect pronounciation: %s but got: %s", tt.pinyin, tt.expected.Pronounciation, actual.Pronounciation)
		}
	}
}
