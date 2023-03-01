import { createRouter, createWebHashHistory, type Router } from "vue-router";

import ViewNotes from "@/views/ViewNotes.vue";
import ViewStats from "@/views/ViewStats.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "notes",
      component: ViewNotes,
    },
    {
      path: "/stats",
      name: "stats",
      component: ViewStats,
    },
  ],
});

export default router as Router;
