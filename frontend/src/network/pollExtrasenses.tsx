import { type GuessTransaction } from "../types/dto/guess_transaction";
import { baseApiClient } from "./client"

export const pollExtrasenses = async () => {
  const result: any[] = (await baseApiClient.get("/poll")).data;
  console.log("pollExtra", result)
  const dtos: GuessTransaction[] = result.map((value) => {
    return {
      guessed_by: value["guessed_by"],
      number: value["number"],
      user_number_id: value["user_number_id"],
    };
  });
  return dtos;
};