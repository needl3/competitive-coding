#include <bits/stdc++.h>
using namespace std;
#define fo(i,n)             for (int i=0; i<n; i++)
#define Fo(i,j,n)           for (int i=j ; i<n ; i++)
#define rfo(i,n)            for (int i=0 ; i>=n ; i--)
#define rFo(i,j,n)          for (int i=j; i>=n; i--)
#define foit(it, x)         for (auto it = x.begin(); it != x.end(); it++)
#define foa(x, a)      for (auto x: a)
#define deb(x)              cout << #x << ": " << x << std::endl;
#define debb(x, y)          cout << #x << ": " << x << "  " << #y << ": " << y << std::endl;
#define debarr(x)           cout << #x << "[ "; for(auto a: x) cout << a << ","; cout << " ]" << endl;
#define ci(x)               cin >> x;
#define ca(a, n)            fo(i, n){int x; cin >> x; a.push_back(x);}
#define ll                  long long
#define all(obj)            obj.begin(), obj.end()
typedef pair<int, int>      pii;
typedef pair<ll, ll>        pll;
typedef vector<int>         vi;
typedef vector<ll>          vl;
typedef vector<vi>          vvi;
typedef vector<vl>          vvl;
typedef vector<pii>         vpii;
typedef vector<string>      vs;
typedef map<int, int>       mii;
typedef map<int, long>      mil;

string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &, char);
class Solution {
    vector<string> chars = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
public:
    vector<string> letterCombinations(string digits, int i=0) {
       if(i >= digits.size()) return {};
	   vs v1;
	   for(auto x: chars[digits[i]-48]) v1.push_back(string()+x);
	   vs v2 = letterCombinations(digits, i+1);
	   vs v3;
	   for(auto x: v1){
		   for(auto y: v2){
			   v3.push_back(string()+x+y);
		   }
	   }
	   if(v2.size()==0) v3 =v1;
	   return v3;
    }
};
