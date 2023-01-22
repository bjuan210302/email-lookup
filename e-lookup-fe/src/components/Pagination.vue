<template>
  <div class="flex flex-col">
    <ul class="flex justify-end mb-1">
      <li class="px-3 py-2 bg-gray-200 text-zinc-600 font-medium rounded-md
        hover:bg-gray-700 hover:text-neutral-200">
        <button @click="emit('changePage', currentPage - 1)" :disabled="currentPage - 1 < 0">Previous</button>
      </li>
      <li class="mx-2 px-3 py-2 bg-gray-200 text-zinc-600 font-medium rounded-md
        hover:bg-gray-700 hover:text-neutral-200">
        <button>{{ currentPage + 1 }}/{{ numberOfPages }}</button>
      </li>
      <li class="px-3 py-2 bg-gray-200 text-zinc-600 font-medium rounded-md
        hover:bg-gray-700 hover:text-neutral-200">
        <button @click="emit('changePage', currentPage + 1)" :disabled="currentPage >= numberOfPages - 1">Next</button>
      </li>
    </ul>
    <small class="text-right text-zinc-500">Total found: {{ totalResults }}</small>
  </div>
</template>

<script setup lang="ts">

defineProps<{
  currentPage: number,
  numberOfPages: number,
  totalResults: number,
}>()

const emit = defineEmits<{
  (e: 'changePage', page: number): void
}>()

</script>