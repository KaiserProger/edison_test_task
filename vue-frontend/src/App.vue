<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <nav class="nav">
    <RouterLink
      to="/history"
      custom
      v-slot="{ href, navigate }"
    >
      <a class="nav-link active" :href="href" @click="navigate">Extrasenses guesses</a>
    </RouterLink>
    <RouterLink
      to="/trust"
      custom
      v-slot="{ href, navigate }"
    >
    <a class="nav-link active" :href="href" @click="navigate">Extrasenses trust</a>
    </RouterLink>
  </nav>

  <RouterView />
</template>
