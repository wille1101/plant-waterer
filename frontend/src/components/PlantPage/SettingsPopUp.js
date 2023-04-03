import React, { useState } from "react";
import axios from "axios";
import "../../css/PopUp.css";
import "../../css/SettingsPopUp.css";

function SettingsPopUp({ token, plant, deletePlant, plants, setPlants }) {
    const [isShowing, setIsShowing] = useState(false);

    const [localPlant, setLocalPlant] = useState({ name: plant.name, latin_name: plant.latin_name, watering_interval: plant.watering_interval});

    const client = axios.create({
        baseURL: "http://localhost:10000/auth"
    })

    const config = {
        headers: {
            Authorization: token,
        },
    };

    const updatePlant = (localPlant) => {
        client.post(`plant/${plant.id}`, {
            name: localPlant.name,
            latin_name: localPlant.latin_name,
            watering_interval: localPlant.watering_interval
        }, config)
        .then(() => {
            const nextPlants = plants.map((p) => {
                if (p.id === plant.id) {
                    const nextWaterWithin = new Date();
                    nextWaterWithin.setDate(nextWaterWithin.getDate() + localPlant.watering_interval);

                    const nextPlant = {...p,
                        name: localPlant.name,
                        latin_name: localPlant.latin_name,
                        watering_interval: localPlant.watering_interval,
                        water_within: nextWaterWithin
                    }
                    return nextPlant;
                } else {
                    return p;
                }
            });
            setPlants(nextPlants);
        });
    }

    const handleChange = (e) => {
        if (e.target.name === "watering_interval") {
            setLocalPlant({ ...localPlant, [e.target.name]: Number(e.target.value) });
        } else {
            setLocalPlant({ ...localPlant, [e.target.name]: e.target.value });
        }
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        updatePlant(localPlant);
        setIsShowing(false);
    }

    return (
        <div>
            <button  className="settings-btn fa fa-gear" onClick={() => setIsShowing(true)}></button>
            {isShowing ? 
            <div className="popup-container settings-popup">
                <div className="popup-body">
                    <h1 className="popup-header">{plant.name}</h1>
                    <form className="edit-plant-form" onSubmit={handleSubmit}>
                        <div className="form-group">
                            <label className="form-label">Plant Name</label>
                            <input type="text" className="form-control" name="name" defaultValue={plant.name} onChange={handleChange} maxLength="40"/>
                        </div>
                        <div className="form-group">
                            <label className="form-label">Latin Plant Name</label>
                            <input type="text" className="form-control" name="latin_name" defaultValue={plant.latin_name} onChange={handleChange} maxLength="40"/>
                        </div>
                        <div className="form-group">
                            <label className="form-label">Watering Interval</label>
                            <input type="number" className="form-control" name="watering_interval" defaultValue={plant.watering_interval} onChange={handleChange} />
                        </div>
                        <div className="btn-container">
                            <button className="delete-btn" onClick={() => deletePlant(plant)}>Delete</button>
                            <button className="close-btn" onClick={() => setIsShowing(false)}>Cancel</button>
                            <button className="save-btn" type="submit" >Save Changes</button>
                        </div>
                    </form>
                </div>
            </div>
            : null }
        </div>
    )
}

export default SettingsPopUp
