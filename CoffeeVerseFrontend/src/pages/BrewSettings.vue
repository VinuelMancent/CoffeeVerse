<template>
  <BrewTable :items="items" @brew="sendBrewToBackend"/>
</template>

<script lang="ts" setup>
import { Ref } from 'vue';
import BrewTable from '../components/BrewTable/BrewTable.vue'
import { BrewSetting, BrewSettings, BrewTypes } from '../components/BrewTable/types'

const items: BrewSettings = [
    {
    brewtype: BrewTypes.Mokkakocher,
    coffee: "unknown",
    grinder: "Comandante C40 MK3",
    grindlevel: 12,
    watertemperature: 94}
]

function sendBrewToBackend(brewsetting: Ref<BrewSetting>){
  const newBrewsetting: BrewSetting = brewsetting.value
  console.log(`uploading ${newBrewsetting.watertemperature} to backend`)
}
</script>
