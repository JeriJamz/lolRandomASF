//trynna make this a non local server. idk how to make it live. And ima make a site so yea. fuck twitter
package main

import(
      "os"
      "fmt"
      "log"
      "net/http"
      "net"
)

func main(){

    service := "53"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)

    http.HandleFunc("/", kyrie)//make sure yall hand check kyrie
    log.Fatal(http.ListenAndServe("Sweatsite:53",nil))

    for{

	conn, err :=  listener.Accept()
	if err != nil{

	    continue

	}
    sum := "I need some for the connection"
    conn.Write([]byte(sum))
    conn.Close()

    }    

}


func kyrie(w http.ResponseWriter, r  *http.Request){//kyrie got handles so he the handler

    fmt.Fprint(w, "The Handler is up. Read the URL", r.URL.Path)   

}


func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "Sum fatal error: %s", err.Error())
        os.Exit(1)

    }

}