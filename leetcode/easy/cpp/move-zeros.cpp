class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int l=0;
        while(l < nums.size() && nums[l] != 0) l++;
        if(l >= nums.size()) return;
        for(int r=l+1; r<nums.size();r++){
            if(nums[r] != 0 && nums[l] == 0){
                swap(nums[r], nums[l]);
                l++;
            }
        }
    }
};
