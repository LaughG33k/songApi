package pkg

import "time"

func Retry(fn func() error, attempt int, timeSleep time.Duration) (err error) {

	for attempt > 0 {

		if err = fn(); err != nil {
			attempt--
			time.Sleep(timeSleep)
			continue
		}

		return nil

	}

	return err
}
