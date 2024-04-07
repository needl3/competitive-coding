class Solution {
	void reverse(vector<int>& v, int start, int end){
		while(start <= end){
			std::swap(v[start], v[end]);
			start++;
			end--;
		}
	}
public:
    void rotate(vector<int>& nums, int k) {
        k = k % nums.size();
		reverse(nums, 0, nums.size()-1);
		reverse(nums, 0, k-1);
		reverse(nums, k, nums.size()-1);
    }
};
