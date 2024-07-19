package buffer_test

import (
	"strconv"
	"strings"
	"testing"
)

func concat() string {
	var str string
	for i := range 100 {
		str += strconv.Itoa(i)
	}
	return str
}

func concatBuilder() string {
	var builder strings.Builder
	for i := range 100 {
		builder.WriteString(strconv.Itoa(i))
	}
	return builder.String()
}

func BenchmarkConcat(b *testing.B) {
	for range b.N {
		concat()
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for range b.N {
		concatBuilder()
	}
}
