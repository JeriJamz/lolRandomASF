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

	goto login

}

func login(){

	fmt.Println("Welcome to the site\nEnter your UserName and Login")

}
