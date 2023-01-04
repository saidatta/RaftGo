package util

const (
	// HEARTBEAT_MS is the heartbeat timer
	HEARTBEAT_MS = 150

	// ELECTION_TIMEOUT_MS is the election timer
	// It must be greater than 5 times the heartbeat timer, and it is recommended to have a 10 times relationship.
	ELECTION_TIMEOUT_MS = 1500

	// SLEEP_DEVIATION_MS is the sleep time
	SLEEP_DEVIATION_MS = 50

	// SCHEDULER_DELAY_MS is the timer start delay
	SCHEDULER_DELAY_MS = 3000

	// EMPTY_TERM is the term value of the empty object
	EMPTY_TERM = -1

	// RETRY_APPEND_MS is the retry log copy
	RETRY_APPEND_MS = 1000

	// LINE_SEP is the line separator
	LINE_SEP = "\n"
)
