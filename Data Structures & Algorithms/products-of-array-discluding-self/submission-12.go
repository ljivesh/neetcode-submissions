func productExceptSelf(nums []int) []int {

    prefix := 1

    output := make([]int, 0)

    for i:=0; i<len(nums); i++ {
        
        output = append(output, prefix)

        prefix = nums[i]*prefix

    }

    suffix := 1;

    for i:=len(nums)-1; i>=0; i-- {

        output[i] = output[i] * suffix

        suffix = suffix * nums[i]        

    }

    return output

}
