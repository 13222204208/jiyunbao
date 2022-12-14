<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联保代码:" prop="userNum">
          <el-input v-model="formData.userNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="小程序类型:" prop="miniType">
          <el-select v-model="formData.miniType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in miniOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="小程序标识:" prop="openid">
          <el-input v-model="formData.openid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实名状态:" prop="realNameState">
          <el-select v-model="formData.realNameState" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in realNameStateOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'MiniUser'
}
</script>

<script setup>
import {
  createMiniUser,
  updateMiniUser,
  findMiniUser
} from '@/api/miniUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const miniOptions = ref([])
const realNameStateOptions = ref([])
const formData = ref({
            nickname: '',
            avatar: '',
            userNum: '',
            miniType: undefined,
            openid: '',
            phone: '',
            realNameState: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMiniUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reminiUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    miniOptions.value = await getDictFunc('mini')
    realNameStateOptions.value = await getDictFunc('realNameState')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMiniUser(formData.value)
               break
             case 'update':
               res = await updateMiniUser(formData.value)
               break
             default:
               res = await createMiniUser(formData.value)
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
