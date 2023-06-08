import { fetchTransactionHistory } from "@/network/transactions";
import type { TransactionHistory } from "@/types/dto/transaction_history";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useHistoryStore = defineStore('history', () => {
  const transactions = ref<TransactionHistory[]>([]);
  async function fetch() {
    transactions.value = await fetchTransactionHistory();
  }

  return { transactions, fetch };
});
