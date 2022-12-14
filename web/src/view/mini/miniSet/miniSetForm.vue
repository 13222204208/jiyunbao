<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="小程序类型:" prop="miniType">
          <el-select v-model="formData.miniType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in miniOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="AppId:" prop="appid">
          <el-input v-model="formData.appid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="秘匙:" prop="appSecret">
          <el-input v-model="formData.appSecret" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户号:" prop="mchId">
          <el-input v-model="formData.mchId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付秘钥:" prop="mchSecret">
          <el-input v-model="formData.mchSecret" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公钥:" prop="public">
          <el-input v-model="formData.public" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="私钥:" prop="private">
          <el-input v-model="formData.private" :clearable="true" placeholder="请输入" />
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
  name: 'MiniSet'
}
</script>

<script setup>
import {
  createMiniSet,
  updateMiniSet,
  findMiniSet
} from '@/api/miniSet'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const miniOptions = ref([])
const formData = ref({
            miniType: undefined,
            appid: '',
            appSecret: '',
            mchId: '',
            mchSecret: '',
            public: '',
            private: '',
        })
// 验证规则
const rule = reactive({
               miniType : [{
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
      const res = await findMiniSet({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reminiSet
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    miniOptions.value = await getDictFunc('mini')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMiniSet(formData.value)
               break
             case 'update':
               res = await updateMiniSet(formData.value)
               break
             default:
               res = await createMiniSet(formData.value)
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
