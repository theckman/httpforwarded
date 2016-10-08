package httpforwarded_test

import (
	"fmt"

	"github.com/theckman/go-httpforwarded"
)

func ExampleParse() {
	// build a mock value for the Forwarded HTTP header
	vals := []string{"for=192.0.2.1; proto=http"}

	// parse the fields in to one map
	params, _ := httpforwarded.Parse(vals)

	// print the origin IP address and protocolg
	fmt.Printf("origin: %s | protocol: %s", params["for"][0], params["proto"][0])
	// output: origin: 192.0.2.1 | protocol: http
}

func ExampleFormat() {
	// build a parameter map
	params := map[string][]string{
		"for":   []string{"192.0.2.1", "192.0.2.4"},
		"proto": []string{"http"},
	}

	// format the parameter map
	val, _ := httpforwarded.Format(params)

	fmt.Print(val)
	// output: for=192.0.2.1, for=192.0.2.4; proto=http
}
