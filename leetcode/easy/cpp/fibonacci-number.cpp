class Solution {
public:
    int fib(int n) {
        if(n==0) return 0;
        int a[2] = {1,1};
        for(int i=3;i<=n;i++){
            a[(i+1)%2] = a[0] + a[1];
        }
        return a[(n+1)%2];
		}
};
