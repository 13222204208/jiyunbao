<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商户ID:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="流水号:" prop="reqId">
          <el-input v-model="formData.reqId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发起方商户号:" prop="certId">
          <el-input v-model="formData.certId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="代理商编号:" prop="agtMercId">
          <el-input v-model="formData.agtMercId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户名称:" prop="mercName">
          <el-input v-model="formData.mercName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户简称:" prop="mercShortName">
          <el-input v-model="formData.mercShortName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户类型:" prop="mercType">
          <el-select v-model="formData.mercType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in mercTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="mcc码:" prop="mccCd">
          <el-input v-model="formData.mccCd" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人邮箱:" prop="contactMail">
          <el-input v-model="formData.contactMail" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人:" prop="contactMan">
          <el-input v-model="formData.contactMan" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人电话:" prop="contactPhone">
          <el-input v-model="formData.contactPhone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="客户经理:" prop="cusMgrNm">
          <el-input v-model="formData.cusMgrNm" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="异步通知地址:" prop="notifyUrl">
          <el-input v-model="formData.notifyUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人证件号:" prop="crpCertNo">
          <el-input v-model="formData.crpCertNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人证件类型:" prop="crpCertType">
          <el-input v-model="formData.crpCertType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证件开始日期:" prop="certBgn">
          <el-input v-model="formData.certBgn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证件有效期:" prop="certExpire">
          <el-input v-model="formData.certExpire" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人姓名:" prop="crpNm">
          <el-input v-model="formData.crpNm" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人手机号:" prop="crpPhone">
          <el-input v-model="formData.crpPhone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算账号:" prop="stlAccNo">
          <el-input v-model="formData.stlAccNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户支行联行号:" prop="bankSubCode">
          <el-input v-model="formData.bankSubCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算账户类型:" prop="stlAccType">
          <el-select v-model="formData.stlAccType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in stlAccTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="营业归属省代码:" prop="busProviceCode">
          <el-input v-model="formData.busProviceCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业归属市代码:" prop="busCityCode">
          <el-input v-model="formData.busCityCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业详细地址:" prop="busAddr">
          <el-input v-model="formData.busAddr" :clearable="true" placeholder="请输入" />
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
  name: 'AppCustInfo'
}
</script>

<script setup>
import {
  createAppCustInfo,
  updateAppCustInfo,
  findAppCustInfo
} from '@/api/appCustInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const mercTypeOptions = ref([])
const stlAccTypeOptions = ref([])
const formData = ref({
            uid: 0,
            reqId: '',
            certId: '',
            agtMercId: '',
            mercName: '',
            mercShortName: '',
            mercType: undefined,
            mccCd: '',
            contactMail: '',
            contactMan: '',
            contactPhone: '',
            cusMgrNm: '',
            notifyUrl: '',
            crpCertNo: '',
            crpCertType: '',
            certBgn: '',
            certExpire: '',
            crpNm: '',
            crpPhone: '',
            stlAccNo: '',
            bankSubCode: '',
            stlAccType: undefined,
            busProviceCode: '',
            busCityCode: '',
            busAddr: '',
        })
// 验证规则
const rule = reactive({
               mercName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mercType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mccCd : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               contactMail : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cusMgrNm : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               notifyUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               crpCertNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               crpCertType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               certBgn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               certExpire : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               crpNm : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               crpPhone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               stlAccNo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               bankSubCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               stlAccType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               busProviceCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               busCityCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               busAddr : [{
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
      const res = await findAppCustInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappCustInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mercTypeOptions.value = await getDictFunc('mercType')
    stlAccTypeOptions.value = await getDictFunc('stlAccType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppCustInfo(formData.value)
               break
             case 'update':
               res = await updateAppCustInfo(formData.value)
               break
             default:
               res = await createAppCustInfo(formData.value)
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
