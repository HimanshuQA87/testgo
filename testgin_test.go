package testgo

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("tt", func() {
	It("test func", func() {
		fmt.Println("Testing done")
		Expect(true).To(BeTrue())
	})
})

func Cnew() {
	fmt.Println("Hi Himanshu")
}

func Cnewupdated() {
	fmt.Println("Hi Himanshu Dhiman")
}

func AddNum(num1, num2 int) int {
	return num1 + num2
}
