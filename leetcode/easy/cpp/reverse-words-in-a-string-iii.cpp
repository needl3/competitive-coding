class Solution {
    void reverse(string &s, int l, int r){
        char temp;
        while(l<r){
            temp = s[l];
            s[l] = s[r];
            s[r] = temp;
            l++;
            r--;
        }
    }
public:
    string reverseWords(string s) {
        int start = 0;
        for(int end=0;end<s.size();end++){
            if(s[end]==' ' || end == s.size()-1){
                if(end == s.size()-1)   end++;
                reverse(s,start,end-1);
                start = end+1;
            }
        }
        return s;
    }
};
