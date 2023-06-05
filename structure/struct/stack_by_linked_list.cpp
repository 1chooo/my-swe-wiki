#include <iostream>

typedef struct Node {
    int data;
    struct Node* next;
} Node;

typedef struct Stack {
    Node* top;
} Stack;

void push(Stack* stack, int value) {
    Node* newNode = new Node;
    newNode->data = value;
    newNode->next = stack->top;
    stack->top = newNode;
}

void pop(Stack* stack) {
    if (stack->top != nullptr) {
        Node* temp = stack->top;
        stack->top = stack->top->next;
        delete temp;
    }
}

int peek(Stack* stack) {
    if (stack->top != nullptr) {
        return stack->top->data;
    }
    return -1; // 或者返回一個錯誤碼表示堆疊為空
}

bool isEmpty(Stack* stack) {
    return stack->top == nullptr;
}

void display(Stack* stack) {
    Node* current = stack->top;
    while (current != nullptr) {
        std::cout << current->data << " ";
        current = current->next;
    }
    std::cout << std::endl;
}

int main() {
    Stack stack;
    stack.top = nullptr;

    push(&stack, 10);
    push(&stack, 20);
    push(&stack, 30);

    std::cout << "Elements in stack: ";
    display(&stack);

    std::cout << "Top element: " << peek(&stack) << std::endl;

    pop(&stack);

    std::cout << "Elements in stack after pop: ";
    display(&stack);

    std::cout << "Is stack empty? " << (isEmpty(&stack) ? "Yes" : "No") << std::endl;

    return 0;
}
