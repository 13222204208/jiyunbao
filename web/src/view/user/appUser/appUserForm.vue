<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务机构:" prop="serviceId">
          <el-input v-model.number="formData.serviceId" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联保代码:" prop="joinCode">
          <el-input v-model="formData.joinCode" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户类型:" prop="mchType">
          <el-select v-model="formData.mchType" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in mchTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="商户入驻协议:" prop="agreement">
          <el-input v-model="formData.agreement" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="微信:" prop="wechat">
          <el-input v-model="formData.wechat" :clearable="false" placeholder="请输入" />
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
  name: 'AppUser'
}
</script>

<script setup>
import {
  createAppUser,
  updateAppUser,
  findAppUser
} from '@/api/appUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const mchTypeOptions = ref([])
const formData = ref({
            nickname: '',
            phone: '',
            serviceId: 0,
            joinCode: '',
            mchType: undefined,
            agreement: '',
            avatar: '',
            wechat: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mchTypeOptions.value = await getDictFunc('mchType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppUser(formData.value)
               break
             case 'update':
               res = await updateAppUser(formData.value)
               break
             default:
               res = await createAppUser(formData.value)
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
