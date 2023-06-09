import { type GuessTransaction } from "./guess_transaction";

export interface TransactionHistory {
    correct_number: number;
    extrasense_guesses: GuessTransaction[];
};