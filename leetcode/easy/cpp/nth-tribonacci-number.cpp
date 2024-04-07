class Solution {
public:
    int tribonacci(int n) {
        if(n==0) return 0;
        int a[3] = {1,1,2};
        for(int i=3;i<n;i++){
            a[i%3] = a[0]+a[1]+a[2];
        }
        return a[(n-1)%3];
    }
};
