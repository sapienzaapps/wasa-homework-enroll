<script>
export default {
	props: ['msg', 'details']
}
</script>

<template>
	<div class="alert alert-danger" role="alert">
		{{ msg }}
		<small v-if="details"><br /><br />Details: {{ details }}</small>
	</div>
</template>

<style>
</style>
