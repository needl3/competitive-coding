
class Solution {
    int *mem;
	int dp(vector<int>& nums, int oi = 0){
        int i = oi;
        if(i >= nums.size()) return 0;

        // Get memoized part if exists
        if(mem[i] != -1) return mem[i];

		int m = 0;
		int o = nums[i];
		int t=0;
		while(i < nums.size() && nums[i] == o){
			m += o;
            i++;
		}

        if(i>=nums.size()) return m;

		if(nums[i]-1 == o){
			t = dp(nums, i);
			int r = nums[i];
			while(r == nums[i]) i++;
		}else t = 0;
		m += dp(nums, i);
        mem[oi] = max(t,m);
        return mem[oi];
	}
public:
    int deleteAndEarn(vector<int>& nums) {
        mem = new int[nums.size()];
        for(int i=0;i<nums.size();i++) mem[i] = -1;
		sort(nums.begin(), nums.end());
		return dp(nums);
    }
};
