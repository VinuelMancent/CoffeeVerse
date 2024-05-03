/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import BrewSettings from './../pages/BrewSettings.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: "/BrewSettings", name: "BrewSettings", component: BrewSettings}
  ]
})

export default router
