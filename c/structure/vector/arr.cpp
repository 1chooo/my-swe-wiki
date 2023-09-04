#include <iostream>

class Vector {
private:
    int* array;
    int size;
    int capacity;

public:
    Vector() {
        size = 0;
        capacity = 2;
        array = new int[capacity];
    }

    ~Vector() {
        delete[] array;
    }

    int getSize() {
        return size;
    }

    int getCapacity() {
        return capacity;
    }

    void push_back(int element) {
        if (size == capacity) {
            resize();
        }
        array[size] = element;
        size++;
    }

    int& operator[](int index) {
        return array[index];
    }

    void insert(int index, int element) {
        if (size == capacity) {
            resize();
        }
        for (int i = size; i > index; i--) {
            array[i] = array[i - 1];
        }
        array[index] = element;
        size++;
    }

    void erase(int index) {
        for (int i = index; i < size - 1; i++) {
            array[i] = array[i + 1];
        }
        size--;
    }

    bool empty() {
        return size == 0;
    }

    void clear() {
        size = 0;
    }

private:
    void resize() {
        capacity *= 2;
        int* newArray = new int[capacity];
        for (int i = 0; i < size; i++) {
            newArray[i] = array[i];
        }
        delete[] array;
        array = newArray;
    }
};

int main() {
    Vector numbers;

    // Add elements to the vector
    numbers.push_back(10);
    numbers.push_back(20);
    numbers.push_back(30);

    // Get the size of the vector
    std::cout << "Size of vector: " << numbers.getSize() << std::endl;

    // Access elements in the vector
    std::cout << "Elements in vector: ";
    for (int i = 0; i < numbers.getSize(); i++) {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;

    // Insert an element at a specific position
    numbers.insert(1, 15);

    // Erase an element at a specific position
    numbers.erase(2);

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
