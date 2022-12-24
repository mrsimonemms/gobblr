package e2e_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	BeforeEach(func() {
		fmt.Println("I'm a before each")
	})

	It("should", func() {
		Expect(2).To(Equal(3))
	})
})

// func BeforeEach() {}

// func TestExample(t *testing.T) {
// 	t.Run("hello", func(t *testing.T) {
// 		fmt.Println(222)
// 	})
// }

// func TestExample2(t *testing.T) {
// 	t.Run("hello", func(t *testing.T) {
// 		fmt.Println(333)
// 	})
// }
