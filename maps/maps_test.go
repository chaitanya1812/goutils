package maps

import (
	"testing"
)

func TestKeys(t *testing.T) {
	t.Run("success for map-string-int", func(t *testing.T) {
		testMap := map[string]int{"k1": 1, "k2": 2, "k3": 3}
		expectedKeys := map[string]struct{}{"k1": {}, "k2": {}, "k3": {}}

		keys := Keys(testMap)

		if len(keys) != len(expectedKeys) {
			t.Errorf("Expected %d keys, but, got %d", len(expectedKeys), len(keys))
		}

		for _, k := range keys {
			if _, ok := expectedKeys[k]; !ok {
				t.Errorf("key %s not found", k)
			}
		}
	})

	t.Run("success for map-int-int", func(t *testing.T) {
		testMap := map[int]int{6: 1, 0: 2, 1: 3}
		expectedKeys := map[int]struct{}{6: {}, 1: {}, 0: {}}

		keys := Keys(testMap)

		if len(keys) != len(expectedKeys) {
			t.Errorf("Expected %d keys, but, got %d", len(expectedKeys), len(keys))
		}

		for _, k := range keys {
			if _, ok := expectedKeys[k]; !ok {
				t.Errorf("key %d not found", k)
			}
		}
	})

}
