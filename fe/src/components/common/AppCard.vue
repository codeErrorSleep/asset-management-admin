<!--------------------------------
 - @Author: Ronnie Zhang
 - @LastEditor: Ronnie Zhang
 - @LastEditTime: 2023/12/04 22:51:04
 - @Email: zclzone@outlook.com
 -  | https://isme.top
 --------------------------------->

<template>
  <div class="auto-bg" :class="{ 'card-border': bordered }">
    <slot />
  </div>
</template>

<script setup>
defineProps({
  bordered: Boolean,
})
</script>
