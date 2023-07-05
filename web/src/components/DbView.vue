<template>
  <div class="data" v-if="fields">
    <Table :fields='fields'/>
  </div>
</template>

<script>

import axios from "axios";

import Table from "@/components/Table.vue";

export default {
  name: 'DbView',
  components: {Table},
  data() {
    return {
      fields: null,
    };
  },
  created() {
    this.getDataFromApi()
  },
  methods: {
    async getDataFromApi() {
      const d = await axios.get('/api/v1/log')
      this.fields = d.data
    }
  },
}
</script>

<style scoped>
a {
  cursor: pointer;
  &:hover {
    text-decoration: underline;
  }
}
.data {
  margin-top:20px;
}
table {
  width: 1000px;
  margin: 0 auto;
}
</style>
