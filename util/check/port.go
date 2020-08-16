package check

import (
	"net"
)

func PortAvailable(port string) bool {
	//conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 5*time.Second)
	//if err != nil {
	//	return false
	//}
	//if conn != nil {
	//	defer conn.Close()
	//	return true
	//}
	//return false

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		//fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
		//os.Exit(1)
		return false
	}

	defer ln.Close()
	//fmt.Printf("TCP Port %q is available", port)
	//os.Exit(0)
	return true
}
