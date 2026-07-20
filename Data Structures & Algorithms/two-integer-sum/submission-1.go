func twoSum(nums []int, target int) []int {

    var numMap = make(map[int]int);

    for i:=0; i<len(nums); i++ {

        if numMap[target - nums[i]] != 0 {

            return []int {numMap[target - nums[i]] - 1, i};
        }

        numMap[nums[i]] = i+1;

    }

    return []int {0, 0};      
}
