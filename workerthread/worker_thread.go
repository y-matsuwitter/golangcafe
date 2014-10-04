package workerthread

import "time"

// Main runs an example of worker thread pattern.
func Main() {
	channel := NewChannel(5)
	channel.StartWorkers()
	NewClient("Alice", channel).Run()
	NewClient("Boby", channel).Run()
	NewClient("Chris", channel).Run()
	time.Sleep(time.Minute)
}
