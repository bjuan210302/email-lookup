<template>
  <div class="mb-6 border border-zinc-300 shadow-md rounded-lg" @click="emit('expandEmail')">
    <h3 class="p-4 pb-2 text-gray-900 bg-gray-100 rounded-t-lg font-semibold tracking-wide">{{ subject }}</h3>

    <div class="p-4">
      <div class="mb-3">
        <div class="mb-3">
          <label>From:</label>
          <span class="ml-2 px-2 text-slate-700 bg-slate-100 rounded-sm">{{ from }}</span>
        </div>
        <div class="flex flex-wrap gap-x-2 gap-y-1">
          <label>To:</label>
          <span v-for="t in to" class="px-2 text-slate-700 bg-slate-100 rounded-sm">{{ t }}</span>
        </div>
      </div>

      <div v-if="highlightedContent">
        <p v-for="h in highlightedContent" v-html="h"
          class="px-4 py-6 text-slate-800 tracking-wide rounded-md last-of-type:mb-0 hover:bg-blue-50
          border-b-2 border-gray-100 last:border-0 transition-all">
        </p>
      </div>

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
