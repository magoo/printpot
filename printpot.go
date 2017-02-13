package main

import "net"
import "fmt"
import "bufio"
import "os"
import "bytes"

func main() {

  fmt.Println("Launching printpot...")

  // listen on all printer interface
  ln, _ := net.Listen("tcp", ":9100")

  for {
  // accept connection on port
  conn, _ := ln.Accept()
  // run loop forever (or until ctrl-c)

      scanner := bufio.NewScanner(conn)

      for scanner.Scan() {
		      fmt.Println(scanner.Text()) // Println will add back the final '\n'

          if bytes.Contains([]byte(scanner.Text()), []byte("%%EOF")) {
            //conn.Close()
            break
          }
	    }

      if err := scanner.Err(); err != nil {
		      fmt.Fprintln(os.Stderr, "reading standard input:", err)
	    }
      conn.Close()

    }

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
