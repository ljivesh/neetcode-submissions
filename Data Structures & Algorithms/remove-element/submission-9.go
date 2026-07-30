func removeElement(nums []int, val int) int {
    back:= len(nums) - 1
	i:=0

	for i<=back {

		if nums[i] == val {
			nums[i] = nums[back]
			back--
		} else {
			i++
		}

	}

	return back+1
}
