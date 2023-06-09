<script setup lang="ts">
  import { watchEffect } from 'vue';
  import { useTrustStore } from '@/stores/trust';

  const store = useTrustStore();
  watchEffect(async () => {
    await store.fetch();
  });
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col mt-4">
        <p
          v-for="(value, index) in store.trustLevels"
          :key="index"
        >
          {{ value.extrasenseId }}:{{ value.trust }}
        </p>
      </div>
    </div>
  </div>
</template>