// no_constructor.cpp
// Compile with: cl /EHsc no_constructor.cpp
#include <time.h>

// No constructor
struct TempData
{
  int StationId;
  time_t timeSet;
  double current;
  double maxTemp;
  double minTemp;
};

// Has a constructor
struct TempData2
{
  TempData2(double minimum, double maximum, double cur, int id, time_t t) : stationId{id}, timeSet{t}, current{cur}, maxTemp{maximum}, minTemp{minimum} {}
  int stationId;
  time_t timeSet;
  double current;
  double maxTemp;
  double minTemp;
};

int main()
{
  time_t time_to_set;

  // Member initialization (in order of declaration):
  TempData td{45978, time(&time_to_set), 28.9, 37.0, 16.7};

  // When there's no constructor, an empty brace initializer does
  // value initialization = {0,0,0,0,0}
  TempData td_emptyInit{};

  // Uninitialized = if used, emits warning C4700 uninitialized local variable
  TempData td_noInit;

  // Member declaration (in order of ctor parameters)
  TempData2 td2{16.7, 37.0, 28.9, 45978, time(&time_to_set)};

  return 0;
}