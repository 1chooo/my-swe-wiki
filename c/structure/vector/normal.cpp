#include <iostream>
#include <vector>

int main() {
    // 創建一個空的 vector
    std::vector<int> numbers;

    // 在 vector 尾部添加元素
    numbers.push_back(10);
    numbers.push_back(20);
    numbers.push_back(30);

    // 獲取 vector 的大小
    std::cout << "Size of vector: " << numbers.size() << std::endl;

    // 訪問 vector 中的元素
    std::cout << "Elements in vector: ";
    for (int i = 0; i < numbers.size(); i++) {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;

    // 在指定位置插入元素
    numbers.insert(numbers.begin() + 1, 15);

    // 刪除指定位置的元素
    numbers.erase(numbers.begin() + 2);

    // 檢查 vector 是否為空
    if (numbers.empty()) {
        std::cout << "Vector is empty." << std::endl;
    } else {
        std::cout << "Vector is not empty." << std::endl;
    }

    // 清空 vector
    numbers.clear();

    // 檢查 vector 是否為空
    if (numbers.empty()) {
        std::cout << "Vector is empty." << std::endl;
    } else {
        std::cout << "Vector is not empty." << std::endl;
    }

    return 0;
}
