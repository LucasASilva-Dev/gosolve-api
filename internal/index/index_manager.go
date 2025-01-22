package index

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type indexMonitor struct {
	sync.Mutex
	lastIndexSync time.Time
	file          []*int
	service       IndexService
}

var indexMonitorInstance indexMonitor

type IndexManager struct {
	indexMonitor *indexMonitor
}

// NewIndexManager creates a new IndexManager and returns a pointer to it.
// The IndexManager is responsible for managing the index and updating it
// every hour.  The index is updated in the background in a goroutine.
func NewIndexManager() (*IndexManager, error) {
	// Create a new indexMonitor to monitor and update the index.
	indexMonitorInstance = indexMonitor{
		service: NewIndexService(),
	}

	// Call updateIndexService to update the index.
	err := updateIndexService(time.Now())
	if err != nil {
		return nil, err
	}

	// Start a goroutine to monitor and update the index every hour.
	go startMonitorIndex()

	// Create a new IndexManager and set its indexMonitor to the
	// indexMonitor instance.
	indexManager := &IndexManager{
		indexMonitor: &indexMonitorInstance,
	}

	return indexManager, nil
}

// startMonitorIndex starts a goroutine to monitor the index and update it
// every hour.  This goroutine will run indefinitely until the program is
// terminated.
func startMonitorIndex() {
	// Set the interval to 1 hour.
	interval := 60 * time.Minute
	// Create a ticker that will send events every hour.
	tick := time.Tick(interval)

	// Run an infinite loop to process events.
	for now := range tick {
		// Call updateIndexService to update the index.
		err := updateIndexService(now)
		// If updateIndexService returned an error, log it and continue to the
		// next iteration.
		if err != nil {
			log.Println("[error] Index Monitor - ", err)
			continue
		}
	}
}

// updateIndexService updates the index when called by the index monitor.
func updateIndexService(now time.Time) error {
	file, err := indexMonitorInstance.service.UpdateIndex()

	// Acquire the lock before updating the index.
	indexMonitorInstance.Lock()
	// Update the last index sync time to the current time.
	indexMonitorInstance.lastIndexSync = now
	// Update the index to the new version.
	indexMonitorInstance.file = file
	// Release the lock after updating the index.
	defer indexMonitorInstance.Unlock()
	// Log that the index has been updated.
	log.Info("Index updated")

	return err
}

// Lookup returns the index of the given position in the index.
// If the position is not found, -1 is returned as the index and false is returned
// as the boolean indicating whether the target value was found.
//
// This function is used by the web server to look up the index of a given
// position in the index. It is a synchronous function that blocks until the
// index is updated.
func (im *IndexManager) Lookup(position int) (int, bool) {
	// If the indexManager is nil, return -1 as the index and false as the boolean.
	// This should never happen, but it's a sanity check just in case.
	if im == nil {
		return 0, false
	}

	// Call binarySearch to perform the binary search on the index.
	// The binarySearch function takes the following parameters:
	// - A slice of pointers to integers.
	// - An integer value to search for in the slice.
	//
	// The binarySearch function returns two values:
	// - The index of the value in the slice if it is found.
	// - A boolean indicating whether the target value was found.
	returnValue, found := binarySearch(im.indexMonitor.file, position)

	// If the target value was found, return the index and true as the boolean.
	if found {
		return returnValue, true
	}

	// If the target value was not found, return -1 as the index and false as the boolean.
	return -1, false
}

// binarySearch performs a binary search on the given slice of integers.
// It returns the index and a boolean indicating whether the target value
// was found. If the target value is not found, the closest value within
// 10% tolerance is returned. If no value is within 10% tolerance, -1 is
// returned as the index.
//
// The algorithm works by repeatedly dividing the search space in half.
// If the target value is found, the index of that value is returned with
// a boolean indicating whether the target value was found.
//
// If the target value is not found, the closest value within 10% tolerance
// is returned. If no value is within 10% tolerance, -1 is returned as the index.
func binarySearch(slice []*int, target int) (int, bool) {
	// Initialize the search space to the entire slice
	if len(slice) == 0 {
		return -1, false
	}
	low, high := 0, len(slice)-1

	// Continue the search until the search space is exhausted
	for low <= high {
		// Calculate the midpoint of the search space
		mid := (low + high) / 2
		// If the target value is found, return the index and a boolean indicating
		// whether the target value was found
		if *slice[mid] == target {
			log.Debugf("binarySearch: target found at index %d", mid)
			return mid, true
		}
		// If the target value is not found, narrow the search space to the
		// upper or lower half depending on whether the target value is
		// greater or less than the value at the midpoint
		if *slice[mid] < target {
			log.Debugf("binarySearch: target is greater than value at index %d, moving to upper half", mid)
			low = mid + 1
		} else {
			log.Debugf("binarySearch: target is less than value at index %d, moving to lower half", mid)
			high = mid - 1
		}
	}

	log.Debugf("binarySearch: search space exhausted, target value not found")

	// At this point, the search space is exhausted and the target value
	// was not found. Return the closest value within 10% tolerance
	closestIndex := low
	if closestIndex >= len(slice) {
		closestIndex = len(slice) - 1
	}
	if closestIndex > 0 && absValue(*slice[closestIndex]-target) > absValue(*slice[closestIndex-1]-target) {
		closestIndex--
	}

	log.Debugf("binarySearch: returning closest index %d", closestIndex)

	// If the value at the closest index is within 10% tolerance of the target
	// value, return the index and a boolean indicating whether the target value
	// was found
	if absValue(*slice[closestIndex]-target) <= target/10 {
		log.Debugf("binarySearch: target found within 10%% tolerance at index %d", closestIndex)
		return closestIndex, true
	}

	// If no value is within 10% tolerance, return -1 as the index
	log.Debugf("binarySearch: target not found within 10%% tolerance, returning -1")
	return -1, false
}

// absValue returns the absolute value of x.
func absValue(x int) int {
	// If x is negative, return its negation, which is positive.
	if x < 0 {
		return -x
	}
	// If x is positive, return it unchanged.
	return x
}
