package dto

type ShortRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortResponse struct {
	Shorten string `json:"shorten"`
}