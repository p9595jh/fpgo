package fpgo

type Shell struct {
	V any
}

func New(v any) *Shell {
	return &Shell{v}
}

func (s *Shell) F(f F) *Shell {
	return f(s)
}
