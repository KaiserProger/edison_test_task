<script setup lang="ts">
  import EnterNumber from '@/components/EnterNumber.vue';
  import ExtraSenseCard from '@/components/ExtraSenseCard.vue';
  import { useGuessesStore } from '@/stores/guesses';
  import { watchEffect } from 'vue';

  const store = useGuessesStore();

  watchEffect(async () => {
    await store.poll();
  });

</script>

<template>
  <div class="container">
    <div class="row row-cols-2 row-cols-md-4 g-4 mt-4">
      <ExtraSenseCard
        v-for="guess in store.transactions"
        :number="guess.number"
        :guessed_by="guess.guessed_by"
        :user_number_id="guess.user_number_id"
        :key="guess.guessed_by"
      />
    </div>
    <EnterNumber />
  </div>
</template>