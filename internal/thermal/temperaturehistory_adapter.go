package thermal

func (h *TemperatureHistoryHistory) Latest() (int, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.values) == 0 {
		return 0, false
	}
	return h.values[len(h.values)-1], true
}
func (h *TemperatureHistoryHistory) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.values)
}

func ExportTemperatureHistory(history *TemperatureHistoryHistory) []int {
	values := history.Values()
	out := make([]int, len(values))
	copy(out, values)
	return out
}
