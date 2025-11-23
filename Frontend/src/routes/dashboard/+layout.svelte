<script>
  import { onMount } from 'svelte';
  import { authStore } from '../../lib/auth.js';
  import { goto } from '$app/navigation';

  // onMount runs after the component is first rendered in the browser.
  onMount(() => {
    // Get the current value from the store.
    const isLoggedIn = $authStore.loggedIn;

    // If the user is NOT logged in, redirect them to the home/login page.
    if (!isLoggedIn) {
      goto('/');
    }
  });
</script>

<!-- The <slot /> tag renders the actual page content (dashboard/+page.svelte) -->
<!-- This will only be rendered if the check above passes. -->
<slot />