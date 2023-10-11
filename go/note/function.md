# Go Language - Function

###### Tags: `golang`

### The most common declaration method

Use `func` to define the method
```go
func funcName(var varType) returnType {
    ...

    return ...
}
```

**Example**

```go
func add(x int, y int) int {
    return x + y
}

// or 
func add(x, y int) int {
    return x + y
}
```