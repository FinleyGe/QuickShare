<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  percent: number;
}>();

const percent = computed(() => {
  return props.percent + '%';
});

</script>
<template>
  <div class="process-bar">
    <div class="process-bar-inner">
    </div>
    <div class="process-number">
      {{ percent }}
    </div>
  </div>
</template>
<style lang="scss" scoped>
.process-bar {
  background-color: #e6e6e6;
  display: flex;
  flex-direction: column;
  overflow: inherit;
  height: fit-content;


  .process-bar-inner {
    height: 20px;
    background-color: #0077cc;
    width: v-bind(percent);
  }

  .process-number {
    display: flex;
    justify-content: center;
  }

}
</style>
