<template>
  <table>
    <thead>
    <tr v-if="fields.length > 0">
      <th v-for="[key] in Object.entries(fields[0])" :key='key'>
        {{ key }} <i class="bi bi-sort-alpha-down" aria-label='Sort Icon'></i>
      </th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="item in fields" :key='item'>
      <td v-for="[key, value] in Object.entries(item)" :key='key'>{{ value }}</td>
    </tr>
    </tbody>
  </table>
</template>
<script>

export default {
  name: 'TableView',
  props: {
    fields: {
      type: Array
    }
  },
}
</script>

<style scoped lang="scss">
table {
  background: #000000;
  border-spacing: 1px;
  td, th {
    padding: 10px;
  }
  td {
    background: #1a2c39;
  }
  th {
    background: #1e282f
  }
}

</style>