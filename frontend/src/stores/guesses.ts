import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { GuessTransaction } from '@/types/dto/guess_transaction'
import { pollExtrasenses } from '@/network/pollExtrasenses'

export const useGuessesStore = defineStore('guesses', () => {
  const transactions = ref<GuessTransaction[]>([])
  async function poll(){
    transactions.value = await pollExtrasenses()
  }

  return { transactions, poll }
})
