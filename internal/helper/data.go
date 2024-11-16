package helper

type DataConstraint interface {
	~[]byte | ~string
}

func PrepareData[TData DataConstraint](data TData) []byte {
	var inData []byte

	switch any(data).(type) {
	case string:
		inData = []byte(any(data).(string))
	case []byte:
		inData = any(data).([]byte)
	}

	return inData
}
