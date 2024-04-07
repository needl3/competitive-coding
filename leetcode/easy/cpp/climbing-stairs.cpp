class Solution{
    int memoize[50];
    int recurse(int n){
        if(n < 0) return 0;
        else if(n == 0) return 1;
        else{
            if(memoize[n] != -1) return memoize[n];
            else{
                int tempWays = recurse(n-1) + recurse(n-2);
                memoize[n] = tempWays;
                return tempWays;
            }
        }
    }
public:
    int climbStairs(int n) {
      for(int i=0;i<50;i++){memoize[i] = -1;}
      return recurse(n);
    }
};
