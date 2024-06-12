package tone

func GetLocation(pinyin string) string {
	locations := map[int]string{
		firstTone:  "in front of",
		secondTone: "in the entrance of",
		thirdTone:  "somewhere inside of",
		fourthTone: "in the bathroom of",
		noTone:     "",
	}
	if l, ok := locations[getTone(pinyin)]; ok {
		return l
	}
	return ""
}
