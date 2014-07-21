package main
import "fmt"
import "net"
func main() {

  conn, err := net.Dial("tcp", "google.com:80")
  if err != nil {
    //handle error
  fmt.Println("error")
}
  fmt.Fprintf(conn, "GET / HTTP/1.0\r\\n\r\n")


}
