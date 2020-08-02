package testhash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash                // hash函数
	replicas int                 // 每个节点对应的虚拟节点个数
	keys     []uint32            // 有序的列表，从小到大排列,哈希环的数据结构
	hashMap  map[uint32]string   // 虚拟节点-》真实节点
	nodeMap  map[string][]uint32 // 真实节点-》虚拟节点
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[uint32]string),
		nodeMap:  make(map[string][]uint32),
	}

	if fn == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

func (m *Map) AddNodes(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key))
			m.keys = append(m.keys, hash)
			// 虚拟节点-》真实节点
			m.hashMap[hash] = key
			// 真实节点-》虚拟节点
			m.nodeMap[key] = append(m.nodeMap[key], hash)
		}
	}
	sort.Slice(m.keys, func(i, j int) bool {
		return m.keys[i] < m.keys[j]
	})
}

func (m *Map) IsEmpty() bool {
	if len(m.keys) == 0 {
		return true
	}
	return false
}

func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	hash := m.hash([]byte(key))
	// Search函数，函数返回true时，返回对应的i，i在[0,n)之间，另外，如果返回值为i,代表f(i+1)也为true，之后所有的i，f(i)都返回true。如果没有找到，则返回n，而不是-1
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

func (m *Map) AddNode(key string) {
	for i := 0; i < m.replicas; i++ {
		hash := m.hash([]byte(strconv.Itoa(i) + key))
		m.keys = append(m.keys, hash)
		// 虚拟节点-》真实节点
		m.hashMap[hash] = key
		// 真实节点-》虚拟节点
		m.nodeMap[key] = append(m.nodeMap[key], hash)
	}
	sort.Slice(m.keys, func(i, j int) bool {
		return m.keys[i] < m.keys[j]
	})
}

func (m *Map) DeleteNode(key string) {
	count := 0
	for i, _ := range m.keys {
		if m.hashMap[m.keys[i]] == key {
			delete(m.hashMap, m.keys[i])
			count++
		} else {
			m.keys[i-count] = m.keys[i]
		}
	}
	if count != m.replicas {
		fmt.Println("删除个数错误")
	}
	delete(m.nodeMap, key)
	m.keys = m.keys[0 : len(m.keys)-count]
}

func (m *Map) List() {
	if m.IsEmpty() {
		fmt.Println("该系统还没有节点，用add命令添加节点")
	} else {
		fmt.Println("节点：")
		fmt.Print("[ ")
		for node := range m.nodeMap {
			fmt.Printf("%s ", node)
		}
		fmt.Printf("]\n")
	}
}
