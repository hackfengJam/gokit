package istrings

// StringList2Map String List转 String Map
func StringList2Map(stringList []string) (vMap map[string]struct{}) {
	vMap = make(map[string]struct{}, len(stringList))
	for _, v := range stringList {
		vMap[v] = struct{}{}
	}
	return vMap
}
