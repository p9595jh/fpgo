package fpgo

// standard function to be a connection of the pipeline
type F func(*Shell) *Shell

// parameterized function which has 1 generic parameter
type P1[P, R any] func(P) R

// parameterized function which has 2 generic parameters
type P2[P1, P2, R any] func(P1, P2) R

// parameterized function which has index and 1 generic parameter
type Pi1[P, R any] func(int, P) R

// parameterized function which has index and 2 generic parameters
type Pi2[P1, P2, R any] func(int, P1, P2) R

// parameterized function which has index and 2 generic parameters and returns void
type Pi2V[P1, P2 any] func(int, P1, P2)

type A func(any)
