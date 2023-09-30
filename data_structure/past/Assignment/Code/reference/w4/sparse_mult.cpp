#include <functional>
#include <iostream>
#include <limits>
#include <sstream>
#include <string>
#include <tuple>
#include <utility>
#include <vector>

struct Node {
  int row;
  int col;
  int val;
  Node()
      : row(), col(), val() {}
  Node(const int &row, const int &column, const int &value)
      : row(row), col(column), val(value) {}
  void set(const int &r, const int &c, const int &v)
  {
    row = r;
    col = c;
    val = v;
  }
  void set(Node other)
  {
    row = other.row;
    col = other.col;
    val = other.val;
  }
};

inline void get_row_col(int &row, int &column)
{
  std::cin >> row >> column;
  std::cin.clear();
  std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
}

inline void init_arr(std::string &str, std::stringstream &ss, const int &row, const int &column, std::vector<Node> &arr)
{
  int buf;
  arr.emplace_back(Node());
  for (int i{}; i < row; ++i) {
    int cnt{};
    getline(std::cin, str);
    ss << str;
    while (ss >> buf) {
      if (buf != 0)
        arr.emplace_back(Node(i, cnt, buf));
      ++cnt;
    }
    ss.str("");
    ss.clear();
  }
  arr[0].set(row, column, arr.size() - 1);
}

std::vector<Node> fast_transpose(const std::vector<Node> &arr)
{
  const int &max_col{ arr[0].col - 1 };
  std::vector<int> row_terms(max_col + 1), start_position(max_col + 1);
  std::vector<Node> answ(arr.size());
  answ[0].set(arr[0].col, arr[0].row, arr[0].val);
  if (arr[0].val > 0) {
    for (int i{ 1 }; i < arr[0].val + 1; ++i)
      ++row_terms[arr[i].col];

    start_position[0] = 1;
    for (int i{ 1 }; i < arr[0].col; ++i)
      start_position[i] = start_position[i - 1] + row_terms[i - 1];

    for (int i{ 1 }; i < arr[0].val + 1; ++i) {
      answ[start_position[arr[i].col]].set(arr[i].col, arr[i].row, arr[i].val);
      ++start_position[arr[i].col];
    }
  }
  return answ;
}

std::vector<Node> sparse_mult(const std::vector<Node> &a, const std::vector<Node> &b)
{
  if (a[0].col != b[0].col) {
    puts("Invalid Matrices");
    exit(1);
  }

  int a_index{ 1 }, b_index{ 1 }, row_begin{ 1 }, curr_row{}, curr_col, curr_sum{};
  std::vector<Node> answ;
  answ.emplace_back(Node());
  while (a_index < a[0].val + 1) {
    curr_col = b[1].row;
    b_index = 1;
    while (b_index < b[0].val + 2) {
      if (a[a_index].row != curr_row) {
        if (curr_sum != 0)
          answ.emplace_back(Node(curr_row, curr_col, curr_sum)), curr_sum = 0;
        a_index = row_begin;
        while (b[b_index].row == curr_col)
          ++b_index;
        curr_col = b[b_index].row;
      }
      else if (b[b_index].row != curr_col) {
        if (curr_sum != 0)
          answ.emplace_back(Node(curr_row, curr_col, curr_sum)), curr_sum = 0;
        a_index = row_begin;
        curr_col = b[b_index].row;
      }
      else {
        if (a[a_index].col > b[b_index].col)
          ++b_index;
        else if (a[a_index].col < b[b_index].col)
          ++a_index;
        else {
          curr_sum += (a[a_index].val * b[b_index].val);
          ++a_index, ++b_index;
        }
      }
    }    // end while ( b_index < b[0].val + 1 )
    while (a[a_index].row == curr_row)
      ++a_index;
    row_begin = a_index, curr_row = a[a_index].row;
  }    // end while ( a_index < a[0].val + 1 )
  answ[0].set(a[0].row, b[0].row, answ.size() - 1);

  return answ;
}

int main()
{
  std::string str;
  std::stringstream ss;

  int row, column;
  get_row_col(row, column);
  std::vector<Node> arr1;
  init_arr(str, ss, row, column, arr1);

  get_row_col(row, column);
  std::vector<Node> arr2;
  init_arr(str, ss, row, column, arr2);

  std::vector<Node> t_arr2 = fast_transpose(arr2);
  t_arr2.emplace_back(Node(arr2[0].col, 0, 0));

  std::vector<Node> answ = sparse_mult(arr1, t_arr2);
  for (const auto &ele : answ) {
    std::cout << ele.row << " " << ele.col << " " << ele.val << '\n';
  }
}