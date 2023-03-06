/*! 
 Build based on gin-vue-admin 
 Time : 1677739392000 */
import{g as e,f as l,d as a,c as t,u as o}from"./appStore.fc215848.js";import{f as u,a as d,g as r}from"./format.7121bf90.js";import{r as s,a as n,b as i,o as p,c,d as m,e as v,w as f,h as b,t as g,z as V,F as h,x as y,m as _,A as w,i as A}from"./index.533d8bbb.js";import"./date.34b11046.js";import"./dictionary.f9434c9f.js";import"./dictionary.a38b2399.js";import"./sysDictionary.dfe03bab.js";const U={class:"gva-search-box"},C={class:"gva-table-box"},k={class:"gva-pagination"},z={class:"dialog-footer"},j={name:"AppStore"},x=Object.assign(j,{setup(j){const x=s([]),N=s({uid:0,storeName:"",storeAvatar:"",category:"",storePhone:"",storeAddress:"",longitude:"",latitude:"",notice:"",close:"",service:"",agreement:"",status:void 0}),P=n({}),S=s(),D=s(1),I=s(0),T=s(10),B=s([]),F=s({}),M=()=>{F.value={}},O=()=>{D.value=1,T.value=10,E()},R=e=>{T.value=e,E()},q=e=>{D.value=e,E()},E=async()=>{const l=await e({page:D.value,pageSize:T.value,...F.value});0===l.code&&(B.value=l.data.list,I.value=l.data.total,D.value=l.data.page,T.value=l.data.pageSize)};E();(async()=>{x.value=await r("storeState")})();const G=s([]),H=e=>{G.value=e};s(!1);const J=s(""),K=async e=>{0===(await a({ID:e.ID})).code&&(A({type:"success",message:"删除成功"}),1===B.value.length&&D.value>1&&D.value--,E())},L=s(!1),Q=()=>{L.value=!1,N.value={uid:0,storeName:"",storeAvatar:"",category:"",storePhone:"",storeAddress:"",longitude:"",latitude:"",notice:"",close:"",service:"",agreement:"",status:void 0}},W=async()=>{var e;null==(e=S.value)||e.validate((async e=>{if(!e)return;let l;switch(J.value){case"create":default:l=await t(N.value);break;case"update":l=await o(N.value)}0===l.code&&(A({type:"success",message:"创建/更改成功"}),Q(),E())}))};return(e,a)=>{const t=i("el-date-picker"),o=i("el-form-item"),r=i("el-input"),s=i("el-button"),n=i("el-form"),A=i("el-table-column"),j=i("el-table"),E=i("el-pagination"),G=i("el-option"),X=i("el-select"),Y=i("el-dialog");return p(),c("div",null,[m("div",U,[v(n,{inline:!0,model:F.value,class:"demo-form-inline"},{default:f((()=>[v(o,{label:"创建时间"},{default:f((()=>[v(t,{modelValue:F.value.startCreatedAt,"onUpdate:modelValue":a[0]||(a[0]=e=>F.value.startCreatedAt=e),type:"datetime",placeholder:"开始时间"},null,8,["modelValue"]),b(" — "),v(t,{modelValue:F.value.endCreatedAt,"onUpdate:modelValue":a[1]||(a[1]=e=>F.value.endCreatedAt=e),type:"datetime",placeholder:"结束时间"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店名称"},{default:f((()=>[v(r,{modelValue:F.value.storeName,"onUpdate:modelValue":a[2]||(a[2]=e=>F.value.storeName=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店电话"},{default:f((()=>[v(r,{modelValue:F.value.storePhone,"onUpdate:modelValue":a[3]||(a[3]=e=>F.value.storePhone=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),v(o,null,{default:f((()=>[v(s,{size:"small",type:"primary",icon:"search",onClick:O},{default:f((()=>[b("查询")])),_:1}),v(s,{size:"small",icon:"refresh",onClick:M},{default:f((()=>[b("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),m("div",C,[v(j,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:B.value,"row-key":"ID",onSelectionChange:H},{default:f((()=>[v(A,{align:"left",label:"日期",width:"180"},{default:f((e=>[b(g(V(u)(e.row.CreatedAt)),1)])),_:1}),v(A,{align:"left",label:"门店用户",prop:"uid",width:"120"}),v(A,{align:"left",label:"门店名称",prop:"storeName",width:"120"}),v(A,{align:"left",label:"主营品类",prop:"category",width:"120"}),v(A,{align:"left",label:"门店电话",prop:"storePhone",width:"120"}),v(A,{align:"left",label:"门店地址",prop:"storeAddress",width:"120"}),v(A,{align:"left",label:"门店公告",prop:"notice",width:"120"}),v(A,{align:"left",label:"结算模式",prop:"close",width:"120"}),v(A,{align:"left",label:"服务设置",prop:"service",width:"120"}),v(A,{align:"left",label:"合同协议中心",prop:"agreement",width:"120"}),v(A,{align:"left",label:"合作状态",prop:"status",width:"120"},{default:f((e=>[b(g(V(d)(e.row.status,x.value)),1)])),_:1}),v(A,{align:"left",label:"按钮组"},{default:f((e=>[v(s,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:a=>(async e=>{const a=await l({ID:e.ID});J.value="update",0===a.code&&(N.value=a.data.reappStore,L.value=!0)})(e.row)},{default:f((()=>[b("变更")])),_:2},1032,["onClick"]),v(s,{type:"primary",link:"",icon:"delete",size:"small",onClick:l=>{return a=e.row,void w.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{K(a)}));var a}},{default:f((()=>[b("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),m("div",k,[v(E,{layout:"total, sizes, prev, pager, next, jumper","current-page":D.value,"page-size":T.value,"page-sizes":[10,30,50,100],total:I.value,onCurrentChange:q,onSizeChange:R},null,8,["current-page","page-size","total"])])]),v(Y,{modelValue:L.value,"onUpdate:modelValue":a[17]||(a[17]=e=>L.value=e),"before-close":Q,title:"弹窗操作"},{footer:f((()=>[m("div",z,[v(s,{size:"small",onClick:Q},{default:f((()=>[b("取 消")])),_:1}),v(s,{size:"small",type:"primary",onClick:W},{default:f((()=>[b("确 定")])),_:1})])])),default:f((()=>[v(n,{model:N.value,"label-position":"right",ref_key:"elFormRef",ref:S,rules:P,"label-width":"80px"},{default:f((()=>[v(o,{label:"门店用户:",prop:"uid"},{default:f((()=>[v(r,{modelValue:N.value.uid,"onUpdate:modelValue":a[4]||(a[4]=e=>N.value.uid=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店名称:",prop:"storeName"},{default:f((()=>[v(r,{modelValue:N.value.storeName,"onUpdate:modelValue":a[5]||(a[5]=e=>N.value.storeName=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店头像:",prop:"storeAvatar"},{default:f((()=>[v(r,{modelValue:N.value.storeAvatar,"onUpdate:modelValue":a[6]||(a[6]=e=>N.value.storeAvatar=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"主营品类:",prop:"category"},{default:f((()=>[v(r,{modelValue:N.value.category,"onUpdate:modelValue":a[7]||(a[7]=e=>N.value.category=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店电话:",prop:"storePhone"},{default:f((()=>[v(r,{modelValue:N.value.storePhone,"onUpdate:modelValue":a[8]||(a[8]=e=>N.value.storePhone=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店地址:",prop:"storeAddress"},{default:f((()=>[v(r,{modelValue:N.value.storeAddress,"onUpdate:modelValue":a[9]||(a[9]=e=>N.value.storeAddress=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"经度:",prop:"longitude"},{default:f((()=>[v(r,{modelValue:N.value.longitude,"onUpdate:modelValue":a[10]||(a[10]=e=>N.value.longitude=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"纬度:",prop:"latitude"},{default:f((()=>[v(r,{modelValue:N.value.latitude,"onUpdate:modelValue":a[11]||(a[11]=e=>N.value.latitude=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"门店公告:",prop:"notice"},{default:f((()=>[v(r,{modelValue:N.value.notice,"onUpdate:modelValue":a[12]||(a[12]=e=>N.value.notice=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"结算模式:",prop:"close"},{default:f((()=>[v(r,{modelValue:N.value.close,"onUpdate:modelValue":a[13]||(a[13]=e=>N.value.close=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"服务设置:",prop:"service"},{default:f((()=>[v(r,{modelValue:N.value.service,"onUpdate:modelValue":a[14]||(a[14]=e=>N.value.service=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"合同协议中心:",prop:"agreement"},{default:f((()=>[v(r,{modelValue:N.value.agreement,"onUpdate:modelValue":a[15]||(a[15]=e=>N.value.agreement=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),v(o,{label:"合作状态:",prop:"status"},{default:f((()=>[v(X,{modelValue:N.value.status,"onUpdate:modelValue":a[16]||(a[16]=e=>N.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:!0},{default:f((()=>[(p(!0),c(h,null,y(x.value,((e,l)=>(p(),_(G,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{x as default};