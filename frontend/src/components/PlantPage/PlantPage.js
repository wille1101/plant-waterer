import React from "react";
import Header from "./Header.js";
import PlantList from "./PlantList.js";


function PlantPage({ setToken, username, token }) {
    const signOut = () => {
        setToken("");
        localStorage.clear();
    }

    return (
        <div className="plantpage">
            <Header setToken={setToken} username={username} signOut={signOut} />
            <PlantList token={token} signOut={signOut}/>
        </div>
    )
}

export default PlantPage;
