<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="请求地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="接口类型:" prop="way">
          <el-select v-model="formData.way" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in requestUrlOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AppUrl'
}
</script>

<script setup>
import {
  createAppUrl,
  updateAppUrl,
  findAppUrl
} from '@/api/appUrl'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const requestUrlOptions = ref([])
const formData = ref({
            url: '',
            way: undefined,
        })
// 验证规则
const rule = reactive({
               url : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               way : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppUrl({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappUrl
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    requestUrlOptions.value = await getDictFunc('requestUrl')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppUrl(formData.value)
               break
             case 'update':
               res = await updateAppUrl(formData.value)
               break
             default:
               res = await createAppUrl(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
