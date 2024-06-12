package pinyin

func getPronounciation(x int) string {
	pro := map[int]string{
		1: "pronounce a as uo",
		2: "tongue on roof of the mouth",
		3: "tongue behind top teeth",
		4: "Initial I is pronounced with y",
		5: "tongue behind bottom teeth",
	}
	if p, ok := pro[x]; ok {
		return p
	}
	return ""
}
