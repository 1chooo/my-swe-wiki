#include <functional>
#include <iostream>
#include <stack>
#include <string>
#include <utility>
#include <vector>

struct Node {
  int r, c;    // rc 語音
  Node() = default;
  Node(int r, int c)
      : r{ r }, c{ c } {};
  Node &operator=(const Node &other)
  {
    this->r = other.r, this->c = other.c;
    return *this;
  }
};

namespace alg {
  constexpr std::pair<int, int> Move[4]{ { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };

  namespace State {
    enum __State : uint32_t {
      GROUND = 0,
      WALL = 1,
      DONE = 2,
      EXPLORE = 3
    };
  }

  auto dfs = [](std::vector<std::vector<int>> &maze, std::stack<Node> &tree, std::stack<Node> &answ, const int &end_y, const int &end_x) -> bool {
    bool state_flag{ true }, goal_flag{ false };
    //std::cout << "goal: (" << end_y << "," << end_x << ")\n";
    while (!tree.empty()) {
      state_flag = true;
      const auto temp = tree.top();
      //std::cout << "temp: (" << temp.r << "," << temp.c << ")\n";
      for (const auto &dir : alg::Move) {
        //std::cout << '(' << temp.r + dir.first << "," << temp.c + dir.second << ")\n";
        if (maze[temp.r + dir.first][temp.c + dir.second] == State::GROUND) {
          maze[temp.r][temp.c] = State::EXPLORE;
          //std::cout << "emplace: (" << temp.r + dir.first << "," << temp.c + dir.second << ")\n";
          tree.emplace(Node(temp.r + dir.first, temp.c + dir.second));
          state_flag = false;

          if (std::equal_to<int>{}(temp.r + dir.first, end_y) && std::equal_to<int>{}(temp.c + dir.second, end_x))    // get goal
            goal_flag = true;

          break;
        }
      }

      if (state_flag) {
        maze[temp.r][temp.c] = State::EXPLORE;
        tree.pop();
      }

      if (goal_flag) {
        while (!tree.empty()) {
          answ.push(std::move(tree.top()));
          tree.pop();
        }

        break;
      }
    }

    return goal_flag;
  };

}    // namespace alg

int main()
{
  int row, col;
  std::cin >> row >> col;
  constexpr int start_x{ 1 }, start_y{ 1 };
  std::vector<std::vector<int>> maze(row + 2, std::vector<int>(col + 2, 1));

  for (int i{ 1 }; i < row + 1; ++i) {
    for (int j{ 1 }; j < col + 1; ++j)
      std::cin >> maze[i][j];
  }

  std::stack<Node> tree, answ;
  tree.emplace(Node{ start_y, start_x });    // start point
  bool state_flag{ true }, goal_flag{ false };

  if (!alg::dfs(maze, tree, answ, row, col))
    puts("Can't reach the exit!");
  else {
    while (!answ.empty()) {
      std::cout << '(' << answ.top().r - 1 << ',' << answ.top().c - 1 << ") ";
      answ.pop();
    }
  }

  //for ( const auto &x : maze ) {
  //    for ( const auto &e : x ) std::cout << e << " ";
  //    puts( "" );
  //}

  return 0;
}