<template>
  <BaseModal :close="close" >

    <!-- AUTH -->
    <form>
      <div class="mb-6">
        <label for="zinc-user" class="block mb-2 text-sm font-medium text-gray-900">ZincSearch user</label>
        <input type="text" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-md block w-full p-2.5"
          id="zinc-user" required v-model="zUser">
      </div>

      <div class="mb-6">
        <label for="zinc-password" class="block mb-2 text-sm font-medium text-gray-900">ZincSearch password</label>
        <input type="text" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-md block w-full p-2.5"
          id="zinc-password" required v-model="zPass">
      </div>

      <button type="button" class="w-full px-5 py-2.5 border-solid border-2 border-blue-400 text-gray-900 text-sm text-center
        font-medium rounded-md hover:border-blue-500" @click="authenticate">Authenticate</button>
    </form>

    <!-- SEARCH CONFIG -->
    <form v-if:="isAuthed" class="mt-8">
      <div class="mb-6">
        <label for="zinc-index" class="block mb-2 text-sm font-medium text-gray-900">Select an
          option</label>
        <select class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-md block w-full p-2.5"
          id="zinc-index" required v-model="searchConfig.zIndex">
          <option v-for="indexName, index in existingIndexes" :value="indexName">{{
            indexName
          }}</option>
        </select>
      </div>

      <div class="mb-6">
        <label for="results-per-page" class="block mb-2 text-sm font-medium text-gray-900">Results per page</label>
        <div class="flex flex-row h-10 w-full">
          <button class="h-full w-20 bg-gray-300 text-2xl text-gray-600
            hover:text-gray-700 hover:bg-slate-300 rounded-l cursor-pointer" type="button"
            @click="() => searchConfig.resultsPerPage--" :disabled="searchConfig.resultsPerPage < 2">
            &minus;
          </button>
          <input type="number" class="flex items-center w-full bg-gray-300 text-md text-center text-gray-700 font-semibold
            cursor-default hover:text-black focus:text-black" v-model.number="searchConfig.resultsPerPage" />
          <button class="h-full w-20 bg-gray-300 text-2xl text-gray-600
            hover:text-gray-700 hover:bg-slate-300 rounded-r cursor-pointer" type="button"
            @click="() => searchConfig.resultsPerPage++">
            &plus;
          </button>
        </div>
      </div>

      <button class="w-full px-5 py-2.5 bg-blue-700 text-white text-sm text-center font-medium rounded-md
        hover:bg-blue-800" @click="saveConfig">Save</button>
    </form>

  </BaseModal>
</template>

<script setup lang="ts">
import { getIndexes, SearchConfig, zAuthenticate } from "../../utils/utils";
import { ref } from 'vue'
import BaseModal from "./BaseModal.vue";

const isAuthed = ref<boolean>(sessionStorage.getItem('auth') != undefined)
const zUser = ref<string>(sessionStorage.getItem('zUser') || 'admin')
const zPass = ref<string>(sessionStorage.getItem('zPass') || 'Complexpass#123')
const existingIndexes = ref<string[]>([])

const { close, searchConfig } = defineProps<{
  close: () => void;
  searchConfig: SearchConfig;
}>()

if (isAuthed.value) {
  getIndexes().then(indexes => {
    existingIndexes.value = indexes
  })
}

const emit = defineEmits<{
  (e: 'newConfig', config: SearchConfig): void
}>()

const authenticate = async () => {
  try {
    const indexes = await zAuthenticate(zUser.value, zPass.value)
    existingIndexes.value = indexes as string[]
    searchConfig.zIndex = (indexes as string[])[0]
    isAuthed.value = true
  } catch (e) {
    isAuthed.value = false
    throw e
  }
}

const saveConfig = () => {
  emit('newConfig', searchConfig)
  close()
}
</script>

<style scoped>
input[type='number']::-webkit-inner-spin-button,
input[type='number']::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>