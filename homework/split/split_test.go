package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Splitstr("qwdqevaqwda", "wd")
	want := []string{"q", "qevaq", "a"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got:%v\n", got, want)
	}
}

//测试组 将几个相同的对同一个函数的测试组合起来
func TestSS(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	tests := []testCase{
		testCase{"awcasada", "a", []string{"wc", "s", "d"}},
		testCase{"rvwrasd", "r", []string{"vw", "asd"}},
		testCase{"qwdqevaqwda", "wd", []string{"q", "qevaq", "a"}},
	}
	for _, test := range tests {
		got := Splitstr(test.str, test.sep)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want:%v but got:%v\n", got, test.want)
		}
	}
}

//子测试 其实就是把切片换成map 这样就可以把每一个测试都显示在命令行中了
func TestMap(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	tests := map[string]testCase{
		"case1": {"awcasada", "a", []string{"wc", "s", "d"}},
		"case2": {"rvwrasd", "r", []string{"vw", "asd"}},
		"case3": {"qwdqevaqwda", "wd", []string{"q", "qevaq", "a"}},
		"case4": {"qwd", "", []string{"q", "w", "d"}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := Splitstr(test.str, test.sep)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want:%v but got:%v\n", got, test.want)
			}
		})
	}
}

//如果想跑测试组里面的某一个测试 就运行
//go test -run=TestMap/case2
//测试覆盖率
//go test -cover

//性能基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Splitstr("qwdqevaqwda", "wd")
	}
}

//go test -bench=Split
//go test -bench=Split -benchmem

//函数比较测试 统一函数执行不同的次数所需要的时间进行比较
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

// go test -bench=Fib2
