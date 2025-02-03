package dto

type CreateTravelPathDto struct {
	UserID        string      `json:"user_id"`
	StartLocation LocationDto `json:"start_location"`
	EndLocation   LocationDto `json:"end_location"`
}

type LocationDto struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
