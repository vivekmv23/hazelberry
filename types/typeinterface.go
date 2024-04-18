package types

type Type interface {
	InitAndValidate() error
	IsEmpty() bool
}

type KeyValue interface {
	getKey() string
	getValue() string
}
