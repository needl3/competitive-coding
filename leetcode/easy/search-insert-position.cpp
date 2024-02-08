class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int l=0, r = nums.size()-1, m;
        while(l<r){
            m = l + (r-l)/2;
            if(nums[m] < target) l = m+1;
            else if(nums[m] > target) r = m-1;
            else return m;
        }
        if(nums[l] < target) return l+1;
        else return l;
    }
};
