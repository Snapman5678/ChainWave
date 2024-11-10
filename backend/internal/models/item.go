package models

import "github.com/google/uuid"

// Item struct
type Item struct {
	Id              uuid.UUID `form:"id"`
	BusinessAdminId uuid.UUID `form:"business_admin_id"`
	Name            string    `form:"name"`
	Description     string    `form:"description"`
	Price           float64   `form:"price"`
	Weight          float64   `form:"weight"`
	Dimensions      string    `form:"dimensions"`
	Category        string    `form:"category"`
	Quantity        int       `form:"quantity"`
	ImageURL        string    `form:"image_url"`
}

// ItemWithDetail struct includes business admin and location details
type ItemWithDetail struct {
	Id                       uuid.UUID `json:"id"`
	BusinessAdminId          uuid.UUID `json:"business_admin_id"`
	Name                     string    `json:"name"`
	Description              string    `json:"description"`
	Price                    float64   `json:"price"`
	Weight                   float64   `json:"weight"`
	Dimensions               string    `json:"dimensions"`
	Category                 string    `json:"category"`
	Quantity                 int       `json:"quantity"`
	ImageURL                 string    `json:"image_url"`
	BusinessAdminCompanyName string    `json:"business_admin_company_name"`
	BusinessAdminContactInfo string    `json:"business_admin_contact_info"`
	LocationAddress          string    `json:"location_address"`
	LocationCity             string    `json:"location_city"`
	LocationState            string    `json:"location_state"`
}
