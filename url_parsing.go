package main

import (
	"fmt"
	"net"
	"net/url"
)

//URLs provide a uniform way to locate resources. Here's how to parse URLs in go


func main(){
//we'll parse this exmple url, which includes a scheme authentication infom,host,port,query params, and queryfragment.

	s:= "postgres://user:pass@localhost.com:5432/path?k=v#f"
	u, err := url.Parse(s) //parse url and ensure no errors.
	if err != nil{
		panic(err)
	}
	fmt.Println(u.Scheme) //accesing the straightforward scheme.
	fmt.Println(u.User) //accesing user with password
	fmt.Println(u.User.Username()) //accessing user only
	p, _:=u.User.Password() //accessing password only
	fmt.Println(p)
	fmt.Println(u.Host) //printing the whole host info with port
	host,port,_ :=net.SplitHostPort(u.Host) //split host and port
	fmt.Println(host) //printing host only
	fmt.Println(port)//printing port only

	fmt.Println(u.Path) // accessing the exact path in url /.
	fmt.Println(u.Fragment)//accessing the fragment #.
	fmt.Println(u.RawQuery) // prints raw query
	m,_ := url.ParseQuery(u.RawQuery) //parsing query parameter into map
	fmt.Println(m)
	fmt.Println(m["k"][0])//access index o 

}