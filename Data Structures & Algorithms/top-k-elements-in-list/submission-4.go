func topKFrequent(nums []int, k int) []int {

	var frequencyMap = make(map[int]int);

	for i:=0; i<len(nums); i++ {
		frequencyMap[nums[i]]++;
	}

	var bucket = make([][]int, len(nums)+1);

	for key, value:= range frequencyMap {
		bucket[value] = append(bucket[value], key);
	}


	var output = make([]int, 0);

	for i:= len(bucket)-1; i>=0; i-- {


		if len(bucket[i]) == 0 {
			continue;
		}
		

		output = append(output, bucket[i]...);
		



		if(len(output)==k) {
			break;
		}

	}

	return output;
	
}
