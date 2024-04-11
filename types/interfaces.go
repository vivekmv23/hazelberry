package types

type Type interface {
	InitAndValidate() error
}

type Credentials interface {
	getId() string
	getPass() string
}
