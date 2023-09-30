#include <iostream>

using namespace std;

// Defining the linked list node class
class Node {
public:
    int data;
    Node* next;

    Node(int data) {
        this->data = data;
        this->next = NULL;
    }
};

// Class to implement the linked list data structure
class LinkedList {
public:
    Node* head;

    LinkedList() {
        this->head = NULL;
    }

    // Function to insert a new node at the front of the linked list
    void insertFront(int data) {
        Node* newNode = new Node(data);
        newNode->next = head;
        head = newNode;
    }

    // Function to insert a new node at the end of the linked list
    void insertEnd(int data) {
        Node* newNode = new Node(data);
        if (head == NULL) {
            head = newNode;
            return;
        }
        Node* lastNode = head;
        while (lastNode->next != NULL) {
            lastNode = lastNode->next;
        }
        lastNode->next = newNode;
    }

    // Function to delete the first occurrence of a node with the given data value
    void deleteNode(int data) {
        Node* temp = head;
        Node* prev = NULL;
        if (temp != NULL && temp->data == data) {
            head = temp->next;
            delete temp;
            return;
        }
        while (temp != NULL && temp->data != data) {
            prev = temp;
            temp = temp->next;
        }
        if (temp == NULL) {
            cout << "Node not found" << endl;
            return;
        }
        prev->next = temp->next;
        delete temp;
    }

    // Function to print the linked list
    void printList() {
        Node* temp = head;
        while (temp != NULL) {
            cout << temp->data << " ";
            temp = temp->next;
        }
        cout << endl;
    }
};

// Driver code to test the linked list implementation
int main() {
    LinkedList list;

    list.insertFront(2);
    list.insertFront(1);
    list.insertEnd(3);
    list.insertEnd(4);

    cout << "Linked list before deletion: ";
    list.printList();

    list.deleteNode(2);

    cout << "Linked list after deletion: ";
    list.printList();

    return 0;
}
