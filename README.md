# Subtle Data Race in Go with sync.Mutex

This repository demonstrates a subtle data race in a Go program that uses a `sync.Mutex` for synchronization. The program appears to correctly use the mutex, but a closer look reveals a flaw leading to inconsistent results.

## Bug Description
The main function launches multiple goroutines, each appending a number to the shared `data` slice. A `sync.Mutex` is used to protect the append operation.  However, the race condition occurs due to the implicit capture of `i` in the anonymous function.  The goroutines may not see the final value of `i` before `wg.Done()` is called, resulting in race conditions.

## Solution
The solution demonstrates how to fix this by passing `i` as a separate variable to prevent the race condition.