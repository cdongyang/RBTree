package rbtree_test

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleInterface() {
	var null Iterator
	var root *node
	var iter Iterator = root
	var NULL = root
	fmt.Println(iter, root)
	fmt.Println(iter == nil, root == nil, iter == root)
	fmt.Println(reflect.DeepEqual(iter, root), reflect.DeepEqual(iter, nil))
	fmt.Println(iter == null, root == null, null == nil)
	fmt.Println(iter == NULL, root == NULL, NULL == nil)
	fmt.Println(iter == (*node)(nil))
	switch iter.(type) {
	case *node:
		fmt.Println("*node")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("other")
	}
	//== judge type first,then judge value
	// Output:
	//<nil> <nil>
	//false true true
	//true false
	//false false true
	//true true true
	//true
	//*node
}

func BenchmarkInterface(b *testing.B) {
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = i
		_ = iface.(int)
	}
}

func BenchmarkInt(b *testing.B) {
	var it int
	for i := 0; i < b.N; i++ {
		it = i
		_ = it
	}
}

func testFunc(a int) int {
	return a
}

type testStruct struct {
	a int
}

func (t *testStruct) testFunc(a int) int {
	return a
}

func (t testStruct) testFunc1(a int) int {
	return a
}

type testInterface interface {
	testFunc(int) int
}

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFunc(0)
	}
}

func BenchmarkStructFunc(b *testing.B) {
	var t testStruct
	for i := 0; i < b.N; i++ {
		t.testFunc1(0)
	}
}

func BenchmarkStructPoiterFunc(b *testing.B) {
	var t = &testStruct{}
	for i := 0; i < b.N; i++ {
		t.testFunc(0)
	}
}

func BenchmarkStructPoiterFunc1(b *testing.B) {
	var t = &testStruct{}
	for i := 0; i < b.N; i++ {
		t.testFunc1(0)
	}
}

func BenchmarkInterfaceFunc(b *testing.B) {
	var iface testInterface = &testStruct{}
	for i := 0; i < b.N; i++ {
		iface.testFunc(0)
	}
}