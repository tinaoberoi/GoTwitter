package feed

import (
	"proj1/lock"
)

// Feed represents a user's twitter feed
// You will add to this interface the implementations as you complete them.
type Feed interface {
	Add(body string, timestamp float64)
	Remove(timestamp float64) bool
	Contains(timestamp float64) bool
	GetFeedList() []*JsonFeed
}

// feed is the internal representation of a user's twitter feed (hidden from outside packages)
// You CAN add to this structure but you cannot remove any of the original fields. You must use
// the original fields in your implementation. You can assume the feed will not have duplicate posts
type feed struct {
	start *post // a pointer to the beginning post
	// ctxt  SharedContext
	l *lock.RWCondRWLock
}

type JsonFeed struct {
	Body      string  `json:"body"`
	Timestamp float64 `json:"timestamp"`
}

// post is the internal representation of a post on a user's twitter feed (hidden from outside packages)
// You CAN add to this structure but you cannot remove any of the original fields. You must use
// the original fields in your implementation.
type post struct {
	body      string  // the text of the post
	timestamp float64 // Unix timestamp of the post
	next      *post   // the next post in the feed
}

// NewPost creates and returns a new post value given its body and timestamp
func newPost(body string, timestamp float64, next *post) *post {
	return &post{body, timestamp, next}
}

// NewFeed creates a empy user feed
func NewFeed() Feed {
	return &feed{start: nil, l: lock.NewRWLock()}
}

// Add inserts a new post to the feed. The feed is always ordered by the timestamp where
// the most recent timestamp is at the beginning of the feed followed by the second most
// recent timestamp, etc. You may need to insert a new post somewhere in the feed because
// the given timestamp may not be the most recent.
func (f *feed) Add(body string, timestamp float64) {
	// fmt.Println("Here")
	f.l.Lock()
	new_post := newPost(body, timestamp, nil)
	// fmt.Println("Here")
	if f.start == nil {
		f.start = new_post
		f.l.Unlock()
		return
	}

	if f.start.timestamp <= timestamp {
		new_post.next = f.start
		f.start = new_post
		f.l.Unlock()
		return
	}

	pred := f.start
	curr := f.start.next

	for curr != nil && curr.timestamp > timestamp {
		pred = curr
		curr = curr.next
	}

	new_post.next = curr
	pred.next = new_post
	f.l.Unlock()

}

func (f *feed) GetFeedList() []*JsonFeed {
	var msg_list []*JsonFeed

	head := f.start

	for head != nil {
		var node JsonFeed
		node.Body = head.body
		node.Timestamp = head.timestamp

		msg_list = append(msg_list, &node)
		head = head.next
	}

	return msg_list
}

// Remove deletes the post with the given timestamp. If the timestamp
// is not included in a post of the  feed then the feed remains
// unchanged. Return true if the deletion was a success, otherwise return false
func (f *feed) Remove(timestamp float64) bool {
	f.l.Lock()
	if f.start == nil {
		f.l.Unlock()
		return false
	}

	if f.start.timestamp == timestamp {
		f.start = f.start.next
		f.l.Unlock()
		return true
	}

	pred := f.start
	curr := pred.next

	for curr != nil {
		if curr.timestamp == timestamp {
			pred.next = curr.next
			f.l.Unlock()
			return true
		}
		pred = curr
		curr = curr.next
	}

	f.l.Unlock()
	return false

}

// Contains determines whether a post with the given timestamp is
// inside a feed. The function returns true if there is a post
// with the timestamp, otherwise, false.
func (f *feed) Contains(timestamp float64) bool {
	f.l.RLock()
	// f.start.mtx.Lock()

	curr := f.start
	// curr.mtx.Lock()
	for curr != nil {
		if curr.timestamp == timestamp {
			f.l.RUnlock()
			return true
		}
		curr = curr.next
	}

	f.l.RUnlock()
	return false
}

// func main() {
// 	postInfo := [3]int{18}
// 	feed := NewFeed()

// 	//Add 20 posts to the feed
// 	for _, num := range postInfo {
// 		body := strconv.Itoa(num)
// 		fmt.Println("something")
// 		fmt.Println(body)
// 		feed.Add(body, float64(num))
// 		// fmt.Println(feed.Contains(float64(num)))
// 	}
// 	// fmt.Println(feed.Contains(2))
// }
