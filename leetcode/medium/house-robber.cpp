#include <bits/stdc++.h>
class Solution {
    map<int, int> memoized;
public:
int dp(vector<int>& nums, int cpt){
                if(cpt < 0) return 0;
                if(memoized.find(cpt) != memoized.end()) return memoized.at(cpt);
                
                int maxValue = max(dp(nums, cpt-1), dp(nums, cpt-2)+nums[cpt]);
                memoized[cpt] = maxValue;
                return maxValue;
        }
        int rob(vector<int>& nums){
                int m = dp(nums, nums.size()-1);
                return m;
        }
};
