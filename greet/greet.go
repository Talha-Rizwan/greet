
package greet

import (
	"fmt"
	
)



func CalculateHash (stringToHash string) string{

	fmt.Printf("changes are made in some block of chain \n")

	// fmt.Printf("String Received : %s \n \n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
	

}
