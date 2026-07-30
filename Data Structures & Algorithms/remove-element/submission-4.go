func removeElement(nums []int, val int) int {
    
	nonValCount:=0

	backPointer:= len(nums)-1
	
	for i:=0; i<len(nums); i++ {

		if i > backPointer {
			break
		} 
		
		if nums[i] == val {

			if nums[backPointer] == val {

				for i<backPointer && nums[backPointer] == val  {
					nums[backPointer] = 0
					backPointer--


				}


			}



				nums[i] = nums[backPointer]
				nums[backPointer] = 0 
				backPointer--


			

		}

		if backPointer < i {
			break
		}

		nonValCount++

	}
	
	return nonValCount

}
