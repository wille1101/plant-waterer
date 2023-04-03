import React, { useState } from "react";
import "../../css/LoginPage.css";
import SignInForm from "./SignInForm.js";
import RegisterForm from "./RegisterForm.js";

function LoginPage({ setToken, setPermUsername }) {
    const [activeIndex, setActiveIndex] = useState(0);

    return (
        <>
        <div className="header-container">
            <div className="header">
                <img className="icon-img" src="/favicon.ico" style={{width: 40 + 'px'}, {height: 40 + 'px'}}/>
                <h1>PlantWaterer</h1>
            </div>
        </div>
        <SignInForm isActive={activeIndex === 0} switchForm={() => setActiveIndex(1)} setToken={setToken} setPermUsername={setPermUsername} />
        <RegisterForm isActive={activeIndex === 1} switchForm={() => setActiveIndex(0)} setToken={setToken} setPermUsername={setPermUsername} />
        </>
    )
}

export default LoginPage;
