package main

/*
Implement the class TweetCounts that supports two methods:
https://leetcode.com/problems/tweet-counts-per-frequency/
1. recordTweet(string tweetName, int time)

Stores the tweetName at the recorded time (in seconds).
2. getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime)

Returns the total number of occurrences for the given tweetName per minute, hour, or day (depending on freq) starting from the startTime (in seconds) and ending at the endTime (in seconds).
freq is always minute, hour or day, representing the time interval to get the total number of occurrences for the given tweetName.
The first time interval always starts from the startTime, so the time intervals are [startTime, startTime + delta*1>,  [startTime + delta*1, startTime + delta*2>, [startTime + delta*2, startTime + delta*3>, ... , [startTime + delta*i, min(startTime + delta*(i+1), endTime + 1)> for some non-negative number i and delta (which depends on freq).

Example:

Input
["TweetCounts","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency"]
[[],["tweet3",0],["tweet3",60],["tweet3",10],["minute","tweet3",0,59],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]

Output
[null,null,null,null,[2],[2,1],null,[4]]

Explanation
TweetCounts tweetCounts = new TweetCounts();
tweetCounts.recordTweet("tweet3", 0);
tweetCounts.recordTweet("tweet3", 60);
tweetCounts.recordTweet("tweet3", 10);                             // All tweets correspond to "tweet3" with recorded times at 0, 10 and 60.
tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 59); // return [2]. The frequency is per minute (60 seconds), so there is one interval of time: 1) [0, 60> - > 2 tweets.
tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 60); // return [2, 1]. The frequency is per minute (60 seconds), so there are two intervals of time: 1) [0, 60> - > 2 tweets, and 2) [60,61> - > 1 tweet.
tweetCounts.recordTweet("tweet3", 120);                            // All tweets correspond to "tweet3" with recorded times at 0, 10, 60 and 120.
tweetCounts.getTweetCountsPerFrequency("hour", "tweet3", 0, 210);  // return [4]. The frequency is per hour (3600 seconds), so there is one interval of time: 1) [0, 211> - > 4 tweets.

*/

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

type TweetCounts struct {
	tweets map[string]*LinkedList
}

func Constructor() TweetCounts {
	return TweetCounts{
		tweets: make(map[string]*LinkedList),
	}
}

func (t *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, ok := t.tweets[tweetName]; !ok {
		t.tweets[tweetName] = &LinkedList{}
	}

	t.tweets[tweetName].Insert(NewNode(time))
}

func genNumber() int {
	return rand.Intn(100)
}

type Interval struct {
	end   int
	start int
	count int
}

func makeIntervals(freq string, startTime int, endTime int) []*Interval {
	var delta int

	switch freq {
	case "minute":
		delta = 60
	case "hour":
		delta = 3600
	case "day":
		delta = 86400
	default:
		panic("invalid freq time")
	}
	rem := endTime - startTime
	x := float64(rem) / float64(delta)

	var size int
	if x == math.Floor(x) {
		size = int(math.Floor(x)) + 1
	} else {
		size = int(math.Ceil(x))
	}

	idx := 0
	intervals := make([]*Interval, int(size))

	// While there is time, slice into new interval
	for ; idx < size; idx++ {
		start := startTime + (idx * delta)
		end := start + delta
		intervals[idx] = &Interval{
			start: start,
			end:   end,
			count: 0,
		}
		rem = rem - delta
	}

	return intervals
}

func (t *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	intervals := makeIntervals(freq, startTime, endTime)
	tweets, ok := t.tweets[tweetName]
	fmt.Println(len(intervals))
	result := make([]int, len(intervals))

	fmt.Println(intervals[0].start, intervals[0].end)
	if !ok {
		return result
	}

	ref := tweets.head

	if ref == nil || len(intervals) == 0 {
		return result
	}

	intervalIdx := 0
	var currInterval *Interval

	for ; ref != nil; ref = ref.next {
		currInterval = intervals[intervalIdx]

		if ref.data < currInterval.start {
			continue
		}

		// Fast forward index to correct interval
		for ref.data >= currInterval.end {
			if currInterval.count > 0 {
				result[intervalIdx] = currInterval.count
			}
			intervalIdx++
			if intervalIdx >= len(intervals) {
				break
			}
			currInterval = intervals[intervalIdx]
		}

		if intervalIdx >= len(intervals) {
			break
		}

		currInterval.count++
	}

	// Handle last case
	if intervalIdx < len(intervals) {
		result[intervalIdx] = intervals[intervalIdx].count
	}

	return result
}

func NewNode(i int) *Node {
	return &Node{data: i}
}

func main() {
	log.SetOutput(os.Stderr)
	tweetCounts := Constructor()
	tweetCounts.RecordTweet("tweet3", 0)
	tweetCounts.RecordTweet("tweet3", 60)
	tweetCounts.RecordTweet("tweet3", 10)
	tweetCounts.RecordTweet("tweet3", 120)
	tweetCounts.GetTweetCountsPerFrequency("minute", "tweet2", 78, 5374)
}

// Node represents a linked list node
type Node struct {
	data int
	next *Node
}

// LinkedList is a single linked list re
type LinkedList struct {
	head *Node
}

func (l *LinkedList) print() {
	for ref := l.head; ref != nil; ref = ref.next {
		fmt.Print(fmt.Sprintf("%v -> ", ref.data))
	}
}

func (l *LinkedList) Insert(n *Node) {
	// If head is empty, set as head
	if l.head == nil {
		l.head = n
		return
	}

	// Do we need to swap with head
	if n.data < l.head.data {
		tmpHead := l.head
		l.head = n
		n.next = tmpHead
		return
	}

	ref := l.head

	for ref.next != nil {
		next := ref.next

		if n.data < next.data {
			// insert n where next should be
			ref.next = n
			n.next = next
			return
		}
		ref = ref.next
	}

	// Append to the back
	if ref.next != nil {
		panic("What the shit!!")
	}

	ref.next = n
}
