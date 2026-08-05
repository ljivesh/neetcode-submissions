
type MyHashMapValue struct {
    key int
    val int
}


type MyHashMap struct {
    buckets [][]MyHashMapValue
    entries int
}


func isPrime (n int) bool {

    if n<=1 {
        return false
    }

    if n == 2 {
        return true
    }

    if n%2 == 0 {
        return false
    }

    for i:=3; i*i <= n; i+=2 {
        if n%i ==0 {
            return false
        }
    }

    return true    

}

func findNearestPrime(n int) int {
    
    if isPrime(n) {
        return n
    }

    offset := 1
    lower := n-offset
    upper := n+offset 

    
    for {

        if isPrime(lower) {
            return lower
        }

        if isPrime(upper) {
            return upper
        }

        lower -= offset
        upper +=offset
    }

}  

func Constructor() MyHashMap {
    return MyHashMap{buckets: make([][]MyHashMapValue, 769), entries:0} 
}

func (hm *MyHashMap) hash(x int) int {
    return x%len(hm.buckets)
}

func (hm *MyHashMap) reValidateMap() {

    currentLength := len(hm.buckets)

    loadFactor := float64(hm.entries) / float64(currentLength)

    if loadFactor >= 0.75 {

        newLength := findNearestPrime(2*currentLength)

        newMap := MyHashMap{buckets: make([][]MyHashMapValue, newLength), entries:0} 

        for _, bucket := range hm.buckets {

            for _, data := range  bucket {
             
                newMap.Put(data.key, data.val)  
                
            }

        }

        *hm = newMap
    }
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
    hm.entries++


    hm.reValidateMap()

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
            hm.entries--
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