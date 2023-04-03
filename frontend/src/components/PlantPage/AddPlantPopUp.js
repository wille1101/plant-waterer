import React, { useState } from "react";
import axios from "axios";
import "../../css/PopUp.css";
import "../../css/AddPlantPopUp.css";

function AddPlantPopUp({ token, plants, setPlants }) {
    const [isShowing, setIsShowing] = useState(false);

    const [plant, setPlant] = useState({ name: '', latin_name: 'Unknown Latin Name', watering_interval: 7});

    const client = axios.create({
        baseURL: "http://localhost:10000/auth"
    })

    const config = {
        headers: {
            "Authorization": token,
        },
    };

    const addPlant = (name) => {
        client.post('plant', {
                name: plant.name,
                latin_name: plant.latin_name,
                watering_interval: plant.watering_interval
            }, config)
            .then((response) => {
                const currentDate = new Date().toISOString().slice(0, 10);
                const nextWaterWithin = new Date();
                nextWaterWithin.setDate(nextWaterWithin.getDate() + plant.watering_interval);

                setPlants([...plants, {
                    id: response.data.id,
                    name: plant.name,
                    latin_name: plant.latin_name,
                    last_watered: currentDate,
                    watering_interval: plant.watering_interval,
                    water_within: nextWaterWithin
                }]);
                //refreshPlants();
            });
    };

    const handleChange = (e) => {
        if (e.target.name === "watering_interval") {
            setPlant({ ...plant, [e.target.name]: Number(e.target.value) });
        } else {
            setPlant({ ...plant, [e.target.name]: e.target.value });
        }
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        addPlant(plant);
        setIsShowing(false);
    }

    return (
        <div className="add-plant-container">
        <button  className="add-btn fa fa-plus" onClick={() => setIsShowing(true)}>
        </button>
        {isShowing ? 
        <div className="popup-container addplant-popup">
            <div className="popup-body">
                <h1 className="popup-header">Add a plant</h1>
                <form className="edit-plant-form" onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label className="form-label">Plant Name</label>
                        <input type="text" className="form-control" name="name" onChange={handleChange} maxLength="40"/>
                    </div>
                    <div className="form-group">
                        <label className="form-label">Latin Plant Name</label>
                        <input type="text" className="form-control" name="latin_name" onChange={handleChange} maxLength="40"/>
                    </div>
                    <div className="form-group">
                        <label className="form-label">Watering Interval</label>
                        <input type="number" className="form-control" name="watering_interval" onChange={handleChange} />
                    </div>
                    <div className="btn-container">
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

export default AddPlantPopUp
