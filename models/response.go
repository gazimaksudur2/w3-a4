package models

type PropertyListResponse struct {
	Result Result `json:"Result"`
}

type Result struct {
	Count int                `json:"Count"`
	Items []PropertyResponse `json:"Items"`
}

type PropertyResponse struct {
	ID        string       `json:"ID"`
	Feed      int          `json:"Feed"`
	Published bool         `json:"Published"`
	GeoInfo   GeoInfo      `json:"GeoInfo"`
	Property  PropertyInfo `json:"Property"`
}

type GeoInfo struct {
	City        string   `json:"City"`
	Country     string   `json:"Country"`
	CountryCode string   `json:"CountryCode"`
	Name        string   `json:"Name"`
	LocationID  string   `json:"LocationID"`
	State       string   `json:"State"`
	StateAbbr   string   `json:"StateAbbr"`
	Lat         float64  `json:"Lat"`
	Lon         float64  `json:"Lon"`
	Breadcrumbs []string `json:"Breadcrumbs"`
}

type PropertyInfo struct {
	Name         string         `json:"Name"`
	Slug         string         `json:"Slug"`
	PropertyType string         `json:"PropertyType"`
	Price        float64        `json:"Price"`
	ReviewScore  float64        `json:"ReviewScore"`
	StarRating   int            `json:"StarRating"`
	Counts       PropertyCounts `json:"Counts"`
	Amenities    []string       `json:"Amenities"`
	Image        ImageInfo      `json:"Image"`
}

type ImageInfo struct {
	Count  int      `json:"Count"`
	Images []string `json:"Images"`
}

type PropertyCounts struct {
	Bedroom   int     `json:"Bedroom"`
	Bathroom  int     `json:"Bathroom"`
	Reviews   float64 `json:"Reviews"`
	Occupancy int     `json:"Occupancy"`
}
