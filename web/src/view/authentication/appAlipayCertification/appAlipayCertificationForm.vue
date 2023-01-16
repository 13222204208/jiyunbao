<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="流水号:" prop="reqId">
          <el-input v-model="formData.reqId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户号:" prop="mercId">
          <el-input v-model="formData.mercId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户名称:" prop="mercName">
          <el-input v-model="formData.mercName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件号码:" prop="contactCorpId">
          <el-input v-model="formData.contactCorpId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件手机号:" prop="contactsTel">
          <el-input v-model="formData.contactsTel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证件持有人类型:" prop="idHolderType">
          <el-input v-model="formData.idHolderType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人证件有效开始日期:" prop="idValidDateBegin">
          <el-input v-model="formData.idValidDateBegin" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人证件有效结束日期:" prop="idValidDateEnd">
          <el-input v-model="formData.idValidDateEnd" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业执照有效开始日期:" prop="busLincenceBegin">
          <el-input v-model="formData.busLincenceBegin" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业执照有效结束日期:" prop="busLincenceEnd">
          <el-input v-model="formData.busLincenceEnd" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人类型:" prop="merContactType">
          <el-select v-model="formData.merContactType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in merContactTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="小微经营类型:" prop="managementType">
          <el-select v-model="formData.managementType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in managementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="门店名称:" prop="storeName">
          <el-input v-model="formData.storeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件类型:" prop="idTypeCd">
          <el-input v-model="formData.idTypeCd" :clearable="true" placeholder="请输入" />
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
  name: 'AppAlipayCertification'
}
</script>

<script setup>
import {
  createAppAlipayCertification,
  updateAppAlipayCertification,
  findAppAlipayCertification
} from '@/api/appAlipayCertification'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const merContactTypeOptions = ref([])
const managementTypeOptions = ref([])
const formData = ref({
            reqId: '',
            mercId: '',
            mercName: '',
            contactCorpId: '',
            contactsTel: '',
            idHolderType: '',
            idValidDateBegin: '',
            idValidDateEnd: '',
            busLincenceBegin: '',
            busLincenceEnd: '',
            merContactType: undefined,
            managementType: undefined,
            storeName: '',
            idTypeCd: '',
        })
// 验证规则
const rule = reactive({
               mercId : [{
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
      const res = await findAppAlipayCertification({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappAlipayCertification
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    merContactTypeOptions.value = await getDictFunc('merContactType')
    managementTypeOptions.value = await getDictFunc('managementType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppAlipayCertification(formData.value)
               break
             case 'update':
               res = await updateAppAlipayCertification(formData.value)
               break
             default:
               res = await createAppAlipayCertification(formData.value)
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
