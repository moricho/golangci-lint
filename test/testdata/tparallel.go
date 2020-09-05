//args: -Etparallel
package testdtest

import (
	"fmt"
	"testing"
)

func call(name string) {
	fmt.Println(name)
}

func setup(name string) func() {
	fmt.Printf("setup: %s\n", name)
	return func() {
		fmt.Println("clean up finished")
	}
}

func Test_Func1(t *testing.T) { // ERROR "Test_Func1 should call t.Parallel on the top level"
	teardown := setup("Test_Func1")
	t.Cleanup(teardown)

	t.Run("Func1_Sub1", func(t *testing.T) {
		call("Func1_Sub1")
		t.Parallel()
	})

	t.Run("Func1_Sub2", func(t *testing.T) {
		call("Func1_Sub2")
		t.Parallel()
	})
}

func Test_Func2(t *testing.T) { // ERROR "Test_Func2's sub tests should call t.Parallel"
	teardown := setup("Test_Func2")
	t.Cleanup(teardown)

	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "Func2_Sub1",
		},
		{
			name: "Func2_Sub2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call(tt.name)
		})
	}
}

func Test_Func3(t *testing.T) { // ERROR "Test_Func3 should use t.Cleanup"
	teardown := setup("Test_Func3")
	defer teardown()

	t.Parallel()

	t.Run("Func3_Sub1", func(t *testing.T) {
		t.Parallel()
		call("Func3_Sub1")
	})

	t.Run("Func3_Sub2", func(t *testing.T) {
		call("Func3_Sub2")
	})
}
