#include <iostream>
#include <limits>
#include <sstream>
#include <string>
#include <utility>
#include <vector>

struct add_poly {
  std::vector<std::pair<int, int>> a, b, ans;
};

int main()
{
  int num1, num2, coeff, exp;
  add_poly ADD;
  std::stringstream sstr;
  std::string str;

  std::cin >> num1;
  std::cin.clear();
  std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
  ADD.a.reserve(num1);

  for (int i{}; i < num1; ++i) {
    getline(std::cin, str);
    sstr << str;
    sstr >> coeff >> exp;
    ADD.a.emplace_back(std::make_pair(coeff, exp));
    sstr.str("");
    sstr.clear();
  }

  std::cin >> num2;
  std::cin.clear();
  std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
  ADD.b.reserve(num2);

  for (int i{}; i < num2; ++i) {
    getline(std::cin, str);
    sstr << str;
    sstr >> coeff >> exp;
    ADD.b.emplace_back(std::make_pair(coeff, exp));
    sstr.str("");
    sstr.clear();
  }

  int index_a{}, index_b{};
  while (true) {
    if (!(index_a >= num1 || index_b >= num2)) {
      if (ADD.a[index_a].second > ADD.b[index_b].second) {
        ADD.ans.push_back(std::move(ADD.a[index_a]));
        ++index_a;
      }
      else if (ADD.a[index_a].second < ADD.b[index_b].second) {
        ADD.ans.push_back(std::move(ADD.b[index_b]));
        ++index_b;
      }
      else {
        ADD.ans.emplace_back(
          std::make_pair(ADD.a[index_a].first + ADD.b[index_b].first,
                         ADD.a[index_a].second));
        ++index_a, ++index_b;
      }
    }
    else if (!(index_a >= num1)) {
      ADD.ans.push_back(std::move(ADD.a[index_a]));
      ++index_a;
    }
    else if (!(index_b >= num2)) {
      ADD.ans.push_back(std::move(ADD.b[index_b]));
      ++index_b;
    }
    else
      break;
  }

  for (const auto &p : ADD.ans) {
    if (p.first != 0)
      std::cout << p.first << " " << p.second << " ";
  }
  return 0;
}