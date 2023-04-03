import React from "react";
import "../../css/Header.css";

function Header({ setToken, username, signOut }) {

    return (
        <div className="header-container">
            <div className="header">
                <img className="icon-img" src="/favicon.ico" style={{width: 40 + 'px'}, {height: 40 + 'px'}}/>
                <h1>PlantWaterer</h1>
                <div className="user-container">
                    <p>Logged in as: {username}</p>
                    <button className="sign-out-button" onClick={signOut}>Sign out</button>
                </div>
            </div>
        </div>
    )
}

export default Header;
