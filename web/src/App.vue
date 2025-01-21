<template>
  <div
    id="app"
    class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800"
  >
    <el-config-provider :locale="zhCn">
      <router-view />
      <button @click="changeLanguage('en')">English</button>
      <button @click="changeLanguage('ru')">Русский</button>
      <button @click="changeLanguage('zh')">中文</button>
    </el-config-provider>
  </div>
</template>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
//import ru from 'element-plus/dist/locale/ru.mjs'
  import { useAppStore } from '@/pinia'

  useAppStore()
  defineOptions({
    name: 'App'
  })

</script>
<script>
export default {
  methods: {
    changeLanguage(lang) {
      this.$i18n.locale = lang; // Смена языка
    },
  },
};
</script>
<style lang="scss">
  // 引入初始化样式
  #app {
    height: 100vh;
    overflow: hidden;
    font-weight: 400 !important;
  }
  .el-button {
    font-weight: 400 !important;
  }

  .gva-body-h {
    min-height: calc(100% - 3rem);
  }

  .gva-container {
    height: calc(100% - 2.5rem);
  }
  .gva-container2 {
    height: calc(100% - 4.5rem);
  }
</style>
