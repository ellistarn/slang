package time

import "time"

// Time interfaces that may be mocked by tests.
var (
	Now   = time.Now
	Since = time.Since
	Until = time.Until
	Unix  = time.Unix
)
