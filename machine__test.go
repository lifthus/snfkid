package snofkid

import "time"

func absoluteDiff(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

func generateSnowflakesWithMilliSecsGenerated(mch *SnowflakeMachine) (snowflakes []int64, actualMillis []int64) {
	for i := 0; i < MaxSequence; i++ {
		now := time.Now().UnixMilli()
		sfid, _ := mch.New()
		snowflakes = append(snowflakes, sfid)
		actualMillis = append(actualMillis, now)
	}
	return
}

func generateSnowflakesFor500ms(mch *SnowflakeMachine) (snowflakes []int64) {
	done := make(chan interface{})
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				nsf, _ := mch.New()
				snowflakes = append(snowflakes, nsf)
			}
		}
	}()
	time.Sleep(time.Millisecond * 500)
	done <- struct{}{}
	return
}

func groupSnowflakesByRawTimestamp(snowflakes []int64) map[int64][]int64 {
	sfgroup := make(map[int64][]int64)
	for _, sfid := range snowflakes {
		ts := RawTimestamp(sfid)
		if _, ok := sfgroup[ts]; !ok {
			sfgroup[ts] = make([]int64, 0)
		}
		sfgroup[ts] = append(sfgroup[ts], sfid)
	}
	return sfgroup
}
