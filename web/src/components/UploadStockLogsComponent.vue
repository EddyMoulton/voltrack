<template>
  <section>
    <b-field>
      <b-upload v-model="dropFiles" multiple drag-drop>
        <section class="section">
          <div class="content has-text-centered">
            <p>
              <b-icon icon="upload" size="is-large"> </b-icon>
            </p>
            <p>Drop your files here or click to upload</p>
          </div>
        </section>
      </b-upload>
    </b-field>

    <div class="tags">
      <span
        v-for="(file, index) in dropFiles"
        :key="index"
        class="tag is-primary"
      >
        {{ file.name }}
        <button
          class="delete is-small"
          type="button"
          @click="deleteDropFile(index)"
        ></button>
      </span>
    </div>
    <b-button type="is-primary" :disabled="actionInProgress" @click="add">
      Add
    </b-button>
  </section>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import { StockLogViewModel } from '../models/StockLogViewModel';
import { ApiClient } from '../api/apiClient';
import Papa from 'papaparse';

@Component({
  name: 'UploadStockLogsComponent'
})
export default class UploadStockLogsComponent extends Vue {
  private dropFiles: File[] = [];
  private fileData: StockLogViewModel[] = [];
  private apiClient: ApiClient;
  private actionInProgress = false;

  constructor() {
    super();
    this.apiClient = new ApiClient('http://localhost:3000/api');
  }

  deleteDropFile(index: number) {
    this.dropFiles.splice(index, 1);
  }

  @Watch('dropFiles')
  onDropFilesChanged(value: File[]) {
    this.fileData = [];

    value.forEach((file) => {
      Papa.parse(file, {
        header: true,
        dynamicTyping: true,
        complete: (result) => {
          if (result) {
            this.fileData.push(...(result.data as StockLogViewModel[]));
          }
        }
      });
    });
  }

  async add() {
    this.actionInProgress = true;
    if (this.fileData && this.fileData.length) {
      try {
        const result = await this.apiClient.AddStockLogs(this.fileData);
        this.dropFiles = [];
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        (this.$parent as any).close();
        this.$buefy.toast.open({
          message: 'Added ' + result.data.addedEntries + ' logs',
          type: 'is-success'
        });
      } catch (e) {
        this.$buefy.toast.open({
          message: 'Failed to add logs',
          type: 'is-danger'
        });
        console.error('Failed to add logs');
        console.error(e);
      } finally {
        this.actionInProgress = false;
      }
    }
  }
}
</script>
