<script>
  import 'animate.css';
  import 'bulma/css/bulma.css';
  import Faucet from './Faucet.svelte';
  import Header from './Header.svelte';
</script>

<svelte:head>
  <link
    href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
    rel="stylesheet"
  />
</svelte:head>
<svelte:body />
<Header />
<Faucet />

<style windi:preflights:global windi:safelist:global>
  :global(body) {
    background: #f6f7f8;
  }
</style>
