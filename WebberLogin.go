import("fmt"
       "os"
       "net/http"
       "net"
)

package main

/**********************Exsiting User********************/

type Usr struct{

    UsrName UserName
    PassWrd PassWord
    Email []Email
}

type UsrName struct{

    UserName string
    Email []Email

}

type PassWrd struct{

    Password string

}

type Email stuct{

    Address string

}

/**********************New User**************************/

type NewUser struct{

    nuEmail nuEmail[]
    nusrName nuUserName
    nuPasswrd nuPassWord

}

type nuEmail struct{

    Address string

}

type nusrName struct{

    nuUserName string

}

type nuPasswrd struct{

    nuPassWord string

}


/*******LOGIC*******/

func main(){

    goto operator
    goto login

}

func login(){

    fmt.Println("Welcome to the site\nEnter your UserName and Login")

}

func operator(){

    service := "53"
    tcpAddr, err := new.ListenTCP("tcp",service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    http.HandleFunc("/")
    log.Fatal(htttp.ListenAndServe("SweatSite/Login:53",nil))


}

func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "Sum Fatal Error: %s", err.Error())
	os.Exit(1)

    }

}
