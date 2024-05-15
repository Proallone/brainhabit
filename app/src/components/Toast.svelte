<!-- Toast.svelte -->
<script>
    export let message;
  </script>
  
  <div class="toast">
    <p>{message}</p>
  </div>
  
  <style>
    .toast {
      position: fixed;
      bottom: 20px;
      left: 50%;
      transform: translateX(-50%);
      background-color: #ff6347; /* Tomato color */
      color: #fff;
      padding: 10px 20px;
      border-radius: 5px;
      box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
    }
  </style>
  