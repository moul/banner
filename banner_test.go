package banner_test

import (
	"fmt"
	"testing"

	"moul.io/banner"
)

func ExampleInline() {
	fmt.Println("start of banner")
	fmt.Println(banner.Inline("hey world."))
	fmt.Println("end of banner")
	// Output:
	// start of banner
	//  _                                    _     _
	// | |_   ___  _  _   __ __ __ ___  _ _ | | __| |
	// | ' \ / -_)| || |  \ V  V // _ \| '_|| |/ _` | _
	// |_||_|\___| \_, |   \_/\_/ \___/|_|  |_|\__,_|(_)
	//             |__/
	// end of banner
}

func TestInline(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"jjj", `
   _    _    _
  (_)  (_)  (_)
  | |  | |  | |
 _/ | _/ | _/ |
|__/ |__/ |__/
`},
		{"j j", `
   _      _
  (_)    (_)
  | |    | |
 _/ |   _/ |
|__/   |__/
`},
		{"j", `
   _
  (_)
  | |
 _/ |
|__/
`},
		{"@?!", `
 ___  ___  ___
|__ \|__ \|__ \
  /_/  /_/  /_/
 (_)  (_)  (_)
`},
		{"ccc", `
 __  __  __
/ _|/ _|/ _|
\__|\__|\__|
`},
		{" ", `

`},
		{"", `

`},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output := banner.Inline(test.input)
			expected := test.expected[1 : len(test.expected)-1]
			if expected != output {
				t.Log("output: \n" + output)
				t.Log("expected: \n" + expected)
				t.Errorf("output differs")
			}
		})
	}
}
