package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "crypto/rand"
    "math/big"
)

func randInt (size int64) int64 {
    max := big.NewInt(size)
    n, err := rand.Int(rand.Reader, max)
    if err != nil {
        log.Fatal(err)
    }
    return n.Int64() + 1
}

func resolveExpression (expression string) int64 {
    expS := strings.Split(expression, "d")
    number, _ := strconv.ParseInt(expS[0], 10, 64)
    dieSize, _ := strconv.ParseInt(expS[1], 10, 64)

    subTotal := []int64{}
    var i int64 = 0
    for ; i != number; i++ {
        subTotal = append(subTotal, randInt(dieSize))
    }
    fmt.Printf("%s%d%s + ", BLUE, subTotal, RESET)
    return arraySum(subTotal)
}

func arraySum (array []int64) int64 {
    var sum int64 = 0
    for i, _ := range array {
        sum = sum + array[i]
    }
    return sum
}

const GREEN = "\x1b[32m"
const BLUE = "\x1b[34m"
const RESET = "\x1b[0m"

func main() {
    if len(os.Args) <= 1 {
        fmt.Printf("usage: ./roll <dice expression>\nexample: ./roll 3d10+4d6+8\n")
        os.Exit(1)
    }

    dicePool := strings.Split(os.Args[1], "+")
    sum := []int64{}
    for _, element := range dicePool {
        if strings.Contains(element, "d") {
            sum = append(sum, resolveExpression(element))
        } else {
            num, _ := strconv.ParseInt(element, 10, 64)
            fmt.Printf("%d + ", num)
            sum = append(sum, num)
        }
    }
    fmt.Printf("\b\b= %s%d%s\n", GREEN, arraySum(sum), RESET)
}
