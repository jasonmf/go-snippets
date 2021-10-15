package main

// Sleep sleeps for specified duration or when the context is cancelled.
// If it's cancelled, the timer is cleaned up properly
func Sleep(ctx context.Context, d time.Duration) {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
	case <-timer.C:
	}
	if !timer.Stop() {
		<-timer.C
	}
}
