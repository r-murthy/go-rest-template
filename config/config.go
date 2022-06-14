package config

const (
	defaultCustomersPageSize = 5
	defaultItemsPageSize
	defaultInvoicesSize
)

// APIConfig is a stores API configuration to be used across the application.
var APIConfig Config = Config{
	CustomersPageSize: defaultCustomersPageSize,
	ItemsPageSize:     defaultItemsPageSize,
	InvoicesPageSize:  defaultInvoicesSize,
}

// Config is a container of api configuration data.
type Config struct {
	Port              string `yaml:"port"`
	PGConn            string `yaml:"pg_conn"`
	StaticDir         string `yaml:"static_dir"`
	AuthSecret        string `yaml:"authsecret"`
	FBServiceFile     string `yaml:"service_file_location"`
	CustomersPageSize int    `yaml:"customers_page_size"`
	ItemsPageSize     int    `yaml:"items_page_size"`
	InvoicesPageSize  int    `yaml:"invoices_page_size"`
}
