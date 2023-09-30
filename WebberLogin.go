import("fmt"
       "os"
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

    Email Email[]
    usrName UserName
    Passwrd PassWord

}


func main(){

	goto login

}

func login(){

	fmt.Println("Welcome to the site\nEnter your UserName and Login")

}
