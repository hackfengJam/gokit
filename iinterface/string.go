package iinterface

// StringList2InterfaceList StringList to InterfaceList
func StringList2InterfaceList(strList []string) (ifList []interface{}) {
	ifList = make([]interface{}, len(strList))
	for _, str := range strList {
		ifList = append(ifList, str)
	}
	return
}
