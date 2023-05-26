import { baseApiClient } from "./client";

export const checkAnswers = async (number_: number) => {
    await baseApiClient.get("/check", {
        params: {
            "number": number_,
        }
    })
};