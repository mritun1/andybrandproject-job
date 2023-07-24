package models

import "time"

// -------------------------------
// STRUCT FOR COLLECTION FIELDS
// -------------------------------
type Users struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	Dob         time.Time `json:"dob"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
