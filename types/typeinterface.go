package types

type Type interface {
	InitAndValidate() error
	IsEmpty() bool
}

type KeyValue interface {
	GetKey() string
	GetValue() string
	IsDisabled() bool
}
