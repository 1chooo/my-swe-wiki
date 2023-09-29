#include <deque>
#include <functional>
#include <iostream>
#include <limits>
#include <sstream>
#include <string>
#include <utility>

namespace mes {
  void sel_sort(std::deque<std::pair<std::string, int>> &arr, const int &num)
  {
    int max_index{};

    for (int i{}; i < num; ++i) {
      max_index = i;

      for (int j = i + 1; j < num; ++j)
        if (std::greater<int>{}(arr[j].second, arr[max_index].second))
          max_index = j;

      std::pair<std::string, int> tmp = arr[max_index];
      while (max_index > i) {
        arr[max_index] = arr[max_index - 1];
        --max_index;
      }
      arr[i] = tmp;
    }
  }
}    // namespace mes

int main()
{
  int num;
  std::string buffer;
  std::stringstream ssbuf;
  std::deque<std::pair<std::string, int>> arr;
  while (getline(std::cin, buffer)) {
    ssbuf << buffer;
    ssbuf >> num;
    ssbuf.str("");
    ssbuf.clear();
    arr.resize(num);
    for (int i{}; i < num; ++i) {
      getline(std::cin, buffer);
      ssbuf << buffer;
      ssbuf >> arr[i].first >> arr[i].second;

      buffer.clear();
      ssbuf.str("");
      ssbuf.clear();
    }

    mes::sel_sort(arr, num);
    for (const auto &x : arr) {
      std::cout << x.first << '\n';
    }

    arr.clear();
  }

  return 0;
}