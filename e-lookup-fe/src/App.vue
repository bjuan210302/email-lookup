<template>

  <!-- LEFT -->
  <div class="basis-3/5 max-h-screen overflow-y-auto">

    <!-- SEARCH HEADER -->
    <div class="sticky top-0 px-5 py-5 bg-gray-100">
      <div class="flex justify-between items-center">
        <div class="relative min-w-[50%]">
          <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
          <input type="text" @keyup.enter="() => search()" v-model="termSearch"
            class="block p-3 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-md border border-gray-300"
            placeholder="Search for content">
          <button @click="() => search()"
            class="text-white absolute right-1 bottom-1 top-1 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-md text-sm px-4 py-2">Search</button>
        </div>
        <Pagination :currentPage="currentPage" :resultsPerPage="resultsPerPage" :numberOfPages="numberOfPages"
          :totalResults="totalResults" @changePage="(newPage) => updatePage(newPage)" />
      </div>
    </div>

    <!-- RESULTS -->
    <div class="overflow-y-auto p-3">
      <SmallEmail v-for="(item, index) in results" :key="item._id" :_id="item._id" :subject="item.subject"
        :from="item.from" :to="item.to" :highlightedContent="item.highlight" @expandEmail="expandEmail(index)" />
    </div>

  </div>

  <!-- EXPANDED EMAIL-->
  <div class="basis-2/5 top-0 p-4 px-5 bg-gray-100 overflow-y-auto">
    <div v-if="!selectedEmail._id">
      Please select an email to view full content
    </div>
    <ExpandedEmail v-else :id="selectedEmail._id" :subject="selectedEmail.subject" :from="selectedEmail.from"
      :to="selectedEmail.to" :date="selectedEmail.date" :content="selectedEmail.content" />
  </div>
</template>

<script setup lang="ts">

import Pagination from "./components/Pagiantion.vue"
import SmallEmail from "./components/EmailMiniView.vue"
import ExpandedEmail from "./components/EmailExpandedView.vue"
import { ref } from "vue";
import { Email, makeQueryRequest } from "./utils/utils"

const termSearch = ref("")
const results = ref<Email[]>([])
const selectedEmail = ref<Email>({} as Email)

const currentPage = ref(0)
const resultsPerPage = ref(10)
const numberOfPages = ref(0)
const totalResults = ref(0)

const search = async (resetPage: boolean = true) => {
  const { hits, totalHits } = await makeQueryRequest(termSearch.value, currentPage.value, resultsPerPage.value)
  console.log(hits)
  let calcPages = Math.trunc(totalHits / resultsPerPage.value)
  if (totalHits % resultsPerPage.value !== 0) {
    calcPages++
  }

  totalResults.value = totalHits
  results.value = hits
  numberOfPages.value = calcPages
  if (resetPage) currentPage.value = 0
}

const expandEmail = (index: number) => {
  selectedEmail.value = results.value[index]
}

const updatePage = (newPAge: number) => {
  currentPage.value = newPAge
  search(false)
}

</script>
