<template>
  <h2 class="py-9 text-2xl text-center tracking-wider font-semibold text-gray-700">{{ subject }}</h2>
  <div class="p-3 pt-0 text-zinc-500">
    <div class="m-4">
      <label class="px-2 text-slate-500 bg-slate-200 rounded-sm">Date:</label>
      {{ date }}
    </div>
    <div class="m-4">
      <label class="px-2 text-slate-500 bg-slate-200 rounded-sm">Subject:</label>
      {{ subject }}
    </div>
    <div class="m-4">
      <label class="px-2 text-slate-500 bg-slate-200 rounded-sm">From:</label>
      <span class="ml-2 px-2 text-zinc-500 bg-zinc-200 border-2 rounded-sm">{{ from }}</span>
    </div>
    <div class="m-4 flex flex-wrap gap-2">
      <label class="px-2 text-slate-500 bg-slate-200 rounded-sm">To:</label>
      <span v-for="t in to" class="px-2 text-zinc-500 bg-zinc-200 rounded-sm">{{ t }}</span>
    </div>
    <div class="m-4 pt-5 border-solid border-t-2 border-gray-300">
      <h3 class="text-zinc-700 text-lg mb-1">Message</h3>
      <p class="text-slate-700">
        {{ content }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  id: string,
  date: string,
  subject: string,
  from: string,
  to: Array<string>,
  content: string,
}>()
</script>