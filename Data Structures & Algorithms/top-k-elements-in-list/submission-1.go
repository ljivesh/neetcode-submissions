func topKFrequent(nums []int, k int) []int {

	var frequencyMap = make(map[int]int);

	for i:=0; i<len(nums); i++ {
		frequencyMap[nums[i]]++;
	}

	type kv struct {
		key int
		val int
	}

	var structArray []kv;

	for key, val:= range frequencyMap {
		structArray = append(structArray, kv{key, val});
	}

	sort.Slice(structArray, func (i int, j int) bool {
		return structArray[i].val > structArray[j].val;
	});

	var output = make([]int, 0)

	for _, val:= range structArray {
		output = append(output, val.key);
	}

	return output[:k];
	
}
