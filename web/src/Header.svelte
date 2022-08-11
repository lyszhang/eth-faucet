<script>
  import { onMount } from 'svelte';

  const navigation = [
    {
      name: 'Guide',
      description: '',
      isExternal: true,
      link: 'https://guide.scroll.io/',
    },
    {
      name: 'Faucet',
      description: 'Receive test ETH and USDC tokens on L1 from Faucet.',
      link: 'https://faucet.scroll.io/',
      isActive: true,
    },
    {
      name: 'Bridge',
      description:
        'Use Bridge (our Hop fork) to transfer tokens between L1 and L2.',
      link: 'https://bridge.scroll.io/',
    },
    {
      name: 'Swap',
      description:
        'Use Swap (our Uniswap fork) to swap tokens and supply liquidity on L2.',
      link: 'https://swap.scroll.io/',
    },
    {
      name: 'Rollup Explorer',
      description:
        'See L2 blocks being committed to L1 and finalized with zkEVM proofs on our Rollup Explorer.',
      link: 'https://rollupscan.scroll.io/',
    },
    {
      name: 'L1 Block Explorer',
      description: 'See your L1 transactions on our fork of Blockscout.',
      link: 'https://l1scan.scroll.io/',
      isExternal: true,
    },
    {
      name: 'L2 Block Explorer',
      description: 'See your L2 transactions on our fork of Blockscout.',
      link: 'https://l2scan.scroll.io/',
      isExternal: true,
    },
  ];

  $: open = false;

  function handleOpen() {
    open = !open;
    console.log(open);
  }
</script>

<div class="w-full bg-white">
  <nav
    class="container relative flex flex-wrap items-center justify-between p-8 mx-auto lg:justify-between xl:px-0"
  >
    <div class="flex flex-wrap items-center justify-between w-full lg:w-auto">
      <a
        href="https://prealpha.scroll.io/"
        class="flex items-center space-x-2 text-2xl font-medium text-gray-800 dark:text-gray-100"
      >
        <img
          src="https://scroll.io/img/logo_with_text.png"
          alt="logo"
          width={96}
          height={39}
          class="cursor-pointer"
        />
      </a>
      <svg
        class="w-6 h-6 fill-curren lg:hidden"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        on:click={handleOpen}
      >
        {#if open}
          <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M18.278 16.864a1 1 0 0 1-1.414 1.414l-4.829-4.828-4.828 4.828a1 1 0 0 1-1.414-1.414l4.828-4.829-4.828-4.828a1 1 0 0 1 1.414-1.414l4.829 4.828 4.828-4.828a1 1 0 1 1 1.414 1.414l-4.828 4.829 4.828 4.828z"
          />
        {:else}
          <path
            fillRule="evenodd"
            d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
          />
        {/if}
      </svg>
      {#if open}
        <div class="flex flex-wrap w-full my-5 lg:hidden">
          {#each navigation as tab}
            <a
              target={tab.isExternal ? '_blank' : '_self'}
              href={tab.link}
              class={`w-full px-2  py-2 -ml-4 text-gray-500 rounded-md hover:text-indigo-500 focus:text-indigo-500 focus:bg-indigo-100 focus:outline-none  ${
                tab.isActive ? 'bg-indigo-100 text-indigo-500' : ''
              } xl:px-4`}
            >
              {tab.name}{' '}
              {#if tab.isExternal}
                <svg
                  class="align-text-bottom w-[20px] h-[20px] inline-block fill-current"
                  focusable="false"
                  aria-hidden="true"
                  viewBox="0 0 24 24"
                  data-testid="OpenInNewIcon"
                  ><path
                    d="M19 19H5V5h7V3H5c-1.11 0-2 .9-2 2v14c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2v-7h-2v7zM14 3v2h3.59l-9.83 9.83 1.41 1.41L19 6.41V10h2V3h-7z"
                  /></svg
                >
              {/if}
            </a>
          {/each}
        </div>
      {/if}
    </div>
    <div class="hidden text-center lg:flex lg:items-center">
      <ul class="items-center justify-end flex-1 list-none  lg:flex">
        {#each navigation as tab}
          <li class=" nav__item xl:mr-3 " key={tab.name}>
            <a
              target={tab.isExternal ? '_blank' : '_self'}
              href={tab.link}
              class={`flex inline-block px-[8px]  py-[8px] text-lg font-normal text-gray-800 no-underline rounded-md  hover:text-indigo-500 focus:text-indigo-500  focus:outline-none xl:px-[16px] ${
                tab.isActive ? 'bg-indigo-100 text-indigo-500' : ''
              }`}
            >
              {tab.name}{' '}
              {#if tab.isExternal}
                <svg
                  class="align-text-bottom w-[27px] h-[27px] inline-block fill-current"
                  focusable="false"
                  aria-hidden="true"
                  viewBox="0 0 24 24"
                  data-testid="OpenInNewIcon"
                  ><path
                    d="M19 19H5V5h7V3H5c-1.11 0-2 .9-2 2v14c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2v-7h-2v7zM14 3v2h3.59l-9.83 9.83 1.41 1.41L19 6.41V10h2V3h-7z"
                  /></svg
                >
              {/if}
            </a>
          </li>
        {/each}
      </ul>
    </div>
    <div class="hidden xl:mr-3 space-x-4 lg:flex nav__item" />
  </nav>
</div>

<style>
  /* windicss layer base */
  *,
  ::before,
  ::after {
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    border-width: 0;
    border-style: solid;
    border-color: #e5e7eb;
  }
  * {
    --tw-ring-inset: var(--tw-empty, /*!*/ /*!*/);
    --tw-ring-offset-width: 0px;
    --tw-ring-offset-color: #fff;
    --tw-ring-color: rgba(59, 130, 246, 0.5);
    --tw-ring-offset-shadow: 0 0 #0000;
    --tw-ring-shadow: 0 0 #0000;
    --tw-shadow: 0 0 #0000;
  }
  :root {
    -moz-tab-size: 4;
    -o-tab-size: 4;
    tab-size: 4;
  }
  :-moz-focusring {
    outline: 1px dotted ButtonText;
  }
  :-moz-ui-invalid {
    box-shadow: none;
  }
  ::moz-focus-inner {
    border-style: none;
    padding: 0;
  }
  ::-webkit-inner-spin-button,
  ::-webkit-outer-spin-button {
    height: auto;
  }
  ::-webkit-search-decoration {
    -webkit-appearance: none;
  }
  ::-webkit-file-upload-button {
    -webkit-appearance: button;
    font: inherit;
  }
  a {
    color: inherit;
    text-decoration: inherit;
  }
  svg {
    display: block;
  }
  ul {
    list-style: none;
    margin: 0;
    padding: 0;
  }
  /* windicss layer components */

  /* windicss layer utilities */
  .container {
    width: 100%;
  }
  @media (min-width: 640px) {
    .container {
      max-width: 640px !important;
    }
  }
  @media (min-width: 768px) {
    .container {
      max-width: 768px !important ;
    }
  }
  @media (min-width: 1024px) {
    .container {
      max-width: 1024px !important;
    }
  }
  @media (min-width: 1280px) {
    .container {
      max-width: 1280px !important;
    }
  }
  @media (min-width: 1536px) {
    .container {
      max-width: 1536px !important;
    }
  }
  .space-x-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(0.5rem * var(--tw-space-x-reverse));
    margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
  }
  .space-x-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(1rem * var(--tw-space-x-reverse));
    margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
  }
  .bg-\[\#3f3238\] {
    --tw-bg-opacity: 1;
    background-color: rgba(63, 50, 56, var(--tw-bg-opacity));
  }
  .focus\:bg-indigo-100:focus {
    --tw-bg-opacity: 1;
    background-color: rgba(224, 231, 255, var(--tw-bg-opacity));
  }
  .bg-indigo-100 {
    --tw-bg-opacity: 1;
    background-color: rgba(224, 231, 255, var(--tw-bg-opacity));
  }
  .dark .dark\:focus\:bg-gray-800:focus {
    --tw-bg-opacity: 1;
    background-color: rgba(31, 41, 55, var(--tw-bg-opacity));
  }
  .bg-white {
    --tw-bg-opacity: 1;
    background-color: rgba(255, 255, 255, var(--tw-bg-opacity));
  }
  .bg-\[\#f6f7f8\] {
    --tw-bg-opacity: 1;
    background-color: rgba(246, 247, 248, var(--tw-bg-opacity));
  }
  .\!border-transparent {
    border-color: transparent !important;
  }
  .rounded-md {
    border-radius: 0.375rem;
  }
  .rounded {
    border-radius: 0.25rem;
  }
  .\!rounded {
    border-radius: 0.25rem !important;
  }
  .cursor-pointer {
    cursor: pointer;
  }
  .block {
    display: block;
  }
  .inline-block {
    display: inline-block;
  }
  .flex {
    display: -webkit-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
  }
  .hidden {
    display: none;
  }
  .flex-col {
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    -webkit-flex-direction: column;
    flex-direction: column;
  }
  .flex-wrap {
    -ms-flex-wrap: wrap;
    -webkit-flex-wrap: wrap;
    flex-wrap: wrap;
  }
  .items-center {
    -webkit-box-align: center;
    -ms-flex-align: center;
    -webkit-align-items: center;
    align-items: center;
  }
  .justify-end {
    -webkit-box-pack: end;
    -ms-flex-pack: end;
    -webkit-justify-content: flex-end;
    justify-content: flex-end;
  }
  .justify-between {
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    -webkit-justify-content: space-between;
    justify-content: space-between;
  }
  .flex-1 {
    -webkit-box-flex: 1;
    -ms-flex: 1 1 0%;
    -webkit-flex: 1 1 0%;
    flex: 1 1 0%;
  }
  .font-medium {
    font-weight: 500;
  }
  .font-normal {
    font-weight: 400;
  }
  .font-bold {
    font-weight: 700;
  }
  .h-6 {
    height: 1.5rem;
  }
  .text-2xl {
    font-size: 1.5rem;
    line-height: 2rem;
  }
  .text-lg {
    font-size: 1.125rem;
    line-height: 1.75rem;
  }
  .text-16px {
    font-size: 16px;
    line-height: 1;
  }
  .text-28px {
    font-size: 28px;
    line-height: 1;
  }
  .text-12px {
    font-size: 12px;
    line-height: 1;
  }
  .list-none {
    list-style-type: none;
  }
  .mx-auto {
    margin-left: auto;
    margin-right: auto;
  }
  .my-5 {
    margin-top: 1.25rem;
    margin-bottom: 1.25rem;
  }
  .mr-100px {
    margin-right: 100px;
  }
  .ml-auto {
    margin-left: auto;
  }
  .-ml-4 {
    margin-left: -1rem;
  }
  .mb-20px {
    margin-bottom: 20px;
  }
  .ml-4px {
    margin-left: 4px;
  }
  .mb-2px {
    margin-bottom: 2px;
  }
  .ml-12px {
    margin-left: 12px;
  }
  .min-h-400px {
    min-height: 400px;
  }
  .min-h-\[calc\(100vh_-_544px\)\] {
    min-height: calc(100vh - 544px);
  }
  .focus\:outline-none:focus {
    outline: 2px solid transparent;
    outline-offset: 2px;
  }
  .p-8 {
    padding: 2rem;
  }
  .py-60px {
    padding-top: 60px;
    padding-bottom: 60px;
  }
  .px-2 {
    padding-left: 0.5rem;
    padding-right: 0.5rem;
  }
  .py-1 {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
  }
  .py-2 {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }
  .py-16px {
    padding-top: 16px;
    padding-bottom: 16px;
  }
  .px-20px {
    padding-left: 20px;
    padding-right: 20px;
  }
  .px-4px {
    padding-left: 4px;
    padding-right: 4px;
  }
  .py-20px {
    padding-top: 20px;
    padding-bottom: 20px;
  }
  .pt-6 {
    padding-top: 1.5rem;
  }
  .tab\.name {
    -moz-tab-size: 4;
    -o-tab-size: 4;
    tab-size: 4;
  }
  .tab\.isExternal {
    -moz-tab-size: 4;
    -o-tab-size: 4;
    tab-size: 4;
  }
  .tab\.link {
    -moz-tab-size: 4;
    -o-tab-size: 4;
    tab-size: 4;
  }
  .tab\.isActive {
    -moz-tab-size: 4;
    -o-tab-size: 4;
    tab-size: 4;
  }
  .relative {
    position: relative;
  }
  .shadow {
    --tw-shadow: 0 1px 3px 0 rgb(0 0 0/0.1), 0 1px 2px -1px rgb(0 0 0/0.1);
    --tw-shadow-colored: 0 1px 3px 0 var(--tw-shadow-color),
      0 1px 2px -1px var(--tw-shadow-color);
    -webkit-box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000),
      var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000),
      var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
  .fill-current {
    fill: currentColor;
  }
  .text-center {
    text-align: center;
  }
  .text-white {
    --tw-text-opacity: 1;
    color: rgba(255, 255, 255, var(--tw-text-opacity));
  }
  .text-gray-800 {
    --tw-text-opacity: 1;
    color: rgba(31, 41, 55, var(--tw-text-opacity));
  }
  .dark .dark\:text-gray-100 {
    --tw-text-opacity: 1;
    color: rgba(243, 244, 246, var(--tw-text-opacity));
  }
  .text-gray-500 {
    --tw-text-opacity: 1;
    color: rgba(107, 114, 128, var(--tw-text-opacity));
  }
  .hover\:text-indigo-500:hover {
    --tw-text-opacity: 1;
    color: rgba(99, 102, 241, var(--tw-text-opacity));
  }
  .focus\:text-indigo-500:focus {
    --tw-text-opacity: 1;
    color: rgba(99, 102, 241, var(--tw-text-opacity));
  }
  .dark .dark\:text-gray-300 {
    --tw-text-opacity: 1;
    color: rgba(209, 213, 219, var(--tw-text-opacity));
  }
  .text-indigo-500 {
    --tw-text-opacity: 1;
    color: rgba(99, 102, 241, var(--tw-text-opacity));
  }
  .dark .dark\:text-gray-200 {
    --tw-text-opacity: 1;
    color: rgba(229, 231, 235, var(--tw-text-opacity));
  }
  .text-\[\#00000073\] {
    --tw-text-opacity: 0.45098039215686275;
    color: rgba(0, 0, 0, var(--tw-text-opacity));
  }
  .text-\[\#000\] {
    --tw-text-opacity: 1;
    color: rgba(0, 0, 0, var(--tw-text-opacity));
  }
  .text-\[\#77838f\] {
    --tw-text-opacity: 1;
    color: rgba(119, 131, 143, var(--tw-text-opacity));
  }
  .capitalize {
    text-transform: capitalize;
  }
  .hover\:underline:hover {
    -webkit-text-decoration-line: underline;
    text-decoration-line: underline;
  }
  .no-underline {
    text-decoration: none;
  }
  .align-text-bottom {
    vertical-align: text-bottom;
  }
  .\!align-text-bottom {
    vertical-align: text-bottom !important;
  }
  .w-full {
    width: 100%;
  }
  .w-6 {
    width: 1.5rem;
  }
  @media (min-width: 768px) {
    .md\:block {
      display: block;
    }
    .md\:hidden {
      display: none;
    }
    .md\:flex-row {
      -webkit-box-orient: horizontal;
      -webkit-box-direction: normal;
      -ms-flex-direction: row;
      -webkit-flex-direction: row;
      flex-direction: row;
    }
    .md\:mb-0 {
      margin-bottom: 0px;
    }
    .md\:mt-0 {
      margin-top: 0px;
    }
    .md\:py-100px {
      padding-top: 100px;
      padding-bottom: 100px;
    }
  }
  @media (min-width: 1024px) {
    .lg\:flex {
      display: -webkit-box;
      display: -ms-flexbox;
      display: -webkit-flex;
      display: flex;
    }
    .lg\:hidden {
      display: none;
    }
    .lg\:items-center {
      -webkit-box-align: center;
      -ms-flex-align: center;
      -webkit-align-items: center;
      align-items: center;
    }
    .lg\:justify-between {
      -webkit-box-pack: justify;
      -ms-flex-pack: justify;
      -webkit-justify-content: space-between;
      justify-content: space-between;
    }
    .lg\:pt-0 {
      padding-top: 0px;
    }
    .lg\:w-auto {
      width: auto;
    }
  }
  @media (min-width: 1280px) {
    .xl\:mr-3 {
      margin-right: 0.75rem;
    }
    .xl\:px-0 {
      padding-left: 0px;
      padding-right: 0px;
    }
    .xl\:px-4 {
      padding-left: 1rem;
      padding-right: 1rem;
    }
  }
  .container {
    max-width: unset !important;
  }
</style>
