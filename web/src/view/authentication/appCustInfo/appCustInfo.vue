<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="流水号">
         <el-input v-model="searchInfo.reqId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商户名称">
         <el-input v-model="searchInfo.mercName" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="商户类型" prop="mercType">
            <el-select v-model="searchInfo.mercType" clearable placeholder="请选择" @clear="()=>{searchInfo.mercType=undefined}">
              <el-option v-for="(item,key) in mercTypeOptions" :key="key" :label="item.label" :value="item.value" />
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
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="商户ID" prop="uid" width="80" />
        <el-table-column align="left" label="客户号" prop="custId" width="160" />
        <el-table-column align="left" label="流水号" prop="reqId" width="180" />
        <el-table-column align="left" label="发起方商户号" prop="certId" width="120" />
        <el-table-column align="left" label="代理商编号" prop="agtMercId" width="120" />
        <el-table-column align="left" label="商户名称" prop="mercName" width="120" />
        <el-table-column align="left" label="商户简称" prop="mercShortName" width="120" />
        <el-table-column align="left" label="商户类型" prop="mercType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.mercType,mercTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="mcc码" prop="mccCd" width="120" />
        <el-table-column align="left" label="联系人邮箱" prop="contactMail" width="120" />
<!--        <el-table-column align="left" label="联系人" prop="contactMan" width="120" />
        <el-table-column align="left" label="联系人电话" prop="contactPhone" width="120" /> -->
        <el-table-column align="left" label="客户经理" prop="cusMgrNm" width="120" />
        <el-table-column align="left" label="异步通知地址" prop="notifyUrl" width="120" />
        <el-table-column align="left" label="法人证件号" prop="crpCertNo" width="120" />
        <el-table-column align="left" label="法人证件类型" prop="crpCertType" width="120" />
        <el-table-column align="left" label="证件开始日期" prop="certBgn" width="120" />
        <el-table-column align="left" label="证件有效期" prop="certExpire" width="120" />
        <el-table-column align="left" label="法人姓名" prop="crpNm" width="120" />
        <el-table-column align="left" label="法人手机号" prop="crpPhone" width="120" />
        <el-table-column align="left" label="结算账号" prop="stlAccNo" width="120" />
        <el-table-column align="left" label="开户支行联行号" prop="bankSubCode" width="120" />
        <el-table-column align="left" label="结算账户类型" prop="stlAccType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.stlAccType,stlAccTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="营业归属省代码" prop="busProviceCode" width="120" />
        <el-table-column align="left" label="营业归属市代码" prop="busCityCode" width="120" />
        <el-table-column align="left" label="营业详细地址" prop="busAddr" width="120" />
        <el-table-column fixed="right" align="left" label="操作" width="220" >
            <template #default="scope">
				      <el-button type="primary" link icon="edit" size="small" class="table-button" @click="addCustInfoApplyFunc(scope.row)">资料上送</el-button>
					  
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateAppCustInfoFunc(scope.row)">资料修改</el-button>
			
			            <el-button type="primary" link icon="edit" size="small" class="table-button"  @click="toUploadImg(scope.row)">图片上送</el-button>

                  <el-button type="primary" link icon="edit" size="small" class="table-button"  @click="auditCustInfoApplyFunc(scope.row)">资料确认</el-button>

                  <el-button type="primary" link icon="edit" size="small" class="table-button"  @click="queryCustApplyFunc(scope.row)">状态查询</el-button>
						
                  <el-button type="primary" link icon="edit" size="small" class="table-button"  @click="changeMercBaseInfoFunc(scope.row)">变更申请</el-button>

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

    
    <el-dialog
    v-model="dialogFormBankCode"
    title="查询行号"
    width="30%"
  >
   
  <el-form :model="formBankData" label-position="right" ref="elFormRef" :rules="bankRule" label-width="100px">

    <el-form-item label="地区编码:"  prop="cityCode	" >
          <el-input v-model="formBankData.cityCode" :clearable="true"  placeholder="请输入地区编码比如 1000" >
          <template #append>
            <el-button :icon="Search"><router-link target="_blank" :to="{path:'/layout/set/appAreaInfo'}">编码查询</router-link></el-button>
          </template>
          </el-input>
        </el-form-item>
        <el-form-item label="银行名称:"  prop="openBankName" >
          <el-input v-model="formBankData.openBankName" :clearable="true"  placeholder="请输入银行名称" />
        </el-form-item>


    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormBankCode = false">取消</el-button>
        <el-button type="primary" @click="getBankCode">
          查询
        </el-button>
      </span>
    </template>
  </el-dialog>



    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="150px">
        <!-- <el-form-item label="商户ID:"  prop="uid" >
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
<!--        <el-form-item label="流水号:"  prop="reqId" >
          <el-input v-model="formData.reqId" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
<!--        <el-form-item label="发起方商户号:"  prop="certId" >
          <el-input v-model="formData.certId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="代理商编号:"  prop="agtMercId" >
          <el-input v-model="formData.agtMercId" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="商户名称:"  prop="mercName" >
          <el-input v-model="formData.mercName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户简称:"  prop="mercShortName" >
          <el-input v-model="formData.mercShortName" :clearable="true"  placeholder="若商户名称长度超过20,商户简称必传" />
        </el-form-item>
        <el-form-item label="商户类型:"  prop="mercType" >
          <el-select v-model="formData.mercType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in mercTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="mcc码:"  prop="mccCd" >
          <el-input v-model="formData.mccCd" :clearable="true"  placeholder="请输入" width="500" >  
            <template #append>
            <el-button :icon="Search"  @click="getMcc" >MCC码查询</el-button>
          </template> 
        </el-input>
        </el-form-item>
        <el-form-item label="联系人邮箱:"  prop="contactMail" >
          <el-input v-model="formData.contactMail" :clearable="true"  placeholder="请输入" />
        </el-form-item>
<!--        <el-form-item label="联系人:"  prop="contactMan" >
          <el-input v-model="formData.contactMan" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人电话:"  prop="contactPhone" >
          <el-input v-model="formData.contactPhone" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="客户经理:"  prop="cusMgrNm" >
          <el-input v-model="formData.cusMgrNm" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="异步通知地址:"  prop="notifyUrl" >
          <el-input v-model="formData.notifyUrl" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="法人证件号:"  prop="crpCertNo" >
          <el-input v-model="formData.crpCertNo" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="开户证件号:"  prop="openCertNo" >
          <el-input v-model="formData.openCertNo" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="法人证件类型:"  prop="crpCertType" >
          <el-input v-model="formData.crpCertType" :clearable="true"  placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="证件开始日期:"  prop="certBgn" >
          <el-input v-model="formData.certBgn" :clearable="true"  placeholder="证件开始日期yyyyMMdd" />
        </el-form-item>
        <el-form-item label="证件有效期:"  prop="certExpire" >
          <el-input v-model="formData.certExpire" :clearable="true"  placeholder="日期格式yyyyMMdd ，如果为长期或者永久，请填值“29991231”" />
        </el-form-item>
        <el-form-item label="法人姓名:"  prop="crpNm" >
          <el-input v-model="formData.crpNm" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人手机号:"  prop="crpPhone" >
          <el-input v-model="formData.crpPhone" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算账号:"  prop="stlAccNo" >
          <el-input v-model="formData.stlAccNo" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户支行联行号:"  prop="bankSubCode" >
          <el-input v-model="formData.bankSubCode" :clearable="true"  placeholder="请输入" >
            <template #append>
            <el-button :icon="Search"  text @click="dialogFormBankCode = true" >行号查询</el-button>
          </template> 
          </el-input>
        </el-form-item>
        <el-form-item label="结算账户类型:"  prop="stlAccType" >
          <el-select v-model="formData.stlAccType" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in stlAccTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="营业归属省代码:"  prop="busProviceCode" >
          <el-input v-model="formData.busProviceCode" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业归属市代码:"  prop="busCityCode" >
          <el-input v-model="formData.busCityCode" :clearable="true"  placeholder="请输入" />
        </el-form-item>

        <el-form-item label="营业归属区(县)代码:"  prop="busAreaCode" >
          <el-input v-model="formData.busAreaCode" :clearable="true"  placeholder="城市下有区(县必填)，城市下无区(县)则不上送" />
        </el-form-item>

        <el-form-item label="营业详细地址:"  prop="busAddr" >
          <el-input v-model="formData.busAddr" :clearable="true"  placeholder="请输入" />
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
  name: 'AppCustInfo'
}
</script>

<script setup>
import {
  createAppCustInfo,
  deleteAppCustInfo,
  deleteAppCustInfoByIds,
  updateAppCustInfo,
  findAppCustInfo,
  getAppCustInfoList,
  addCustInfoApply,
  auditCustInfoApply,//资料确认
  queryCustApply,//商户状态查询
  changeMercBaseInfo,//基本信息变更申请
} from '@/api/appCustInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { Search } from '@element-plus/icons-vue'
import {useRouter} from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import {  inject } from 'vue-demi'

const reload = inject('reload')

const router=useRouter()
// 自动化生成的字典（可能为空）以及字段
const mercTypeOptions = ref([])
const stlAccTypeOptions = ref([])
const formBankData = ref({
      cityCode:'',
      openBankName:'',
})

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
        openCertNo: '',
        crpCertType: '00',
        certBgn: '',
        certExpire: '',
        crpNm: '',
        crpPhone: '',
        stlAccNo: '',
        bankSubCode: '',
        stlAccType: undefined,
        busProviceCode: '',
        busCityCode: '',
        busAreaCode: '',
        busAddr: '',
        })

        //验证行号查询数据
        const bankRule = reactive({
          cityCode:[{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
          openBankName:[{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],        
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
  const table = await getAppCustInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    mercTypeOptions.value = await getDictFunc('mercType')
    stlAccTypeOptions.value = await getDictFunc('stlAccType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

const dialogFormBankCode = ref(false)
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
            deleteAppCustInfoFunc(row)
        })
    }
//跳转到图片上送页面
const toUploadImg = (row) => {
	console.log("图片上送",row)
	router.push({
		name:"appUploadImg",
		query:{
			id:row.ID
		}
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
      const res = await deleteAppCustInfoByIds({ ids })
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

// 资料上送
const addCustInfoApplyFunc = async(row) => {
  let res = await addCustInfoApply(row)
  console.log("结果确认",res)
  if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: res.msg
                })
          }
          reload()
}


///基本信息变更申请
const changeMercBaseInfoFunc = async(row) => {
  //付值 客户号
  let res = await changeMercBaseInfo({"custId":row.custId})
  console.log("结果确认",res)
  if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: res.msg
                })
          }
}
//资料确认
const auditCustInfoApplyFunc = async(row) => {
        //资料确认
        let res = await auditCustInfoApply(row)
        console.log("结果确认",res)
        if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
          }
          reload()
}

//商户状态查询 
const queryCustApplyFunc = async(row) => {
        //资料确认
       const result = await queryCustApply(row)
        console.log("结果确认",result.code)
        if (result.code === 0) {
                ElMessage({
                  type: 'success',
                  message: result.msg
                })
              }
    
}


// 更新行
const updateAppCustInfoFunc = async(row) => {
    const res = await findAppCustInfo({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reappCustInfo
        dialogFormVisible.value = true
    }
}

//获取行号
const getBankCode = () => {

  if(formBankData.value.cityCode == ""){
    ElMessage({
          type: 'warning',
          message: '请填写地区编码'
        })
        return
  }

  if(formBankData.value.openBankName == ""){
    ElMessage({
          type: 'warning',
          message: '请填写银行名称'
        })
        return
  }
  window.open(import.meta.env.VITE_BASE_PATH+"/sdk/request/findBankNameAndBankCode.php?cityCode="+formBankData.value.cityCode+"&openBankName="+formBankData.value.openBankName,"_blank")
}

//获取mcc码
const getMcc = () => {
  let mercType = formData.value.mercType
  console.log("类型值 ",mercType)
  //商户类型,1-个体工商户、 2-企业、 3-小微，
  if(mercType == 3){
    mercType = 1
  }else if(mercType == 4){
    mercType = 2
  }else{
    mercType = 3
  }

  window.open(import.meta.env.VITE_BASE_PATH+"/sdk/request/queryMccList.php?merc_type="+mercType,"_blank")
}


// 删除行
const deleteAppCustInfoFunc = async (row) => {
    const res = await deleteAppCustInfo({ ID: row.ID })
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
        openCertNo: '',
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
        busAreaCode: '',
        busAddr: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>