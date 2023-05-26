import { GuessTransaction } from "../types/dto/guess_transaction";
import { TransactionHistory } from "../types/dto/transaction_history";
import { baseApiClient } from "./client"

export const transactionHistory = async () => {
  let result: any[] = (await baseApiClient.get("/transactions")).data;
  console.log('history', result);
  let transactionHistory: TransactionHistory[] = result.map((value) => {
    return {
      correct_number: value["correct_number"],
      extrasense_guesses: value["extrasense_guesses"].map((x: GuessTransaction) => {
        return {
          guessed_by: x.guessed_by,
          number: x.number,
          user_number_id: x.user_number_id,
        }
      })
    };
  });
  console.log(transactionHistory);
  sessionStorage.setItem("transaction-history", JSON.stringify(transactionHistory));
};
