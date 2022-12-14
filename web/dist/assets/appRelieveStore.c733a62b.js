/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{g as e,f as a,d as l,c as t,u as o}from"./appRelieveStore.64335204.js";import{f as u,a as s,g as d}from"./format.215352f7.js";import{r as i,a as r,b as n,o as c,c as p,d as v,e as m,w as f,h as b,t as g,A as y,F as h,x as w,m as V,C as _,i as C}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const k={class:"gva-search-box"},z={class:"gva-table-box"},I={class:"gva-pagination"},j={class:"dialog-footer"},x={name:"AppRelieveStore"},S=Object.assign(x,{setup(x){const S=i([]),A=i({uid:0,storeId:0,contents:"",status:void 0}),U=r({contents:[{required:!0,message:"",trigger:["input","blur"]}]}),D=i(),R=i(1),T=i(0),B=i(10),F=i([]),M=i({}),q=()=>{M.value={}},O=()=>{R.value=1,B.value=10,H()},E=e=>{B.value=e,H()},G=e=>{R.value=e,H()},H=async()=>{const a=await e({page:R.value,pageSize:B.value,...M.value});0===a.code&&(F.value=a.data.list,T.value=a.data.total,R.value=a.data.page,B.value=a.data.pageSize)};H();(async()=>{S.value=await d("relieveStore")})();const J=i([]),K=e=>{J.value=e};i(!1);const L=i(""),N=async e=>{0===(await l({ID:e.ID})).code&&(C({type:"success",message:"删除成功"}),1===F.value.length&&R.value>1&&R.value--,H())},P=i(!1),Q=()=>{P.value=!1,A.value={uid:0,storeId:0,contents:"",status:void 0}},W=async()=>{var e;null==(e=D.value)||e.validate((async e=>{if(!e)return;let a;switch(L.value){case"create":default:a=await t(A.value);break;case"update":a=await o(A.value)}0===a.code&&(C({type:"success",message:"创建/更改成功"}),Q(),H())}))};return(e,l)=>{const t=n("el-date-picker"),o=n("el-form-item"),d=n("el-button"),i=n("el-form"),r=n("el-table-column"),C=n("el-table"),x=n("el-pagination"),H=n("el-input"),J=n("el-option"),X=n("el-select"),Y=n("el-dialog");return c(),p("div",null,[v("div",k,[m(i,{inline:!0,model:M.value,class:"demo-form-inline"},{default:f((()=>[m(o,{label:"创建时间"},{default:f((()=>[m(t,{modelValue:M.value.startCreatedAt,"onUpdate:modelValue":l[0]||(l[0]=e=>M.value.startCreatedAt=e),type:"datetime",placeholder:"开始时间"},null,8,["modelValue"]),b(" — "),m(t,{modelValue:M.value.endCreatedAt,"onUpdate:modelValue":l[1]||(l[1]=e=>M.value.endCreatedAt=e),type:"datetime",placeholder:"结束时间"},null,8,["modelValue"])])),_:1}),m(o,null,{default:f((()=>[m(d,{size:"small",type:"primary",icon:"search",onClick:O},{default:f((()=>[b("查询")])),_:1}),m(d,{size:"small",icon:"refresh",onClick:q},{default:f((()=>[b("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),v("div",z,[m(C,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:F.value,"row-key":"ID",onSelectionChange:K},{default:f((()=>[m(r,{type:"selection",width:"55"}),m(r,{align:"left",label:"日期",width:"180"},{default:f((e=>[b(g(y(u)(e.row.CreatedAt)),1)])),_:1}),m(r,{align:"left",label:"用户",prop:"uid",width:"120"}),m(r,{align:"left",label:"门店",prop:"storeId",width:"120"}),m(r,{align:"left",label:"解除原因",prop:"contents",width:"120"}),m(r,{align:"left",label:"状态",prop:"status",width:"120"},{default:f((e=>[b(g(y(s)(e.row.status,S.value)),1)])),_:1}),m(r,{align:"left",label:"按钮组"},{default:f((e=>[m(d,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:l=>(async e=>{const l=await a({ID:e.ID});L.value="update",0===l.code&&(A.value=l.data.reappRelieveStore,P.value=!0)})(e.row)},{default:f((()=>[b("变更")])),_:2},1032,["onClick"]),m(d,{type:"primary",link:"",icon:"delete",size:"small",onClick:a=>{return l=e.row,void _.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{N(l)}));var l}},{default:f((()=>[b("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),v("div",I,[m(x,{layout:"total, sizes, prev, pager, next, jumper","current-page":R.value,"page-size":B.value,"page-sizes":[10,30,50,100],total:T.value,onCurrentChange:G,onSizeChange:E},null,8,["current-page","page-size","total"])])]),m(Y,{modelValue:P.value,"onUpdate:modelValue":l[6]||(l[6]=e=>P.value=e),"before-close":Q,title:"弹窗操作"},{footer:f((()=>[v("div",j,[m(d,{size:"small",onClick:Q},{default:f((()=>[b("取 消")])),_:1}),m(d,{size:"small",type:"primary",onClick:W},{default:f((()=>[b("确 定")])),_:1})])])),default:f((()=>[m(i,{model:A.value,"label-position":"right",ref_key:"elFormRef",ref:D,rules:U,"label-width":"80px"},{default:f((()=>[m(o,{label:"用户:",prop:"uid"},{default:f((()=>[m(H,{modelValue:A.value.uid,"onUpdate:modelValue":l[2]||(l[2]=e=>A.value.uid=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(o,{label:"门店:",prop:"storeId"},{default:f((()=>[m(H,{modelValue:A.value.storeId,"onUpdate:modelValue":l[3]||(l[3]=e=>A.value.storeId=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(o,{label:"解除原因:",prop:"contents"},{default:f((()=>[m(H,{modelValue:A.value.contents,"onUpdate:modelValue":l[4]||(l[4]=e=>A.value.contents=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(o,{label:"状态:",prop:"status"},{default:f((()=>[m(X,{modelValue:A.value.status,"onUpdate:modelValue":l[5]||(l[5]=e=>A.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:!0},{default:f((()=>[(c(!0),p(h,null,w(S.value,((e,a)=>(c(),V(J,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{S as default};
