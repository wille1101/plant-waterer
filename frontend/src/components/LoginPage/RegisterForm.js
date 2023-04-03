import React, { useState } from "react";
import axios from "axios";

function RegisterForm({ isActive, switchForm, setToken, setPermUsername }) {
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();
    const [confirmPassword, setConfirmPassword] = useState();

    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const client = axios.create({
        baseURL: "http://localhost:10000"
    })

    const registerUser = async (credentials) => {
        return (
        client.post('registerUser', {
            user_name: credentials.username,
            password: credentials.password,
            confirm_password: credentials.confirmPassword
        })
        .then((res) => {
            setError(false);
            return false;
        })
        .catch((error) => {
            setError(true);
            switch (error.response.data.errorType) {
                case "passwordsNotMatching": {
                    setErrorMessage("The password and confirm password don't match!");
                    break;
                }
                case "usernameTaken": {
                    setErrorMessage("The entered user name is already taken!");
                    break;
                }
                default: {
                    setErrorMessage(`Unknown error occured: ${error.response.data.message}`);
                    break;
                }
            };
            return true;
        })
        )
    }

    const loginUser = async (credentials) => {
        return (
        client.post('login', {
            user_name: credentials.username,
            password: credentials.password
        })
        .then(response => response.data)
        .catch((error) => {
            if (error.response) {
                setError(true);
            }
        })

        )
    }

    const handleSubmit = async (e) => {
        e.preventDefault();

        let resError = false
        resError = await registerUser ({
            username,
            password,
            confirmPassword
        });

        if (!resError) {
            const token = await loginUser ({
                username,
                password
            });

            if (token !== undefined) {
                setToken(token);
                setPermUsername(username);
            }
        }

    }

    return (
        <>
        { isActive ?
        <div className="register-wrapper">
            <h1>Register</h1>
                <form className="login-form register-form" onSubmit={handleSubmit}>
                    <div className="switch-form-text">
                        Already registered?{" "}
                        <span className="switch-form-link" onClick={switchForm}>
                            Log In
                        </span>
                    </div>
                    <div className="form-group">
                        <label className="form-label">Username</label>
                        {errorMessage.includes("user name") ?
                            <input type="text" className="form-control error-form-input" onChange={e => setUserName(e.target.value)} />
                            :
                            <input type="text" className="form-control" onChange={e => setUserName(e.target.value)} />
                        }
                    </div>
                    <div className="form-group">
                        <label className="form-label">Password</label>
                        <input type="password" className="form-control" onChange={e => setPassword(e.target.value)} />
                    </div>
                    <div className="form-group">
                        <label className="form-label">Confirm Password</label>
                        {errorMessage.includes("password") ?
                            <input type="password" className="form-control error-form-input" onChange={e => setConfirmPassword(e.target.value)} />
                            :
                            <input type="password" className="form-control" onChange={e => setConfirmPassword(e.target.value)} />
                        }
                    </div>
                    <div className="btn-container">
                        <button className="log-in-btn" type="submit" >Register</button>
                    </div>
                </form>
        {error ? <p className="error-message">{errorMessage}</p>
        : null }
        </div>
        : null }
        </>
    )
}

export default RegisterForm;
