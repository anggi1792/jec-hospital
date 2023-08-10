package jecconfiguration

//	Declare Database Configuration Entity
type DbConfiguration struct {
	Host       string
	Port       string
	Dbname     string
	Dbuser     string
	Dbpassword string
	Sslmode    string
}
