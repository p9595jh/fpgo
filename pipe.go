package fpgo

func Pipe[I, O any](input I, funcs ...F) O {
	s := &Shell{input}
	for _, f := range funcs {
		s = f(s)
	}
	return s.V.(O)
}

func ProcessingPipe[I, O any](input I, processors [2]A, funcs ...F) O {
	if processors[0] == nil {
		processors[0] = func(a any) {}
	}
	if processors[1] == nil {
		processors[1] = func(a any) {}
	}
	s := &Shell{input}
	for _, f := range funcs {
		processors[0](s.V)
		s = f(s)
		processors[1](s.V)
	}
	return s.V.(O)
}

func Once[I, O any](input I, f F) O {
	return f(&Shell{input}).V.(O)
}
