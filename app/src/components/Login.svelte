<!-- Login.svelte -->
<script>
	let email = '';
	let password = '';
	let errorMessage = '';

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
			console.log('Login successful', data);
		} catch (error) {
			// Handle login errors
			errorMessage = 'Login failed. Please check your credentials and try again.';
			console.error('Login error:', error.message);
		}

		// Reset the form fields after submission
		email = '';
		password = '';
	}
</script>

<div class="login-page">
	<div class="login-container">
		<h2>Login</h2>
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
		height: 90vh;
		background-color: #f0f5e7; /* Set a background color for the whole page */
	}

	.login-container {
		max-width: 400px;
		padding: 20px;
		border: 2px solid #ccc;
		border-radius: 10px;
		background-color: #fff;
		box-shadow: 0 0 50px rgba(0, 0, 0, 0.2); /* Add a subtle box shadow */
	}

	h2 {
		text-align: center;
		margin-bottom: 20px;
	}

	.form-group {
		margin-bottom: 15px;
		margin-right: 20px;
	}

	label {
		display: block;
		font-weight: bold;
	}

	input[type='email'],
	input[type='password'] {
		width: 100%;
		padding: 10px;
		border: 1px solid #ccc;
		border-radius: 3px;
	}

	button {
		width: 100%;
		padding: 10px;
		border: none;
		border-radius: 3px;
		background-color: #2ecc71; /* Use the primary color for the button */
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
