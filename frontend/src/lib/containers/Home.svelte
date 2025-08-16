<script lang="ts">
    import { ScanDirectory } from '../../../wailsjs/go/services/FileService';
    import { SelectDirectory } from '../../../wailsjs/go/main/App';
    import { loaderStore } from '../stores/loaderstore';
  import Loader from '../components/Loader/Loader.svelte';

    let files: string[] = [];
    let selectedDirectory: string = "No directory selected";
    let isScanning: boolean = false;
    let errorMessage: string = "";

    async function chooseDirectory() {
        try {
            const dir = await SelectDirectory();
            console.log(dir);
            if (!dir) {
                return; // User cancelled the dialog
            }

            selectedDirectory = dir;
            isScanning = true;
            
            errorMessage = "";
            files = []; // Clear previous results
            loaderStore.set(true);
            const result = await ScanDirectory(dir);
            files = result || [];

        } catch (err) {
            console.error("Error during directory scan:", err);
            errorMessage = `Error: ${err}`;
            selectedDirectory = "Error scanning directory.";
            loaderStore.set(false);
        } finally {
            isScanning = false;
            loaderStore.set(false);
        }
    }
</script>

<div class="home-container">
    <header>
        <h2>Dashboard</h2>
        <p>Select a directory to begin indexing your files.</p>
    </header>

    <div class="scan-controls">
        <button on:click={chooseDirectory} disabled={isScanning}>
            {#if isScanning}
                <Loader />
            {:else}
                Select Scan Directory
            {/if}

        </button>
        <div class="selected-path">
            <strong>Selected:</strong>
            <span>{selectedDirectory}</span>
        </div>
    </div>

    {#if errorMessage}
        <div class="error-box">
            {errorMessage}
        </div>
    {/if}

    <div class="results-container">
        {#if isScanning}
            <p>Scanning for files, please wait...</p>
        {:else if files.length > 0}
            <h3>Found {files.length} files:</h3>
            <ul class="file-list">
                {#each files as file}
                    <li>{file}</li>
                {/each}
            </ul>
        {/if}
    </div>
</div>

<style>
    .home-container header { margin-bottom: 2rem; }
    h2 { font-size: 1.8rem; font-weight: 500; margin-bottom: 0.5rem; }
    .scan-controls { background-color: #1f2d40; padding: 1.5rem; border-radius: 8px; display: flex; align-items: center; gap: 1rem; }
    button { background-color: #3b82f6; color: white; border: none; padding: 0.8rem 1.5rem; border-radius: 6px; font-size: 1rem; font-weight: 500; cursor: pointer; transition: background-color 0.2s; }
    button:hover { background-color: #2563eb; }
    button:disabled { background-color: #555; cursor: not-allowed; }
    .selected-path { background-color: #1b2636; padding: 0.8rem 1rem; border-radius: 6px; flex-grow: 1; font-family: 'Courier New', Courier, monospace; font-size: 0.9rem; }
    .selected-path span { color: #a0b0c0; }
    .results-container { margin-top: 2rem; }
    .file-list { background-color: #1f2d40; border-radius: 8px; padding: 1rem; max-height: 400px; overflow-y: auto; list-style: none; }
    .file-list li { padding: 0.5rem 1rem; border-bottom: 1px solid #2a3b51; font-size: 0.9rem; }
    .file-list li:last-child { border-bottom: none; }
    .error-box { background-color: #4d1f1f; color: #ffdddd; border: 1px solid #8c2b2b; padding: 1rem; border-radius: 8px; margin-top: 1.5rem; }
</style>
