type MyHashSet struct {
    buckets [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{
        buckets: make([][]int, 1000),
    }
}

func (s *MyHashSet) hash(key int) int {
    return key % len(s.buckets)
}

func (s *MyHashSet) Add(key int) {
    idx := s.hash(key)

    for _, v := range s.buckets[idx] {
        if v == key {
            return
        }
    }

    s.buckets[idx] = append(s.buckets[idx], key)
}

func (s *MyHashSet) Remove(key int) {
    idx := s.hash(key)

    bucket := s.buckets[idx]

    for i, v := range bucket {
        if v == key {
            s.buckets[idx] = append(bucket[:i], bucket[i+1:]...)
            return
        }
    }
}

func (s *MyHashSet) Contains(key int) bool {
    idx := s.hash(key)

    for _, v := range s.buckets[idx] {
        if v == key {
            return true
        }
    }

    return false
}