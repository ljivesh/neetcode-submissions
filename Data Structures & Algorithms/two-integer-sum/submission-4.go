func twoSum(nums []int, target int) []int {

    var numMap = make(map[int]int);

    for i:=0; i<len(nums); i++ {

        if prevIdx, ok:= numMap[nums[i]]; ok {

            return []int {prevIdx, i};
        }

        numMap[target - nums[i]] = i;

    }

    return []int {0, 0};      
}
