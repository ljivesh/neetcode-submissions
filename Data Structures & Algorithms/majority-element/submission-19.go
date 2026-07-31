func majorityElement(nums []int) int {
    
	candidate:=nums[0]
	counter:=0

	for _, elem := range nums {
		
		if elem == candidate {
			counter++
		}  else {
			counter--
		}

		if counter == 0 {
			candidate = elem
			counter++
		}

	}

	return candidate

}
