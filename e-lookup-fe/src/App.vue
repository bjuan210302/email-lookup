<template>
  <div class="basis-3/5 max-h-screen overflow-y-auto">
    <div class="px-5 py-5 bg-gray-100">
      <div class="flex justify-between items-center p-5">
        <div>
          <input type="text" v-model="termSearch">
          <h1 class="text-xl underline font-bold">
            term {{ termSearch }}
          </h1>
          <button @click="search">ssearch</button>
          <label>
            Total found: {{ total }}
          </label>
        </div>
        <Pagination />
      </div>
    </div>

    <!-- RESULTS -->
    <div class="overflow-y-auto p-3">
      <SmallEmail v-for="(item, index) in results" :key="item._id" :_id="item._id" :subject="item.subject"
        :from="item.from" :to="item.to" :highlightedContent="item.highlight[0]" @expandEmail="expandEmail(index)" />
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

<script lang="ts">

import { ref } from "vue";
import { Email, makeQueryRequest } from "./utils/utils"
import Pagination from "./components/Pagiantion.vue"
import SmallEmail from "./components/EmailMiniView.vue"
import ExpandedEmail from "./components/EmailExpandedView.vue"

export default {
  components: {
    Pagination,
    SmallEmail,
    ExpandedEmail,
  },
  setup() {
    const termSearch = ref("")
    const total = ref("0")
    const results = ref<Email[]>([])
    const selectedEmail = ref<Email>({} as Email)

    const currentPage = ref(0)
    const resultsPerPage = ref(5)
    const numberOfPages = ref(0)

    const search = async () => {
      const { hits, totalHits } = await makeQueryRequest(termSearch.value)
      console.log(hits)
      let calcPages = Math.trunc(totalHits / resultsPerPage.value)
      if (totalHits % resultsPerPage.value !== 0) {
        calcPages++
      }

      results.value = hits
      numberOfPages.value = calcPages
      currentPage.value = 0
    }

    const expandEmail = (index: number) => {
      selectedEmail.value = results.value[index]
    }

    const paginator = (page: number) => {

    }

    return {
      termSearch,
      total,
      results,
      selectedEmail,
      currentPage,
      maxResult: resultsPerPage,
      maxPage: numberOfPages,
      search,
      expandEmail,
      changePage: paginator
    }

  }
}

</script>
