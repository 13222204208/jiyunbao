<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="手机号">
         <el-input v-model="searchInfo.phone" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="机构名称">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增机构</el-button>
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
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="机构名称" prop="name" width="120" />
        <el-table-column align="left" label="机构类型" prop="way" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.way,mchTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="负责人姓名" prop="principal" width="120" />
        <el-table-column align="left" label="身份证号" prop="card" width="120" />
        <el-table-column align="left" label="邮箱" prop="mail" width="120" />
        <el-table-column align="left" label="所属地区" prop="area" width="120" />
        <el-table-column align="left" label="详细地区" prop="address" width="120" />
        <el-table-column align="left" label="资料状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,serviceStateOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="资质状态" prop="certification" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.certification,serviceCertificationOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateAppFacilitatingAgencyFunc(scope.row)">变更</el-button>
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
        <el-form-item label="手机号:"  prop="phone" >
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="机构名称:"  prop="name" >
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="机构类型:"  prop="way" >
          <el-select v-model="formData.way" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in mchTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人姓名:"  prop="principal" >
          <el-input v-model="formData.principal" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:"  prop="card" >
          <el-input v-model="formData.card" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:"  prop="mail" >
          <el-input v-model="formData.mail" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属地区:"  prop="area" >
          <el-input v-model="formData.area" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="详细地区:"  prop="address" >
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资料状态:"  prop="status" >
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in serviceStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="资质状态:"  prop="certification" >
          <el-select v-model="formData.certification" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in serviceCertificationOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AppFacilitatingAgency'
}
</script>

<script setup>
import {
  createAppFacilitatingAgency,
  deleteAppFacilitatingAgency,
  deleteAppFacilitatingAgencyByIds,
  updateAppFacilitatingAgency,
  findAppFacilitatingAgency,
  getAppFacilitatingAgencyList
} from '@/api/appFacilitatingAgency'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getAppFacilitatingAgencyList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    mchTypeOptions.value = await getDictFunc('mchType')
    serviceStateOptions.value = await getDictFunc('serviceState')
    serviceCertificationOptions.value = await getDictFunc('serviceCertification')
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
            deleteAppFacilitatingAgencyFunc(row)
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
      const res = await deleteAppFacilitatingAgencyByIds({ ids })
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
const updateAppFacilitatingAgencyFunc = async(row) => {
    const res = await findAppFacilitatingAgency({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reappFacilitatingAgency
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppFacilitatingAgencyFunc = async (row) => {
    const res = await deleteAppFacilitatingAgency({ ID: row.ID })
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
