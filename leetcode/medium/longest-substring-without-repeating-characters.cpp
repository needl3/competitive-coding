class Solution {
    map<char, int> m;
public:
    int lengthOfLongestSubstring(string s) {
        int l=0;
        int mx=0;
        for(int r=0;r<s.size();r++){
            if(m.find(s[r]) != m.end()){
                if(s[r] == s[l]){
                    if(mx < r-l) mx = r-l;
                    m[s[r]] = r;
                    l++;
                }else{
                    if(mx < r-m[s[r]]) mx = r-m[s[r]];
                    while(l < m[s[r]]){
                        m.erase(s[l]);
                        l++;
                    }
                    m[s[r]] = r;
                    l++;
                }
            }else{
                m[s[r]] = r;
                if(r-l+1 > mx) mx = r-l+1;
            }
        }
        return mx;
    }
};
