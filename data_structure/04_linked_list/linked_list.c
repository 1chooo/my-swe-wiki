#include <stdio.h>
#include <stdlib.h>

// Defining the linked list node structure
struct Node {
    int data;
    struct Node* next;
};

// Function to insert a new node at the front of the linked list
void insertFront(struct Node** head, int data) {
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = *head;
    *head = newNode;
}

// Function to insert a new node at the end of the linked list
void insertEnd(struct Node** head, int data) {
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = NULL;
    if (*head == NULL) {
        *head = newNode;
        return;
    }
    struct Node* lastNode = *head;
    while (lastNode->next != NULL) {
        lastNode = lastNode->next;
    }
    lastNode->next = newNode;
}

// Function to delete the first occurrence of a node with the given data value
void deleteNode(struct Node** head, int data) {
    struct Node* temp = *head;
    struct Node* prev = NULL;
    if (temp != NULL && temp->data == data) {
        *head = temp->next;
        free(temp);
        return;
    }
    while (temp != NULL && temp->data != data) {
        prev = temp;
        temp = temp->next;
    }
    if (temp == NULL) {
        printf("Node not found\n");
        return;
    }
    prev->next = temp->next;
    free(temp);
}

// Function to print the linked list
void printList(struct Node* head) {
    while (head != NULL) {
        printf("%d ", head->data);
        head = head->next;
    }
    printf("\n");
}

// Driver code to test the linked list implementation
int main() {
    struct Node* head = NULL;

    insertFront(&head, 2);
    insertFront(&head, 1);
    insertEnd(&head, 3);
    insertEnd(&head, 4);

    printf("Linked list before deletion: ");
    printList(head);

    deleteNode(&head, 2);

    printf("Linked list after deletion: ");
    printList(head);

    return 0;
}
