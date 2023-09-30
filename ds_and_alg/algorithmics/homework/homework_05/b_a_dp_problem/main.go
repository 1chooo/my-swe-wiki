package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)

    defer writer.Flush()

    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    for i := 0; i < n; i++ {
        scanner.Scan()
        S := scanner.Text()

        num := 1
        sign := 1
        x0 := 0
        x1 := 0

        j := 0
        for j < len(S) && S[j] != '=' {
            if S[j] == '+' {
                sign = 1
                num = 1
            } else if S[j] == '-' {
                sign = -1
                num = 1
            } else {
                num = 0
                if S[j] == 'x' {
                    num = 1
                } else {
                    for j < len(S) && '0' <= S[j] && S[j] <= '9' {
                        num = num*10 + int(S[j]-'0')
                        j++
                    }
                    j-- // S[i] will be processed next time
                }

                if S[j] == 'x' {
                    x1 += num * sign
                } else {
                    x0 -= num * sign
                }
            }
            j++
        }

        for j < len(S) {
            if S[j] == '+' {
                sign = 1
                num = 1
            } else if S[j] == '-' {
                sign = -1
                num = 1
            } else {
                num = 0
                if S[j] == 'x' {
                    num = 1
                } else {
                    for j < len(S) && '0' <= S[j] && S[j] <= '9' {
                        num = num*10 + int(S[j]-'0')
                        j++
                    }
                    j-- // S[i] will be processed next time
                }

                if S[j] == 'x' {
                    x1 -= num * sign
                } else {
                    x0 += num * sign
                }
            }
            j++
        }

        if x1 == 0 {
            if x0 == 0 {
                fmt.Fprintln(writer, "IDENTITY")
            } else {
                fmt.Fprintln(writer, "IMPOSSIBLE")
            }
        } else {
            result := float64(x0) / float64(x1)
            fmt.Fprintln(writer, result)
        }
    }
}
