class Solution {
public:
    int majorityElement(vector<int>& nums) {
        
		int candidate = nums[0];
		int counter = 0;

		for(int i=0; i<nums.size(); i++) {

			if (counter == 0) {
				candidate = nums[i];
			}

			if (candidate == nums[i]) {
				counter++;
			}

			else {
				counter--;
			}

		}

		return candidate;


    }
};