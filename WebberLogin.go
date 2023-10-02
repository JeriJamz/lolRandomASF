package main

import("fmt"
       "os"
       "net/http"
       "net"
       "strings"
       "encoding/gob"
)

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

type Email struct{
    
    Address string

}

/**********************New User**************************/

type NewUser struct{

    nuEmail []nuEmail
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

    operator()
    login()

}

func login(){

    fmt.Println("Welcome to the site\nEnter your UserName and Login")
    scanln(k)//working on communication btw html and server
    if k == strings.ToLower("Yes"){

	NewUserLogin()

    }
}

func operator(){

    service := "53"
    tcpAddr, err := new.ListenTCP("tcp",service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    http.HandleFunc("/")
    log.Fatal(htttp.ListenAndServe("SweatSite/Login:53",nil))
    Kyrie()

}

func Kyrie(){

    ServeHTTP(w, http.ResponseWriter, r *http.Request)

}

func NewUserLogin(){

    NewUser := Usr{nusrName:nuUserName{usrName:"",nuEmail:""},
		   {nuPasswrd:nuPassWord{Password:""}}}
    
    WriteNewUsrLogin()

    func WriteNewUsrLogin(fileName string, key interface{}){

        fmt.Println("Hello New User. Please Type in your User Name")
        outFile, err := os.Write("Login.txt")
        checkError(err)
        encoder := gob.NewEncoder(outFile)
        err = encoder.Encode(key)
        checkError(err)
        fmt.Println("Please Enter Your UserName:")
        newUsrName := scanln()
        outFile := newUsrName
        fmt.Println("Please Enter Your Password")
        newPassword := scanln()
        outFile := newPassword
        outFile.Close()

    }


    WriteNewUsrLogin("NewUser.gob",NewUser)

}


func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "Sum Fatal Error: %s", err.Error())
	os.Exit(1)

    }

}
