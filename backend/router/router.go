package router

import (
    "github.com/gorilla/mux"

    "github.com/wille1101/plant-waterer/backend/routes"
    "github.com/wille1101/plant-waterer/backend/middlewares"
)

// Router creates a new router with the required routes and handler functions and returns it.
func Router() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/", routes.HomePage)

    r.HandleFunc("/login", routes.OptionsPOSTOK).Methods("OPTIONS")
    r.HandleFunc("/login", routes.GenerateLoginToken).Methods("POST")

    r.HandleFunc("/registerUser", routes.OptionsPOSTOK).Methods("OPTIONS")
    r.HandleFunc("/registerUser", routes.RegisterUser).Methods("POST")

    r.HandleFunc("/auth/plants", routes.OptionsGETOK).Methods("OPTIONS")
    r.HandleFunc("/auth/plant", routes.OptionsPOSTOK).Methods("OPTIONS")
    r.HandleFunc("/auth/plant/{plantId}", routes.OptionsDELETEAndPOSTOK).Methods("OPTIONS")
    r.HandleFunc("/auth/plant/{plantId}/water", routes.OptionsPOSTOK).Methods("OPTIONS")
    r.HandleFunc("/auth/plant/{plantId}/water-history", routes.OptionsGETOK).Methods("OPTIONS")

    authRoutes := r.PathPrefix("/auth").Subrouter()

    authRoutes.HandleFunc("/plants", routes.GetPlants).Methods("GET")

    authRoutes.HandleFunc("/plant", routes.CreatePlant).Methods("POST")

    authRoutes.HandleFunc("/plant/{plantId}", routes.DeletePlant).Methods("DELETE")

    authRoutes.HandleFunc("/plant/{plantId}", routes.UpdatePlant).Methods("POST")

    authRoutes.HandleFunc("/plant/{plantId}/water", routes.WaterPlant).Methods("POST")

    authRoutes.HandleFunc("/plant/{plantId}/water-history", routes.GetWaterHistory).Methods("GET")

    authRoutes.Use(middlewares.CheckToken)

    return r
}
