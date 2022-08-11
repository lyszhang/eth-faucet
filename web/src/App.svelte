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
  /* windicss(hmr:1:all-modules) */
  /* windicss layer base */
  @font-face {
    font-family: 'TT Norms';
    src: url('/fonts/TTNorms-Bold.woff2'), url('/fonts/TTNorms-Bold.ttf'),
      url('/fonts/TTNorms-Bold.otf');
    font-style: normal;
    font-weight: bold;
  }

  @font-face {
    font-family: 'TT Norms';
    src: url('/fonts/TTNorms-Regular.woff2'), url('/fonts/TTNorms-Regular.ttf'),
      url('/fonts/TTNorms-Regular.otf');
    font-style: normal;
    font-weight: normal;
  }

  @font-face {
    font-family: 'TT Norms';
    src: url('/fonts/TTNorms-Medium.woff2'), url('/fonts/TTNorms-Medium.ttf'),
      url('/fonts/TTNorms-Medium.otf');
    font-style: normal;
    font-weight: 500;
  }
  :global(body) {
    background: #f6f7f8;
    font-family: 'TT Norms', ui-sans-serif, system-ui, -apple-system,
      BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,
      'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji',
      'Segoe UI Symbol', 'Noto Color Emoji' !important;
    letter-spacing: normal;
    -webkit-font-smoothing: initial;
  }
</style>
