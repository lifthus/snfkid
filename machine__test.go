package snofkid

import "time"

func absoluteDiff(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

func generateSnowflakesWithMilliSecsGenerated(mch *SnowflakeMachine) (snowflakes []int64, actualMillis []int64) {
	for i := 0; i < MaxSequence+1; i++ {
		now := time.Now().UnixMilli()
		sfid, err := mch.New()
		if err != nil {
			continue
		}
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
				nsf, err := mch.New()
				if err != nil {
					continue
				}
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

func HasDuplicates(snowflakes []int64) bool {
	sfset := make(map[int64]interface{})
	for _, sf := range snowflakes {
		if _, ok := sfset[sf]; ok {
			return true
		}
		sfset[sf] = struct{}{}
	}
	return false
}
