/*
Sheret Static Web Server
(c) Ethan Piliavin

A tiny, simple static file web server.

*/
package main

import (
	"flag"
	"io/ioutil"
	"log"
    "os"
    "io"
	"net/http"
    "strings"
    "fmt"
    "path/filepath"
)

const appname string ="Sheret"
const appdesc string ="Static Web Server"
const version string = "1.0"

func loggingHandler(h http.Handler, quiet bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        if !quiet {
            
            log.Printf("%s %s %s", 
                r.RemoteAddr, r.Method, r.URL.Path)
                
            if r.Method == "POST" {
            
                r.ParseForm()
                log.Println("---- POST Data: ------------------------")
                for k, v := range r.Form {
                    log.Printf("%s \t = \t %s", k, strings.Join(v, ""))
                }
                log.Printf("---- End POST Data. %d Fields Received --", len(r.Form))
            }    
        }           
        
		h.ServeHTTP(w, r)     
	})
}

//Check if file exists or can be created
func IsValid(fp string) bool {
  // Check if file already exists
  if _, err := os.Stat(fp); err == nil {
    return true
  } else {
  	 fmt.Println(err)
  }

  // Attempt to create it
  var d []byte
  if err := ioutil.WriteFile(fp, d, 0644); err == nil {
    os.Remove(fp) // And delete it
    return true
  } else {
  	 fmt.Println(err)
  }
  return false
}

func main() {

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "%s v%s - %s\n", appname, version, appdesc)
        fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", filepath.Base(os.Args[0]))
        fmt.Fprintf(os.Stderr, "Parameters:\n\n")
        flag.PrintDefaults()
    }

	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "directory to serve from")
	quiet := flag.Bool("q", false, "suppress all logging")
	logfile := flag.String("f", "./sheret.log", "log to disk path [./sheret.log]")
	
	flag.Parse()

    if IsValid(*logfile) {
    
        logfile, err := os.OpenFile(*logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            log.Fatalln("Failed to open log file.", ":", err)
        }
        
        multiLog := io.MultiWriter(logfile, os.Stdout) 
        log.SetOutput(multiLog)
    } else {
    	 fmt.Fprintf(os.Stderr, "Unable to create logfile: %s\n", *logfile)
    	 os.Exit(1)
    }
    
    log.SetFlags(log.LstdFlags)
	
    http.Handle("/", loggingHandler(http.FileServer(http.Dir(*directory)),*quiet))
	
    log.Printf("%s v%s serving %s on HTTP port: %s\n", appname, version, *directory, *port)
	log.Printf("-- Press CTRL-C to terminate --\n")
   	log.Fatal(http.ListenAndServe(":"+*port, nil))
}