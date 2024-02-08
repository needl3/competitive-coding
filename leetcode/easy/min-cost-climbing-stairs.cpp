class Solution{
    int *m;
    int dp(vector<int>& nums, int i){
        if(i >= nums.size()) return 0;

        if(m[i] != -1) return m[i];

        int cost1 = dp(nums, i+1);
		if(i+1 >= nums.size()) cost1 = 1000;
        int cost2 = dp(nums, i+2);

        m[i] = min(cost1, cost2) + nums[i];

        return m[i];
    }
public:
    int minCostClimbingStairs(vector<int>& cost) {
        m = new int[cost.size()];
        for(int i=0;i<cost.size();i++) m[i] = -1;
        return min(dp(cost, 0), dp(cost, 1));
    }
};
