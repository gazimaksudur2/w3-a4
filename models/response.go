package models

type PropertyListResponse struct {
	Result Result `json:"Result"`
}

type Result struct {
	Count int `json:"Count"`
	Items []PropertyResponse `json:"Items"`
}

type PropertyResponse struct {
	ID string `json:"ID"`
	GeoInfo GeoInfo `json:"GeoInfo"`
	Property PropertyInfo `json:"Property"`
}

type GeoInfo struct {
	Country string `json:"Country"`
	State string `json:"State"`
	City string `json:"City"`
	Latitude float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Breadcrumbs []string `json:"Breadcrumbs"`
}

type PropertyInfo struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
	Price float64 `json:"Price"`
	Bedroom int `json:"Bedroom"`
	Bathroom int `json:"Bathroom"`
	Rating float64 `json:"Rating"`
	Reviews int `json:"Reviews"`
	Amenities []string `json:"Amenities"`
	Images []string `json:"Images"`
}