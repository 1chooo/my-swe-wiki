#include <iostream>

typedef struct Node {
    int data;
    struct Node* left;
    struct Node* right;
} Node;

Node* createNode(int value) {
    Node* newNode = new Node;
    newNode->data = value;
    newNode->left = nullptr;
    newNode->right = nullptr;
    return newNode;
}

void inorderTraversal(Node* root) {
    if (root == nullptr) {
        return;
    }
    inorderTraversal(root->left);
    std::cout << root->data << " ";
    inorderTraversal(root->right);
}

void preorderTraversal(Node* root) {
    if (root == nullptr) {
        return;
    }
    std::cout << root->data << " ";
    preorderTraversal(root->left);
    preorderTraversal(root->right);
}

void postorderTraversal(Node* root) {
    if (root == nullptr) {
        return;
    }
    postorderTraversal(root->left);
    postorderTraversal(root->right);
    std::cout << root->data << " ";
}

int main() {
    // 手動建構一個二元樹
    Node* root = createNode(1);
    root->left = createNode(2);
    root->right = createNode(3);
    root->left->left = createNode(4);
    root->left->right = createNode(5);

    std::cout << "Inorder traversal: ";
    inorderTraversal(root);
    std::cout << std::endl;

    std::cout << "Preorder traversal: ";
    preorderTraversal(root);
    std::cout << std::endl;

    std::cout << "Postorder traversal: ";
    postorderTraversal(root);
    std::cout << std::endl;

    return 0;
}
