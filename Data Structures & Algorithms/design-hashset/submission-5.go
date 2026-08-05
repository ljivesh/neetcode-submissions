type MyHashSet struct {
    buckets [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{buckets: make([][]int, 100)}
}

func (this *MyHashSet) getHashKey(x int) int {
    return x%len(this.buckets) 
}

func (this *MyHashSet) Add(key int) {
    contains := this.Contains(key)

    if contains {
        return
    }

    hash:=this.getHashKey(key)
    bucket := this.buckets[hash]

    this.buckets[hash] = append(bucket, key)

}

func (this *MyHashSet) Remove(key int) {
    
    contains := this.Contains(key)

    if !contains {
        return
    }

    hash:=this.getHashKey(key)

    bucket := this.buckets[hash]

    pos :=0

    for i, _ := range bucket {
        if bucket[i] == key {
            pos = i
            break
        }
    }

    bucket = append(bucket[:pos], bucket[pos+1:]...)

    this.buckets[hash] = bucket   

}

func (this *MyHashSet) Contains(key int) bool {
    
    hash := this.getHashKey(key)

    bucket := this.buckets[hash]

    for _, elem := range bucket {
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
 