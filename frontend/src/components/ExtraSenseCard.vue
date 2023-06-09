<script setup lang="ts">
  import type { GuessTransaction } from '@/types/dto/guess_transaction';

  const props = defineProps<GuessTransaction>();
</script>

<template>
  <div class="col">
    <div class="card mx-2" style="width: 18rem;">
      <div class="card-body">
        <div class="card-title">Extrasense {{ props.guessed_by.slice(0, 9) }}</div>
        <div class="card-text">"Number is {{ props.number }}"</div>
      </div>
    </div>
  </div>
</template>