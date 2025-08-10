<script lang="ts">
  import Home from './lib/Home.svelte';
</script>

<main class="app-container">
  <div class="sidebar">
    <h1 class="brand">Katalog</h1>
    <nav>
      <ul>
        <li class="active"><a href="#">Home</a></li>
        <li><a href="#">Tags</a></li>
        <li><a href="#">Settings</a></li>
      </ul>
    </nav>
  </div>
  <div class="content">
    <Home />
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
      'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
      sans-serif;
    background-color: #1b2636;
    color: #e0e0e0;
  }

  .app-container {
    display: flex;
    height: 100vh;
  }

  .sidebar {
    width: 220px;
    background-color: #1f2d40;
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
  }

  .brand {
    font-size: 1.5rem;
    font-weight: 600;
    margin-bottom: 2rem;
    color: #ffffff;
  }

  nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  nav li a {
    display: block;
    padding: 0.75rem 1rem;
    border-radius: 6px;
    text-decoration: none;
    color: #a0b0c0;
    margin-bottom: 0.5rem;
    transition: background-color 0.2s;
  }

  nav li a:hover {
    background-color: #2a3b51;
  }
  
  nav li.active a {
    background-color: #3b82f6;
    color: white;
  }

  .content {
    flex-grow: 1;
    padding: 2rem;
    overflow-y: auto;
  }
</style>
