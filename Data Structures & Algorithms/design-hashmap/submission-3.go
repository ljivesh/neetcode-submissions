type MyHashMapValue struct {
    key int
    val int
}


type MyHashMap struct {
    buckets [][]MyHashMapValue
}

func Constructor() MyHashMap {
    return MyHashMap{buckets: make([][]MyHashMapValue, 769)} 
}

func (hm *MyHashMap) hash(x int) int {
    return x%len(hm.buckets)
}

func (hm *MyHashMap) Put(key int, value int) {
    
    hashedIdx := hm.hash(key)

    bucket := hm.buckets[hashedIdx]

    for idx, data := range bucket {
        if data.key == key {
            bucket[idx].val = value 
            return
        }
    }

    hm.buckets[hashedIdx] = append(bucket, MyHashMapValue{key: key, val: value})

}

func (hm *MyHashMap) Get(key int) int {
    hashedIdx := hm.hash(key)

    bucket := hm.buckets[hashedIdx]
    
    for _, data := range bucket {
        if data.key == key {
            return data.val
        }
    }

    return -1
}

func (hm *MyHashMap) Remove(key int) {
    hashedIdx := hm.hash(key)

    bucket := hm.buckets[hashedIdx]

    for idx, data := range bucket {
        if data.key == key {
            hm.buckets[hashedIdx] = append(bucket[:idx], bucket[idx+1:]...)
            return
        }
    }

    return
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */