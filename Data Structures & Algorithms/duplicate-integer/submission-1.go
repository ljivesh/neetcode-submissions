func hasDuplicate(nums []int) bool {
    
   seenMap:= make(map[int]bool);

   for i:= 0; i<len(nums); i++ {
        if(seenMap[nums[i]]) {
            return true
        }

        seenMap[nums[i]] = true;
   }
    
    return false;

}
