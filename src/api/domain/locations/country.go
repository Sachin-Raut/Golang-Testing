package locations

//Country is
type Country struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	TimeZone       string `json:"time_zone"`
	GeoInformation GeoInformation
	State          []State
}

//GeoInformation is
type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

//GeoLocation is
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//State is
type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
