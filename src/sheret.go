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
  "github.com/icza/gox/osx" // for OpenDefault
)

const appname string ="Sheret"
const appdesc string ="Static Web Server"
const version string = "1.2"

func loggingHandler(h http.Handler, quiet bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        if !quiet {
            
            log.Printf("%s %s %s", 
                r.RemoteAddr, r.Method, r.RequestURI)
                
            r.ParseForm()

            if len(r.Form) >= 1 {

              //Enumerate URL Parameters
              if len(r.URL.Query()) >= 1 {
                  for k, v := range r.URL.Query(){
                      log.Printf("\t URL:\t%s\t = \t%s", k, strings.Join(v, ""))
                  }
              }
              
              //Enumerate POSTed data 
              if r.Method == "POST" {
                  for k, v := range r.PostForm {
                      log.Printf("\tPOST:\t%s\t = \t%s", k, strings.Join(v, ""))
                  }
              }               
              
              log.Printf("---- End Data. %d Values Received ----", len(r.Form))
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
	nobrowser := flag.Bool("b", false, "disable open system browser on launch")

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

	if !*quiet {
    log.Printf("%s v%s serving %s on HTTP port: %s\n", appname, version, *directory, *port)
		log.Printf("-- Press CTRL-C to terminate --\n")
	}

	http.Handle("/", loggingHandler(http.FileServer(http.Dir(*directory)),*quiet))
	
	if !*nobrowser {
		osx.OpenDefault("http://localhost:"+*port)
	}
  log.Fatal(http.ListenAndServe(":"+*port, nil))

}
