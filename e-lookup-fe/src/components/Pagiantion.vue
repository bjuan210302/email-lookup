<template>
  <ul class="flex">
    <li class="mx-1 px-3 py-2 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-sm">
      <button class="flex items-center" @click="changePage(currentPage - 1)"
        :disabled="currentPage - 1 < 0">Previous</button>
    </li>
    <li class="mx-1 px-3 py-2 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-sm">
      <button>{{ currentPage+ 1}}</button>
    </li>
    <li class="mx-1 px-3 py-2 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-sm">
      <button class="flex items-center" @click="changePage(currentPage + 1)"
        :disabled="currentPage == maxPage - 1">Next</button>
    </li>
  </ul>
</template>

<script lang="ts">

import { ref } from "vue";

export default {
  setup() {
    const currentPage = ref(0)
    const maxResult = ref(5)
    const maxPage = ref(0)

    const paginator = (page: number) => {

    }

    return {
      currentPage,
      maxResult,
      maxPage,
      changePage: paginator
    }

  }
}

</script>