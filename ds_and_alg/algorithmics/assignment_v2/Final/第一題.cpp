#include<iostream>
#include<cmath>
#include<string>
#include<map>
#include<queue>
using namespace std;
#define N 3
#define N2 9

struct Puzzle {
	int f[N2];
	int space;
	string path;
	
	bool operator < (const Puzzle &p) const {
		for(int i = 0; i < N2; i++) {
			if(f[i] == p.f[i]) continue;
			return f[i] > p.f[i];
		}
		return false;
	}
};

static const int dx[4] = {-1, 0, 1, 0};
static const int dy[4] = {0, 1, 0, -1};
static const char dir[4] = {'u', 'l', 'd', 'r'};

bool isTarget(Puzzle p,Puzzle t) {
	for(int i = 1; i < N2+1; i++){
		if(p.f[i-1]==9)
 			cout<<0;
		else
			cout<<p.f[i-1];
		if(i%3==0)
			cout<<endl;
		else
                	cout<<" ";
	}
	
	for(int i = 0; i < N2; i++)
		if(p.f[i] != t.f[i] ) return false;
	return true;
}

string bfs(Puzzle s,Puzzle t) {
	queue<Puzzle> Q;
	map<Puzzle, bool> V;
	Puzzle u, v;
	s.path = "";
	Q.push(s);
	V[s] = true;
	
	while(!Q.empty() ) {
		u = Q.front(); Q.pop();
		if(isTarget(u,t)) return u.path;
		int sx = u.space / N;
		int sy = u.space % N;
		for(int r = 0; r < 4; r++) {
			int tx = sx + dx[r];
			int ty = sy + dy[r];
			if(tx < 0 || ty < 0 || tx >= N || ty >= N) continue;
			v = u;
			swap(v.f[u.space], v.f[tx*N + ty]);
			v.space = tx * N + ty;
			if(!V[v]) {
				V[v] = true;
				v.path += dir[r];
				Q.push(v);
			}
		}
	}
	
	return "unsolvable";
}

int main(){
	Puzzle in,out;
	
	for(int i = 0; i < N2; i++){
		cin>>in.f[i];
		if(in.f[i] == 0) {
			in.f[i] = N2; //set space
			in.space = i;
		}
	}
	for(int i = 0; i < N2; i++){
		cin>>out.f[i];
		if(out.f[i] == 0) {
			out.f[i] = N2; //set space
			out.space = i;
		}
	}
	string ans = bfs(in,out);
	//cout<<ans.size()<<endl;
	
	return 0;
}