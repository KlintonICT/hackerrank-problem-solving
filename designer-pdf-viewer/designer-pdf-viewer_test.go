package designer_pdf_viewer

import (
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := designerPdfViewer(tc.data.h, tc.data.word)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
