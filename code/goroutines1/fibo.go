package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"	
)

// This implementation is from: https://www.thepolyglotdeveloper.com/2016/12/fibonacci-sequence-printed-golang/
func FibonacciLoop(n int) int {
    f := make([]int, n+1, n+2)
    if n < 2 {
        f = f[0:2]
    }
    f[0] = 0
    f[1] = 1
    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}
// End

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Type a term number or other character to finish: ")
		nints, _ := reader.ReadString('\n')
		nint, errInt := strconv.ParseInt(strings.TrimSpace(nints), 10, 64)
		if errInt != nil {
			break
		}
		fmt.Printf("You typed: %d and the term is %d\n", nint,FibonacciLoop(int(nint)))
	}
	
}