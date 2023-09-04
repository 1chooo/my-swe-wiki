#include <iostream>

typedef struct Node {
    int data;
    struct Node* next;
} Node;

typedef struct Queue {
    Node* front;
    Node* rear;
} Queue;

void enqueue(Queue* queue, int value) {
    Node* newNode = new Node;
    newNode->data = value;
    newNode->next = nullptr;
    if (queue->rear == nullptr) {
        queue->front = newNode;
        queue->rear = newNode;
    } else {
        queue->rear->next = newNode;
        queue->rear = newNode;
    }
}

void dequeue(Queue* queue) {
    if (queue->front != nullptr) {
        Node* temp = queue->front;
        queue->front = queue->front->next;
        delete temp;
        if (queue->front == nullptr) {
            queue->rear = nullptr;
        }
    }
}

void display(Queue* queue) {
    Node* current = queue->front;
    while (current != nullptr) {
        std::cout << current->data << " ";
        current = current->next;
    }
    std::cout << std::endl;
}

int main() {
    Queue queue;
    queue.front = nullptr;
    queue.rear = nullptr;

    enqueue(&queue, 10);
    enqueue(&queue, 20);
    enqueue(&queue, 30);

    display(&queue);

    dequeue(&queue);

    display(&queue);

    return 0;
}
