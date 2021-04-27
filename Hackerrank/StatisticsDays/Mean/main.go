package main

import (
   // "bufio"
    "fmt"
   // "io"
    //"os"
    //"strconv"
    //"strings"
)

/*
 * Complete the 'weightedMean' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY X
 *  2. INTEGER_ARRAY W
 */

func weightedMean(X []int32, W []int32) {
    // Write your code here
    
    var sumTot float32 = 0
    var weightSum float32 = 0


    for i:=0; i < len(X); i ++{
        sumTot += float32(X[i] * W[i])
		weightSum += float32(W[i])
    }
    
    fmt.Printf("%0.1f\n", float32(sumTot/weightSum))
	
 

}

func main() {
		
	X := []int32{10, 40, 30, 50, 20, 10, 40, 30, 50, 20}
	W := []int32{1,2,3,4,5,6,7,8,9,10}



	weightedMean(X,W)


}
