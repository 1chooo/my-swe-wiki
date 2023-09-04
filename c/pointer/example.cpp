#include <iostream>

int main() {
    int x = 5;
    int* ptr; // Create an integer pointer

    ptr = &x; // Point the pointer to the address of x

    // Display the value of x, address of x, value of ptr, and value pointed by ptr
    std::cout << "Value of x: " << x << std::endl;
    std::cout << "Address of x: " << &x << std::endl;
    std::cout << "Value of ptr: " << ptr << std::endl;
    std::cout << "Value pointed by ptr: " << *ptr << std::endl;

    *ptr = 10; // Modify the value of x using the pointer

    std::cout << "New value of x: " << x << std::endl;

    int y = 20;
    ptr = &y; // Update the pointer to point to the address of y

    std::cout << "Value of y: " << y << std::endl;
    std::cout << "Address of y: " << &y << std::endl;
    std::cout << "Value of ptr: " << ptr << std::endl;
    std::cout << "Value pointed by ptr: " << *ptr << std::endl;

    return 0;
}
