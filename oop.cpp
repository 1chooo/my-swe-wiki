#include <iostream>
#include <cstdlib>

using namespace std;

class Student {

  protected:
    int rollNumber;

  public:
    void setRoll(int);
    void getRoll(void);
};

void Student::setRoll(int roll) {
  rollNumber = roll;
}

void Student::getRoll(){
  cout << "The roll number is " << rollNumber << endl;
}

class Exam : public Student {

  protected:
    float math, physics;

  public:
    void setMarks(float, float);
    void getMarks(void);
};

void Exam::setMarks(float mathMark, float physicsMark) {
  math = mathMark;
  physics = physicsMark;
}

void Exam::getMarks() {
  cout << "The marks obtained in maths are " << math << endl;
  cout << "The marks obtained in physics are " << physics << endl;
}

class Result : public Exam {
  float percentage;

  public:
    void display(void);
};

void Result::display() {

  getRoll();
  getMarks();
  percentage = (math + physics) / 2;

  cout << "The total percentage is " << percentage << endl;
}

int main(void) {
  Result result;

  result.setRoll(12);
  result.setMarks(78.8, 62);
  result.display();

  return 0;
}

