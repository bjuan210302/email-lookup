<template>
  <div>
    <div class="bg-slate-300 p-5">
      <h2 class="text-xl text-center tracking-wider font-semibold">{{ subject }}</h2>
    </div>
    <div class="m-4 text-zinc-500">
      <label class="px-2 text-zinc-800 bg-slate-200 rounded-sm">Date:</label>
      {{ date }}
    </div>
    <div class="m-4 text-zinc-500 ">
      <label class="px-2 text-zinc-800 bg-slate-200 rounded-sm">Subject:</label>
      {{ subject }}
    </div>
    <div class="m-4 text-zinc-500">
      <label class="px-2 text-zinc-800 bg-slate-200 rounded-sm">From:</label>
      {{ from }}
    </div>
    <div class="m-4 text-zinc-500">
      <label class="px-2 text-zinc-800 bg-slate-200 rounded-sm">To:</label>
      {{ to.join('; ') }}
    </div>
    <div class="m-4 mt-6">
      <h3 class="font-bold">Message:</h3>
      <p>
        {{ content }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  id: string,
  date: string,
  subject: string,
  from: string,
  to: Array<string>,
  content: string,
}>()
</script>