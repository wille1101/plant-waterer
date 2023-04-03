import React from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import PlantPage from "./components/PlantPage/PlantPage.js";
import LoginPage from "./components/LoginPage/LoginPage.js";
import useToken from "./components/LoginPage/useToken.js";
import useUsername from "./components/LoginPage/useUsername.js";



function App() {
    const {token, setToken} = useToken();
    const {username, setUsername} = useUsername();

    if (!token) {
        return (
            <BrowserRouter>
                <Routes>
                    <Route path="/login" element={<LoginPage setToken={setToken} setPermUsername={setUsername} />} />
                    <Route path="/*" element={<Navigate to="/login" replace />} />
                </Routes>
            </BrowserRouter>
        );
    }

    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<PlantPage setToken={setToken} username={username} token={token} />} />
                    <Route path="/*" element={<Navigate to="/" replace />} />
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export default App;
