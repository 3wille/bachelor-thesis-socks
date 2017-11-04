package main

import (
  "log"
  "net"
  "./socks5"
)

// main function of the executable
func main() {
  // try to get the interface
  ief, err := net.InterfaceByName("eth0")
  if err != nil {
    log.Fatal(err)
  }
  // try to retrieve the addresses configured on above interface
  addrs, err := ief.Addrs()
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("A: %v", addrs)

  srv := socks5.New()

  // This callback is executed after receiving the authentication message of the
  // client. It sets the User and Password field on the Connection Object.
  srv.AuthUsernamePasswordCallback = func(c *socks5.Conn, username, password []byte) error {
    user_string := string(username)
    password_string := string(password)

    // OPTIONAL: block some username
    // if user != "guest" {
    //   return socks5.ErrAuthenticationFailed
    // }

    // OPTIONAL: print received username and password
    // log.Printf("Welcome %v!", user)
    // log.Printf("Password: %v", string(password))

    // save User and Password on the Connection
    c.User = user_string
    c.Password = password_string
    return nil
  }

  // This callback is executed when a CONNECT command was received. The target
  // might be changed or an error return.
  srv.HandleConnectFunc(func(c *socks5.Conn, host string) (newHost string, err error) {
    // if host == "example.com:80" {
    //   return host, socks5.ErrConnectionNotAllowedByRuleset
    // }
    // if user, ok := c.Data.(string); ok {
    //   log.Printf("%v connecting to %v", user, host)
    // }

    // just return the received target host
    return host, nil
  })

  // This callback is executed when the client or target closes the connection.
  srv.HandleCloseFunc(func(c *socks5.Conn) {
    // if user, ok := c.Data.(string); ok {
    //   log.Printf("Goodbye %v!", user)
    // }
  })

  // Start the proxy on a specified address:port
  srv.ListenAndServe(":12345")
}
