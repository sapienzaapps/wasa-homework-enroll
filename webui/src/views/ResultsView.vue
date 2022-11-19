<script>
import LogModal from "../components/LogModal.vue";
export default {
	components: {LogModal},
	data: function() {
		return {
			errormsg: null,
			detailedmsg: null,
			log: null,
			loading: false,
			results: null,
		}
	},
	methods: {
		async openGitLog(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/results/" + id + "/git");
				this.log = response.data;
				const modal = new bootstrap.Modal(document.getElementById('logviewer'));
				modal.show();
			} catch (e) {
				if (e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
			this.loading = false;
		},
		async openOpenAPILog(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/results/" + id + "/openapi");
				this.log = response.data;
				const modal = new bootstrap.Modal(document.getElementById('logviewer'));
				modal.show();
			} catch (e) {
				if (e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
			this.loading = false;
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/results/");
				this.results = response.data;
			} catch (e) {
				if (e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Homework results</h1>
			<LoadingSpinner :loading="loading"></LoadingSpinner>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg" :details="detailedmsg"></ErrorMsg>

		<LogModal id="logviewer" :log="log"></LogModal>

		<p>
			All scores are within 0-30 range.<br />
			Click on the Git commit ID to retrieve the git command output.<br />
			Click on the score value to get the detailed evaluation.
		</p>

		<table class="table" v-if="results !== null">
			<thead>
			<tr>
				<th scope="col">Student ID</th>
				<th scope="col">Git commit ID</th>
				<th scope="col">OpenAPI</th>
				<th scope="col">Go</th>
				<th scope="col">Vue.js</th>
				<th scope="col">Docker</th>
				<th scope="col">Last check</th>
			</tr>
			</thead>
			<tbody>
			<tr v-for="r in this.results">
				<td>{{ r.studentID }}</td>
				<td v-if="r.hash === ''" colspan="5" @click="openGitLog(r.studentID)">
					(empty repository or error in git-clone)
				</td>
				<td v-if="r.hash !== ''" @click="openGitLog(r.studentID)"><pre>{{ r.hash }}</pre></td>
				<td v-if="r.hash !== ''" @click="openOpenAPILog(r.studentID)">{{ r.openAPI }}</td>
				<td v-if="r.hash !== ''">- <!-- {{ r.go }} --></td>
				<td v-if="r.hash !== ''">- <!-- {{ r.vue }} --></td>
				<td v-if="r.hash !== ''">- <!-- {{ r.docker }} --></td>
				<td>{{ r.lastCheck }}</td>
			</tr>
			</tbody>
		</table>
	</div>
</template>

<style>
</style>
