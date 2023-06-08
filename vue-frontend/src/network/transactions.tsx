import { type GuessTransaction } from "../types/dto/guess_transaction";
import { type TransactionHistory } from "../types/dto/transaction_history";
import { baseApiClient } from "./client"

export const fetchTransactionHistory = async () => {
  const result: any[] = (await baseApiClient.get("/transactions")).data;
  console.log('history', result);
  const transactionHistories: TransactionHistory[] = result.map((value) => {
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
  return transactionHistories;
};
