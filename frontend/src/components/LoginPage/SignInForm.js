import React, { useState } from "react";
import axios from "axios";

function SignInForm({ isActive, switchForm, setToken, setPermUsername }) {
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();

    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const loginUser = async (credentials) => {
        const client = axios.create({
            baseURL: "http://localhost:10000"
        })

        return (
        client.post('login', {
            user_name: credentials.username,
            password: credentials.password
        })
        .then(response => response.data)
        .catch((error) => {
            setError(true);
            switch (error.response.data.errorType) {
                case "invalidLogin": {
                    setErrorMessage("The entered login information is invalid!");
                    break;
                }
                default: {
                    setErrorMessage(`Unknown error occured: ${error.response.data.message}`);
                    break;
                }
            };
        })

        )
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        const token = await loginUser ({
            username,
            password
        });

        if (token !== undefined) {
            setToken(token);
            setPermUsername(username);
        }
    }

    return (
        <>
        { isActive ?
        <div className="login-wrapper">
            <h1>Log In</h1>
                <form className="login-form" onSubmit={handleSubmit}>
                    <div className="switch-form-text">
                        Not registered yet?{" "}
                        <span className="switch-form-link" onClick={switchForm}>
                            Sign Up
                        </span>
                    </div>
                    <div className="form-group">
                        <label className="form-label">Username</label>
                        {errorMessage.includes("invalid") ?
                            <input type="text" className="form-control error-form-input" onChange={e => setUserName(e.target.value)} />
                            :
                            <input type="text" className="form-control" onChange={e => setUserName(e.target.value)} />
                        }
                    </div>
                    <div className="form-group">
                        <label className="form-label">Password</label>
                        {errorMessage.includes("invalid") ?
                            <input type="password" className="form-control error-form-input" onChange={e => setPassword(e.target.value)} />
                            :
                            <input type="password" className="form-control" onChange={e => setPassword(e.target.value)} />
                        }
                    </div>
                    <div className="btn-container">
                        <button className="log-in-btn" type="submit" >Log In</button>
                    </div>
                </form>
        {error ? <p className="error-message">{errorMessage}</p>
        : null }
        </div>
        : null }
        </>
    )
}

export default SignInForm;
