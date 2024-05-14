<script>
	import { page } from '$app/stores';
	import Header from '../components/Header.svelte';
</script>
<div>
    <Header/>
    <h1>{$page.status}: {$page.error?.message}</h1>
</div>
