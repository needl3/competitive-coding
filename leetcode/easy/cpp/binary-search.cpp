#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int search(vector<int> &nums, int target) {
    int l = 0, r = nums.size() - 1, m;
    while (l <= r) {
      m = l + (r - l) / 2;
      if (target < nums[m])
        r = m - 1;
      else if (target > nums[m])
        l = m + 1;
      else
        return m;
    }
    return -1;
  }
};
