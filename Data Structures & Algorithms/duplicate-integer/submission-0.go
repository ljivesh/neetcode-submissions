func hasDuplicate(nums []int) bool {
    
    hasDuplicate:= false;

    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            if(nums[i] == nums[j]) {
                hasDuplicate = true;
                break;
            }
        }

        if(hasDuplicate) {
         
            break;
        };
    }

    return hasDuplicate;

}
