#include <iostream>
#include <vector>

int main() {
    // Create an empty vector
    std::vector<int> numbers;

    // Add elements to the vector
    numbers.push_back(10);
    numbers.push_back(20);
    numbers.push_back(30);

    // Get the size of the vector
    std::cout << "Size of vector: " << numbers.size() << std::endl;

    // Access elements in the vector
    std::cout << "Elements in vector: ";
    for (int i = 0; i < numbers.size(); i++) {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;

    // Insert an element at a specific position
    numbers.insert(numbers.begin() + 1, 15);

    // Erase an element at a specific position
    numbers.erase(numbers.begin() + 2);

    // Check if the vector is empty
    if (numbers.empty()) {
        std::cout << "Vector is empty." << std::endl;
    } else {
        std::cout << "Vector is not empty." << std::endl;
    }

    // Clear the vector
    numbers.clear();

    // Check if the vector is empty
    if (numbers.empty()) {
        std::cout << "Vector is empty." << std::endl;
    } else {
        std::cout << "Vector is not empty." << std::endl;
    }

    return 0;
}
