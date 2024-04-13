package types

type Type interface {
	InitAndValidate() error
}

type KeyValue interface {
	getKey() string
	getValue() string
}
