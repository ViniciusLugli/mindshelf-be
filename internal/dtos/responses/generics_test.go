package responses

import "testing"

func TestNewPaginatedResponseMapsItemsAndMetadata(t *testing.T) {
	resp := NewPaginatedResponse([]int{1, 2, 3}, func(v int) string {
		return string(rune('0' + v))
	}, 9, 2, 3, 3)

	if len(resp.Data) != 3 {
		t.Fatalf("expected 3 items, got %d", len(resp.Data))
	}

	if resp.Data[0] != "1" || resp.Data[1] != "2" || resp.Data[2] != "3" {
		t.Fatalf("unexpected mapped data: %#v", resp.Data)
	}

	if resp.Total != 9 || resp.Page != 2 || resp.Limit != 3 || resp.Total_pages != 3 {
		t.Fatalf("unexpected pagination metadata: %#v", resp)
	}
}
