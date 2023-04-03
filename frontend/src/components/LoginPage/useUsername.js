import { useState } from 'react';

function useUsername() {

    const getUsername = () => {
        const username = localStorage.getItem("username");
        return username
    };

    const [username, setUsername] = useState(getUsername());


    const saveUsername = (username) => {
        localStorage.setItem("username", username);
        setUsername(username);
    }

    return {
        setUsername: saveUsername,
        username
    }

}

export default useUsername;
