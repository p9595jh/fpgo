# fpgo

## Summary

Simple implementation of `Functional Programming` in Golang.

It uses generic so Go1.18+ needed.

## Description

Array or object is wrapped by `Shell`, and you can handle this using some functions.

All the fp functions return a function which has `*Shell` parameter and return type that named `F`, and this `F` is connecting the functional programming.

## Types

Before explaining the functions, there are custom types to represent more easier.

```
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
```

## Functions

Some of the functions are similar to javascript's.

```
Map[T, U any](convert Pi1[*T, *U]) F
Filter[T any](f Pi1[*T, bool]) F
Slice[T any](idxes ...int) F
Reduce[T, U any](reducer Pi2[*T, *U, *U]) F
Sort[T any](sorter P2[*T, *T, bool]) F
ArrayCopy[T any]() F
Func[T, U any](f P1[T, U]) F
Some[T any](f Pi1[*T, bool]) F
Every[T any](f Pi1[*T, bool]) F
ForEach[T any](f Pi2V[*T, []T]) F
Reverse[T any]() F
Append[T any](elements []T) F
MapOne[T, U any](f P1[*T, *U]) F
```

There are 2 ways to connect those functions:

- using pipe
- chaining

## Pipe

`Pipe` can handle the input using parameterized functions (`F`).

```
// `I` means input type and `O` means output type
Pipe[I, O any](input I, funcs ...F) O

// if you want to reverse and convert int array to string array,
stringArray := fpgo.Pipe[[]int, []string](
    []int{1, 2, 3},
    fpgo.Reverse[int](),
    fpgo.Map[int, string](func(idx int, element *int) *string) {
        s := fmt.Sprintf("No.%d", *element)
        return &s
    }
)
fmt.Println(stringArray) // [No.3 No.2 No.1]
```

You can also preprocess and postprocess by using `ProcessingPipe`.

```
// processors array is like {preprocessor, postprocessor}
ProcessingPipe[I, O any](input I, processors [2]A, funcs ...F) O

// if you want to see before and after the processing of each pipe work,
fpgo.ProcessingPipe[[]int, []int](
    []int{1, 2, 3},
    [2]fpgo.A{
        func(a any) {
            fmt.Printf("before: %v\n", a)
        },
        func(a any) {
            fmt.Printf("after: %v\n", a)
        },
    },
    functions...,
)
```

## Chaining

Chaining can be executed by the method `F` of `Shell`.

A way to reproduce the example of `Pipe` is like below.

```
stringArray := fpgo.New([]int{1, 2, 3}).
                    F(fpgo.Reverse[int]()).
                    F(fpgo.Map[int, string](func(idx int, element *int) *string) {
                        s := fmt.Sprintf("No.%d", *element)
                        return &s
                    }).V.([]string)
fmt.Println(stringArray) // [No.3 No.2 No.1]
```

## How to use

```
go get github.com/p9595jh/fpgo
```

```
import "github.com/p9595jh/fpgo"
```
