import axios from "axios";

export const baseApiClient = axios.create({
    baseURL: 'http://backend:8080/',
    timeout: 1000,
    withCredentials: true,
});
