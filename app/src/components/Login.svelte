<!-- Login.svelte -->
<script>
	import { setToken } from '$lib/stores/auth';
	import { goto } from '$app/navigation';
	import Toast from '../components/Toast.svelte';
	let email = '';
	let password = '';
	let errorMessage = '';
	let showError = false;

	async function handleSubmit() {
		try {
			const response = await fetch('http://localhost:8080/api/auth/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email: email, password: password })
			});

			if (!response.ok) {
				throw new Error('Login failed');
			}

			// Login successful, you can redirect the user or handle the response accordingly
			const data = await response.json();
			setToken(data.token);
			goto('/habits');

			// console.log('Login successful', data);
		} catch (error) {
			// Handle login errors
			errorMessage = 'Login failed. Please check your credentials and try again';
			console.log(errorMessage)

			showError = true;
			setTimeout(() => {
				showError = false;
			}, 3000); // Hide toast after 3 seconds
		}

		// Reset the form fields after submission
		email = '';
		password = '';
	}
</script>

<div class="login-page">
	<div class="login-container">
		<h2>Login</h2>
		{#if showError}
		<Toast message={errorMessage} />
	  	{/if}
		<form on:submit|preventDefault={handleSubmit}>
			<div class="form-group">
				<label for="email">Email:</label>
				<input type="email" id="email" bind:value={email} required />
			</div>
			<div class="form-group">
				<label for="password">Password:</label>
				<input type="password" id="password" bind:value={password} required />
			</div>
			<button type="submit">Login</button>
		</form>
	</div>
</div>

<style>
	.login-page {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 95vh;
		background-color: #f0f5e7; /* Set a background color for the whole page */
	}

	.login-container {
		max-width: 400px;
		padding: 30px;
		border: 2px unset #ccc;
		border-radius: 10px;
		background-color: #fff;
		box-shadow: 0 0 20px rgba(0, 120, 0, 0.2); /* Add a subtle box shadow */
	}

	h2 {
		text-align: center;
		font-weight: 1000;
		font-family:monospace;
		margin-bottom: 20px;
	}

	.form-group {
		margin-bottom: 15px;
	}

	label {
		display: block;
		font-weight: bold;
	}

	input[type='email'],
	input[type='password'] {
		width: 100%;
		padding: 10px;
		border: 1px solid #1ebb4d;
		border-radius: 3px;
	}

	button {
		width: 100%;
		padding: 10px;
		border: none;
		border-radius: 3px;
		background-color: rgb(46, 204, 113); /* Use the primary color for the button */
		color: #fff;
		cursor: pointer;
	}

	button:hover {
		background-color: #27ae60; /* Darken the color on hover */
	}

	button:active {
		background-color: #27ae60; /* Darken the color when clicked */
		transform: translateY(1px);
	}
</style>
