<template>
  <v-container class="fill-height">
    <v-responsive
      class="align-top fill-height mx-auto"
    >
      <v-data-table :headers="headers" :items="items">
        <template v-slot:top>
          <v-sheet>
            <h1 class="text-center">
            Coffee Brewing Overview
            </h1>
            <v-dialog 
              style="width: 80%;"
              class="alight-top fill-height mx-auto"
              >
              <template v-slot:activator="{ props: activatorProps }">
                <v-btn
                  v-bind="activatorProps"
                  text="Open Dialog"
                  class="align-center float-end mx-auto"
                >
              Add new Entry
              </v-btn>
              </template>

              <template v-slot:default="{ isActive }">
                <v-card title="Füge eine neue Brühung hinzu">

                  <v-card-actions>
                  <AddBrewForm/>
                    <v-spacer></v-spacer>
                    <v-btn
                      text="Abbrechen"
                      @click="isActive.value = false"
                    ></v-btn>
                  </v-card-actions>
                </v-card>
              </template>
            </v-dialog>
            
          </v-sheet>
        </template>
      </v-data-table>
    </v-responsive>
  </v-container>
</template>

<script setup lang="ts">
import type { PropType } from 'vue'
import { BrewSettings } from './types'
import AddBrewForm from '../AddBrewForm/AddBrewForm.vue';

const props = defineProps({
  items: { type: Object as PropType<BrewSettings>, required: true },
})

const headers = [
  {
    title: "Brühmethode",
    key: "Brewtype"
  },
  {
    title: "Kaffee",
    key: "coffee"
  },
  {
    title: "Mühle",
    key: "grinder"
  },
  {
    title: "Mahlstufe (in Klicks)",
    key: "grindlevel"
  },
  {
    title: "Wassertemperatur (in °C)",
    key: "watertemperature"
  }
]
</script>
