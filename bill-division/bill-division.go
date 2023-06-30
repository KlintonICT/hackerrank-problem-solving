package bill_division

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	sum := int32(0)

	for i, element := range bill {
		if int32(i) != k {
			sum += element
		}
	}

	divide := sum / 2

	if b > divide {
		fmt.Println(b - divide)
	} else {
		fmt.Println("Bon Appetit")
	}
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Bill Division")

	for _, tc := range tc.getTestCases() {
		fmt.Printf("TestCase: %d; data: %v; result: \n", tc.id, tc.data)
		bonAppetit(tc.data.bill, tc.data.k, tc.data.b)
	}
}
