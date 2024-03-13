import { computed } from "@vue/reactivity";
import { defineStore } from "pinia";
import { ref } from 'vue'

export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)

  const increment = () => {
    count.value++
  }

  return { count, increment }
})