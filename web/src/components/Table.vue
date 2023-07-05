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
  name: 'Table',
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
  td, th {
    background: #0d3049;
    padding: 10px;
  }
}

</style>