// Package network provides the TCP authoritative server and client networking.
package network

// Server represents the authoritative game server.
type Server struct {
	Address  string
	Port     int
	TickRate int
}

// NewServer creates and returns a new Server.
func NewServer(address string, port, tickRate int) *Server {
	return &Server{
		Address:  address,
		Port:     port,
		TickRate: tickRate,
	}
}

// Start begins the server listen loop.
func (s *Server) Start() error {
	// Skeleton: server start logic
	return nil
}

// Stop gracefully shuts down the server.
func (s *Server) Stop() error {
	// Skeleton: server stop logic
	return nil
}

// Client represents a game client connection.
type Client struct {
	ServerAddress string
	ServerPort    int
}

// NewClient creates and returns a new Client.
func NewClient(address string, port int) *Client {
	return &Client{
		ServerAddress: address,
		ServerPort:    port,
	}
}

// Connect establishes a connection to the server.
func (c *Client) Connect() error {
	// Skeleton: client connect logic
	return nil
}

// Disconnect closes the connection to the server.
func (c *Client) Disconnect() error {
	// Skeleton: client disconnect logic
	return nil
}
