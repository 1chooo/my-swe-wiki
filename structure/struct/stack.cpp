#include <iostream>

typedef struct Node {
    int data;
    struct Node* next;
} Node;

void push(Node** top, int value) {
    Node* newNode = new Node;
    newNode->data = value;
    newNode->next = *top;
    *top = newNode;
}

void pop(Node** top) {
    if (*top != nullptr) {
        Node* temp = *top;
        *top = (*top)->next;
        delete temp;
    }
}

void display(Node* top) {
    Node* current = top;
    while (current != nullptr) {
        std::cout << current->data << " ";
        current = current->next;
    }
    std::cout << std::endl;
}

int main() {
    Node* top = nullptr;

    push(&top, 10);
    push(&top, 20);
    push(&top, 30);

    display(top);

    pop(&top);

    display(top);

    return 0;
}
