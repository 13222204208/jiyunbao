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
        <el-table-column align="left" label="商户进件编号" prop="custInfoId" width="150" />
        <el-table-column align="left" label="流水号" prop="reqId" width="150" />
        <el-table-column align="left" label="图片类别" prop="picType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.picType,uploadImgTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="图片地址" prop="imgUrl" width="150" > 
          <template #default="scope">
          <el-image
            style="width: 100px; height: 100px"
            :src="'https://jiyunbao.vvv5g.com/api/'+scope.row.imgUrl"
            fit="cover"
          />
        </template>
        </el-table-column>
        <el-table-column align="left" label="入网申请流水号" prop="sysFlowId" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateAppUploadImgFunc(scope.row)">变更图片</el-button>
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
<!--        <el-form-item label="商户进件编号:"  prop="custInfoId" >
          <el-input v-model.number="formData.custInfoId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="流水号:"  prop="reqId" >
          <el-input v-model="formData.reqId" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="图片类别:"  prop="picType" >
          <el-select v-model="formData.picType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in uploadImgTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
    <!--    <el-form-item label="图片地址:"  prop="imgUrl" >
          <el-input v-model="formData.imgUrl" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
		
		 <el-form-item label="图片:">
					<div
					  class="user-headpic-update"
					  :style="{
								'background-image': `url(${
								  path
								})`,
								'background-repeat': 'no-repeat',
								'background-size': 'cover',
								'width':'200px',
								'height':'200px',
							  }"
					>
					</div>
				
					<el-upload
					  :action="`${urlPath}/fileUploadAndDownload/upload`"
					  :before-upload="checkFile"
					  :headers="{ 'x-token': userStore.token }"
					  :on-error="uploadError"
					  :on-success="uploadSuccess"
					  :show-file-list="false"
					  class="upload-btn"
					>
					  <el-button size="small" type="primary">上传图片</el-button>
					</el-upload>
					
		        </el-form-item>
				
<!--        <el-form-item label="入网申请流水号:"  prop="sysFlowId" >
          <el-input v-model="formData.sysFlowId" :clearable="true"  placeholder="请输入" />
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
  name: 'AppUploadImg'
}
</script>

<script setup>
import {
  createAppUploadImg,
  deleteAppUploadImg,
  deleteAppUploadImgByIds,
  updateAppUploadImg,
  findAppUploadImg,
  getAppUploadImgList
} from '@/api/appUploadImg'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRoute, useRouter} from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
const urlPath = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const path = ref("")

const route = useRoute()
const router = useRouter()


console.log("获取的参数",route.query.id) 

router.isReady().then(() => {
  formData.value.custInfoId = route.query.id
})

//图片上传
const enterImg = async(url) => {
  console.log('图片地址',url)
  path.value = import.meta.env.VITE_BASE_API+'/'+url
  formData.value.imgUrl = url
}

const fullscreenLoading = ref(false)
const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 0.5
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩未见上传图片大小不能超过 500KB，请使用压缩上传')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}
const uploadSuccess = (res) => {
  fullscreenLoading.value = false
  console.log("图片",res)
  if (res.code === 0) {
	  var url = res.data.file.url
	  path.value = import.meta.env.VITE_BASE_API+'/'+url
	  console.log("图片地址",path.value)
	  formData.value.imgUrl = url
	  
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      getTableData()
    }
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}
const chooseImgRef = ref(null)
const openChooseImg = () => {
  chooseImgRef.value.open()
}



// 自动化生成的字典（可能为空）以及字段
const uploadImgTypeOptions = ref([])
const formData = ref({
        custInfoId: route.query.id,
        reqId: '',
        picType: undefined,
        imgUrl: '',
        sysFlowId: '',
        })

// 验证规则
const rule = reactive({
               picType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               imgUrl : [{
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
  searchInfo.value.custInfoId = route.query.id
  const table = await getAppUploadImgList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    uploadImgTypeOptions.value = await getDictFunc('uploadImgType')
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
            deleteAppUploadImgFunc(row)
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
      const res = await deleteAppUploadImgByIds({ ids })
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
const updateAppUploadImgFunc = async(row) => {
    const res = await findAppUploadImg({ ID: row.ID })

    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reappUploadImg
        console.log("数据详情",res.data)
        path.value = import.meta.env.VITE_BASE_API+'/'+res.data.reappUploadImg.imgUrl


        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAppUploadImgFunc = async (row) => {
    const res = await deleteAppUploadImg({ ID: row.ID })
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
        custInfoId: route.query.id,
        reqId: '',
        picType: undefined,
        imgUrl: '',
        sysFlowId: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
			 formData.value.custInfoId = parseInt(formData.value.custInfoId )
			 console.log("接收的参数",route.query.id)
			 
              let res
              switch (type.value) {
                case 'create':
                  res = await createAppUploadImg(formData.value)
                  break
                case 'update':
                  res = await updateAppUploadImg(formData.value)
                  break
                default:
                  res = await createAppUploadImg(formData.value)
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
