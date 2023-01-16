<template>
  <div class="flex flex-col">
    <ul class="flex justify-end">
      <li class="px-3 py-1 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-sm">
        <button @click="emit('changePage', currentPage - 1)" :disabled="currentPage - 1 < 0">Previous</button>
      </li>
      <li class="mx-2 px-3 py-1 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-sm">
        <button>{{ currentPage + 1 }}/{{ numberOfPages }}</button>
      </li>
      <li class="px-3 py-1 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-sm">
        <button @click="emit('changePage', currentPage + 1)" :disabled="currentPage == numberOfPages - 1">Next</button>
      </li>
    </ul>
    <small class="text-right text-gray-500">Total found: {{ totalResults }}</small>
  </div>
</template>

<script setup lang="ts">

defineProps<{
  currentPage: number,
  resultsPerPage: number,
  numberOfPages: number,
  totalResults: number,
}>()

const emit = defineEmits<{
  (e: 'changePage', page: number): void
}>()

</script>