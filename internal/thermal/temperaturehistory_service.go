package thermal

import "sync"

type TemperatureHistoryHistory struct {
	mu     sync.RWMutex
	values []int
}

func (h *TemperatureHistoryHistory) Add(value int) {
	h.mu.Lock()
	h.values = append(h.values, value)
	h.mu.Unlock()
}
func (h *TemperatureHistoryHistory) Values() []int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.values
}
