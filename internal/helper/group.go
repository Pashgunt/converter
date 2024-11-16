package helper

type GroupConstraint interface {
	~[]string | ~string
}

func PrepareGroups[TGroups GroupConstraint](groups TGroups) []string {
	var inGroups []string

	switch any(groups).(type) {
	case string:
		inGroups = []string{any(groups).(string)}
	case []string:
		inGroups = any(groups).([]string)
	}

	return inGroups
}
