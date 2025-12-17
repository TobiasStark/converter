<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  let isDragging = false;
  let fileInput: HTMLInputElement;

  let selectedFile: File | null = null;


  function handleDragEnter() {
    isDragging = true;
  }

  function handleDragLeave() {
    isDragging = false;
  }


  function handleDrop(event: DragEvent) {
    isDragging = false;
    if (event.dataTransfer && event.dataTransfer.files.length > 0) {
      selectedFile = event.dataTransfer.files[0];
    }
  }


  function handleFileSelect() {
    if (fileInput && fileInput.files && fileInput.files.length > 0) {
      selectedFile = fileInput.files[0];
    }
  }


  async function uploadFile() {
    if (!selectedFile) {
      alert("Please choose a file.");
      return;
    }

    const endpoint = 'http://localhost:8080/api/v1/upload';

    const formData = new FormData();
    formData.append('file', selectedFile);

    try {
      const response = await fetch(endpoint, {
        method: 'POST',
        body: formData
      });

      if (response.ok) {
        const result = await response.json();
        console.log('Successfully uploaded:', result);
        dispatch('uploadSuccess', { filename: selectedFile.name, data: result });
        selectedFile = null;
        fileInput.value = '';
      } else {
        const errorText = await response.text();
        throw new Error(`Upload failed: ${response.status} - ${errorText}`);
      }
    } catch (error) {
      console.error('Error uploading file:', error);
      alert(`Error uploading file: ${error.message}`);
      dispatch('uploadFailure', { error: error.message });
    }
  }

  
  function clearSelection() {
    selectedFile = null;
    if (fileInput) {
      fileInput.value = '';
    }
  }
</script>

<input
  type="file"
  bind:this={fileInput}
  on:change={handleFileSelect}
  hidden
/>

<div class="uploader-container">
  <div
    class="dropzone"
    class:dragging={isDragging}
    on:dragenter|preventDefault={handleDragEnter}
    on:dragleave|preventDefault={handleDragLeave}
    on:dragover|preventDefault
    on:drop|preventDefault={handleDrop}
    on:click={() => fileInput.click()}
    on:keydown={(e) => {
      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault();
        fileInput.click();
      }
    }}
    role="button"
    tabindex="0"
  >
    {#if selectedFile}
      <p class="file-info">Your file: <strong>{selectedFile.name}</strong></p>
      <button class="clear-button" on:click|stopPropagation={clearSelection}>
        Choose a different file (X)
      </button>
    {:else}
      <p class="instruction">
        Drag and drop a file here, or click to select.
      </p>
      {#if isDragging}
        <div class="drag-overlay">Drop file here</div>
      {/if}
    {/if}
  </div>

  <button
    class="upload-button"
    on:click={uploadFile}
    disabled={!selectedFile}
  >
    Upload file
  </button>
</div>

<style>
  .uploader-container {
    display: flex;
    flex-direction: column;
    gap: 15px;
    padding: 20px;
    border-radius: 8px;
    background-color: #f8f8f8;
    max-width: 400px;
    margin: 20px auto;
  }

  .dropzone {
    position: relative;
    border: 3px dashed #ccc;
    padding: 30px;
    text-align: center;
    cursor: pointer;
    transition: all 0.2s ease-in-out;
    border-radius: 6px;
    min-height: 120px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .dropzone.dragging {
    border-color: #007bff;
    background-color: #e6f7ff;
    box-shadow: 0 0 10px rgba(0, 123, 255, 0.5);
  }

  .instruction {
    margin: 0;
    color: #666;
  }

  .file-info {
    margin: 0 0 10px 0;
    color: #333;
  }

  .upload-button {
    padding: 10px 15px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.2s;
  }

  .upload-button:hover:not(:disabled) {
    background-color: #0056b3;
  }

  .upload-button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }

  .clear-button {
    background: none;
    border: 1px solid #dc3545;
    color: #dc3545;
    padding: 5px 10px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .clear-button:hover {
    background-color: #dc3545;
    color: white;
  }

  .drag-overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 123, 255, 0.9);
    color: white;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1.2em;
    font-weight: bold;
    border-radius: 6px;
  }
</style>