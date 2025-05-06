package main
import "fmt"
const NMAX int = 100

func main() {
	var A, B, C [NMAX]int
	var nA, nB, nC int

	fmt.Scan(&nA)
	for i := 0; i < nA; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Scan(&nB)
	for i := 0; i < nB; i++ {
		fmt.Scan(&B[i])
	}

	for i := 0; i < nA; i++ {
		for j := 0; j < nB; j++ {
			if A[i] == B[j] {
				duplicate := false
				for k := 0; k < nC; k++ {
					if C[k] == A[i] {
						duplicate = true
						break
					}
				}
				if !duplicate {
					C[nC] = A[i]
					nC++
				}
			}
		}
	}
	for i := 0; i < nC; i++ {
		fmt.Print(C[i], " ")
	}
}

