<template>
    <v-form v-model="valid" style="width: 80%;">
        <v-container>
            <v-row>
                <v-select
                v-model="myBrewSetting.brewtype"
                label="Brühmethode"
                hide-details
                required
                :items="Object.values(BrewTypes)"
                class="formelement"
                />
            </v-row>
            <v-row>
                <v-text-field
                    v-model="myBrewSetting.coffee"
                    label="Kaffee"
                    required
                    class="formelement"
                ></v-text-field>
            </v-row>
            <v-row>
                <v-text-field
                    v-model="myBrewSetting.grinder"
                    label="Mühle"
                    required
                    class="formelement"
                ></v-text-field>
            </v-row>
            <v-row>
                <v-text-field
                    v-model="myBrewSetting.grindlevel"
                    label="Mahlstufe (in Klicks)"
                    required
                    class="formelement"
                ></v-text-field>
            </v-row>
            <v-row>
                <v-text-field
                    v-model="myBrewSetting.watertemperature"
                    label="Wassertemperatur (in °C)"
                    required
                    class="formelement"
                    type="number"
                ></v-text-field>
            </v-row>
            <v-row>
                <v-btn @click="emitBrewSetting">
                Brühvorgang speichern
                </v-btn>
            </v-row>
        </v-container>
    </v-form>
</template>

<script setup lang="ts">
import { BrewSetting, BrewTypes } from '../BrewTable/types'
import { ref, Ref } from 'vue'

const emit = defineEmits(['submit'])

let myBrewSetting: Ref<BrewSetting> = ref({
    brewtype: BrewTypes.Filter,
    coffee: "",
    grinder: "",
    grindlevel: 0,
    watertemperature: 0
})

let valid: boolean = true

function emitBrewSetting(){
    console.log(`emitting brewsetting from form to parent: ${myBrewSetting.value.coffee}`)
    emit("submit", myBrewSetting)
}
</script>

<style>
    .formelement{
        width: 100%;
        padding-left: 3%;
        margin-left: 3%;
        margin-bottom: 0.5em;
    }
</style>