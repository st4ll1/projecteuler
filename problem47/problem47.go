package main

import (
    "fmt"
	"github.com/kavehmz/prime"
)

func main(){
	var digits int = 4 
	var max uint64 = 10
		
	for i := 1; i < 8; i++ {		
		max *= 10
	}
	
	primes := prime.Primes(max);
	
	var dpf int = 0;
	for i := 1; uint64(i) < max; i++ {
		var factors int = 0
		var currentNumber uint64 = uint64(i)
		for p := 0 ; p < len(primes); p++ {
			if currentNumber%primes[p] != 0 {
				continue
			}
			factors++
			for currentNumber%primes[p] == 0 {				
				currentNumber /= primes[p]				
			}			
			if currentNumber == 1 {
				if factors == digits{
					dpf++
				} else {
					dpf = 0
				}
				break
			}		
		}
		if dpf == digits{
    		fmt.Println("Solution Problem 47: ",i-digits+1)	
			break	
		}
	}	
	
	
}
