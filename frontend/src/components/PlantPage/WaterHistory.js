import React, { useState, useEffect } from "react";
import axios from "axios";
import "../../css/WaterHistory.css";
import Calender from 'react-github-contribution-calendar';

function WaterHistory({ token, plant }) {
    const client = axios.create({
        baseURL: "http://localhost:10000/auth"
    })

    const [isShowing, setIsShowing] = useState(false);
    const [refresh, setRefresh] = useState(false);
    const [timeStamps, setTimeStamps] = useState([]);

    const config = {
        headers: {
            "Authorization": token,
        },
    };

    useEffect(() => {
        if (isShowing) {
            client.get(`plant/${plant.id}/water-history`, config).then((response) => {
                setTimeStamps(response.data);
            });
        }
    }, [plant, refresh]);

    const handleClick = () => {
        setRefresh(true);
        setIsShowing(!isShowing);
    }

    const values = {};

    if (timeStamps != null) {
        timeStamps.map((timeStamp) => (
            values[`${timeStamp.TimeStamp.slice(0, 10)}`] = 4
        ));
    }

    var panelColors = [
    '#EEEEEE',
    '#028fb3',
    '#028fb3',
    '#028fb3',
    '#028fb3'
    ];

    let until = new Date().toISOString().slice(0, 10);

    return (
        <div className="water-history-container">
            { !isShowing ? 
            <p  className="water-history-btn" onClick={() => handleClick()}>
                <i className="arrow down"></i>
            </p>
            : null }
            { isShowing ?
            <>
            <p  className="water-history-btn" onClick={() => setIsShowing(!isShowing)}>
                <i className="arrow up"></i>
            </p>
            <div className="calender-container">
                <Calender values={values} until={until} panelColors={panelColors}/>
            </div>
            </>
            : null }
        </div>
    )
}

export default WaterHistory
