#include <iostream>
#include <vector>

class Stack {
private:
    std::vector<int> data;

public:
    bool empty() {
        return data.empty();
    }

    void push(int value) {
        data.push_back(value);
    }

    void pop() {
        if (!empty()) {
            data.pop_back();
        }
    }

    int top() {
        if (!empty()) {
            return data.back();
        }
        return -1; // Or throw an exception for an empty stack
    }
};

int main() {
    Stack stack;

    stack.push(10);
    stack.push(20);
    stack.push(30);

    std::cout << "Top element: " << stack.top() << std::endl;

    stack.pop();

    std::cout << "Top element after pop: " << stack.top() << std::endl;

    return 0;
}
