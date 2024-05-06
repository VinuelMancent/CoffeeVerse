<template>
  <BrewTable :items="items"/>
</template>

<script lang="ts" setup>
import BrewTable from '../components/BrewTable/BrewTable.vue'
import { BrewSettings, BrewTypes } from '../components/BrewTable/types'

const items: BrewSettings = [
    {
    brewtype: BrewTypes.Mokkakocher,
    coffee: "unknown",
    grinder: "Comandante C40 MK3",
    grindlevel: 12,
    watertemperature: 94}
]
</script>
