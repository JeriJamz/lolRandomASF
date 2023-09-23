//trynna make this a non local server. idk how to make it live. And ima make a site so yea. fuck twitter
package main

import(
      "fmt"
      "log"
      "net/http"
)

func main(){

    http.HandleFunc("/", kyrie)//make sure yall hand check kyrie
    log.Fatal(http.ListenAndServe("Sweatsite:53",nil))
}


func kyrie(w http.ResponseWriter, r  *http.Request){//kyrie got handles so he the handler

    fmt.Fprint(w, "The Handler is up. Read the URL", r.URL.Path)   

}