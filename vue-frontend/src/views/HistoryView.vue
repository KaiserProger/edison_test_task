<script setup lang="ts">
  import { useHistoryStore } from '@/stores/history';
  import { watchEffect } from 'vue';

  const store = useHistoryStore();

  watchEffect(async () => {
    await store.fetch();
  })
</script>

<template>
  <div class="container">
    <div class="row mt-2">
      <div class="col">
        <table class="table" v-if="store.transactions.length > 0">
          <thead>
            <tr>
              <th>Number</th>
              <th
                style="word-wrap: break-word; min-width: 160px; max-width: 160px;"
                v-for="(value, index) in store.transactions[0].extrasense_guesses"
                :key="index"
              >
                {{ value.guessed_by.slice(0, 9) }}
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(transaction, index) in store.transactions"
              :key="index"
            >
              <td>{{ transaction.correct_number }}</td>
              <td
                v-for="(guess, _index) in transaction.extrasense_guesses"
                :key="_index"
              >
                {{ guess.number }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>