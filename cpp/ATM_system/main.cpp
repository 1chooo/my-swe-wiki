#include <iostream>
#include <cstdlib>
#include <string.h>
#include <stdlib.h>

using namespace std;

class Bank {

private  :
  string name;
  long long accnumber;
  char type[10];
  long long amount = 0;
  long long tot = 0;

public :
  void setValue() {
    cout << "Enter name:\n";
    cin.ignore();

    getline(cin, name);

    cout << "Enter Account Number:\n";
    cin >> accnumber;
    cout << "Enter Account type:\n";
    cin >> type;
    cout << "Enter Balance:\n";
    cin >> tot;
  }

  void showData() {
    cout << "Name:" << name << endl;
    cout << "Account No.:" << accnumber << endl;
    cout << "Account type:" << type << endl;
    cout << "Balance:" << tot << endl;
  }

  void deposit() {
    cout << "\nEnter amount to be Deposited\n";
    cin >> amount;
  }

  void showBalance() {
    tot += amount;
    cout << "\nTotal balance is: " << tot << endl;
  }

  void withdraw() {
    int a, avai_balance;

    cout << "Enter amount to withdraw\n";
    cin >> a;
    avai_balance = tot - a;
    cout << "Avaliable Balance is: " << avai_balance << endl;
  }
};

int main(void) {
  Bank bank;

  int choice;

  while(1) {
    cout << "\n-------WELCOME-------\n\n";
    cout << "Enter Your Choice\n";
    cout << "\t1. Enter name, Account number, Account type\n";
    cout << "\t2. Balance Enquiry\n";
    cout << "\t3. Deposit Money\n";
    cout << "\t4. Show Total Money\n";
    cout << "\t5. Withdraw Money\n";
    cout << "\t6. Cancel\n";

    cin >> choice;

    switch (choice) {
      case 1:
        bank.setValue();
        break;
      case 2:
        bank.showData();
        break;
      case 3:
        bank.deposit();
        break;
      case 4:
        bank.showBalance();
        break;
      case 5:
        bank.withdraw();
        break;
      case 6:
        exit(1);
        break;
      default:
        cout << "\nInvalid choice\n";
        break;
    }
  }

  return 0;
}