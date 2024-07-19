package hashmap

import "slices"

// 200000 以内倍增的质数列表. 作为演示代码足够了
var primes = [...]uintptr{7, 17, 37, 79, 163, 331, 673, 1361, 2729, 5471, 10949, 21911, 43853, 87719, 175447}

// nextPrime 在 primes 里寻找下一个质数.
// 用质数做桶数量, 可以最大程度减少哈希冲突的概率, 使哈希后的数据分布更加均匀
func nextPrime(n uintptr) uintptr {
	i, _ := slices.BinarySearch(primes[:], n)
	return primes[i]
}

// 最大载荷因子的 10 倍
const maxLoadFactor10 = 8

// KeyValue 是一个键值对
type KeyValue[K comparable, V any] struct {
	Key   K // 键
	Value V // 值
}

// bucket 是哈希表的一个桶
type bucket[K comparable, V any] []*KeyValue[K, V]

// HashMap 是一个演示用途的哈希表实现
type HashMap[K comparable, V any] struct {
	hashFunc func(K) uintptr // 哈希函数 f(key) -> hash
	buckets  []bucket[K, V]  // 桶
}

// New 创建一个以 K-V 为键值对的哈希表
// hash 参数用以计算 K 的哈希值
func New[K comparable, V any](hash func(K) uintptr) *HashMap[K, V] {
	return &HashMap[K, V]{
		hashFunc: hash,
		buckets:  make([]bucket[K, V], primes[0]),
	}
}

// bucketIndex 在 m 中寻找 key 对应的桶
// b 为桶指针. 如果找到 i 为 key 在 b 中的索引, 否则 i 为 -1
func (m *HashMap[K, V]) bucketIndex(key K) (b *bucket[K, V], i int) {
	index := m.hashFunc(key) % uintptr(len(m.buckets))
	b = &m.buckets[index]
	i = slices.IndexFunc(*b, func(kv *KeyValue[K, V]) bool { return kv.Key == key })
	return
}

// Get 在 m 中查找 key
// 用 Go map 语法表示为 value, ok = m[key]
func (m *HashMap[K, V]) Get(key K) (value V, ok bool) {
	b, i := m.bucketIndex(key)
	if ok = b != nil; ok {
		value = (*b)[i].Value
	}
	return
}

// Len 返回 m 的长度
// 用 Go map 语法表示为 len(m)
func (m *HashMap[K, V]) Len() (n int) {
	for _, bucket := range m.buckets {
		n += len(bucket)
	}
	return
}

// grow 为 m 扩容
func (m *HashMap[K, V]) grow() {
	// 下一个不小于 当前桶个数*2 的质数
	n := nextPrime(uintptr(len(m.buckets) * 2))
	if n == uintptr(len(m.buckets)) {
		return // 没有更大的质数了
	}
	old := *m                                // 备份当前桶
	m.buckets = make([]bucket[K, V], n)      // 分配更多桶
	old.Range(func(kv KeyValue[K, V]) bool { // 把旧值逐个搬进新桶
		m.Set(kv.Key, kv.Value)
		return true
	})
}

// Set 设置 key 对应的 value
// 用 Go map 语法表示为 m[key] = value
func (m *HashMap[K, V]) Set(key K, value V) {
	if m.Len()*10/len(m.buckets) >= maxLoadFactor10 {
		m.grow() // 扩容
	}
	b, i := m.bucketIndex(key)
	if i != -1 {
		// 覆盖当前值
		(*b)[i].Value = value
	} else {
		// 在桶 b 中添加新的键值对
		(*b) = append((*b), &KeyValue[K, V]{Key: key, Value: value})
	}
}

// Delete 删除 m 中 key 对应的键值对
// 用 Go map 语法表示为 delete(m, key)
func (m *HashMap[K, V]) Delete(key K) {
	b, i := m.bucketIndex(key)
	if i == -1 {
		return // 没有找到
	}
	// 从桶 b 中删除 i 处的元素
	*b = slices.Delete(*b, i, i+1)
}

// Range 遍历 m 中的所有键值对
// 用 Go map 语法表示为
// for k, v := range m { if !f(k, v) { break } }
func (m *HashMap[K, V]) Range(f func(KeyValue[K, V]) bool) {
	for _, bucket := range m.buckets {
		for _, kv := range bucket {
			if !f(*kv) {
				return
			}
		}
	}
}
