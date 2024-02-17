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
				Initial: "qtip (qu-)",
				Final:   "space station (-Ø)",
			},
		},
		{
			"zong",
			Word{
				Initial: "steve gomeZ (zu-)",
				Final:   "england charles house (-eng)",
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
				Initial: "cosmonaut (c-)",
				Final:   "anna’s flat (-a)",
			},
		},
		{
			"cai",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "iceland hut (-ai)",
			},
		},
		{
			"can",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "hanoi flat (-an)",
			},
		},
		{
			"cang",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "angkor flat (-ang)",
			},
		},
		{
			"cao",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "autobahn stop (-ao)",
			},
		},
		{
			"ce",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "1st flat berlin (-e)",
			},
		},
		{
			"cen",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "endelich’s flat (-en)",
			},
		},
		{
			"ceng",
			Word{
				Initial: "cosmonaut (c-)",
				Final:   "england charles house (-eng)",
			},
		},
		{
			"cha",
			Word{
				Initial: "(jesus) christ (ch-)",
				Final:   "anna’s flat (-a)",
			},
		},
		{
			"chai",
			Word{
				Initial: "(jesus) christ (ch-)",
				Final:   "iceland hut (-ai)",
			},
		},
		{
			"chan",
			Word{
				Initial: "(jesus) christ (ch-)",
				Final:   "hanoi flat (-an)",
			},
		},
		{
			"chang",
			Word{
				Initial: "(jesus) christ (ch-)",
				Final:   "angkor flat (-ang)",
			},
		},
		{
			"xuan",
			Word{
				Initial: "xhibit (xu-)",
				Final:   "hanoi flat (-an)",
			},
		},
		{
			"xue",
			Word{
				Initial: "xhibit (xu-)",
				Final:   "1st flat berlin (-e)",
			},
		},
		{
			"xun",
			Word{
				Initial: "xhibit (xu-)",
				Final:   "endelich’s flat (-en)",
			},
		},
		{
			"yong",
			Word{
				Initial: "yung hurn (yu-)",
				Final:   "england charles house (-eng)",
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
				Initial: "zafiro añejo (zhu-)",
				Final:   "osaka flat (-o)",
			},
		},
		{
			"zong",
			Word{
				Initial: "steve gomeZ (zu-)",
				Final:   "england charles house (-eng)",
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
				Initial: "skyler white (shu-)",
				Final:   "subway station (-ei)",
			},
		},
		{
			"shun",
			Word{
				Initial: "skyler white (shu-)",
				Final:   "endelich’s flat (-en)",
			},
		},
		{
			"chuai",
			Word{
				Initial: "chuck mcgill (chu-)",
				Final:   "iceland hut (-ai)",
			},
		},
		{
			"chuan",
			Word{
				Initial: "chuck mcgill (chu-)",
				Final:   "hanoi flat (-an)",
			},
		},
		{
			"chuang",
			Word{
				Initial: "chuck mcgill (chu-)",
				Final:   "angkor flat (-ang)",
			},
		},
		{
			"qie",
			Word{
				Initial: "qualle (qi-)",
				Final:   "1st flat berlin (-e)",
			},
		},
		{
			"qin",
			Word{
				Initial: "qualle (qi-)",
				Final:   "endelich’s flat (-en)",
			},
		},
		{
			"qing",
			Word{
				Initial: "qualle (qi-)",
				Final:   "england charles house (-eng)",
			},
		},
		{
			"qiu",
			Word{
				Initial: "qualle (qi-)",
				Final:   "sis’ house (-ou)",
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
	}
}
