<script>
  import { loaderStore } from '../../stores/loaderstore';

  export let size = "12px";
  export let color = "#3498db";
  export let speed = "0.6s";
  export let overlay = true; // optional full-screen mode

  $: isVisible = $loaderStore; // reactive subscription
</script>

{#if isVisible}
  <div class:overlay>
    <div class="loader" style="--size:{size}; --color:{color}; --speed:{speed}">
      <span></span>
      <span></span>
      <span></span>
    </div>
  </div>
{/if}

<style>
  .loader {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: calc(var(--size) / 2);
  }

  .loader span {
    width: var(--size);
    height: var(--size);
    background: var(--color);
    border-radius: 50%;
    animation: bounce var(--speed) infinite ease-in-out;
  }

  .loader span:nth-child(2) {
    animation-delay: calc(var(--speed) / 3);
  }

  .loader span:nth-child(3) {
    animation-delay: calc(var(--speed) * 2 / 3);
  }

  @keyframes bounce {
    0%, 80%, 100% {
      transform: scale(0.6);
    }
    40% {
      transform: scale(1);
    }
  }

  .overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.4);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
  }
</style>
