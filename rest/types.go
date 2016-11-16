package rest

type Site struct {
	Name        string
	Description string

	// IDs
	ID int
}

type Network struct {
	Address string `json:network_address"`
	// May be map[string]string or map[string][]string
	Attributes   interface{}
	IsIP         bool
	PrefixLength int `json:prefix_length"`
	State        string
	Version      string

	// IDs
	ID       int
	ParentID int `json:parent_id"`
	SiteID   int `json:site_id"`
}

type Device struct {
	// May be map[string]string or map[string][]string
	Attributes interface{}
	Hostname   string

	// IDs
	ID     int
	SiteID int `json:site_id"`
}

type Interface struct {
	Addresses []string
	// May be map[string]string or map[string][]string
	Attributes  interface{}
	Description string
	MacAddress  string `json:"mac_address"`
	Name        string
	Networks    []string
	Speed       int
	Type        int

	// IDs
	ID       int
	Device   int `json:"device"`
	ParentID int `json:"parent_id"`
}

type Circuit struct {
	AEndpoint int `json:"a_endpoint"`
	ZEndpoint int `json:"z_endpoint"`
	Name      string
	// May be map[string]string or map[string][]string
	Attributes interface{}

	//IDs
	ID int
}

type Attribute struct {
	Constraints  Constraints
	Description  string
	Display      bool
	Multi        bool
	Name         string
	Required     bool
	ResourceName string `json:"resource_name"`

	// IDs
	ID     int
	SiteID int `json:site_id"`
}

type Constraints struct {
	Pattern     string
	ValidValues []string `json:"valid_values"`
	AllowEmpty  bool     `json:"allow_empty"`
}

type User struct {
	ID          int
	Email       string
	Permissions map[string]Permission
}

type Permission struct {
	UserID      string `json:"user_id"`
	SiteID      string `json:"site_id"`
	Permissions []string
}
