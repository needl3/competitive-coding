class Solution {
    public:
        vector<int> twoSum(vector<int>& numbers, int target) {
            std::map<int, int> finalKeys;
            for(int i=0;i<numbers.size();i++){
                if(finalKeys.find(target-numbers[i]) != finalKeys.end()){
                    numbers = {finalKeys[target-numbers[i]]+1, i+1};
                    return numbers;
                }else   finalKeys[numbers[i]] = i;
            }
            return {};
        }
};
