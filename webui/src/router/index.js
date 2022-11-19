import {createRouter, createWebHashHistory} from 'vue-router'
import EnrollView from "../views/EnrollView.vue";
import ResultsView from "../views/ResultsView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: ResultsView},
		{path: '/enroll', component: EnrollView},
	]
})

export default router
