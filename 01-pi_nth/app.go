
package main

import (
    "strconv"
    "fmt"
)

func calculatePi(digits int) string {
    var result string;
    digits++

    x := make([]int, digits*3+2)
    r := make([]int, digits*3+2)

    for i := range x {
        x[i] = 20
    }

    for i := 0; i < digits; i++ {
        var carry int
        for j := range x {
            num := len(x) - j - 1
            dem := num * 2 + 1

            x[j] += carry

            q := x[j] / dem
            r[j] = x[j] % dem

            carry = q * num
        }

        if i < digits - 1 {
            result += strconv.Itoa(x[len(x) - 1] / 10)
        }

        r[len(x) - 1] = x[len(x) - 1] % 10

        for j := range x {
            x[j] = r[j] * 10
        }
    }

    return result
}

func getInput() int {
    var i int
    fmt.Print("Please enter the number of digits to calculate: ")
    _, err := fmt.Scanf("%d", &i)
    if err != nil {
        panic(err)
    }
    return i

}

func main() {
    n := getInput()
    fmt.Println("PI: ", calculatePi(n))
}
