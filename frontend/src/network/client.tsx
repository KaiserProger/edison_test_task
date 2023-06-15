import axios from "axios";

export const baseApiClient = axios.create({
    baseURL: "http://80.243.141.49:8080",
    timeout: 1000,
    withCredentials: true,
});
