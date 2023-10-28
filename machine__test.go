package snofkid

import "time"

func generateSnowflakesWithMilliSecsGenerated(mch *SnowflakeMachine) (snowflakes []int64, actualMillis []int64) {
	for i := 0; i < MaxSequence; i++ {
		now := time.Now().UnixMilli()
		sfid := mch.New()
		snowflakes = append(snowflakes, sfid)
		actualMillis = append(actualMillis, now)
	}
	return
}

func absoluteDiff(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}
