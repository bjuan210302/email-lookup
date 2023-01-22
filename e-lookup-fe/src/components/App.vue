<template>
  <ErrorBoundary>

    <!-- LEFT -->
    <div class="basis-3/5 max-h-screen overflow-y-auto">

      <!-- HEADER -->
      <div class="sticky top-0 px-5 py-5 bg-gray-100 border-solid border-b-2 border-gray-200">
        <div class="flex justify-between items-center">

          <div class="flex items-center min-w-[50%]">
            <div class="relative min-w-full mb-1">
              <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                <svg class="w-5 h-5 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
              <input type="text" @keyup.enter="() => search()" v-model="termSearch"
                class="block p-3 pl-10 w-full text-sm text-zinc-900 bg-gray-50 rounded-md border border-gray-300"
                placeholder="Search for content">
              <button @click="() => search()"
                class="text-white absolute right-1 bottom-1 top-1 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-md text-sm px-4 py-2">Search</button>
            </div>
            <button class="inline-flex ml-2" @click="oConfigModal">
              <svg class="stroke-zinc-600" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                <path
                  d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c1 .608 2.296 .07 2.572 -1.065z">
                </path>
                <circle cx="12" cy="12" r="3"></circle>
              </svg>
            </button>
          </div>


          <Pagination :currentPage="currentPage" :numberOfPages="numberOfPages" :totalResults="totalResults"
            @changePage="(newPage) => updatePage(newPage)" />
        </div>
      </div>

      <!-- RESULTS -->
      <div class="overflow-y-auto p-3">

        <div v-if="totalResults === 0" class="absolute top-[50%] left-0 w-3/5
      text-center my-auto text-zinc-300 font-semibold tracking-wider">
          No results found
        </div>

        <SmallEmail v-if="totalResults > 0" v-for="(item, index) in results" :key="item._id" :_id="item._id"
          :subject="item.subject" :from="item.from" :to="item.to" :highlightedContent="item.highlight"
          @expandEmail="expandEmail(index)" />
      </div>


    </div>

    <!-- EXPANDED EMAIL-->
    <div class="basis-2/5 top-0 p-4 px-5 bg-gray-100 overflow-y-auto
    border-solid border-l-2 border-gray-200">

      <span v-if="!selectedEmail._id" class="absolute top-[50%] right-0 w-2/5
    text-center my-auto text-zinc-400 tracking-wider">
        Please select an email to view full content
      </span>

      <ExpandedEmail v-else :id="selectedEmail._id" :subject="selectedEmail.subject" :from="selectedEmail.from"
        :to="selectedEmail.to" :date="selectedEmail.date" :content="selectedEmail.content" />
    </div>

    <!-- CONFIG MODAL-->
    <ConfigModal v-if="sConfigModal" :close="xConfigModal" :searchConfig="searchConfig"
      @newConfig="(c) => updateConfig(c)" />

  </ErrorBoundary>
</template>

<script setup lang="ts">

import Pagination from "./Pagination.vue"
import SmallEmail from "./EmailMiniView.vue"
import ExpandedEmail from "./EmailExpandedView.vue"
import ConfigModal from "./modals/ConfigModal.vue"
import { ref } from "vue";
import { Email, makeQueryRequest, SearchConfig } from "../utils/utils"
import { useModal } from "../utils/hooks";
import ErrorBoundary from "./modals/ErrorBoundary.vue";

const termSearch = ref("")
const results = ref<Email[]>([])
const selectedEmail = ref<Email>({} as Email)

const currentPage = ref(0)
const numberOfPages = ref(0)
const totalResults = ref(0)

const searchConfig = ref({
  zIndex: '',
  resultsPerPage: 10
})

const { showModal: sConfigModal, closeModal: xConfigModal, openModal: oConfigModal }
  = useModal(sessionStorage.getItem('auth') == undefined)

const search = async (resetPage: boolean = true) => {
  if (resetPage) currentPage.value = 0

  try {
    const { hits, totalHits, numPages } = await makeQueryRequest(termSearch.value, currentPage.value, searchConfig.value)
    totalResults.value = totalHits
    results.value = hits
    numberOfPages.value = numPages
  } catch (e) {
    throw e
  }
}

const expandEmail = (index: number) => {
  selectedEmail.value = results.value[index]
}

const updatePage = (newPAge: number) => {
  currentPage.value = newPAge
  search(false)
}

const updateConfig = (config: SearchConfig) => {
  console.log(config)
  searchConfig.value = config;
}


</script>