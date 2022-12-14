<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="机构名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="机构类型:" prop="way">
          <el-select v-model="formData.way" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in mchTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人姓名:" prop="principal">
          <el-input v-model="formData.principal" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:" prop="card">
          <el-input v-model="formData.card" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:" prop="mail">
          <el-input v-model="formData.mail" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属地区:" prop="area">
          <el-input v-model="formData.area" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="详细地区:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资料状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in serviceStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="资质状态:" prop="certification">
          <el-select v-model="formData.certification" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in serviceCertificationOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AppFacilitatingAgency'
}
</script>

<script setup>
import {
  createAppFacilitatingAgency,
  updateAppFacilitatingAgency,
  findAppFacilitatingAgency
} from '@/api/appFacilitatingAgency'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const mchTypeOptions = ref([])
const serviceStateOptions = ref([])
const serviceCertificationOptions = ref([])
const formData = ref({
            phone: '',
            name: '',
            way: undefined,
            principal: '',
            card: '',
            mail: '',
            area: '',
            address: '',
            status: undefined,
            certification: undefined,
        })
// 验证规则
const rule = reactive({
               phone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               way : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               principal : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               certification : [{
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
      const res = await findAppFacilitatingAgency({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappFacilitatingAgency
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mchTypeOptions.value = await getDictFunc('mchType')
    serviceStateOptions.value = await getDictFunc('serviceState')
    serviceCertificationOptions.value = await getDictFunc('serviceCertification')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppFacilitatingAgency(formData.value)
               break
             case 'update':
               res = await updateAppFacilitatingAgency(formData.value)
               break
             default:
               res = await createAppFacilitatingAgency(formData.value)
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
