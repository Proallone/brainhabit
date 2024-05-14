<!-- Header.svelte -->
<script>
</script>

<header class="header">
    <div class="menu-button">
      <!-- Menu button -->
      <!-- Hamburger menu icon SVG -->
      <svg class="icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M4 6H20M4 12H20M4 18H20" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
    </div>
    <h1 class="title">BrainHabit</h1>
  </header>
  
  <style>
    .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 10px;
      background-color: #2ecc71; /* Header background color */
      color: white; /* Header text color */
      position: sticky;
      top: 0;
      z-index: 1000; /* Ensure the header stays on top of other content */
    }
  
    .menu-button {
      cursor: pointer;
    }
  
    .title {
      font-size: 24px;
      margin: 0;
      text-align: center; /* Center the title along the x-axis */
      flex-grow: 1; /* Allow the title to grow and fill the available space */
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; /* Example font family */
    }
  
    .icon {
      width: 24px;
      height: 24px;
    }
  </style>