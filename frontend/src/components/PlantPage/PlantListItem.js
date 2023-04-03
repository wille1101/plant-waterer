import React from "react";
import SettingsPopUp from "./SettingsPopUp.js";
import WaterHistory from "./WaterHistory.js";

function PlantListItem({ token, plant, waterPlant, deletePlant, plants, setPlants }) {
    if (plant.id != null) {
        let waterWithin = Date.parse(plant.water_within);
        let now = Date.now();
        let waterWithinDays = waterWithin - now;
        let diffDays = Math.ceil(waterWithinDays / (1000 * 60 * 60 * 24)); 

        return (
            <li className="plant-list-item">
                <div className="plant">
                    <div className="plant-name-container">
                        <h1 className="plant-name">{plant.name}</h1>
                        <h4 className="latin-name">{plant.latin_name}</h4>
                    </div>
                    <div className="plant-info">
                        { diffDays <= 1 ?
                        <p>Water within: <span className="text-urgent">{diffDays} day</span></p>
                        : 
                        <p>Water within: {diffDays} days</p>
                        }
                        <p>Last watered: {plant.last_watered.slice(0, 10)}</p>
                    </div>
                </div>
                <button className="water-btn" onClick={() => waterPlant(plant)}>
                    <span>Water</span>
                </button>
                <SettingsPopUp 
                    token={token}
                    plant={plant}
                    deletePlant={deletePlant}
                    plants={plants}
                    setPlants={setPlants}/>
                <WaterHistory token={token} plant={plant}/>
            </li>
        )
    }
}

export default PlantListItem;
