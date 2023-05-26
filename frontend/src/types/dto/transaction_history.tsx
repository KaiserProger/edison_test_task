import { GuessTransaction } from "./guess_transaction";

export type TransactionHistory = {
    correct_number: number;
    extrasense_guesses: GuessTransaction[];
};