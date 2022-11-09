package leetcode

import (
	"fmt"
	"sort"
)

type TimeMap struct {
	TimestampValueMap map[string][]TimestampValue
}

type TimestampValue struct {
	Value     string
	Timestamp int
}

func TimeMapConstructor() TimeMap {
	return TimeMap{TimestampValueMap: make(map[string][]TimestampValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.TimestampValueMap[key] == nil {
		this.TimestampValueMap[key] = []TimestampValue{{Value: value, Timestamp: timestamp}}
	} else {
		timestampValue := TimestampValue{Value: value, Timestamp: timestamp}
		this.TimestampValueMap[key] = append(this.TimestampValueMap[key], timestampValue)
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if this.TimestampValueMap[key] == nil || timestamp < this.TimestampValueMap[key][0].Timestamp {
		return ""
	}
	timestampValueList := this.TimestampValueMap[key]
	index := sort.Search(len(timestampValueList), func(i int) bool {
		return timestampValueList[i].Timestamp > timestamp
	})
	return timestampValueList[index-1].Value
}

func Test_TimeBasedKey_ValueStore() {
	timeMap := TimeMapConstructor()
	timeMap.Set("foo", "bar", 1)       // store the key "foo" and value "bar" along with timestamp = 1.
	fmt.Println(timeMap.Get("foo", 1)) // return "bar"
	fmt.Println(timeMap.Get("foo", 3)) // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	timeMap.Set("foo", "bar2", 4)      // store the key "foo" and value "bar2" along with timestamp = 4.
	fmt.Println(timeMap.Get("foo", 4)) // return "bar2"
	fmt.Println(timeMap.Get("foo", 5)) // return "bar2"

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	fmt.Println(timeMap.Get("love", 5))  // ""
	fmt.Println(timeMap.Get("love", 10)) // "high"
	fmt.Println(timeMap.Get("love", 15)) // "high"
	fmt.Println(timeMap.Get("love", 20)) // "low"
	fmt.Println(timeMap.Get("love", 25)) // "low"
}
