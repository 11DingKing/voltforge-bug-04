package thermal

import "testing"

func TestVoltForge04(t *testing.T) {
	history := &TemperatureHistoryHistory{}
	history.Add(38)
	history.Add(42)
	values := history.Values()
	values[0] = 99
	again := history.Values()
	if again[0] != 38 {
		t.Fatalf("stored sample was mutated: %v", again)
	}
}
