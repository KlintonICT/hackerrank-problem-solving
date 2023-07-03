package designer_pdf_viewer

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	max := int32(1)

	for _, char := range word {
		if max < h[char-'a'] {
			max = h[char-'a']
		}
	}

	return max * int32(len(word))
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Designer PDF Viewer")

	for _, tc := range tc.getTestCases() {
		result := designerPdfViewer(tc.data.h, tc.data.word)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
