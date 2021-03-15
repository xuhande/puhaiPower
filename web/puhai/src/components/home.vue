<template>

</template>

<script>
export default {
  name: "home",
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style scoped>
div {
  position: absolute;
  top: 0;
  width: 250px;
  height: 1000px;
}
</style>
