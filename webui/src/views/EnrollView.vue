<script>
export default {
	components: {},
	data: function() {
		return {
			errormsg: null,
			detailedmsg: null,
			loading: false,
			publicKey: null,
			studentInfo: {
				studentId: 0,
				firstName: "",
				lastName: "",
				email: "",
				repoURL: "",
			}
		}
	},
	methods: {
		async enroll() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/enroll/", this.studentInfo);
				this.publicKey = response.data.publicKey;
			} catch (e) {
				if (e.response && e.response.status === 409) {
					this.errormsg = "You are already enrolled. If you think that this is an error, write an e-mail to us.";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
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
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Deliver your homework</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg" :details="detailedmsg"></ErrorMsg>

		<LoadingSpinner :loading="loading"></LoadingSpinner>

		<div v-if="!loading && !publicKey">
			<div class="mb-3">
				<label for="studentId" class="form-label">Student ID</label>
				<input type="number" class="form-control" id="studentId" v-model="studentInfo.studentId" min="1000000" max="9999999">
			</div>
			<div class="mb-3">
				<label for="firstName" class="form-label">First name</label>
				<input type="text" class="form-control" id="firstName" v-model="studentInfo.firstName">
			</div>
			<div class="mb-3">
				<label for="lastName" class="form-label">Last name</label>
				<input type="text" class="form-control" id="lastName" v-model="studentInfo.lastName">
			</div>
			<div class="mb-3">
				<label for="email" class="form-label">Email (use the @studenti.uniroma1.it address)</label>
				<input type="email" class="form-control" id="email" v-model="studentInfo.email">
			</div>
			<div class="mb-3">
				<label for="repoURL" class="form-label">Repository URL</label>
				<input type="url" class="form-control" id="repoURL" v-model="studentInfo.repoURL">
			</div>

			<p>Accepted URLs: Only SSH URL is supported. Examples:</p>
			<ul>
				<li>git@github.com:yourname/yourrepo.git</li>
				<li>git@gitlab.com:yourname/yourrepo.git</li>
				<li>ssh://git@bitbucket.com/yourname/yourrepo.git</li>
			</ul>

			<button type="button" class="btn btn-sm btn-primary" @click="enroll">
				Enroll
			</button>
		</div>
		<div v-if="!loading && publicKey">
			Please configure this public key in your repository:
			<pre class="alert alert-success">{{ publicKey }}</pre>
			<b>Copy and paste this key now! You won't be able to access it again!</b>
		</div>
	</div>
</template>

<style>
</style>
