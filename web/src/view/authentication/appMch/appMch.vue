<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <!-- <el-form-item label="企业名称">
         <el-input v-model="searchInfo.firmName" placeholder="搜索条件" />

        </el-form-item> -->
        <el-form-item label="店铺名称">
         <el-input v-model="searchInfo.storeName" placeholder="搜索条件" />

        </el-form-item>

        <el-form-item label="联系电话">
         <el-input v-model="searchInfo.storeName" placeholder="搜索条件" />

        </el-form-item>

           <el-form-item label="状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in checkStateOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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

        <el-table-column align="center" label="商户ID" prop="uid" width="80" />
        <el-table-column align="left" label="商户昵称" prop="" width="120" />
        <el-table-column align="left" label="企业名称" prop="firmName" width="120" />
        <el-table-column align="left" label="店铺名称" prop="storeName" width="120" />
        <el-table-column align="left" label="商户种类" prop="" width="120" />
        <el-table-column align="left" label="注册手机" prop="" width="120" />
        <el-table-column align="left" label="注册日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="商户类型" prop="" width="120" />
        <el-table-column align="left" label="商户状态" prop="" width="120" />
        <el-table-column align="left" label="结算模式" prop="" width="120" />
        <el-table-column align="left" label="联保代码" prop="" width="120" />

        <el-table-column align="left" label="主营品类" prop="mainType" width="120" />
        <el-table-column align="left" label="从业资格证" prop="" width="120" />
        <el-table-column align="left" label="车辆运输证" prop="" width="120" />
        <el-table-column align="left" label="车辆照片" prop="" width="120" />
        <el-table-column align="left" label="驾驶室照片" prop="" width="120" />
        <el-table-column align="left" label="负责人照片" prop="" width="120" />
        <el-table-column align="left" label="身份证头像面" prop="" width="120" />
        <el-table-column align="left" label="身份证国徽面" prop="" width="120" />
        <el-table-column align="left" label="营业执照" prop="" width="120" />
        <el-table-column align="left" label="许可证" prop="" width="120" />
        <el-table-column align="left" label="其它资料" prop="" width="120" />

<!--        <el-table-column align="left" label="确认佣金" prop="commission" width="120" />
        <el-table-column align="left" label="T+1结算" prop="final" width="120" />
        <el-table-column align="left" label="特约收单" prop="acquirer" width="120" />
        <el-table-column align="left" label="法人" prop="legalPerson" width="120" />
        <el-table-column align="left" label="热线电话" prop="phone" width="120" />
        <el-table-column align="left" label="负责人头像" prop="avatar" width="120" />
        <el-table-column align="left" label="身份证正面" prop="cardFront" width="120" />
        <el-table-column align="left" label="身份证反面" prop="cardReverse" width="120" />
        <el-table-column align="left" label="委托书" prop="entrust" width="120" /> -->
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,checkStateOptions) }}
            </template>
        </el-table-column>
        <el-table-column fixed="right" label="按钮组" width="175">
            <template #default="scope">
              <el-button type="primary" link  size="small" class="table-button" >支付认证</el-button>
              
            <el-button type="primary" link  size="small" class="table-button" @click="updateAppMchFunc(scope.row)">变更</el-button>
            <el-button type="primary" link  size="small" class="table-button" >审核</el-button><br>
            <el-button type="primary" link  size="small" class="table-button" >认证查询</el-button>
            <el-button type="primary" link size="small" @click="deleteRow(scope.row)" class="table-button">删除</el-button>
            <el-button type="primary" link  size="small" class="table-button" >禁用</el-button>
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
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="商户ID:"  prop="uid" >
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="企业名称:"  prop="firmName" >
          <el-input v-model="formData.firmName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺名称:"  prop="storeName" >
          <el-input v-model="formData.storeName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主营品类:"  prop="mainType" >
          <el-input v-model="formData.mainType" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="确认佣金:"  prop="commission" >
          <el-input v-model.number="formData.commission" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="T+1结算:"  prop="final" >
          <el-input v-model.number="formData.final" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="特约收单:"  prop="acquirer" >
          <el-input v-model.number="formData.acquirer" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人:"  prop="legalPerson" >
          <el-input v-model="formData.legalPerson" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="热线电话:"  prop="phone" >
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责人头像:"  prop="avatar" >
          <el-input v-model="formData.avatar" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证正面:"  prop="cardFront" >
          <el-input v-model="formData.cardFront" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证反面:"  prop="cardReverse" >
          <el-input v-model="formData.cardReverse" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="委托书:"  prop="entrust" >
          <el-input v-model="formData.entrust" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:"  prop="remark" >
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in checkStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'AppMch'
}
</script>

<script setup>
import {
  createAppMch,
  deleteAppMch,
  deleteAppMchByIds,
  updateAppMch,
  findAppMch,
  getAppMchList
} from '@/api/appMch'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getAppMchList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    checkStateOptions.value = await getDictFunc('checkState')
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
            deleteAppMchFunc(row)
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
      const res = await deleteAppMchByIds({ ids })
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
const updateAppMchFunc = async(row) => {
    const res = await findAppMch({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reappMch
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppMchFunc = async (row) => {
    const res = await deleteAppMch({ ID: row.ID })
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
