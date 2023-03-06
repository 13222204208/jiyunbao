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

        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmitForbidden">已禁用用户</el-button>
          
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <!-- <div class="gva-btn-list">
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
        </div> -->
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
        <el-table-column fixed="right" label="按钮组" width="195">
            <template #default="scope">
              <el-button type="primary" link  size="small" class="table-button" @click="toCustInfo(scope.row)" >支付认证</el-button>
              
            <el-button type="primary" link  size="small" class="table-button" @click="updateAppMchFunc(scope.row)">审核</el-button><br>
            <!-- <el-button type="primary" link  size="small" class="table-button" >审核</el-button><br> -->
            <el-button type="primary" link  size="small" class="table-button" >认证查询</el-button>
            <el-button type="primary" link size="small" @click="deleteRow(scope.row)" class="table-button">删除</el-button>
            <el-button type="primary" link  size="small" @click="forbiddenRow(scope.row)" class="table-button" >{{scope.row.forbidden == 1 ?"禁用":"取消禁用"}}</el-button>
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
        <!-- <el-form-item label="商户ID:"  prop="uid" >
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="用户名:"  prop="" >
          <!-- <el-input v-model="formData.AppUser.nickname"   placeholder="请输入" /> -->
          {{ formData.AppUser.nickname }}
        </el-form-item>

        <el-form-item label="企业名称:"  prop="firmName" >
          <el-input v-model="formData.firmName"   placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户名称:"  prop="storeName" >
          <el-input v-model="formData.storeName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联保代码:"  prop="" >
          <!-- <el-input v-model="formData.AppUser.joinCode" :clearable="true"  placeholder="请输入" /> -->
          {{ formData.AppUser.joinCode }}
        </el-form-item>

        <el-form-item label="商户类型:"  prop="" >
          <el-radio-group v-model="formData.mchType" class="ml-4">
              <el-radio :label=1>小微商户1.0 </el-radio>
              <el-radio :label=2>小微商户2.0 </el-radio>
              <el-radio :label=3>个人商户</el-radio>
              <el-radio :label=4>企业 </el-radio>
            </el-radio-group>
        </el-form-item>

        <el-form-item label="主营品类:"  prop="mainType" >
          <!-- <el-input v-model="formData.AppUser.Classify.title" :clearable="true"  placeholder="请输入" /> -->
          {{ formData.AppUser.Classify.title }}
        </el-form-item>
        <!-- <el-form-item label="确认佣金:"  prop="commission" >
          <el-input v-model.number="formData.commission" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="注册手机:"  prop="" >
          <!-- <el-input v-model="formData.AppUser.phone" :clearable="true"  placeholder="请输入" /> -->
          {{formData.AppUser.phone}}
        </el-form-item>
        <!-- <el-form-item label="T+1结算:"  prop="final" >
          <el-input v-model.number="formData.final" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="特约收单:"  prop="acquirer" >
          <el-input v-model.number="formData.acquirer" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="结算模式:"  prop="final" >
            <el-radio-group v-model="formData.final" class="ml-4">
              <el-radio :label=1>T+1结算 </el-radio>
              <el-radio :label=0>特约收单 </el-radio>
            </el-radio-group>
        </el-form-item>
        <el-form-item label="法人代表:"  prop="legalPerson" >
          <el-input v-model="formData.legalPerson"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="热线电话:"  prop="phone" >
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入" />
        </el-form-item>

        <el-form-item label="人车合影照:">				
					<el-upload
            class="avatar-uploader"			
					  :action="`${urlPath}fileUploadAndDownload/upload`"
					  :before-upload="checkFile"
					  :headers="{ 'x-token': userStore.token }"
					  :on-error="uploadError"
					  :on-success="uploadSuccess"
					  :show-file-list="false"
           	>
           <img v-if="formData.manAndCar" :src="urlPath+formData.manAndCar" class="avatar" style="width:200px;height:200px"/>
					  <el-button v-else  size="small" type="primary">上传图片</el-button>
					</el-upload>
		    </el-form-item>

        <el-form-item label="从业资格证:">				
					<el-upload
            class="avatar-uploader"			
					  :action="`${urlPath}fileUploadAndDownload/upload`"
					  :before-upload="checkFile"
					  :headers="{ 'x-token': userStore.token }"
					  :on-error="uploadError"
					  :on-success="uploadSuccessTwo"
					  :show-file-list="false"
           	>
           <img v-if="formData.job" :src="urlPath+formData.job" class="avatar" style="width:200px;height:200px"/>
					  <el-button v-else  size="small" type="primary">上传图片</el-button>
					</el-upload>
		    </el-form-item>

        <el-form-item label="车辆运输证:">				
					<el-upload
            class="avatar-uploader"			
					  :action="`${urlPath}fileUploadAndDownload/upload`"
					  :before-upload="checkFile"
					  :headers="{ 'x-token': userStore.token }"
					  :on-error="uploadError"
					  :on-success="uploadSuccessThree"
					  :show-file-list="false"
           	>
           <img v-if="formData.transport" :src="urlPath+formData.transport" class="avatar" style="width:200px;height:200px"/>
					  <el-button v-else  size="small" type="primary">上传图片</el-button>
					</el-upload>
		    </el-form-item>

          <el-form-item label="驾驶室照片:">				
					<el-upload
            class="avatar-uploader"			
					  :action="`${urlPath}fileUploadAndDownload/upload`"
					  :before-upload="checkFile"
					  :headers="{ 'x-token': userStore.token }"
					  :on-error="uploadError"
					  :on-success="uploadSuccessFour"
					  :show-file-list="false"
           	>
           <img v-if="formData.cab" :src="urlPath+formData.cab" class="avatar" style="width:200px;height:200px"/>
					  <el-button v-else  size="small" type="primary">上传图片</el-button>
					</el-upload>
		    </el-form-item>

        <!-- <el-form-item label="负责人头像:"  prop="avatar" >
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
        </el-form-item> -->
        <el-form-item label="备注:"  prop="remark" >
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核:"  prop="status" >
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
import {useRouter} from 'vue-router'

import { useUserStore } from '@/pinia/modules/user'
import { unset } from 'lodash'
const urlPath = ref(import.meta.env.VITE_BASE_API+"/")
const userStore = useUserStore()

const router=useRouter()
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
        manAndCar:"",
        job:"",
        transport:"",
        cab:"",
        remark: '',
        status: undefined,
        forbidden:undefined,
        })

// 验证规则
const rule = reactive({
              
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
const searchInfo = ref({"forbidden":1})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

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
	  // path.value = import.meta.env.VITE_BASE_API+'/'+url
	  // console.log("图片地址",path.value)
	  formData.value.manAndCar = url
	  
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

const uploadSuccessTwo = (res) => {
  fullscreenLoading.value = false
  console.log("图片",res)
  if (res.code === 0) {
	  var url = res.data.file.url
	  formData.value.job = url
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

const uploadSuccessThree = (res) => {
  fullscreenLoading.value = false
  console.log("图片",res)
  if (res.code === 0) {
	  var url = res.data.file.url
	  formData.value.transport = url
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

const uploadSuccessFour = (res) => {
  fullscreenLoading.value = false
  console.log("图片",res)
  if (res.code === 0) {
	  var url = res.data.file.url
	  formData.value.cab = url
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

//跳转到商户进件
const toCustInfo = (row) => {
	router.push({
		name:"appCustInfo",
		query:{
			id:row.ID
		}
	})
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

//查询已经禁用的用户
const onSubmitForbidden = () => {
  page.value = 1
  pageSize.value = 10
  searchInfo.value.forbidden = 2
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

//禁用用户
const forbiddenRow = async(row) => {

  // ElMessageBox.confirm('确定要执行吗?', '提示', {
  //       confirmButtonText: '确定',
  //       cancelButtonText: '取消',
  //       type: 'warning'
  //   }).then(() => {
        const res = await findAppMch({ ID: row.ID })
      console.log("结果",res.code)
      if (res.code === 0) {
          formData.value = res.data.reappMch
          if(res.data.reappMch.forbidden == 1){
            formData.value.forbidden = 2
            searchInfo.value.forbidden = 2
          }else{
            formData.value.forbidden = 1
          }
          console.log("修改的数据",formData.value)
          const r =  await updateAppMch(formData.value)
            if (r.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '更改成功'
                })
                getTableData()
              }
      }
    // })

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
    console.log("测试啊的",formData.value)
    // formData.value = {
    //     uid: 0,
    //     firmName: '',
    //     storeName: '',
    //     mainType: '',
    //     commission: 0,
    //     final: 0,
    //     acquirer: 0,
    //     legalPerson: '',
    //     phone: '',
    //     avatar: '',
    //     cardFront: '',
    //     cardReverse: '',
    //     entrust: '',
    //     manAndCar:"",
    //     job:"",
    //     transport:"",
    //     cab:"",
    //     remark: '',
    //     status: undefined,
    //     }
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
                  
                  console.log("提交的数据",formData.value)
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
