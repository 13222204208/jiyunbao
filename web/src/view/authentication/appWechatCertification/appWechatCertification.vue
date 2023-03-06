<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="商户号">
         <el-input v-model="searchInfo.mercId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="流水号" prop="reqId" width="120" />
        <el-table-column align="left" label="商户号" prop="mercId" width="120" />
        <el-table-column align="left" label="商户名称" prop="mercName" width="120" />
        <el-table-column align="left" prop="qrcodeData"  label="小程序码" width="180">
          <template #default="scope">
            <el-image 
              style="width: 70%;"
              :src="imageUrl+scope.row.qrcodeData"
              >
          </el-image>
          </template>
     
        </el-table-column>
        <!-- <el-table-column align="left" label="证件类型" prop="idTypeCd" width="120" />
        <el-table-column align="left" label="联系人证件正面照片" prop="idFrontImg" width="120" />
        <el-table-column align="left" label="联系人证件反面照片" prop="idBackImg" width="120" />
        <el-table-column align="left" label="联系人证件有效期开始时间" prop="contIdValidDateBegin" width="120" />
        <el-table-column align="left" label="联系人证件有效期结束时间" prop="contIdValidDateEnd" width="120" />
        <el-table-column align="left" label="联系人证件号码" prop="contactCorpId" width="120" /> -->
        <el-table-column align="left" label="法人居住地址" prop="idAddress" width="120" />
        <el-table-column align="left" label="联系人类型" prop="merContactType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.merContactType,merContactTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="小微经营类型" prop="managementType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.managementType,managementTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="门店名称" prop="storeName" width="120" />
        <!-- <el-table-column align="left" label="业务办理授权函" prop="busAutLetterImg" width="120" /> -->
        <el-table-column align="left" fixed="right" label="按钮组" width="240">
            <template #default="scope">

              <el-button type="primary" link icon="edit" size="small" class="table-button" @click="appWechatCertificationFunc(scope.row)">实名认证</el-button>

              <el-button type="primary" link icon="edit" size="small" class="table-button" @click="getAuthMessagesByMercIdFunc(scope.row)">明细信息</el-button>

            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateAppWechatCertificationFunc(scope.row)">变更</el-button><br>



            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="certificationStateFunc(scope.row)">申请状态</el-button>

    

            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="getAuthStateFunc(scope.row)">授权状态</el-button>

            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <!-- <el-form-item label="流水号:"  prop="reqId" >
          <el-input v-model="formData.reqId" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="商户号:"  prop="mercId" >
          <el-input v-model="formData.mercId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="商户名称:"  prop="mercName" >
          <el-input v-model="formData.mercName" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="证件类型:"  prop="idTypeCd" >
          <el-input v-model="formData.idTypeCd" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件正面照片:"  prop="idFrontImg" >
          <el-input v-model="formData.idFrontImg" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件反面照片:"  prop="idBackImg" >
          <el-input v-model="formData.idBackImg" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件有效期开始时间:"  prop="contIdValidDateBegin" >
          <el-input v-model="formData.contIdValidDateBegin" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人证件有效期结束时间:"  prop="contIdValidDateEnd" >
          <el-input v-model="formData.contIdValidDateEnd" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="联系人证件号码:"  prop="contactCorpId" >
          <el-input v-model="formData.contactCorpId" :clearable="true"  placeholder="请输入" />
        </el-form-item>  -->
        <!-- <el-form-item label="法人居住地址:"  prop="idAddress" >
          <el-input v-model="formData.idAddress" :clearable="true"  placeholder="企业商户必填" />
        </el-form-item> -->
        <el-form-item label="联系人类型:"  prop="merContactType" >
          <el-select v-model="formData.merContactType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in merContactTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="小微经营类型:"  prop="managementType" >
          <el-select v-model="formData.managementType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in managementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="门店名称:"  prop="storeName" >
          <el-input v-model="formData.storeName" :clearable="true"  placeholder="小微商户必填" />
        </el-form-item>
        <!-- <el-form-item label="业务办理授权函:"  prop="busAutLetterImg" >
          <el-input v-model="formData.busAutLetterImg" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'AppWechatCertification'
}
</script>

<script setup>
import {
  createAppWechatCertification,
  deleteAppWechatCertification,
  deleteAppWechatCertificationByIds,
  updateAppWechatCertification,
  findAppWechatCertification,
  getAppWechatCertificationList,
  wechatCertification,
  certificationState,
  getAuthMessagesByMercId,
  getAuthState,
} from '@/api/appWechatCertification'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const merContactTypeOptions = ref([])
const managementTypeOptions = ref([])
const formData = ref({
        reqId: '',
        mercId: '',
        mercName: '',
        idTypeCd: '',
        idFrontImg: '',
        idBackImg: '',
        contIdValidDateBegin: '',
        contIdValidDateEnd: '',
        contactCorpId: '',
        idAddress: '',
        merContactType: 65,
        managementType: 1,
        storeName: '',
        busAutLetterImg: '',
        qrcodeData: ''
        })

const imageUrl = ref('data:image/png;base64,')
console.log("图片",imageUrl)
// 验证规则
const rule = reactive({
               mercId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAppWechatCertificationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    merContactTypeOptions.value = await getDictFunc('merContactType')
    managementTypeOptions.value = await getDictFunc('managementType	')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteAppWechatCertificationFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteAppWechatCertificationByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAppWechatCertificationFunc = async(row) => {
    const res = await findAppWechatCertification({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reappWechatCertification
        dialogFormVisible.value = true
    }
}

//获取微信实名认证的申请状态
const certificationStateFunc = async(row) => {
  let applyNo = row.applyNo
  if(applyNo == ""){
    ElMessage({
                type: 'warning',
                message: "申请单编号不能为空"
            })
    return
  }
  const res = await certificationState({ applyNo: applyNo })
    if (res.code === 0) {
      ElMessage({
                type: 'success',
                message: res.msg
            })
    }
}

//明细信息查询
const getAuthMessagesByMercIdFunc = async(row) => {
  const  res = await getAuthMessagesByMercId({mercId:row.mercId})
  if (res.code === 0) {
      ElMessage({
                type: 'success',
                message: res.msg
            })
    }
}

//授权状态查询
const getAuthStateFunc = async(row) => {
  const  res = await getAuthState({mercId:row.mercId})
  if (res.code === 0) {
      ElMessage({
                type: 'success',
                message: res.msg
            })
    }
}

//实名认证
const appWechatCertificationFunc = async(row) => {
    const res = await wechatCertification({"ID":row.ID})
    if (res.code === 0) {
            ElMessage({
                type: 'success',
                message: res.msg
            })
    }
}


// 删除行
const deleteAppWechatCertificationFunc = async (row) => {
    const res = await deleteAppWechatCertification({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        reqId: '',
        mercId: '',
        mercName: '',
        idTypeCd: '',
        idFrontImg: '',
        idBackImg: '',
        contIdValidDateBegin: '',
        contIdValidDateEnd: '',
        contactCorpId: '',
        idAddress: '',
        merContactType: undefined,
        managementType: undefined,
        storeName: '',
        busAutLetterImg: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createAppWechatCertification(formData.value)
                  break
                case 'update':
                  res = await updateAppWechatCertification(formData.value)
                  break
                default:
                  res = await createAppWechatCertification(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
