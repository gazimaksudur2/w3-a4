package models

type PropertyResponse struct {
	ID string `json:"ID"`
	GeoInfo GeoInfo `json:"GeoInfo"`
	Property PropertyInfo `json:"Property"`
}

type GeoInfo struct {
	City string `json:"City"`
}

type PropertyInfo struct {
	Name string `json:"Name"`
}