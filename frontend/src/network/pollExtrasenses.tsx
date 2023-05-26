import { GuessTransaction } from "../types/dto/guess_transaction";
import { baseApiClient } from "./client"

export const pollExtrasenses = async () => {
  let result: any[] = (await baseApiClient.get("/poll")).data;
  console.log("pollExtra", result)
  let dtos: GuessTransaction[] = result.map((value) => {
    return {
      guessed_by: value["guessed_by"],
      number: value["number"],
      user_number_id: value["user_number_id"],
    };
  });
  sessionStorage.setItem("guesses", JSON.stringify(dtos));
  return dtos
}