package pinyin

import (
	"fmt"
	"strings"
)

type Word struct {
	Initial string
	Final   string
}

type Table struct {
	groupA map[string]Word
	groupI map[string]Word
	groupU map[string]Word
	groupV map[string]Word
}

var groupA = `	space station (-Ø)	anna’s flat (-a)	osaka flat (-o)	1st flat berlin (-e)	iceland hut (-ai)	subway station (-ei)	autobahn stop (-ao)	sis’ house (-ou)	hanoi flat (-an)	endelich’s flat (-en)	angkor flat (-ang)	england charles house (-eng)
ghost (Ø-)	er	a	o	e	ai	ei	ao	ou	an	en	ang	eng
bitch (b-)		ba	bo [buo]		bai	bei	bao		ban	ben	bang	beng
police (p-)		pa	po [puo]		pai	pei	pao	pou	pan	pen	pang	peng
monk (m-)		ma	mo [muo]	me	mai	mei	mao	mou	man	men	mang	meng
fred (f-)		fa	fo [fuo]			fei		fou	fan	fen	fang	feng
dostojevski (d-)		da		de	dai	dei	dao	dou	dan	den	dang	deng
travis bickle (t-)		ta		te	tai	tei	tao	tou	tan		tang	teng
nicola tesla (n-)		na		ne	nai	nei	nao	nou	nan	nen	nang	neng
lover(s) (l-)		la	lo	le	lai	lei	lao	lou	lan		lang	leng
goonies (g-)		ga		ge	gai	gei	gao	gou	gan	gen	gang	geng
korean girl (k-)		ka		ke	kai		kao	kou	kan	ken	kang	keng
hitler (h-)		ha		he	hai	hei	hao	hou	han	hen	hang	heng
zhou zhou (zh-)	zhi	zha		zhe	zhai	zhei	zhao	zhou	zhan	zhen	zhang	zheng
(jesus) christ (ch-)	chi	cha		che	chai		chao	chou	chan	chen	chang	cheng
shlomo (sh-)	shi	sha		she	shai	shei	shao	shou	shan	shen	shang	sheng
robocop (r-)	ri			re			rao	rou	ran	ren	rang	reng
zuki (z-)	zi	za		ze	zai	zei	zao	zou	zan	zen	zang	zeng
cosmonaut (c-)	ci	ca		ce	cai		cao	cou	can	cen	cang	ceng
santa claus (s-)	si	sa		se	sai		sao	sou	san	sen	sang	seng`

var groupI = `	space station (-Ø)	anna’s flat (-a)	osaka flat (-o)	1st flat berlin (-e)	iceland hut (-ai)	autobahn stop (-ao)	sis’ house (-ou)	hanoi flat (-an)	endelich’s flat (-en)	angkor flat (-ang)	england charles house (-eng)
yak (y-)	yi	ya	yo	ye	yai	yao	you	yan	yin	yang	ying
bird (bi-)	bi			bie		biao		bian	bin		bing
piglet (pi-)	pi			pie		piao		pian	pin		ping
mink (mi-)	mi			mie		miao	miu	mian	min		ming
dinosaur (di-)	di	dia		die		diao	diu	dian			ding
tiger (ti-)	ti			tie		tiao		tian			ting
nilpferd (ni-)	ni			nie		niao	niu	nian	nin	niang	ning
lion (li-)	li	lia		lie		liao	liu	lian	lin	liang	ling
giraffe (ji-)	ji	jia		jie		jiao	jiu	jian	jin	jiang	jing
qualle (qi-)	qi	qia		qie		qiao	qiu	qian	qin	qiang	qing
x-gf/dragon (xi-)	xi	xia		xie		xiao	xiu	xian	xin	xiang	xing`

var groupU = `	space station (-Ø)	anna’s flat (-a)	osaka flat (-o)	iceland hut (-ai)	subway station (-ei)	hanoi flat (-an)	endelich’s flat (-en)	angkor flat (-ang)	england charles house (-eng)
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
zafiro añejo (zhu-)	zhu	zhua	zhuo	zhuai	zhui	zhuan	zhun	zhuang	zhong
chuck mcgill (chu-)	chu	chua	chuo	chuai	chui	chuan	chun	chuang	chong
skyler white (shu-)	shu	shua	shuo	shuai	shui	shuan	shun	shuang
ruth mcgill (ru-)	ru	rua	ruo		rui	ruan	run		rong
steve gomeZ (zu-)	zu		zuo		zui	zuan	zun		zong
crazy-8 (cu-	cu		cuo		cui	cuan	cun		cong
saul goodman (su-)	su		suo		sui	suan	sun		song`

var groupV = `	space station (-Ø)	1st flat berlin (-e)	hanoi flat (-an)	endelich’s flat (-en)	england charles house (-eng)
yung hurn (yu-)	yu	yue	yuan	yun	yong
necro (nü-)	nü	nüe
lord finesse (lü-)	lü	lüe
jeru the damaja (ju-)	ju	jue	juan	jun	jiong
qtip (qu-)	qu	que	quan	qun	qiong
xhibit (xu-)	xu	xue	xuan	xun	xiong`

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
			table[col] = Word{
				Initial: initial,
				Final:   finals[j],
			}

		}
	}

	return table
}
