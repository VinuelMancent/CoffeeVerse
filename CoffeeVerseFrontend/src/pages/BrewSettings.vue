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

async function sendBrewToBackend(brewsetting: Ref<BrewSetting>){
  const newBrewsetting: BrewSetting = brewsetting.value
  console.log(`uploading ${newBrewsetting.watertemperature} to backend`)

  const headers: Headers = new Headers()
  headers.set('Content-Type', 'application/json')
  // We also need to set the `Accept` header to `application/json`
  // to tell the server that we expect JSON in response
  headers.set('Accept', 'application/json')
  headers.set('Access-Control-Allow-Origin', '*')

  const request: RequestInfo = new Request('http://localhost:8080/insertBrewSetting', {
    // We need to set the `method` to `POST` and assign the headers
    method: 'POST',
    headers: headers,
    // Convert the user object to JSON and pass it as the body
    body: JSON.stringify(newBrewsetting)
  })
  console.log("send brew to backend")
  await fetch(request)
  console.log("sent brew to backend")
}
</script>
