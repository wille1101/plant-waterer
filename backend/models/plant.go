package models

import (
    "github.com/google/uuid"
)

// Plant is the model of a plant object.
type Plant struct {
    ID               int64     `json:"id,omitempty"`
    OwnerId          uuid.UUID `json:"uuid,omitempty"`
    Name             string    `json:"name"`
    LatinName        string    `json:"latin_name"`
    LastWatered      string    `json:"last_watered,omitempty"`
    WateringInterval int64     `json:"watering_interval"`
    WaterWithin      string    `json:"water_within,omitempty"`
}
