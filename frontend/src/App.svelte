<script lang="ts">
  import FileUploader from "./lib/FileUploader.svelte";

  function handleUploadSuccess(event: CustomEvent) {
    const { filename, data } = event.detail;
    alert(`File "${filename}" uploaded successfully! Server: ${JSON.stringify(data)}`);
  }

  function handleUploadFailure(event: CustomEvent) {
    console.error('Upload failed:', event.detail.error);
  }
</script>

<main>
  <FileUploader
    on:uploadSuccess={handleUploadSuccess}
    on:uploadFailure={handleUploadFailure}
  />
</main>