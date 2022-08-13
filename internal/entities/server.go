package entities

type Server struct {
	Address  string
	Port     string
	Users    []Client
	FileName string
}
