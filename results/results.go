package results

import (
	"sort"
	"time"
)

// Result is a structure that encapsulates processable result data
type Result struct {
	Timestamp   time.Time
	RequestTime time.Duration
	Error       error
	Threads     int
}

// ResultSet is a convenience mapping for type []Result
type ResultSet []Result

// Implement Sort interface
func (r ResultSet) Len() int {
	return len(r)
}

func (r ResultSet) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r ResultSet) Less(i, j int) bool {
	return r[i].Timestamp.UnixNano() < r[j].Timestamp.UnixNano()
}

// Reduce reduces the ResultSet into buckets defined by the given interval
func (r ResultSet) Reduce(interval time.Duration) []ResultSet {
	sort.Sort(r)

	start := r[0].Timestamp
	end := r[len(r)-1].Timestamp

	// create the buckets
	bucketCount := getBucketCount(start, end, interval)
	buckets := make([]ResultSet, bucketCount+1)

	for _, result := range r {

		currentBucket := getBucketNumber(result.Timestamp, start, end, interval, bucketCount)
		buckets[currentBucket] = append(buckets[currentBucket], result)
	}

	return buckets
}

func getBucketCount(start time.Time, end time.Time, interval time.Duration) int {
	totalDuration := end.UnixNano() - start.UnixNano()
	return int(totalDuration / int64(interval))
}

func getBucketNumber(current time.Time, start time.Time, end time.Time, interval time.Duration, bucketCount int) int {
	bucketSize := int64(interval)
	curr := end.UnixNano() - current.UnixNano()

	return bucketCount - int(curr/bucketSize)
}
