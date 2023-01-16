<template>
  <div class="mb-4 p-4 border border-gray-500" @click="emit('expandEmail')">
    <div class="text-gray-500 font-bold">{{ subject }}</div>
    <div class="py-1">{{ from }} {{ to }}</div>
    <div class="py-1" v-html="highlightedContent"></div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  _id: string,
  subject: string,
  from: string,
  to: Array<string>,
  highlightedContent: string,
}>()

const emit = defineEmits<{
  (e: 'expandEmail'): void
}>()
</script>