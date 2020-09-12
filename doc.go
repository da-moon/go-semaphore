// Package semaphore implements binary and counting semaphore
// and can be used to limit maximum number of outstanding operation
// based on permitpool github.com/hashicorp/vault/sdk/physical
// and github.com/marusama/semaphore and github.com/hlts2/lock-free
// there is a `slowSemaphore` struct that I used to learn more
// about CAS,sync.Cond and compare performances
// Counting semaphore ( and by extension slowsemaphore ) is broken when changing
// go maxproc
// seems like a pretty shit package
// mutex seems more performant
package semaphore
