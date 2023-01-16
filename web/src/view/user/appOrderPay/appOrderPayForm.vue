<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户号:" prop="mercId">
          <el-input v-model="formData.mercId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="subject">
          <el-input v-model="formData.subject" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付方式:" prop="payType">
          <el-select v-model="formData.payType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in payTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="支付金额:" prop="totalAmount">
          <el-input-number v-model="formData.totalAmount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="买家ID:" prop="buyer_id">
          <el-input v-model="formData.buyer_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="appid:" prop="appid">
          <el-input v-model="formData.appid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付状态:" prop="payState">
          <el-select v-model="formData.payState" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in payStateOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AppOrderPay'
}
</script>

<script setup>
import {
  createAppOrderPay,
  updateAppOrderPay,
  findAppOrderPay
} from '@/api/appOrderPay'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const payTypeOptions = ref([])
const payStateOptions = ref([])
const formData = ref({
            orderNo: '',
            mercId: '',
            subject: '',
            payType: undefined,
            totalAmount: 0,
            buyer_id: '',
            appid: '',
            payState: undefined,
        })
// 验证规则
const rule = reactive({
               payType : [{
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
      const res = await findAppOrderPay({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappOrderPay
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    payTypeOptions.value = await getDictFunc('payType')
    payStateOptions.value = await getDictFunc('payState')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppOrderPay(formData.value)
               break
             case 'update':
               res = await updateAppOrderPay(formData.value)
               break
             default:
               res = await createAppOrderPay(formData.value)
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
