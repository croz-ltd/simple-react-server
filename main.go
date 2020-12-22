package main

import (
	"flag"
	"github.com/iancoleman/strcase"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Flags struct {
	BindAddress    string
	ServeDirectory string
}

func main() {
	flags := Flags{}
	flags.ParseFlags()

	servePath, err := filepath.Abs(flags.ServeDirectory)
	if err != nil {
		log.Fatalf("Path is not valid. %v", err)
	}
	log.Printf("Serving '%s' on HTTP port: %s\n", servePath, flags.BindAddress)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(servePath)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		mux.ServeHTTP(w,r)
	})

	err = http.ListenAndServe(flags.BindAddress, nil)
	if err != nil {
		log.Fatalf("Error while running HTTP server: %v", err)
	} else {
		log.Println("Shutdown gracefully")
	}
}

// Hierarchy of resolving is as follows:
//  1. Default values
//    2. Environment variable
//      3. Runtime argument
func (f *Flags) ParseFlags() {
	stringFlag(&f.BindAddress, "bind-address", ":3000", "The address on which server binds to.")
	stringFlag(&f.ServeDirectory, "serve-directory", "/var/www", "The directory which will be served")
	flag.Parse()
}

func stringFlag(parameter *string, name string, value string, usage string) {
	def := value
	envVal, exists := os.LookupEnv(strcase.ToScreamingSnake(name))
	if exists {
		def = envVal
	}

	flag.StringVar(parameter, name, def, usage)
}
