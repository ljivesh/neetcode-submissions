func majorityElement(nums []int) int {
    
	candidate:=nums[0]
	counter:=0

	for _, elem := range nums {

		if counter == 0 {
			candidate = elem
			counter++

			continue
		}
		
		if elem == candidate {
			counter++
		}  else {
			counter--
		}



	}

	return candidate

}
