package iinterface

// Int64List2InterfaceList Int64List to InterfaceList
func Int64List2InterfaceList(inList []int64) (ifList []interface{}) {
	ifList = make([]interface{}, len(inList))
	for _, str := range inList {
		ifList = append(ifList, str)
	}
	return
}

// UInt64List2InterfaceList UInt64List to InterfaceList
func UInt64List2InterfaceList(inList []uint64) (ifList []interface{}) {
	ifList = make([]interface{}, len(inList))
	for _, str := range inList {
		ifList = append(ifList, str)
	}
	return
}

// Int32List2InterfaceList Int32List to InterfaceList
func Int32List2InterfaceList(inList []int32) (ifList []interface{}) {
	ifList = make([]interface{}, len(inList))
	for _, str := range inList {
		ifList = append(ifList, str)
	}
	return
}

// UInt32List2InterfaceList UInt32List to InterfaceList
func UInt32List2InterfaceList(inList []uint32) (ifList []interface{}) {
	ifList = make([]interface{}, len(inList))
	for _, str := range inList {
		ifList = append(ifList, str)
	}
	return
}

// IntList2InterfaceList IntList to InterfaceList
func IntList2InterfaceList(inList []int) (ifList []interface{}) {
	ifList = make([]interface{}, len(inList))
	for _, str := range inList {
		ifList = append(ifList, str)
	}
	return
}

// UIntList2InterfaceList UIntList to InterfaceList
func UIntList2InterfaceList(inList []uint) (ifList []interface{}) {
	ifList = make([]interface{}, len(inList))
	for _, str := range inList {
		ifList = append(ifList, str)
	}
	return
}
