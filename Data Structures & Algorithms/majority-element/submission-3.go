func majorityElement(nums []int) int {
	
	frequencyMap := make(map[int]int)


	for _,elem := range nums {
		frequencyMap[elem]++
	}

	majorityElement := 0

	for key, value := range frequencyMap {
		if value >= (len(nums)/2)+1 {
			majorityElement = key
		}

	}

	return majorityElement

	
}
