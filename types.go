package main

import "time"

// ApiResponse represents the structure for the entire API response
type ApiResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result Result `json:"result"`
}

// Result represents the structure for the 'result' field in the API response
type Result struct {
	Query    string `json:"query"`
	Total    int    `json:"total"`
	Duration int    `json:"duration"`
	Hits     []Hit  `json:"hits"`
	Links    Links  `json:"links"`
}

// Hit represents the structure for each item in the 'hits' array
type Hit struct {
	Location         Location         `json:"location"`
	IP               string           `json:"ip"`
	AutonomousSystem AutonomousSystem `json:"autonomous_system"`
	Services         []Service        `json:"services"`
	DNS              DNS              `json:"dns"`
	LastUpdatedAt    time.Time        `json:"last_updated_at"`
	OperatingSystem  OperatingSystem  `json:"operating_system,omitempty"`
}

// Service represents the structure for each service in the 'services' array
type Service struct {
	TransportProtocol   string `json:"transport_protocol"`
	ExtendedServiceName string `json:"extended_service_name"`
	ServiceName         string `json:"service_name"`
	Port                int    `json:"port"`
	Certificate         string `json:"certificate,omitempty"`
}

// Location represents the geographical location structure
type Location struct {
	Continent   string      `json:"continent"`
	Coordinates Coordinates `json:"coordinates"`
	Province    string      `json:"province"`
	City        string      `json:"city"`
	CountryCode string      `json:"country_code"`
	PostalCode  string      `json:"postal_code"`
	Timezone    string      `json:"timezone"`
	Country     string      `json:"country"`
}

// Coordinates represents the latitude and longitude coordinates
type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// AutonomousSystem represents the structure for an autonomous system
type AutonomousSystem struct {
	ASN         int    `json:"asn"`
	CountryCode string `json:"country_code"`
	BGPPrefix   string `json:"bgp_prefix"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DNS represents the DNS information structure
type DNS struct {
	ReverseDNS ReverseDNS `json:"reverse_dns"`
}

// ReverseDNS represents the reverse DNS structure
type ReverseDNS struct {
	Names []string `json:"names"`
}

// OperatingSystem represents the operating system information structure
type OperatingSystem struct {
	Product                             string     `json:"product"`
	CPE                                 string     `json:"cpe"`
	Other                               []KeyValue `json:"other"`
	Source                              string     `json:"source"`
	ComponentUniformResourceIdentifiers []string   `json:"component_uniform_resource_identifiers"`
	Part                                string     `json:"part"`
	Vendor                              string     `json:"vendor"`
	Version                             string     `json:"version,omitempty"`
}

// KeyValue represents a generic key-value pair structure
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Links represents the pagination links structure
type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
