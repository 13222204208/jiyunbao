<template>
  <div>

    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
     
        </div>
        <el-table
        :data="tableData"
        row-key="ID"
        >
        <el-table-column align="left" label="ID" width="100" prop="" />
        <el-table-column align="left" label="分类等级" style="color:blue" width="100" prop="level" >
          <template #default="scope">   <el-button type="primary" plain>{{ scope.row.level }}  级</el-button>
</template>
        </el-table-column>

        <el-table-column align="left" label="名称" prop="title" width="120" />
        <!-- <el-table-column align="left" label="父ID" prop="pid" width="120" /> -->
        <el-table-column align="left" label="图标" prop="icon" width="150" > 
          <template #default="scope">
          <el-image
            style="width: 60px; height: 60px"
            :src=" urlPath+'/'+scope.row.icon"
            fit="cover"
          />
        </template>
        </el-table-column>

        <el-table-column align="left" label="状态" width="180">
            <template #default="scope">{{ scope.row.status == "1"?"已显示":"已隐藏" }}</template>
        </el-table-column>

        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
            type="primary" link icon="plus" size="small" 
              @click="addMenu(scope.row)"
            >添加子菜单</el-button>
            <el-button
            type="primary" link icon="edit" size="small" 
              @click="editMenu(scope.row.ID)"
            >编辑</el-button>
            <el-button
            type="primary" link icon="delete" size="small"
              @click="deleteMenu(scope.row.ID)"
            >删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="图标:">
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
            <el-form-item label="显示">
      <el-switch  v-model="formData.status" active-value="1"	inactive-value="0"	 />
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
  name: 'MiniClassify'
}
</script>

<script setup>
import {
  createMiniClassify,
  deleteMiniClassify,
  deleteMiniClassifyByIds,
  updateMiniClassify,
  findMiniClassify,
  getMiniClassifyList,
} from '@/api/miniClassify'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
const urlPath = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const path = ref("")

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        pid: '0',
        icon: "",
        level: 1,
        status:"1",
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(100)
const tableData = ref([])
const searchInfo = ref({})
const dialogTitle = ref('新增分类')
const isEdit = ref(false)


//图片上传
const enterImg = async(url) => {
  console.log('图片地址',url)
  path.value = import.meta.env.VITE_BASE_API+'/'+url
  formData.value.icon = url
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
	  formData.value.icon = url
	  
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
  const table = await getMiniClassifyList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = arraytotree(table.data.list)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
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
            deleteMiniClassifyFunc(row)
        })
    }


const addMenu = (row) => {
  console.log("eeee",row)
  
  dialogTitle.value = '新增分类'
  formData.value.pid = String(row.ID)
  formData.value.level = row.level + 1
  console.log(formData.value)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 修改菜单方法
const editMenu = async(id) => {
  console.log('id',id)
  dialogTitle.value = '编辑分类'
  const res = await findMiniClassify({ ID:id })
  type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reminiClassify
        isEdit.value = true

        path.value = import.meta.env.VITE_BASE_API+'/'+res.data.reminiClassify.icon
        //setOptions()
        dialogFormVisible.value = true
    }

}

const deleteMenu = (ID) => {
  ElMessageBox.confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteMiniClassify({ ID: ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}


const arraytotree = (arr)=> {
                var top = [], sub = [], tempObj = {};
                arr.forEach(function (item) {
                    if (item.pid == 0) { // 顶级分类
                        top.push(item)
                    } else {
                        sub.push(item) // 其他分类
                    }
                    item.children = []; // 默然添加children属性
                    tempObj[item.ID] = item // 用当前分类的id做key，存储在tempObj中
                })
                sub.forEach(function (item) {
                    // 取父级
                    var parent = tempObj[item.pid] || {'children': []}
                    // 把当前分类加入到父级的children中
                    parent.children.push(item)
                })
                return top
            }

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    path.value = ''
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        title: '',
        pid: '0',
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createMiniClassify(formData.value)
          break
        case 'update':
          res = await updateMiniClassify(formData.value)
          break
        default:
          res = await createMiniClassify(formData.value)
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
}
</script>

<style>
</style>