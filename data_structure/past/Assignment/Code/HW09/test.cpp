#include <iostream>
// #include <bits/stdc++.h>
#include <cstring>
using namespace std;

// number of vertices in graph
#define V 4


// void adjacency_matrix()
// {
//   for (int i = 0; i < 5; ++i)
//     for (int j = 0; j < 5; ++j)
//       graph[i][j] = 0;

//   int a, b, w; // 一條邊的端點、另一個端點、邊的權重
//   while (cin >> a >> b >> w)
//     graph[a][b] = w;
// }


// create a 2d array of size 7x7
// for adjacency matrix to represent graph

int main()
{
  // create a 2d array of size 7x7
  // for adjacency matrix to represent graph
  // int G[V][V] = {
  //     {0, 28, 0, 0, 0, 10, 0},
  //     {28, 0, 16, 0, 0, 0, 14},
  //     {0, 16, 0, 12, 0, 0, 0},
  //     {0, 0, 12, 22, 0, 18},
  //     {0, 0, 0, 22, 0, 25, 24},
  //     {10, 0, 0, 0, 25, 0, 0},
  //     {0, 14, 0, 18, 24, 0, 0}};
  int G[V][V] = {
      {0, 3, 3, 1},
      {0, 0, 0, 1},
      {0, 0, 0, 0},
      {0, 0, 0, 0}};

  int edge; // number of edge

  // create an array to check visited vertex
  int visit[V];

  // initialise the visit array to false
  for (int i = 0; i < V; i++)
  {
    visit[i] = false;
  }

  // set number of edge to 0
  edge = 0;

  // the number of edges in minimum spanning tree will be
  // always less than (V -1), where V is the number of vertices in
  // graph

  // choose 0th vertex and make it true
  visit[0] = true;

  int x; //  row number
  int y; //  col number

  // print for edge and weight
  cout << "Edge"
       << " : "
       << "Weight";
  cout << endl;
  int count = 0;
  while (edge < V - 1)
  { // in spanning tree consist the V-1 number of edges

    // For every vertex in the set S, find the all adjacent vertices
    //  , calculate the distance from the vertex selected.
    //  if the vertex is already visited, discard it otherwise
    // choose another vertex nearest to selected vertex.

    int min = INT_MAX;
    x = 0;
    y = 0;

    for (int i = 0; i < V; i++)
    {
      if (visit[i])
      {
        for (int j = 0; j < V; j++)
        {
          if (!visit[j] && G[i][j])
          { // not in selected and there is an edge
            if (min > G[i][j])
            {
              min = G[i][j];
              x = i;
              y = j;
            }
          }
        }
      }
    }
    cout << x << " ---> " << y << " :  " << G[x][y];
    count += G[x][y];
    cout << endl;
    visit[y] = true;
    edge++;
  }

  cout << count << endl;

  return 0;
}