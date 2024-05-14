<!-- Login.svelte -->
<script>
    	import Header from '../../components/Header.svelte';

    let email = '';
    let password = '';
  
    function handleSubmit() {
      // Here you can add your logic for handling the login submission
      // For example, you can send a request to your backend server
      console.log('Email:', email);
      console.log('Password:', password);
      // Reset the form fields after submission
      email = '';
      password = '';
    }
  </script>
  <div>
    <Header/>
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
    .login-container {
      max-width: 400px;
      margin: 0 auto;
      padding: 20px;
      border: 1px solid #ccc;
      border-radius: 5px;
      background-color: #f9f9f9;
    }
  
    h2 {
      text-align: center;
      margin-bottom: 20px;
    }
  
    .form-group {
      margin-bottom: 15px;
    }
  
    label {
      display: block;
      font-weight: bold;
    }
  
    input[type="email"],
    input[type="password"] {
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
      background-color: #007bff;
      color: #fff;
      cursor: pointer;
    }
  
    button:hover {
      background-color: #0056b3;
    }
  
    button:active {
      background-color: #0056b3;
      transform: translateY(1px);
    }
  </style>
  