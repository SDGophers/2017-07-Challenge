# 2017-07-Challenge
Meetup challenge for July, 2017

## Overview

Use Go profiling tools to improve performance of a simple web service app that parses a fairly large tab seprated value data set. Since we are focusing on profiling more than coding, a basic working app is provided.


## Sample Data

Sample data is from last.fm in the form of artist play counts by user.  The data is in tab separated format (tsv).

The data is included in the `data` directory. It has been broken down into some smaller files for development and testing. If you want to use a larger data file, download and uncompress the orignal `lastfm-dataset-360K.tar.gz` to access the full data set (see below).

To pull out <n> number of records, for example 1,000,000, from within the `data` directory (where it is assumed the tar.gz file resides and has been uncompressed):

```
head -n 1000000 lastfm-dataset-360K/usersha1-artmbid-artname-plays.tsv > music-1000000.tsv
```

Source: https://www.upf.edu/web/mtg/lastfm360k

This dataset contains <user, artist, plays> tuples collected from Last.fm API using the `user.getTopArtists()` method.

Data file:

* http://mtg.upf.edu/static/datasets/last.fm/lastfm-dataset-360K.tar.gz (original)
* https://u78367877.dl.dropboxusercontent.com/u/78367877/lastfm-dataset-360K.tar.gz (alternate)

## Resoureces

* [William Kennedy Profiling training material](https://github.com/ardanlabs/gotraining/tree/master/topics/go/profiling)
* [Go Performance Wiki](https://github.com/golang/go/wiki/Performance)

Special thanks to [William Kennedy](https://www.ardanlabs.com/) for providing such comprehensive training material to the world. Go read his most excellent [blog](https://www.goinggo.net/).

## Exercises

### Memory and CPU Profiling

* [Example](https://github.com/ardanlabs/gotraining/tree/master/topics/go/profiling/memcpu)
* Uses benchmarks in `main_test.go` to generate data.
* Use a fairly big data set to generate a good bit of data.

1. `go test -run=xxx -bench . -benchtime 3s -benchmem -memprofile mem.out -cpuprofile cpu.out`
2. `go tool pprof 2017-07-Challenge.test cpu.out`
3. Try different commands, like `top` or `top20`
4. Can it be optimized any more?

### http/pprof Profiling

* [Example](https://github.com/ardanlabs/gotraining/blob/master/topics/go/profiling/pprof/README.md)

1. Uncomment the `http.ListenAndServe(...)` line at the bottom of `main.go` (and the `Printf` statement if you are a stickler for details.
2. Generate a load on the server using some kind of web loading tool (`hey`, `ab`, `seige`, etc.)
3. Raw http/pprof: `http://localhost:4000/debug/pprof`
4. Looking at Heap Profiles: `http://localhost:4000/debug/pprof/heap?debug=1`
5. CPU Profiling: `go tool pprof ./pprof http://localhost:4000/debug/pprof/profile`

### Tracing

* [Example](https://github.com/ardanlabs/gotraining/tree/f5a66e4f7a153e4b4f73dd264b8d86835e45efd9/topics/go/profiling/trace)

1. Uncomment the block between `/*` and `*/ at the top of `main()`
2. Disable the `http.ListenAndServe(...)`
3. Rebuild: `go build main.go`
4. Run it: `./main` - you should see a trace.out file created
5. View the output: `go tool trace trace.out`
6. Click 'View trace'
7. Things to note: heap allocations, garbage collection (GC), timing of goroutines

### Challenges

This is what we can demonstrate.

* Make it more efficient; demonstrate with profiling data
* Introduce "problems" that will make it slower; observe results - basically demonstrate what makes Go inefficient
   * Is allocating on the stack more efficient than allocating on the heap?
   * What data structures are most efficient?
* Add new functionality, including but not limited to:
   * Sort results by artist, listens, number of users
   * User demographics (will need to parse `usersha1-profile.tsv` in the data tar.gz file) - this would be pretty interesting
   * New requests, for instance: Top N artist, Top N users (with demographics), Bottom N artists ("the long tail"), etc.
* Profile initial new functionality, optimize and compare new profile data

