<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商户ID:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="企业名称:" prop="firmName">
          <el-input v-model="formData.firmName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺名称:" prop="storeName">
          <el-input v-model="formData.storeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主营品类:" prop="mainType">
          <el-input v-model="formData.mainType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="确认佣金:" prop="commission">
          <el-input v-model.number="formData.commission" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="T+1结算:" prop="final">
          <el-input v-model.number="formData.final" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="特约收单:" prop="acquirer">
          <el-input v-model.number="formData.acquirer" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人:" prop="legalPerson">
          <el-input v-model="formData.legalPerson" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="热线电话:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责人头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证正面:" prop="cardFront">
          <el-input v-model="formData.cardFront" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证反面:" prop="cardReverse">
          <el-input v-model="formData.cardReverse" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="委托书:" prop="entrust">
          <el-input v-model="formData.entrust" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in checkStateOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AppMch'
}
</script>

<script setup>
import {
  createAppMch,
  updateAppMch,
  findAppMch
} from '@/api/appMch'

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
            firmName: '',
            storeName: '',
            mainType: '',
            commission: 0,
            final: 0,
            acquirer: 0,
            legalPerson: '',
            phone: '',
            avatar: '',
            cardFront: '',
            cardReverse: '',
            entrust: '',
            remark: '',
            status: undefined,
        })
// 验证规则
const rule = reactive({
               firmName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mainType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               commission : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               legalPerson : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               avatar : [{
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
               entrust : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findAppMch({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reappMch
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
               res = await createAppMch(formData.value)
               break
             case 'update':
               res = await updateAppMch(formData.value)
               break
             default:
               res = await createAppMch(formData.value)
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
