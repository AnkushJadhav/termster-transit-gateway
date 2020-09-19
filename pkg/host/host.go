package host

// Host represents a remote server to communicate with
type Host struct {
	addr string
	port int
}

// New creates a new host with address addr and port p
func New(addr string, p int) *Host {
	h := &Host{
		addr: addr,
		port: p,
	}

	return h
}

// String gives a string representation of the host h
func (h *Host) String() string {
	return h.addr + ":" + string(h.port)
}
