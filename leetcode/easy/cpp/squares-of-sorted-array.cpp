class Solution {
public:
  vector<int> sortedSquares(vector<int> &nums) {
    vector<int> v(nums.size());
    int l = 0, r = nums.size() - 1, n = r;
    while (l <= r) {
      int ls = nums[l] * nums[l];
      int rs = nums[r] * nums[r];
      if (ls < rs) {
        v[n] = rs;
        r--;
      } else {
        v[n] = ls;
        l++;
      }
      n--;
    }
    return v;
  }
};
