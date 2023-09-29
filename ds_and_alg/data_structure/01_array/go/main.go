package main

import "fmt"

func main() {
    // Creating an array
    arr := [5]int{1, 2, 3, 4, 5}

    // Accessing elements of the array
    fmt.Println(arr[0]) // Output: 1
    fmt.Println(arr[2]) // Output: 3

    // Getting the length of the array
    length := len(arr)
    fmt.Printf("Length of the array: %d\n", length) // Output: 5

    // Iterating through the array
    fmt.Println("Elements of the array:")
    for i := 0; i < length; i++ {
        fmt.Println(arr[i])
    }
}
