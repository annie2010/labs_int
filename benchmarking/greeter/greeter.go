// Greeter computes a greeting message

package greeter

import (
	"fmt"
)

func greeter1(n string, a int) string {
	const format = "Hello, %s! You are %d old today..."
	return fmt.Sprintf(format, n, a)
}

func greeter2(n string, a int) string {
	// YOUR_CODE
	return ""
}
