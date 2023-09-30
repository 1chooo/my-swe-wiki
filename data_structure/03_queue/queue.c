#include <stdio.h>

#define MAX_SIZE 100

// Creating a queue data structure
struct Queue {
    int arr[MAX_SIZE];
    int front, rear;
};

// Enqueue operation to add an element to the rear of the queue
void enqueue(struct Queue* queue, int data) {
    if (queue->rear == MAX_SIZE - 1) {
        printf("Queue Overflow\n");
        return;
    }
    queue->arr[++queue->rear] = data;
    if (queue->front == -1) {
        queue->front = 0;
    }
}

// Dequeue operation to remove and return the element from the front of the queue
int dequeue(struct Queue* queue) {
    if (queue->front == -1 || queue->front > queue->rear) {
        printf("Queue Underflow\n");
        return -1;
    }
    int data = queue->arr[queue->front++];
    if (queue->front > queue->rear) {
        queue->front = -1;
        queue->rear = -1;
    }
    return data;
}

// Driver code to test the queue implementation
int main() {
    struct Queue queue;
    queue.front = -1;
    queue.rear = -1;

    enqueue(&queue, 1);
    enqueue(&queue, 2);
    enqueue(&queue, 3);

    printf("%d\n", dequeue(&queue));  // Output: 1
    printf("%d\n", dequeue(&queue));  // Output: 2
    printf("%d\n", dequeue(&queue));  // Output: 3
    printf("%d\n", dequeue(&queue));  // Output: Queue Underflow

    return 0;
}
