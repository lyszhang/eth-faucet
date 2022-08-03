<script>
  import { onMount } from 'svelte';
  import { getAddress } from '@ethersproject/address';
  import { setDefaults as setToast, toast } from 'bulma-toast';
  import LC from 'leancloud-storage';
  import { connected, selectedAccount, defaultEvmStores } from 'svelte-web3';
  let address = null,
    hasPermission = false;
  let faucetInfo = {
    account: '0x0000000000000000000000000000000000000000',
    network: 'testnet',
    payout: 1,
  };

  const HAS_LOGINED = 'hasLogined';

  $: if ($connected) {
    localStorage.setItem(HAS_LOGINED, 'logined');
  }

  $: {
    if ($selectedAccount) {
      address = $selectedAccount;
      const query = new LC.Query('Subscriber');
      query.equalTo('walletAddress', address.toLowerCase());
      query.find().then((res) => {
        if (res.length) {
          hasPermission = true;
        } else {
          hasPermission = false;
        }
      });
    } else {
      hasPermission = false;
    }
  }

  $: document.title = `Scroll ${capitalize(faucetInfo.network)} Faucet`;

  onMount(async () => {
    LC.init({
      appId: 'hvIeDclG2pt2nzAdbKWM0qms-MdYXbMMI',
      appKey: 'lKObgvpdxLT2JK839oxSM4Fn',
      serverURL: 'https://leancloud.scroll.io',
      serverURLs: 'https://leancloud.scroll.io',
    });
    autoSetProviderIfLogined();
    const res = await fetch('/api/info');
    faucetInfo = await res.json();
  });

  function autoSetProviderIfLogined() {
    let hasLogined = localStorage.getItem(HAS_LOGINED);
    if (hasLogined) {
      defaultEvmStores.setProvider();
    }
  }

  setToast({
    position: 'bottom-center',
    dismissible: true,
    pauseOnHover: true,
    closeOnClick: false,
    animate: { in: 'fadeIn', out: 'fadeOut' },
  });

  function handleLoginMetamask() {
    defaultEvmStores.setProvider();
  }

  async function handleRequest() {
    try {
      address = getAddress(address);
    } catch (error) {
      toast({ message: error.reason, type: 'is-warning' });
      return;
    }

    let formData = new FormData();
    formData.append('address', address);
    const res = await fetch('/api/claim', {
      method: 'POST',
      body: formData,
    });
    let message = await res.text();
    let type = res.ok ? 'is-success' : 'is-warning';
    toast({ message, type });
  }

  function capitalize(str) {
    const lower = str.toLowerCase();
    return str.charAt(0).toUpperCase() + lower.slice(1);
  }
</script>

<main>
  <section class="hero is-info is-fullheight">
    <div class="hero-head">
      <nav class="navbar">
        <div class="container">
          <div class="navbar-brand">
            <a
              class="navbar-item !hover:bg-transparent "
              href="https://enter.scroll.io/"
            >
              <img src="/logo.png" alt="logo" class="h-30px" />
            </a>
            <a class="navbar-item !pl-0 !hover:bg-transparent " href=".">
              <span><b>Scroll Faucet</b></span>
            </a>
          </div>
          <div id="navbarMenu" class="navbar-menu">
            <div class="navbar-end">
              <span class="navbar-item">
                <a
                  class="button is-white is-outlined"
                  href="https://github.com/scroll-dev/eth-faucet"
                >
                  <span class="icon">
                    <i class="fa fa-github" />
                  </span>
                  <span>View Source</span>
                </a>
              </span>
            </div>
          </div>
        </div>
      </nav>
    </div>

    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-6 is-offset-3">
          <h1 class="title">
            Receive {faucetInfo.payout} ETH & 100 USDC per request
          </h1>
          <h2 class="subtitle">
            Serving from {faucetInfo.account}
          </h2>

          {#if !hasPermission}
            <div
              class="card w-[448px] mt-[24px] mx-auto  bg-white  py-[20px] px-[32px] shadow-md rounded "
            >
              <p class="font-light">
                To prevent faucet botting, you must sign in with MetaMask.
              </p>
              <div
                on:click={handleLoginMetamask}
                class="{$selectedAccount
                  ? 'cursor-not-allowed  opacity-50'
                  : 'cursor-pointer hover:shadow-md'} w-full  py-[16px] px-[24px] flex justify-center items-center rounded border  my-[20px]"
              >
                <img
                  alt="metamask logo"
                  class="w-[60px]"
                  src="/metamask-fox.png"
                />
                <div class="ml-[16px]">
                  <p class="text-[18px] font-bold">MetaMask</p>
                  <p class="text-[16px] font-light mt-[6px]">
                    Connect using browser wallet
                  </p>
                </div>
              </div>

              {#if $selectedAccount}
                <div class="relative py-[12px]">
                  <div class="absolute inset-0 flex items-center">
                    <div class="w-full border-b border-gray-300" />
                  </div>
                  <div class="relative flex justify-center top-[-2px]">
                    <span
                      class="bg-white px-[12px] text-sm text-gray-500 text-[16px] "
                    >
                      no permission, need to sign up
                    </span>
                  </div>
                </div>
                <a
                  href="https://signup.scroll.io/"
                  class="cursor-pointer py-[10px] px-[16px] font-bold rounded border w-[100px] mx-auto mb-[16px] block"
                >
                  Sign Up
                </a>
              {/if}

              <p class=" text-center font-light text-[#595959]">
                And join our communities
              </p>
              <div class="flex justify-center mt-[10px]">
                <a class="mx-[10px]" href="https://twitter.com/Scroll_ZKP">
                  <img src="/twitter.png" alt="twitter logo" class="h-20px" />
                </a>
                <a class="mx-[10px]" href="https://discord.gg/CNzNVt4Feu">
                  <img src="/discord.png" alt="discord logo" class="h-20px" />
                </a>
                <a class="mx-[10px]" href="https://github.com/scroll-tech">
                  <img src="/github.png" alt="github logo" class="h-20px" />
                </a>
              </div>
            </div>
          {:else}
            <p class="control">
              <button
                on:click={handleRequest}
                class="button is-primary  text-[20px]"
              >
                Request
              </button>
            </p>
          {/if}
        </div>
      </div>
    </div>
  </section>
</main>

<style>
  .hero.is-info {
    background: #3f3238;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
  }
  .hero .subtitle {
    padding: 3rem 0;
    line-height: 1.5;
  }
</style>
