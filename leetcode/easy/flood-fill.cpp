#include <bits/stdc++.h>
#include <unordered_set>
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
public:
    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int color) {
    	queue<vi> q;
		set<vector<int>> s;
		q.push({sr,sc});
		int x = sr;
		int y = sc;
		while(!q.empty()){
			x = q.front()[0];
			y = q.front()[1];
			q.pop();

			if(x<0 || x>=image.size() || y<0 || y>=image[0].size()) continue;

			s.insert({x,y});

			if(x+1 < image.size() && image[x+1][y] == image[x][y] && s.find({x+1,y}) == s.end())
				q.push({x+1,y});

			if(x-1 >= 0 && image[x-1][y] == image[x][y] && s.find({x-1,y}) == s.end())
				q.push({x-1,y});

			if(y-1 >= 0 && image[x][y-1] == image[x][y] && s.find({x,y-1}) == s.end())
				q.push({x,y-1});

			if(y+1 < image[0].size() && image[x][y+1] == image[x][y] && s.find({x,y+1}) == s.end())
				q.push({x,y+1});
			image[x][y] = color;
		}
		return image;
    }
};
