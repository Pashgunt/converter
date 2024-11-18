package enum

const (
	GroupNameDetailInit = `^\[(\w+(?:,\w+)*)\]`
	GroupNameDetail     = `\[(\w+(?:,\w+)*)\]`
	GroupPrefixList     = `(?:^|\(|\.|\.\()([\w,]+)`
	GroupPrefix         = `(?:\w+\.)?(\w+)$`
	CamelCase           = `([a-z0-9])([A-Z])`
	SnackCase           = `${1}_${2}`
)
