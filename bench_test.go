package rbtree_test

import (
	"strconv"
	"testing"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/rbtree"
)

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

func BenchmarkAll(b *testing.B) {
	var ns = []int{1e3, 1e4, 1e5, 1e6, 1e7}
	b.Run("setInsert", runWith(benchmarkSetInsert, ns...))
	b.Run("setNoescapeInsert", runWith(benchmarkSetNoescapeInsert, ns...))
	b.Run("setErase", runWith(benchmarkSetErase, ns...))
	b.Run("setFind", runWith(benchmarkSetFind, ns...))
	b.Run("mapInsert", runWith(benchmarkMapInsert, ns...))
	b.Run("mapErase", runWith(benchmarkSetErase, ns...))
	b.Run("mapFind", runWith(benchmarkMapFind, ns...))
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, ns...))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, ns...))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, ns...))
}

func Benchmark(b *testing.B) {
	BenchmarkSet(b)
	BenchmarkMap(b)
	BenchmarkHashMap(b)
}

func BenchmarkSet(b *testing.B) {
	ns := []int{1e1, 1e2, 1e3, 1e4, 1e5, 1e6}
	b.Run("setInsert", runWith(benchmarkSetInsert, 0))
	b.Run("setNoescapeInsert", runWith(benchmarkSetNoescapeInsert, 0))
	b.Run("setErase", runWith(benchmarkSetErase, 0))
	b.Run("setFind", runWith(benchmarkSetFind, 0))
	b.Run("setFindM", runWith(benchmarkSetFindM, ns...))
	b.Run("set", runWith(benchmarkSet, ns...))
}

func BenchmarkSet1E5(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e5))
	b.Run("setNoescapeInsert", runWith(benchmarkSetNoescapeInsert, 1e5))
	b.Run("setErase", runWith(benchmarkSetErase, 1e5))
	b.Run("setFind", runWith(benchmarkSetFind, 1e5))
}

func BenchmarkSet1E6(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e6))
	b.Run("setNoescapeInsert", runWith(benchmarkSetNoescapeInsert, 1e6))
	b.Run("setErase", runWith(benchmarkSetErase, 1e6))
	b.Run("setFind", runWith(benchmarkSetFind, 1e6))
}

func BenchmarkMap(b *testing.B) {
	ns := []int{1e1, 1e2, 1e3, 1e4, 1e5, 1e6}
	b.Run("mapInsert", runWith(benchmarkMapInsert, 0))
	b.Run("mapErase", runWith(benchmarkSetErase, 0))
	b.Run("mapFind", runWith(benchmarkMapFind, 0))
	b.Run("mapFindM", runWith(benchmarkMapFindM, ns...))
	b.Run("map", runWith(benchmarkMap, ns...))
}

func BenchmarkMap1E5(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 1e5))
	b.Run("mapErase", runWith(benchmarkSetErase, 1e5))
	b.Run("mapFind", runWith(benchmarkMapFind, 1e5))
}

func BenchmarkMap1E6(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 1e6))
	b.Run("mapErase", runWith(benchmarkSetErase, 1e6))
	b.Run("mapFind", runWith(benchmarkMapFind, 1e6))
}

func BenchmarkHashMap(b *testing.B) {
	ns := []int{1e1, 1e2, 1e3, 1e4, 1e5, 1e6}
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 0))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 0))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 0))
	b.Run("hashMapFindM", runWith(benchmarkHashMapFindM, ns...))
	b.Run("hashMap", runWith(benchmarkHashMap, ns...))
}

func BenchmarkHashMap1E5(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 1e5))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 1e5))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 1e5))
}

func BenchmarkHashMap1E6(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 1e6))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 1e6))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 1e6))
}

func BenchmarkSetInsert(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 0))
}

func BenchmarkSetErase(b *testing.B) {
	b.Run("setErase", runWith(benchmarkSetErase, 0))
}

func BenchmarkSetFind(b *testing.B) {
	b.Run("setFind", runWith(benchmarkSetFind, 0))
}

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}

func benchmarkSetInsert(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(int(0), rbtree.CompareInt)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = set.Insert(key)
	}
}

func benchmarkSetNoescapeInsert(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(int(0), rbtree.CompareInt)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = set.Insert(rbtree.NoescapeInterface(key))
	}
}

func benchmarkSetErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(int(0), rbtree.CompareInt)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = set.Erase(key)
	}
}

func benchmarkSet(b *testing.B, m int) {
	var keys = make([]int, m)
	var rand = benchRand
	for i := 0; i < m; i++ {
		keys[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var set = rbtree.NewSet(int(0), rbtree.CompareInt)
		for j := 0; j < m; j++ {
			_, _ = set.Insert(rbtree.NoescapeInterface(keys[j]))
		}
		for j := 0; j < m; j++ {
			_ = set.Erase(keys[j])
		}
	}
}

func benchmarkSetFind(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(int(0), rbtree.CompareInt)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = set.Find(key)
	}
}

func benchmarkSetFindM(b *testing.B, m int) {
	var set = rbtree.NewSet(int(0), rbtree.CompareInt)
	var keys = make([]int, m)
	var rand = benchRand
	for i := 0; i < m; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Find(keys[i%m])
	}
}

func benchmarkMapInsert(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = rbtree.NewMap(int(0), false, rbtree.CompareInt)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = mp.Insert(key, true)
	}
}

func benchmarkMapErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = rbtree.NewMap(int(0), false, rbtree.CompareInt)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(keys[i], true)
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = mp.Erase(key)
	}
}

func benchmarkMap(b *testing.B, m int) {
	var mp = rbtree.NewMap(int(0), false, rbtree.CompareInt)
	var keys = make([]int, m)
	var rand = benchRand
	for i := 0; i < m; i++ {
		keys[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < m; j++ {
			_, _ = mp.Insert(rbtree.NoescapeInterface(keys[j]), rbtree.NoescapeInterface(true))
		}
		for j := 0; j < m; j++ {
			_ = mp.Erase(keys[j])
		}
	}
}

func benchmarkMapFind(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = rbtree.NewMap(int(0), false, rbtree.CompareInt)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(keys[i], true)
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = mp.Find(key)
	}
}

func benchmarkMapFindM(b *testing.B, m int) {
	var mp = rbtree.NewMap(int(0), false, rbtree.CompareInt)
	var keys = make([]int, m)
	var rand = benchRand
	for i := 0; i < m; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(keys[i], true)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mp.Find(keys[i%m])
	}
}

func benchmarkHashMapInsert(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = make(map[int]bool)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		mp[key] = true
	}
}

func benchmarkHashMapErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = make(map[int]bool)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		delete(mp, key)
	}
}

func benchmarkHashMap(b *testing.B, m int) {
	var keys = make([]int, m)
	var rand = benchRand
	for i := 0; i < m; i++ {
		keys[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var mp = make(map[int]bool)
		for j := 0; j < m; j++ {
			mp[keys[j]] = true
		}
		for j := 0; j < m; j++ {
			delete(mp, keys[j])
		}
	}
}

func benchmarkHashMapFind(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = make(map[int]bool)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_, _ = mp[key]
	}
}

func benchmarkHashMapFindM(b *testing.B, m int) {
	var mp = make(map[int]bool)
	var keys = make([]int, m)
	var rand = benchRand
	for i := 0; i < m; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = mp[keys[i%m]]
	}
}
