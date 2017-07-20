package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	// "runtime/trace"
)

const DataFile = "data/music-1000000.tsv"
const Port = 4000

const (
	UserId = iota
	ArtistId
	Artist
	Plays
)

type PlayCount struct {
	ArtistId string
	Artist   string
	Users    []string
	Plays    int
}

func NewPlayCount(rec []string) *PlayCount {
	elem := &PlayCount{
		ArtistId: rec[ArtistId],
		Artist:   rec[Artist],
	}

	return elem
}

type MusicData map[string]*PlayCount

var musicData MusicData

// Reads tsv data file into records.
func readData(filename string) (records [][]string, err error) {
	// Verify file can be opened.
	fmt.Printf("Opening data file '%s'...\n", filename)
	f, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer f.Close()

	// Read tsv file with inconsistent quotes
	fmt.Println("Reading data...")
	r := csv.NewReader(f)
	r.Comma = '\t'
	r.LazyQuotes = true
	r.TrimLeadingSpace = true

	var rec []string

	for {
		rec, err = r.Read()
		//fmt.Printf("%7d %s\n", i, err)
		if err == io.EOF {
			return records, nil
		}
		if err == nil {
			records = append(records, rec)
		} else {
			//fmt.Printf("Read error: %s - ignoring\n", err)
		}

	}

	return records, err
}

// Parses tsv records into defined data format.
func parseData(records [][]string) MusicData {
	data := make(MusicData)
	for _, rec := range records {
		plays, _ := strconv.Atoi(rec[Plays])
		//fmt.Printf("%d: %5d %s %s %s\n", i, plays, rec[ArtistId], rec[UserId], rec[Artist])

		elem, ok := data[rec[ArtistId]]
		if !ok {
			elem = NewPlayCount(rec)
		}
		elem.Plays += plays
		elem.Users = append(elem.Users, rec[UserId])
		data[rec[ArtistId]] = elem
	}

	return data
}

func handleGetOne(rw http.ResponseWriter, r *http.Request) {
	var pc *PlayCount
	for _, pc = range musicData {
		fmt.Printf("%#v\n", pc)
		break
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(pc)
}

func main() {
	/* Uncomment to enable tracing tool

	// Create a file to hold tracing data.
	// View output: go tool trace trace.out
	tf, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()

	// Start gathering the tracing data.
	trace.Start(tf)
	defer trace.Stop()
	*/

	records, err := readData(DataFile)
	if err != nil {
		fmt.Printf("Encountered error after %d records: %s", len(records), err)
	}
	fmt.Printf("Read %d records.\n", len(records))

	fmt.Println("Parsing data...")
	musicData = parseData(records)
	fmt.Printf("Parsed %d artist play count records.\n", len(musicData))

	for id, elem := range musicData {
		fmt.Printf("%s: %5d %7d %s\n", id, len(elem.Users), elem.Plays, elem.Artist)
	}

	http.HandleFunc("/getone", handleGetOne)

	// Uncomment to profile web server functionality.
	// fmt.Printf("Listening on port %d...", Port)
	// http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
