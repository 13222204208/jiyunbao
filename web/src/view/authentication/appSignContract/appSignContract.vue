<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
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
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="流水号" prop="reqId" width="180" />
        <!-- <el-table-column align="left" label="发起商户号" prop="certId" width="120" /> -->
        <el-table-column align="left" label="商户号" prop="mercId" width="180" />
        <!-- <el-table-column align="left" label="到账方式" prop="busOpenType" width="120" /> -->
        <el-table-column align="left" label="合同类型" prop="contractType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.contractType,contractTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="签约通知" prop="isSendConMsg" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.isSendConMsg,isSendConMsgOptions) }}
            </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="通知地址" prop="notifyUrl" width="120" /> -->
        <el-table-column align="left" label="客户号" prop="custId" width="180" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateAppSignContractFunc(scope.row)">签约申请</el-button>

            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="stateAppSignContractFunc(scope.row)">签约状态</el-button><br>

            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="balanceQueryFunc(scope.row)">余额查询</el-button>            

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
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="100px">
        <!-- <el-form-item label="流水号:"  prop="reqId" >
          <el-input v-model="formData.reqId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发起商户号:"  prop="certId" >
          <el-input v-model="formData.certId" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="到账方式:"  prop="busOpenType" >
          <!-- <el-input v-model="formData.busOpenType" :clearable="true"  placeholder="请输入" /> -->
          <el-checkbox-group v-model="checkList">
            <el-checkbox label="扫码工作日到账"  size="large" />
            <el-checkbox label="扫码实时到账"  size="large" />
            <el-checkbox label="刷卡工作日到账"  size="large"/>
            <el-checkbox label="刷卡实时到账"   size="large"/>
            <el-checkbox label="D1到账 "  size="large"/>
          </el-checkbox-group>

        </el-form-item>
        <el-form-item label="合同类型:"  prop="contractType" >
          <el-select v-model="formData.contractType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="签约通知:"  prop="isSendConMsg" >
          <el-select v-model="formData.isSendConMsg" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in isSendConMsgOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="通知地址:"  prop="notifyUrl" >
          <el-input v-model="formData.notifyUrl" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="客户号:"  prop="custId" >
          <el-input v-model="formData.custId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
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
  name: 'AppSignContract'
}
</script>

<script setup>
import {
  createAppSignContract,
  deleteAppSignContract,
  deleteAppSignContractByIds,
  updateAppSignContract,
  queryAuthInfo,
  findAppSignContract,
  getAppSignContractList,
  balanceQuery,
} from '@/api/appSignContract'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})


const checkList = ref([])


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
  const table = await getAppSignContractList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    contractTypeOptions.value = await getDictFunc('contractType')
    isSendConMsgOptions.value = await getDictFunc('isSendConMsg')
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
            deleteAppSignContractFunc(row)
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
      const res = await deleteAppSignContractByIds({ ids })
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
const updateAppSignContractFunc = async(row) => {
  
    const res= await updateAppSignContract(row)
    if(res.code === 0){
      ElMessage({
                type: 'success',
                message: res.msg
            })
    }
}

//签约状态查询
const stateAppSignContractFunc = async(row) => {
  
  const res= await queryAuthInfo(row)
  if(res.code === 0){
    ElMessage({
              type: 'success',
              message: res.msg
          })
  }
}

//余额查询
const balanceQueryFunc = async(row) => {
  
  const res= await balanceQuery({"mercId":row.mercId})
  if(res.code === 0){
    let total = "金额："+res.msg
    ElMessageBox.alert(total, '商户总额', {
    // if you want to disable its autofocus
    // autofocus: false,
    confirmButtonText: 'OK',

  })
  }
}

// 删除行
const deleteAppSignContractFunc = async (row) => {
    const res = await deleteAppSignContract({ ID: row.ID })
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
        certId: '',
        busOpenType: '',
        contractType: undefined,
        isSendConMsg: undefined,
        notifyUrl: '',
        custId: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
      console.log("提交的数据",checkList.value)
      let c = checkList.value
      let arr = []

      for (var i=0;i<c.length;i++)
      { 
        console.log("数据： ",c[i])
        if(c[i] =="扫码工作日到账"){
          arr.push("00")
        }
        if(c[i] == "扫码实时到账"){
          console.log("第二 ",c[i])
          arr.push("01")
        }
        
        if(c[i] == "刷卡工作日到账"){
          arr.push("10")
        }
        
        if(c[i] == "刷卡实时到账"){
          arr.push("11")
        } 
        
        if(c[i] == "D1到账"){
          arr.push("20")
        }
      }

      if(arr.length > 1){
        formData.value.busOpenType = arr.join("|")
      }else{
        formData.value.busOpenType = arr.toString()
      }

      console.log("挚友",formData.value)
  
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
