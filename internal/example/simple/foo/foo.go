package foo

type Foo struct {
	Example string
}

func New(example string) Foo {
	return Foo{
		Example: example,
	}
}
