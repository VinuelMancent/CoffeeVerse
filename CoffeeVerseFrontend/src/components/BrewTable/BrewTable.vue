<template>
  <v-container class="fill-height">
    <v-responsive
      class="align-centerfill-height mx-auto"
    >
        <v-data-table :items="items"></v-data-table>
    </v-responsive>
  </v-container>
</template>

<script setup lang="ts">
  const items = []
</script>
