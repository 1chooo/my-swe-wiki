#include <stdio.h>
#include <stdbool.h>

#define INITIAL_CAPACITY 10

struct Stack {
    int* arr;
    int top;
    int capacity;
};

void initialize(struct Stack* stack) {
    stack->capacity = INITIAL_CAPACITY;
    stack->arr = (int*)malloc(sizeof(int) * stack->capacity);
    stack->top = -1;
}

void destroy(struct Stack* stack) {
    free(stack->arr);
}

bool isEmpty(struct Stack* stack) {
    return stack->top == -1;
}

bool isFull(struct Stack* stack) {
    return stack->top == stack->capacity - 1;
}

void expandStack(struct Stack* stack) {
    stack->capacity *= 2;
    stack->arr = (int*)realloc(stack->arr, sizeof(int) * stack->capacity);
}

void push(struct Stack* stack, int data) {
    if (isFull(stack)) {
        expandStack(stack);
    }
    stack->arr[++stack->top] = data;
}

int pop(struct Stack* stack) {
    if (isEmpty(stack)) {
        printf("Stack Underflow\n");
    }
    return stack->arr[stack->top--];
}

int peek(struct Stack* stack) {
    if (isEmpty(stack)) {
        printf("Stack is empty\n");
        return -1; 
    }
    return stack->arr[stack->top];
}

void clear(struct Stack* stack) {
    stack->top = -1;
}

int main() {
    struct Stack stack;
    initialize(&stack);

    push(&stack, 1);
    push(&stack, 2);
    push(&stack, 3);

    printf("%d\n", pop(&stack)); 
    printf("%d\n", pop(&stack));  
    printf("%d\n", peek(&stack)); 
    clear(&stack); 
    printf("%d\n", pop(&stack));

    destroy(&stack); 

    return 0;
}
