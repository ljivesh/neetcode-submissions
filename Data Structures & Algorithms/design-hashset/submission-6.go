type MyHashSet struct {
    buckets [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{buckets: make([][]int, 769)}
}

func (this *MyHashSet) getHashedIdx(x int) int {
    return x%len(this.buckets) 
}

func (this *MyHashSet) Add(key int) {

    hash:=this.getHashedIdx(key)

    for _, elem := range this.buckets[hash] {
        if elem == key {
            return
        }
    }

    this.buckets[hash] = append(this.buckets[hash], key)

}

func (this *MyHashSet) Remove(key int) {
    
    hash:=this.getHashedIdx(key)

    for i, _ := range this.buckets[hash] {
        if this.buckets[hash][i] == key {
            this.buckets[hash] = append(this.buckets[hash][:i], this.buckets[hash][i+1:]...)
            return
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    
    hash := this.getHashedIdx(key)

    for _, elem := range this.buckets[hash] {
        if elem == key {
            return true
        }
    } 

    return false

}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 