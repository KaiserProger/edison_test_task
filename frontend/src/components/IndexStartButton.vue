<script setup lang="ts">
  import { init } from '@/network/init';
import router from '../router/index';

  const handler = async () => {
    await init();
    router.push('/poll');
  };
</script>


<template>
  <button @click="handler()" type="button" class="btn btn-primary">
    Start your way!
  </button>
</template>