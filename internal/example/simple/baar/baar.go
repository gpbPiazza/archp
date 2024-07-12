package baar

type Baar struct {
	Example string
}

func New(example string) Baar {
	return Baar{
		Example: example,
	}
}
