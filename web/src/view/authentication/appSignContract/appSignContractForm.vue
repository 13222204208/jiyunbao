<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="流水号:" prop="reqId">
          <el-input v-model="formData.reqId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发起商户号:" prop="certId">
          <el-input v-model="formData.certId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="到账方式:" prop="busOpenType">
          <el-input v-model="formData.busOpenType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同类型:" prop="contractType">
          <el-select v-model="formData.contractType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="签约通知:" prop="isSendConMsg">
          <el-select v-model="formData.isSendConMsg" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in isSendConMsgOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="通知地址:" prop="notifyUrl">
          <el-input v-model="formData.notifyUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="客户号:" prop="custId">
          <el-input v-model="formData.custId" :clearable="true" placeholder="请输入" />
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
  name: 'AppSignContract'
}
</script>

<script setup>
import {
  createAppSignContract,
  updateAppSignContract,
  findAppSignContract
} from '@/api/appSignContract'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const contractTypeOptions = ref([])
const isSendConMsgOptions = ref([])
const formData = ref({
            reqId: '',
            certId: '',
            busOpenType: '',
            contractType: undefined,
            isSendConMsg: undefined,
            notifyUrl: '',
            custId: '',
        })
// 验证规则
const rule = reactive({
               contractType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isSendConMsg : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               custId : [{
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
      const res = await findAppSignContract({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappSignContract
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    contractTypeOptions.value = await getDictFunc('contractType')
    isSendConMsgOptions.value = await getDictFunc('isSendConMsg')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppSignContract(formData.value)
               break
             case 'update':
               res = await updateAppSignContract(formData.value)
               break
             default:
               res = await createAppSignContract(formData.value)
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
