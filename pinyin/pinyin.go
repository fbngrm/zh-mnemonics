package pinyin

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
)

type Word struct {
	Initial        string
	Final          string
	Pronounciation string
}

type Table struct {
	groupA map[string]Word
	groupI map[string]Word
	groupU map[string]Word
	groupV map[string]Word
}

var groupA = `	space station (-Ø)	anna’s flat (-a)	osaka flat (-o)	1st flat berlin (-e)	iceland hut (-ai)	subway station (-ei)	autobahn stop (-ao)	sis’ house (-ou)	hanoi flat (-an)	endelich’s flat (-en)	angkor flat (-ang)	charles house (-eng)
ghost (Ø-)	er	a	o	e	ai	ei	ao	ou	an	en	ang	eng
bitch (b-)		ba	bo|1		bai	bei	bao		ban	ben	bang	beng
police (p-)		pa	po|1		pai	pei	pao	pou	pan	pen	pang	peng
monk (m-)		ma	mo|1	me	mai	mei	mao	mou	man	men	mang	meng
fred (f-)		fa	fo|1			fei		fou	fan	fen	fang	feng
dostojevski (d-)		da		de	dai	dei	dao	dou	dan	den	dang	deng
travis bickle (t-)		ta		te	tai	tei	tao	tou	tan		tang	teng
nicola tesla (n-)		na		ne	nai	nei	nao	nou	nan	nen	nang	neng
lover(s) (l-)		la	lo	le	lai	lei	lao	lou	lan		lang	leng
goonies (g-)		ga		ge	gai	gei	gao	gou	gan	gen	gang	geng
korean girl (k-)		ka		ke	kai		kao	kou	kan	ken	kang	keng
hitler (h-)		ha		he	hai	hei	hao	hou	han	hen	hang	heng
zhou zhou (zh-)	zhi|2	zha|2		zhe|2	zhai|2	zhei|2	zhao|2	zhou|2	zhan|2	zhen|2	zhang|2	zheng|2
(jesus) christ (ch-)	chi|2	cha|2		che|2	chai|2		chao|2	chou|2	chan|2	chen|2	chang|2	cheng|2
shlomo (sh-)	shi|2	sha|2		she|2	shai|2	shei|2	shao|2	shou|2	shan|2	shen|2	shang|2	sheng|2
robocop (r-)	ri|2			re|2			rao|2	rou|2	ran|2	ren|2	rang|2	reng|2
zuki (z-)	zi|3	za|3		ze|3	zai|3	zei|3	zao|3	zou|3	zan|3	zen|3	zang|3	zeng|3
cosmonaut (c-)	ci|3	ca|3		ce|3	cai|3		cao|3	cou|3	can|3	cen|3	cang|3	ceng|3
santa claus (s-)	si	sa		se	sai		sao	sou	san	sen	sang	seng`

var groupI = `	space station (-Ø)	anna’s flat (-a)	osaka flat (-o)	1st flat berlin (-e)	iceland hut (-ai)	autobahn stop (-ao)	sis’ house (-ou)	hanoi flat (-an)	endelich’s flat (-en)	angkor flat (-ang)	charles house (-eng)
yak (y-)	yi|4	ya	yo	ye	yai	yao	you	yan	yin	yang	ying
bird (bi-)	bi			bie		biao		bian	bin		bing
piglet (pi-)	pi			pie		piao		pian	pin		ping
mink (mi-)	mi			mie		miao	miu	mian	min		ming
dinosaur (di-)	di	dia		die		diao	diu	dian			ding
tiger (ti-)	ti			tie		tiao		tian			ting
nilpferd (ni-)	ni			nie		niao	niu	nian	nin	niang	ning
lion (li-)	li	lia		lie		liao	liu	lian	lin	liang	ling
giraffe (ji-)	ji|5	jia|5		jie|5		jiao|5	jiu|5	jian|5	jin|5	jiang|5	jing|5
qualle (qi-)	qi|5	qia|5		qie|5		qiao|5	qiu|5	qian|5	qin|5	qiang|5	qing|5
x-gf/dragon (xi-)	xi|5	xia|5		xie|5		xiao|5	xiu|5	xian|5	xin|5	xiang|5	xing|5`

var groupU = `	space station (-Ø)	anna’s flat (-a)	osaka flat (-o)	iceland hut (-ai)	subway station (-ei)	hanoi flat (-an)	endelich’s flat (-en)	angkor flat (-ang)	charles house (-eng)
walter white (w-)	wu	wa	wo	wai	wei	wan	wen	wang	weng
betsy kettleman (bu-)	bu
paige novick (pu-)	pu
mike ermentraut (mu-)	mu
francesca liddy (fu-)	fu
don eladio (du-)	du		duo		dui	duan	dun		dong
tuco salamanca (tu-)	tu		tuo		tui	tuan	tun		tong
nacho varga (nu-)	nu		nuo			nuan			nong
lalo salamanca (lu-)	lu		luo			luan	lun		long
gus fring (gu-)	gu	gua	guo	guai	gui	guan	gun	guang	gong
kim wexler (ku-)	ku	kua	kuo	kuai	kui	kuan	kun	kuang	kong
huell babineaux (hu-)	hu	hua	huo	huai	hui	huan	hun	huang	hong
zafiro añejo (zhu-)	zhu|2	zhua|2	zhuo|2	zhuai|2	zhui|2	zhuan|2	zhun|2	zhuang|2	zhong|2
chuck mcgill (chu-)	chu|2	chua|2	chuo|2	chuai|2	chui|2	chuan|2	chun|2	chuang|2	chong|2
skyler white (shu-)	shu|2	shua|2	shuo|2	shuai|2	shui|2	shuan|2	shun|2	shuang|2
ruth mcgill (ru-)	ru|2	rua|2	ruo|2		rui|2	ruan|2	run|2		rong|2
steve gomeZ (zu-)	zu|3		zuo|3		zui|3	zuan|3	zun|3		zong|3
crazy-8 (cu-)	cu|3		cuo|3		cui|3	cuan|3	cun|3		cong|3
saul goodman (su-)	su		suo		sui	suan	sun		song`

var groupV = `	space station (-Ø)	1st flat berlin (-e)	hanoi flat (-an)	endelich’s flat (-en)	charles house (-eng)
yung hurn (yu-)	yu	yue	yuan	yun	yong
necro (nü-)	nü	nüe
lord finesse (lü-)	lü	lüe
jeru the damaja (ju-)	ju|5	jue|5	juan|5	jun|5	jiong|5
qtip (qu-)	qu|5	que|5	quan|5	qun|5	qiong|5
xhibit (xu-)	xu|5	xue|5	xuan|5	xun|5	xiong|5`

func NewTable() *Table {
	return &Table{
		groupA: getTable(groupA),
		groupI: getTable(groupI),
		groupU: getTable(groupU),
		groupV: getTable(groupV),
	}
}

func (t *Table) Find(pinyin string) (Word, error) {

	if w, ok := t.groupA[pinyin]; ok {
		return w, nil
	}
	if w, ok := t.groupI[pinyin]; ok {
		return w, nil
	}
	if w, ok := t.groupU[pinyin]; ok {
		return w, nil
	}
	if w, ok := t.groupV[pinyin]; ok {
		return w, nil
	}
	return Word{}, fmt.Errorf("not found: %s", pinyin)
}

func getTable(group string) map[string]Word {
	table := make(map[string]Word)
	rows := strings.Split(group, "\n")
	finals := strings.Split(rows[0], "\t")

	for i, row := range rows {
		// skip the header row
		if i == 0 {
			continue
		}

		// split the row into columns
		cols := strings.Split(row, "\t")
		initial := cols[0]

		for j, col := range cols {
			// skip the first column as it represents the row header
			if j == 0 {
				continue
			}
			if col == "" {
				continue
			}
			pronounciation := ""
			pinyin := col
			if strings.Contains(col, "|") {
				parts := strings.Split(col, "|")
				if len(parts) != 2 {
					slog.Error("get pronounciation", "word", col, "len", len(parts))
					continue
				}
				pinyin = parts[0]
				i, err := strconv.Atoi(parts[1])
				if err != nil {
					slog.Error("get pronounciation", "word", col, "err", err.Error())
					continue
				}
				pronounciation = getPronounciation(i)
			}
			table[pinyin] = Word{
				Initial:        initial,
				Final:          finals[j],
				Pronounciation: pronounciation,
			}

		}
	}
	return table
}
