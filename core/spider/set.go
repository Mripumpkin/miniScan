package spider

import (
	"strings"
	"sync"
)

// StrSet 定义字符串集合的结构体
type StrSet struct {
	mu    sync.RWMutex // 读写锁，用于保护并发访问
	items map[string]struct{}
}

// NewStrSet 创建一个新的字符串集合实例
func NewStrSet() *StrSet {
	return &StrSet{
		items: make(map[string]struct{}),
	}
}

// Add 添加元素到集合中
func (s *StrSet) Add(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item] = struct{}{} // 空结构体节省内存
}

// Remove 从集合中移除元素
func (s *StrSet) Remove(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, item)
}

// Contains 检查集合中是否包含某个元素
func (s *StrSet) Contains(item string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.items[item]
	return exists
}

// Size 返回集合的大小
func (s *StrSet) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// Slice 返回集合中的所有元素（以切片形式）
func (s *StrSet) Slice() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]string, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// Clear 清空集合
func (s *StrSet) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[string]struct{})
}

func (set *StrSet) ContainsI(item string) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.items {
		if strings.EqualFold(k, item) {
			return true
		}
	}
	return false
}
