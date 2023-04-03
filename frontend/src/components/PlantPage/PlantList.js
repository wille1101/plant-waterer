import React, { useState, useEffect } from "react";
import axios from "axios";
import "../../css/PlantList.css";
import PlantListItem from "./PlantListItem.js";
import AddPlantPopUp from "./AddPlantPopUp.js";

function PlantList({ token, signOut }) {
    const [plants, setPlantsArray] = useState([{}]);

    const client = axios.create({
        baseURL: "http://localhost:10000/auth"
    })

    const config = {
        headers: {
            "Authorization": token,
        },
    };

    useEffect(() => {
        client.get("plants", config)
        .then((response) => {
            setPlants(response.data);
        })
        .catch((error) => {
            if (error.response.data.errorType === "tokenExpired") {
                signOut()
            }
        });
    }, []);

    const setPlants = (plants) => {
        const nextPlants = [...plants];
        nextPlants.sort((a, b) => Date.parse(a.water_within) - Date.parse(b.water_within));
        setPlantsArray(nextPlants);
    }

    const deletePlant = (plant) => {
        client.delete(`plant/${plant.id}`, config);
        setPlants(
            plants.filter((tempPlant) => {
                return tempPlant.id !== plant.id;
            })
        );

    }

    const waterPlant = (plant) => {
        let now = new Date().toISOString().slice(0, 10);
        let last_watered = plant.last_watered.slice(0, 10);

        if (now !== last_watered) {
            client.post(`plant/${plant.id}/water`, {}, config)
            .then(() => {
                const nextPlants = plants.map((p) => {
                    if (p.id === plant.id) {
                        const nextWaterWithin = new Date();
                        nextWaterWithin.setDate(nextWaterWithin.getDate() + p.watering_interval);
                        const nextPlant = {...p, last_watered: now, water_within: nextWaterWithin};

                        return nextPlant;
                    } else {
                        return p;
                    }
                });
                setPlants(nextPlants);
            });
        }
    }

    return (
        <div className="plant-list-container">
            <ul className="plant-list">
                { plants != null ? plants.map((plant) => (
                    <PlantListItem
                        key={`plant-list-item-${plant.id}`}
                        token={token}
                        plant={plant}
                        waterPlant={waterPlant}
                        deletePlant={deletePlant}
                        plants={plants}
                        setPlants={setPlants}
                    />
                )) 
                :
                    <p className="no-plants-text">Add a plant to get started</p>
                }
            </ul>
            <AddPlantPopUp token={token} plants={plants} setPlants={setPlants} />
        </div>
    )
}

export default PlantList
