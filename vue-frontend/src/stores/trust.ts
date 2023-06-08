import { fetchTrustLevels } from "@/network/trust";
import { type TrustLevel } from "@/types/dto/trust_level";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useTrustStore = defineStore('trust', () => {
  const trustLevels = ref<TrustLevel[]>([]);
  async function fetch(){
    trustLevels.value = await fetchTrustLevels();
  }

  return { trustLevels, fetch }
});