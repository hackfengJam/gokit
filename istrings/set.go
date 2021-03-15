package istrings

// StringList2Set String List 转 String List Set
func StringList2Set(stringList []string) (noRepeatedList []string) {
	noRepeatedList = make([]string, 0 ,len(stringList))

	vMap := StringList2Map(stringList)
	for v := range vMap {
		noRepeatedList = append(noRepeatedList, v)
	}
	return noRepeatedList
}

