import "fmt"

resp, err := http.Get('https://blockchain.info/pt/ticker')

fmt.Printf(resp)
