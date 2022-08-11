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
  <section class="hero  h-[calc(100vh-110px)] flex justify-center align-middle">
    <div class="mt-[-10vh]">
      <div class="container has-text-centered">
        <div class="column is-6 is-offset-3">
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
                      Signup for waitlist
                    </span>
                  </div>
                </div>
                <a
                  href="https://signup.scroll.io/"
                  class="button my-[8px]  rounded-md px-[16px] py-[8px] text-[16px] bg-indigo-100  text-indigo-500  !border-none focus:text-white hover:(bg-indigo-100 text-indigo-500 opacity-80)"
                >
                  Sign Up
                </a>
              {/if}

              <p class=" text-center font-light text-[#595959]">
                Connect with us!
              </p>
              <div class="flex justify-center mt-[10px]">
                <a class="mx-[10px]" href="https://twitter.com/Scroll_ZKP">
                  <img src="/twitter.png" alt="twitter logo" class="h-20px" />
                </a>
                <a class="mx-[10px]" href="https://discord.gg/s84eJSdFhn">
                  <img src="/discord.png" alt="discord logo" class="h-20px" />
                </a>
                <a class="mx-[10px]" href="https://github.com/scroll-tech">
                  <img src="/github.png" alt="github logo" class="h-20px" />
                </a>
              </div>
            </div>
          {:else}
            <div
              class="card  mx-auto  bg-white  py-[60px] px-[32px] shadow-md rounded "
            >
              <h1 class="title !mb-50px text-[26px] xl:text-[32px]">
                Request {faucetInfo.payout} ETH & 100 USDC
              </h1>
              <p class="control">
                <button
                  on:click={handleRequest}
                  class="button  rounded-md px-[16px] py-[8px] text-[20px] bg-indigo-100 text-indigo-500  !border-none focus:text-white hover:(bg-indigo-100 text-indigo-500 opacity-80)"
                >
                  Request
                </button>
              </p>
            </div>
          {/if}
        </div>
      </div>
    </div>
  </section>
</main>
