type MyHashSet[T comparable] struct {
    mp map[T]struct{}
}

func Constructor() MyHashSet[int] {
    return MyHashSet[int]{mp: make(map[int]struct{})}
}

func (this *MyHashSet[T]) Add(key T) {
    this.mp[key] = struct{}{}
}

func (this *MyHashSet[T]) Remove(key T) {
    delete(this.mp, key)
}

func (this *MyHashSet[T])Contains(key T) bool {
    _, exists := this.mp[key]

    return exists

}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 