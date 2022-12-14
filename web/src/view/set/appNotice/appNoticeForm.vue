<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公告内容:" prop="contents">
          <el-input v-model="formData.contents" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公告类型:" prop="way">
          <el-select v-model="formData.way" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in noticeOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AppNotice'
}
</script>

<script setup>
import {
  createAppNotice,
  updateAppNotice,
  findAppNotice
} from '@/api/appNotice'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const noticeOptions = ref([])
const formData = ref({
            contents: '',
            way: undefined,
        })
// 验证规则
const rule = reactive({
               contents : [{
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
      const res = await findAppNotice({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappNotice
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    noticeOptions.value = await getDictFunc('notice')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppNotice(formData.value)
               break
             case 'update':
               res = await updateAppNotice(formData.value)
               break
             default:
               res = await createAppNotice(formData.value)
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
