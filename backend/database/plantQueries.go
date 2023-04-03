package database

import (
    "log"

    "github.com/wille1101/plant-waterer/backend/models"

    "github.com/google/uuid"
)

// GetPlantsQuery runs a select query on the database, returning all plants with the given
// UUID.
func GetPlantsQuery(uuid uuid.UUID) ([]models.Plant, error) {
    db := createConnection()
    defer db.Close()

    var plants []models.Plant

    sqlQuery := "SELECT * FROM plants WHERE owner_id = $1 ORDER BY water_within ASC"
    rows, err := db.Query(sqlQuery, uuid)
    if (err != nil) {
        log.Fatalf("Unable to execute query! %v : %s", err, sqlQuery)
    }
    defer rows.Close()

    for rows.Next() {
        var plant models.Plant

        err = rows.Scan(&plant.ID, &plant.OwnerId, &plant.Name, &plant.LatinName, &plant.LastWatered, &plant.WateringInterval, &plant.WaterWithin)
        if (err != nil) {
            log.Fatalf("Unable to scan database rows. %v", err)
        }

        plants = append(plants, plant)
    }

    return plants, err
}

// InsertPlantQuery runs an insert query on the database, inserting the provided plant model with the
// given UUID.
func InsertPlantQuery(plant models.Plant, uuid uuid.UUID) int64 {
    db := createConnection()
    defer db.Close()

    sqlQuery := "INSERT INTO plants (owner_id, name, latin_name, watering_interval) VALUES ($1, $2, $3, $4) RETURNING plant_id"

    var plantId int64

    err := db.QueryRow(sqlQuery, uuid, plant.Name, plant.LatinName, plant.WateringInterval).Scan(&plantId)
    if (err != nil) {
        log.Fatalf("Unable to execute insert query")
    }

    return plantId
}

// DeletePlantQuery deletes a plant from the database with the given plant id and owner UUID.
func DeletePlantQuery(plantId int64, uuid uuid.UUID) {
    db := createConnection()
    defer db.Close()

    sqlQuery := "DELETE FROM plants WHERE plant_id = $1 AND owner_id = $2"

    _, err := db.Exec(sqlQuery, plantId, uuid)
    if (err != nil) {
        log.Fatalf("Unable to delete plant in database. %v", err)
    }

}

// UpdatePlantQuery updates a plant in the database with the given plant model and owner UUID.
func UpdatePlantQuery(plantId int64, plant models.Plant, uuid uuid.UUID) {
    db := createConnection()
    defer db.Close()

    sqlQuery := "UPDATE plants SET name=$3, latin_name=$4, watering_interval=$5 WHERE plant_id=$1 AND owner_id = $2"

    _, err := db.Exec(sqlQuery, plantId, uuid, plant.Name, plant.LatinName, plant.WateringInterval)
    if (err != nil) {
        log.Fatalf("Unable to update plant %v", err)
    }

}

// WaterPlantQuery updates the last_watered attribute of a plant to the current date.
func WaterPlantQuery(plantId int64, uuid uuid.UUID) {
    db := createConnection()
    defer db.Close()

    sqlQuery := "UPDATE plants SET last_watered = now() WHERE plant_id = $1 AND owner_id = $2"

    _, err := db.Exec(sqlQuery, plantId, uuid)
    if (err != nil) {
        log.Fatalf("Unable to water plant in database. %v", err)
    }

}

// GetWaterHistoryQuery queries the database of all watering timestamps for the given plant id and owner UUID.
func GetWaterHistoryQuery(plantId int64, uuid uuid.UUID) ([]models.TimeStamp, error) {
    db := createConnection()
    defer db.Close()

    var timeStamps []models.TimeStamp

    sqlQuery := "SELECT timestamp FROM waterings WHERE plant_id = $1 AND owner_id = $2"

    rows, err := db.Query(sqlQuery, plantId, uuid)
    if (err != nil) {
        log.Fatalf("Unable to water plant in database. %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var timeStamp models.TimeStamp

        err = rows.Scan(&timeStamp.TimeStamp)
        if (err != nil) {
            log.Fatalf("Unable to scan database rows. %v", err)
        }

        timeStamps = append(timeStamps, timeStamp)
    }

    return timeStamps, err
}
