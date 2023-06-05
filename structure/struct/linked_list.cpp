#include <iostream>

typedef struct Node {
    int data;
    struct Node* next;
} Node;

void insert(Node** head, int value) {
    Node* newNode = new Node;
    newNode->data = value;
    newNode->next = *head;
    *head = newNode;
}

void display(Node* head) {
    Node* current = head;
    while (current != nullptr) {
        std::cout << current->data << " ";
        current = current->next;
    }
    std::cout << std::endl;
}

int main() {
    Node* head = nullptr;

    insert(&head, 10);
    insert(&head, 20);
    insert(&head, 30);

    display(head);

    return 0;
}
