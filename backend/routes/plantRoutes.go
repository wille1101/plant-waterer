package routes

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "strconv"

    "github.com/wille1101/plant-waterer/backend/models"
    "github.com/wille1101/plant-waterer/backend/auth"
    db "github.com/wille1101/plant-waterer/backend/database"

    "github.com/gorilla/mux"
)

// GetPlants handles the endpoint which returns all plants.
func GetPlants(w http.ResponseWriter, r *http.Request) {
    uuid, err := auth.GetUUIDFromToken(r.Header.Get("Authorization"))

    plants, err := db.GetPlantsQuery(uuid)
    if (err != nil) {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Database error")
        return
    }

    log.Println("Route accessed: GetPlants")

    data, err := json.Marshal(plants)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(data)
}

// CreatePlant handles the endpoint which creates a new plant.
func CreatePlant(w http.ResponseWriter, r *http.Request) {
    uuid, err := auth.GetUUIDFromToken(r.Header.Get("Authorization"))
    var plant models.Plant

    d := json.NewDecoder(r.Body)
    d.DisallowUnknownFields()
    err = d.Decode(&plant)
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    log.Println("Route accessed: CreatePlant")

    plantId := db.InsertPlantQuery(plant, uuid)

    res := response {
        ID: plantId,
        Message: "Plant inserted",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(data)
}

// DeletePlant handles the endpoint which deletes a plant.
func DeletePlant(w http.ResponseWriter, r *http.Request) {
    uuid, err := auth.GetUUIDFromToken(r.Header.Get("Authorization"))
    params := mux.Vars(r)
    plantId, err := strconv.Atoi(params["plantId"])
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    log.Println("Route accessed: DeletePlant")

    db.DeletePlantQuery(int64(plantId), uuid)

    res := response {
        ID: int64(plantId),
        Message: "Plant deleted",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(data)
}

// UpdatePlant handles the endpoint which updates a plant.
func UpdatePlant(w http.ResponseWriter, r *http.Request) {
    uuid, err := auth.GetUUIDFromToken(r.Header.Get("Authorization"))
    params := mux.Vars(r)
    plantId, err := strconv.Atoi(params["plantId"])
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    log.Println("Route accessed: UpdatePlant")

    var plant models.Plant

    d := json.NewDecoder(r.Body)
    d.DisallowUnknownFields()
    err = d.Decode(&plant)
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    db.UpdatePlantQuery(int64(plantId), plant, uuid)

    res := response {
        ID: int64(plantId),
        Message: "Plant updated",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(data)
}

// WaterPlant handles the endpoint which updates a plant's last_watered attribute.
func WaterPlant(w http.ResponseWriter, r *http.Request) {
    uuid, err := auth.GetUUIDFromToken(r.Header.Get("Authorization"))
    params := mux.Vars(r)
    plantId, err := strconv.Atoi(params["plantId"])
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    log.Println("Route accessed: WaterPlant")

    db.WaterPlantQuery(int64(plantId), uuid)

    res := response {
        ID: int64(plantId),
        Message: "Plant watered",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(data)

}

// GetWaterHistory handles the endpoint which returns a plants watering time stamps.
func GetWaterHistory(w http.ResponseWriter, r *http.Request) {
    uuid, err := auth.GetUUIDFromToken(r.Header.Get("Authorization"))
    params := mux.Vars(r)
    plantId, err := strconv.Atoi(params["plantId"])
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    waterHistory, err := db.GetWaterHistoryQuery(int64(plantId), uuid)
    if (err != nil) {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Database error")
        return
    }

    log.Println("Route accessed: GetWaterHistory")

    data, err := json.Marshal(waterHistory)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(data)
}

