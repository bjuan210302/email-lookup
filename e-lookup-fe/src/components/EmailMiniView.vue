<template>
  <div class="mb-4 p-4 border border-gray-500" @click="emit('expandEmail')">
    <div class="text-zinc-500 font-bold">{{ subject }}</div>
    <div class="py-1">{{ from }} {{ to }}</div>
    <div class="py-1" v-if="highlightedContent">
      <p v-for="h in highlightedContent" v-html="h"></p>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  _id: string,
  subject: string,
  from: string,
  to: Array<string>,
  highlightedContent: string[] | null,
}>()

const emit = defineEmits<{
  (e: 'expandEmail'): void
}>()
</script>