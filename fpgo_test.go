package fpgo_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/p9595jh/fpgo"
)

func TestPipe(t *testing.T) {
	v := fpgo.Pipe[[]int, *int](
		[]int{1, 3, 2, 4, 5},
		fpgo.Sort(func(i1, i2 *int) bool {
			return *i1 > *i2
		}),
		fpgo.Map(func(_ int, t *int) *string {
			s := fmt.Sprintf("abc%d", *t)
			return &s
		}),
		fpgo.Slice[string](1, 4),
		fpgo.Reverse[string](),
		fpgo.Reduce(func(i int, s1, s2 *string) *string {
			s := fmt.Sprintf("%s %s", *s1, *s2)
			return &s
		}),
		fpgo.Func(func(s *string) []string {
			*s = strings.Trim(*s, " ")
			return strings.Split(*s, " ")
		}),
		fpgo.Map(func(_ int, t *string) *int {
			i, _ := strconv.ParseInt((*t)[3:], 10, 32)
			i2 := int(i)
			return &i2
		}),
		fpgo.Reduce(func(i int, t *int, u *int) *int {
			*u += *t
			return u
		}),
	)
	t.Log(*v)
}

func TestProcessing(t *testing.T) {
	res := fpgo.ProcessingPipe[[]int, []int](
		[]int{5, 4, 1, 3, 2},
		[2]fpgo.A{
			func(a any) {
				t.Log(a)
			},
		},
		fpgo.Sort(func(t1, t2 *int) bool {
			return *t1 < *t2
		}),
		fpgo.Reverse[int](),
		fpgo.At[int](2),
		fpgo.MapOne(func(i *int) *[]int {
			return &[]int{*i}
		}),
		fpgo.Append([]int{10, 20}),
	)
	t.Log(res)
}

func TestChain(t *testing.T) {
	s := fpgo.New([]int{2, 4, 1, 5, 3}).
		F(fpgo.Append([]int{10, 11})).
		F(fpgo.Reverse[int]()).
		F(fpgo.Map(func(i1 int, i2 *int) *int {
			*i2 += i1
			return i2
		})).
		F(fpgo.Filter(func(i1 int, i2 *int) bool {
			return i1%2 == 1
		})).
		F(fpgo.ForEach(func(i1 int, i2 *int, i3 []int) {
			t.Log(*i2)
		})).
		F(fpgo.At[int](-1)).
		F(fpgo.MapOne(func(i *int) *string {
			s := fmt.Sprintf("*%d*", *i)
			return &s
		})).V.(string)
	t.Log(s)
}
