package install

import "fmt"

// Help .
func Help() {
	fmt.Println(`
USAGE
    bench install
ARGUMENT
    OPTION  
OPTION
EXAMPLES
	go run cmd/bench/main.go install
DESCRIPTION
    init env`)
}

// Run .
func Run() {
	initAdminAccount()
}
