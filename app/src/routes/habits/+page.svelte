<!-- Table.svelte -->
<script>
    // Define sample habit records with sequential IDs
    let habits = [
      { id: 1, name: 'Exercise', frequency: 'Daily', streak: 10 },
      { id: 2, name: 'Reading', frequency: 'Weekly', streak: 5 },
      { id: 3, name: 'Meditation', frequency: 'Daily', streak: 15 },
      { id: 4, name: 'Writing', frequency: 'Weekly', streak: 8 },
      { id: 5, name: 'Running', frequency: 'Daily', streak: 20 },
      { id: 6, name: 'Drawing', frequency: 'Weekly', streak: 12 },
      { id: 7, name: 'Coding', frequency: 'Daily', streak: 30 },
      { id: 8, name: 'Cooking', frequency: 'Weekly', streak: 9 },
      { id: 9, name: 'Swimming', frequency: 'Daily', streak: 25 },
      { id: 10, name: 'Gardening', frequency: 'Weekly', streak: 6 },
      { id: 11, name: 'Yoga', frequency: 'Daily', streak: 18 },
      { id: 12, name: 'Singing', frequency: 'Weekly', streak: 11 },
      { id: 13, name: 'Walking', frequency: 'Daily', streak: 22 },
      { id: 14, name: 'Dancing', frequency: 'Weekly', streak: 7 },
      { id: 15, name: 'Cycling', frequency: 'Daily', streak: 17 },
      { id: 16, name: 'Painting', frequency: 'Weekly', streak: 13 },
      { id: 17, name: 'Gaming', frequency: 'Daily', streak: 28 },
      { id: 18, name: 'Listening to Music', frequency: 'Weekly', streak: 10 },
      { id: 19, name: 'Watching Movies', frequency: 'Daily', streak: 23 },
      { id: 20, name: 'Traveling', frequency: 'Weekly', streak: 14 },
    ];
  
    // Sorting
    let sortBy = null;
    let sortAsc = true;
  
    function sortByColumn(column) {
      if (sortBy === column) {
        sortAsc = !sortAsc;
      } else {
        sortBy = column;
        sortAsc = true;
      }
  
      habits = [...habits.sort((a, b) => {
        const aValue = typeof a[column] === 'string' ? a[column].toLowerCase() : a[column];
        const bValue = typeof b[column] === 'string' ? b[column].toLowerCase() : b[column];
        return sortAsc ? aValue - bValue : bValue - aValue;
      })];
    }
  </script>
  
  <div class="table-container">
    <a href="/login">Login</a>
    <table class="table">
      <thead>
        <tr>
          <th on:click={() => sortByColumn('id')}>ID</th>
          <th on:click={() => sortByColumn('name')}>Name</th>
          <th on:click={() => sortByColumn('frequency')}>Frequency</th>
          <th on:click={() => sortByColumn('streak')}>Streak</th>
        </tr>
      </thead>
      <tbody>
        {#each habits as habit (habit.id)}
          <tr>
            <td>{habit.id}</td>
            <td>{habit.name}</td>
            <td>{habit.frequency}</td>
            <td>{habit.streak}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  
  <style>
    .table-container {
      height: 400px; /* Set the desired height */
      overflow-y: auto; /* Enable vertical scrolling */
      scrollbar-color: #94d2a2 #f0f5e7; /* Custom scrollbar colors (thumb and track) for Firefox */
    }
  
    .table-container::-webkit-scrollbar {
      width: 16px; /* Width of the scrollbar for WebKit browsers */
    }
  
    .table-container::-webkit-scrollbar-thumb {
      background-color: #94d2a2; /* Color of the scrollbar thumb */
      border-radius: 8px; /* Border radius of the scrollbar thumb */
    }
  
    .table {
      width: 100%;
      border-collapse: collapse;
      border-radius: 0.5rem;
      overflow: hidden;
      box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
      background-color: #f0f5e7; /* Light pastel green background */
    }
  
    th, td {
      padding: 12px 15px;
      text-align: left;
      border-bottom: 1px solid #d9e4c7; /* Lighter border color */
      font-family: Arial, sans-serif;
    }
  
    th {
      background-color: #94d2a2; /* Pastel green header background */
      font-weight: bold;
      color: #333; /* Dark text color */
      text-transform: uppercase;
      cursor: pointer;
      position: sticky;
      top: 0;
      z-index: 1;
    }
  
    tr:hover {
      background-color: #e6f7ec; /* Light green hover background */
    }
  
    tbody tr:nth-child(even) {
      background-color: #f0f5e7; /* Light pastel green background for even rows */
    }
  
    tbody tr:nth-child(odd) {
      background-color: #ebf3e0; /* Lighter pastel green background for odd rows */
    }
  
    td:first-child, th:first-child {
      border-top-left-radius: 0.3rem;
      border-bottom-left-radius: 0.3rem;
    }
  
    td:last-child, th:last-child {
      border-top-right-radius: 0.3rem;
      border-bottom-right-radius: 0.3rem;
    }
  </style>
  