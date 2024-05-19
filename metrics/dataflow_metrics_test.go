package metrics

import "testing"

func TestRunMetricsService(t *testing.T) {
	if err := RunMetricsService("0.0.0.0:20004"); err != nil {
		t.Errorf("RunMetricsService() error = %v, wantErr %v", err, err)
	}
}
