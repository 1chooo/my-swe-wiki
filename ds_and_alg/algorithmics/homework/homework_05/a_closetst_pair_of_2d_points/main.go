package main

import (
    "fmt"
    "math"
    "sort"
)

type point struct {
    x, y float64
}

func distance(a, b point) float64 {
    return math.Sqrt(math.Pow((a.x - b.x), 2) + math.Pow((a.y - b.y), 2))
}

// // Compare by Y
// func cmpY(a, b point) bool {
//     return a.y < b.y
// }

// // Compare by X
// func cmpX(a, b point) bool {
//     return a.x < b.x
// }

func closestPoint(s []point, low, high int, rec []point) float64 {
    var d1, d2, d3, d float64
    var mid, i, j, index int
    P := make([]point, high-low+1)
    temp1 := make([]point, 2)
    temp2 := make([]point, 2)

    if high-low == 1 {
        rec[0] = s[low]
        rec[1] = s[high]
        return distance(s[low], s[high])
    } else if high-low == 2 {
        d1 = distance(s[low], s[low+1])
        d2 = distance(s[low+1], s[high])
        d3 = distance(s[low], s[high])
        if d1 < d2 && d1 < d3 {
            rec[0] = s[low]
            rec[1] = s[low+1]
            return d1
        } else if d2 < d3 {
            rec[0] = s[low+1]
            rec[1] = s[high]
            return d2
        } else {
            rec[0] = s[low]
            rec[1] = s[high]
            return d3
        }
    } else {
        mid = (low + high) / 2
        d1 = closestPoint(s, low, mid, rec)
        temp1[0] = rec[0]
        temp1[1] = rec[1]
        d2 = closestPoint(s, mid+1, high, rec)
        temp2[0] = rec[0]
        temp2[1] = rec[1]
        if d1 < d2 {
            d = d1
            rec[0] = temp1[0]
            rec[1] = temp1[1]
        } else {
            d = d2
            rec[0] = temp2[0]
            rec[1] = temp2[1]
        }

        index = 0
        for i = mid; i >= low && (s[mid].x-s[i].x) < d; i-- {
            P[index] = s[i]
            index++
        }
        for i = mid + 1; i <= high && (s[i].x-s[mid].x) < d; i++ {
            P[index] = s[i]
            index++
        }
        sort.Slice(P[:index], func(i, j int) bool {
            return P[i].y < P[j].y
        })
        for i = 0; i < index; i++ {
            for j = i + 1; j < i+7 && j < index; j++ {
                if (P[j].y - P[i].y) >= d {
                    break
                } else {
                    d3 = distance(P[i], P[j])
                    if d3 < d {
                        rec[0] = P[i]
                        rec[1] = P[j]
                        d = d3
                    }
                }
            }
        }
        return d
    }
}

func main() {
    var p [100]point
    var n, m int
    var minDist float64

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&m)
        for j := 0; j < m; j++ {
            fmt.Scan(&p[j].x, &p[j].y)
        }
        sort.Slice(p[:m], func(i, j int) bool {
            return p[i].x < p[j].x
        })

        index := make([]point, 2)
        minDist = closestPoint(p[:m], 0, m-1, index)

        fmt.Printf("%.3f\n", minDist)
    }
}
