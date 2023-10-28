package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lifthus/snofkid"
)

var (
	epoch int64
)

func main() {
	epochF := flag.Int64("e", snofkid.TwitterEpoch, "epoch to use")
	machineF := flag.Int64("m", 0, "machine ID")
	portF := flag.String("p", "", "port to listen on")

	certF := flag.String("cert", "", "TLS certificate file")
	keyF := flag.String("key", "", "TLS key file")
	flag.Parse()

	epoch = *epochF

	switch {
	case *portF != "":
		log.Fatal(StartSnowflakeServer(*portF, *certF, *keyF))
	default:
		mch, err := snofkid.NewMachine(*epochF, *machineF)
		if err != nil {
			log.Fatalf("failed to create machine: %v", err)
		}
		sfid, err := mch.New()
		if err != nil {
			log.Fatalf("failed to generate snowflake: %v", err)
		}
		fmt.Println(sfid)
	}
}

func StartSnowflakeServer(port string, cert, key string) error {
	mux := http.NewServeMux()

	mchs := make(map[int64]*snofkid.SnowflakeMachine)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mid, err := getMachineIDFromFirstSegment(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var mch *snofkid.SnowflakeMachine
		var ok bool
		if mch, ok = mchs[mid]; !ok {
			mch, err = snofkid.NewMachine(epoch, mid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		sfid, err := mch.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%d", sfid)
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("Snowflake server listening on port:", port)
	if cert != "" && key != "" {
		return server.ListenAndServeTLS(cert, key)
	}
	return server.ListenAndServe()

}

func getMachineIDFromFirstSegment(path string) (int64, error) {
	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")
	if len(segs) != 1 {
		return 0, fmt.Errorf("invalid machine ID")
	}
	return strconv.ParseInt(segs[0], 10, 64)
}
