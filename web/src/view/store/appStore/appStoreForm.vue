<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="门店用户:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店名称:" prop="storeName">
          <el-input v-model="formData.storeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店头像:" prop="storeAvatar">
          <el-input v-model="formData.storeAvatar" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主营品类:" prop="category">
          <el-input v-model="formData.category" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店电话:" prop="storePhone">
          <el-input v-model="formData.storePhone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店地址:" prop="storeAddress">
          <el-input v-model="formData.storeAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="经度:" prop="longitude">
          <el-input v-model="formData.longitude" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="纬度:" prop="latitude">
          <el-input v-model="formData.latitude" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店公告:" prop="notice">
          <el-input v-model="formData.notice" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算模式:" prop="close">
          <el-input v-model="formData.close" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务设置:" prop="service">
          <el-input v-model="formData.service" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同协议中心:" prop="agreement">
          <el-input v-model="formData.agreement" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合作状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in storeStateOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AppStore'
}
</script>

<script setup>
import {
  createAppStore,
  updateAppStore,
  findAppStore
} from '@/api/appStore'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const storeStateOptions = ref([])
const formData = ref({
            uid: 0,
            storeName: '',
            storeAvatar: '',
            category: '',
            storePhone: '',
            storeAddress: '',
            longitude: '',
            latitude: '',
            notice: '',
            close: '',
            service: '',
            agreement: '',
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppStore({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappStore
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    storeStateOptions.value = await getDictFunc('storeState')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppStore(formData.value)
               break
             case 'update':
               res = await updateAppStore(formData.value)
               break
             default:
               res = await createAppStore(formData.value)
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
