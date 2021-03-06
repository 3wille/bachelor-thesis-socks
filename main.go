package main

import (
	"log"
	"net"

	"./socks5"
	"github.com/jinzhu/configor"
)

// config file schema
var Config = struct {
	Prefixes      []string `required: true`
	SingleAddress bool     `default:false`
}{}

// main function of the executable
func main() {
	// load config with above schema
	configor.Load(&Config, "config.yml")
	log.Printf("config: %#v", Config.Prefixes)

	// parse strings from config file to IP objects
	var prefixes []net.IP
	for _, prefix := range Config.Prefixes {
		prefixIP, _, err := net.ParseCIDR(prefix)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		prefixes = append(prefixes, prefixIP)
	}

	srv := socks5.New(prefixes)

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
		if !Config.SingleAddress {
			c.User = user_string
			c.Password = password_string
		} else {
			c.User = "A"
			c.Password = "B"
		}
		return nil
	}

	// This callback is executed when no authentication was provided.
	// It just accepts the incoming connection and uses the same IP address
	// for each request. This is used for debugging purposes.
	srv.AuthNoAuthenticationRequiredCallback = nil
	// srv.AuthNoAuthenticationRequiredCallback = func(c *socks5.Conn) error {
	// 	c.User = "no_auth"
	// 	c.Password = "no_auth"
	// 	return nil
	// }

	// This callback is executed when a CONNECT command was received. The target
	// might be changed or an error returned.
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
