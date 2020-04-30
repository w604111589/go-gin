package example

import (
	"log"
	"testing"
)

func Test_StringPlus(t *testing.T) {
	StringPlus()
	t.Log("记录测试时间")
}

func Benchmark_StringPlus(B *testing.B) {
	log.Println("测试次数:", B.N)
	for i := 0; i < B.N; i++ {
		StringPlus()
	}

}

func Test_StringJoin(t *testing.T) {
	StringJoin()
	t.Log("记录测试时间")
}

func Benchmark_StringJoin(B *testing.B) {
	log.Println("测试次数:", B.N)
	for i := 0; i < B.N; i++ {
		StringJoin()
	}
}

func Test_StringBuilder(t *testing.T) {
	StringBuilder()
	t.Log("记录测试时间")
}

func Test_StringBuffer(t *testing.T) {
	StringBuffer()
	t.Log("记录测试时间")
}

func Test_StringFmt(t *testing.T) {
	StringFmt()
	t.Log("记录测试时间")
}
