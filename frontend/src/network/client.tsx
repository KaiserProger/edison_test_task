import axios from "axios";

export const baseApiClient = axios.create({
    baseURL: process.env.REACT_APP_SERVER_IP,
    timeout: 1000,
    withCredentials: true,
});
