import { useState } from 'react';

function useToken() {

    const getToken = () => {
        const userToken = localStorage.getItem("userToken");
        return userToken
    };

    const [token, setToken] = useState(getToken());


    const saveToken = (userToken) => {
        localStorage.setItem("userToken", userToken.token);
        setToken(userToken.token);
    }

    return {
        setToken: saveToken,
        token
    }

}

export default useToken;
