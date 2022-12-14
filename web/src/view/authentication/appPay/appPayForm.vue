<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业执照原件:" prop="businessLicense">
          <el-input v-model="formData.businessLicense" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户许可证:" prop="openCard">
          <el-input v-model="formData.openCard" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户名称:" prop="mchName">
          <el-input v-model="formData.mchName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户简称:" prop="mchForShort">
          <el-input v-model="formData.mchForShort" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户地址:" prop="mchAddress">
          <el-input v-model="formData.mchAddress" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人姓名:" prop="legalPerson">
          <el-input v-model="formData.legalPerson" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人电话:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人身份证号:" prop="legalCard">
          <el-input v-model="formData.legalCard" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证正面:" prop="cardFront">
          <el-input v-model="formData.cardFront" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证反面:" prop="cardReverse">
          <el-input v-model="formData.cardReverse" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证有效期:" prop="cardValidity">
          <el-input v-model="formData.cardValidity" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户类型:" prop="mchType">
          <el-input v-model="formData.mchType" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in checkStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人姓名:" prop="linkmanName">
          <el-input v-model="formData.linkmanName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人电话:" prop="linkmanPhone">
          <el-input v-model="formData.linkmanPhone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人邮箱:" prop="linkmanMail">
          <el-input v-model="formData.linkmanMail" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人通讯地址:" prop="linkmanAddress">
          <el-input v-model="formData.linkmanAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行卡正面图片:" prop="bankFrontImg">
          <el-input v-model="formData.bankFrontImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算卡类型:" prop="closeType">
          <el-input v-model="formData.closeType" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算帐户卡号:" prop="closeCardNum">
          <el-input v-model="formData.closeCardNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算帐户户名:" prop="closeCardName">
          <el-input v-model="formData.closeCardName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行预留手机号:" prop="bankPhone">
          <el-input v-model="formData.bankPhone" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行:" prop="openBank">
          <el-input v-model="formData.openBank" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行编号支行:" prop="openBankNum">
          <el-input v-model="formData.openBankNum" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行名称支行:" prop="openBankName">
          <el-input v-model="formData.openBankName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行省市区:" prop="openBankCity">
          <el-input v-model="formData.openBankCity" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="起始结算金额:" prop="startMoney">
          <el-input v-model.number="formData.startMoney" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人结算:" prop="legalPersonClose">
          <el-input v-model="formData.legalPersonClose" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="非法人姓名:" prop="notLegalPersonName">
          <el-input v-model="formData.notLegalPersonName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="非法人身份证号码:" prop="notLegalPersonCardNum">
          <el-input v-model="formData.notLegalPersonCardNum" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="非法人结算授权书:" prop="notLegalPersonAuthImg">
          <el-input v-model="formData.notLegalPersonAuthImg" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="被授权人身份证正面:" prop="authCardFrontImg">
          <el-input v-model="formData.authCardFrontImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="被授权人身份证反面:" prop="authCardReverseImg">
          <el-input v-model="formData.authCardReverseImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="费率:" prop="rate">
          <el-input v-model="formData.rate" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="协议类型:" prop="agreementType">
          <el-input v-model="formData.agreementType" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店门头照:" prop="storeHeadImg">
          <el-input v-model="formData.storeHeadImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店内景照:" prop="storeImg">
          <el-input v-model="formData.storeImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门店收银台照:" prop="storeCashierImg">
          <el-input v-model="formData.storeCashierImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="租赁合同:" prop="contractImg">
          <el-input v-model="formData.contractImg" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其它资料:" prop="otherImg">
          <el-input v-model="formData.otherImg" :clearable="false" placeholder="请输入" />
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
  name: 'AppPay'
}
</script>

<script setup>
import {
  createAppPay,
  updateAppPay,
  findAppPay
} from '@/api/appPay'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const checkStateOptions = ref([])
const formData = ref({
            uid: 0,
            businessLicense: '',
            openCard: '',
            mchName: '',
            mchForShort: '',
            mchAddress: '',
            legalPerson: '',
            phone: '',
            legalCard: '',
            cardFront: '',
            cardReverse: '',
            cardValidity: '',
            mchType: '',
            status: undefined,
            linkmanName: '',
            linkmanPhone: '',
            linkmanMail: '',
            linkmanAddress: '',
            bankFrontImg: '',
            closeType: '',
            closeCardNum: '',
            closeCardName: '',
            bankPhone: '',
            openBank: '',
            openBankNum: '',
            openBankName: '',
            openBankCity: '',
            startMoney: 0,
            legalPersonClose: '',
            notLegalPersonName: '',
            notLegalPersonCardNum: '',
            notLegalPersonAuthImg: '',
            authCardFrontImg: '',
            authCardReverseImg: '',
            rate: '',
            agreementType: '',
            storeHeadImg: '',
            storeImg: '',
            storeCashierImg: '',
            contractImg: '',
            otherImg: '',
        })
// 验证规则
const rule = reactive({
               businessLicense : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               openCard : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mchName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mchAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               legalPerson : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               phone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               legalCard : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cardFront : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cardReverse : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cardValidity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mchType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               linkmanName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               linkmanPhone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               linkmanMail : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               linkmanAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               bankFrontImg : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               closeType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               closeCardNum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               bankPhone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               openBankNum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               openBankName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               openBankCity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startMoney : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               storeHeadImg : [{
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
      const res = await findAppPay({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappPay
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    checkStateOptions.value = await getDictFunc('checkState')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAppPay(formData.value)
               break
             case 'update':
               res = await updateAppPay(formData.value)
               break
             default:
               res = await createAppPay(formData.value)
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
