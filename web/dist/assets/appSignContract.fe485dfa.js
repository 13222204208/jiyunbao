/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{g as e,d as l,u as a,q as t,a as s,c as o}from"./appSignContract.c703b0b7.js";import{f as u,a as i,g as n}from"./format.215352f7.js";import{r as d,a as r,b as c,o as p,c as v,d as m,e as g,w as f,h as b,t as y,A as h,F as C,x as w,m as _,C as k,i as z}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const V={class:"gva-search-box"},I={class:"gva-table-box"},T={class:"gva-btn-list"},S=m("p",null,"确定要删除吗？",-1),x={style:{"text-align":"right","margin-top":"8px"}},j={class:"gva-pagination"},U={class:"dialog-footer"},M={name:"AppSignContract"},q=Object.assign(M,{setup(M){const q=d([]),A=d([]),D=d({reqId:"",certId:"",busOpenType:"",contractType:void 0,isSendConMsg:void 0,notifyUrl:"",custId:""}),O=r({contractType:[{required:!0,message:"",trigger:["input","blur"]}],isSendConMsg:[{required:!0,message:"",trigger:["input","blur"]}],custId:[{required:!0,message:"",trigger:["input","blur"]}]}),B=d(),F=d(1),R=d(0),E=d(10),G=d([]),H=d({}),J=d([]),K=()=>{H.value={}},L=()=>{F.value=1,E.value=10,Q()},N=e=>{E.value=e,Q()},P=e=>{F.value=e,Q()},Q=async()=>{const l=await e({page:F.value,pageSize:E.value,...H.value});0===l.code&&(G.value=l.data.list,R.value=l.data.total,F.value=l.data.page,E.value=l.data.pageSize)};Q();(async()=>{q.value=await n("contractType"),A.value=await n("isSendConMsg")})();const W=d([]),X=e=>{W.value=e},Y=d(!1),Z=async()=>{const e=[];if(0===W.value.length)return void z({type:"warning",message:"请选择要删除的数据"});W.value&&W.value.map((l=>{e.push(l.ID)}));0===(await l({ids:e})).code&&(z({type:"success",message:"删除成功"}),G.value.length===e.length&&F.value>1&&F.value--,Y.value=!1,Q())},$=d(""),ee=async e=>{0===(await s({ID:e.ID})).code&&(z({type:"success",message:"删除成功"}),1===G.value.length&&F.value>1&&F.value--,Q())},le=d(!1),ae=()=>{$.value="create",le.value=!0},te=()=>{le.value=!1,D.value={reqId:"",certId:"",busOpenType:"",contractType:void 0,isSendConMsg:void 0,notifyUrl:"",custId:""}},se=async()=>{var e;null==(e=B.value)||e.validate((async e=>{console.log("提交的数据",J.value);let l,t=J.value,s=[];for(var u=0;u<t.length;u++)console.log("数据： ",t[u]),"扫码工作日到账"==t[u]&&s.push("00"),"扫码实时到账"==t[u]&&(console.log("第二 ",t[u]),s.push("01")),"刷卡工作日到账"==t[u]&&s.push("10"),"刷卡实时到账"==t[u]&&s.push("11"),"D1到账"==t[u]&&s.push("20");if(s.length>1?D.value.busOpenType=s.join("|"):D.value.busOpenType=s.toString(),console.log("挚友",D.value),e){switch($.value){case"create":default:l=await o(D.value);break;case"update":l=await a(D.value)}0===l.code&&(z({type:"success",message:"创建/更改成功"}),te(),Q())}}))};return(e,l)=>{const s=c("el-date-picker"),o=c("el-form-item"),n=c("el-button"),d=c("el-form"),r=c("el-popover"),M=c("el-table-column"),Q=c("el-table"),$=c("el-pagination"),oe=c("el-checkbox"),ue=c("el-checkbox-group"),ie=c("el-option"),ne=c("el-select"),de=c("el-input"),re=c("el-dialog");return p(),v("div",null,[m("div",V,[g(d,{inline:!0,model:H.value,class:"demo-form-inline"},{default:f((()=>[g(o,{label:"创建时间"},{default:f((()=>[g(s,{modelValue:H.value.startCreatedAt,"onUpdate:modelValue":l[0]||(l[0]=e=>H.value.startCreatedAt=e),type:"datetime",placeholder:"开始时间"},null,8,["modelValue"]),b(" — "),g(s,{modelValue:H.value.endCreatedAt,"onUpdate:modelValue":l[1]||(l[1]=e=>H.value.endCreatedAt=e),type:"datetime",placeholder:"结束时间"},null,8,["modelValue"])])),_:1}),g(o,null,{default:f((()=>[g(n,{size:"small",type:"primary",icon:"search",onClick:L},{default:f((()=>[b("查询")])),_:1}),g(n,{size:"small",icon:"refresh",onClick:K},{default:f((()=>[b("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),m("div",I,[m("div",T,[g(n,{size:"small",type:"primary",icon:"plus",onClick:ae},{default:f((()=>[b("新增")])),_:1}),g(r,{visible:Y.value,"onUpdate:visible":l[4]||(l[4]=e=>Y.value=e),placement:"top",width:"160"},{reference:f((()=>[g(n,{icon:"delete",size:"small",style:{"margin-left":"10px"},disabled:!W.value.length,onClick:l[3]||(l[3]=e=>Y.value=!0)},{default:f((()=>[b("删除")])),_:1},8,["disabled"])])),default:f((()=>[S,m("div",x,[g(n,{size:"small",type:"primary",link:"",onClick:l[2]||(l[2]=e=>Y.value=!1)},{default:f((()=>[b("取消")])),_:1}),g(n,{size:"small",type:"primary",onClick:Z},{default:f((()=>[b("确定")])),_:1})])])),_:1},8,["visible"])]),g(Q,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:G.value,"row-key":"ID",onSelectionChange:X},{default:f((()=>[g(M,{align:"left",label:"日期",width:"180"},{default:f((e=>[b(y(h(u)(e.row.CreatedAt)),1)])),_:1}),g(M,{align:"left",label:"流水号",prop:"reqId",width:"180"}),g(M,{align:"left",label:"商户号",prop:"mercId",width:"180"}),g(M,{align:"left",label:"合同类型",prop:"contractType",width:"120"},{default:f((e=>[b(y(h(i)(e.row.contractType,q.value)),1)])),_:1}),g(M,{align:"left",label:"签约通知",prop:"isSendConMsg",width:"120"},{default:f((e=>[b(y(h(i)(e.row.isSendConMsg,A.value)),1)])),_:1}),g(M,{align:"left",label:"客户号",prop:"custId",width:"180"}),g(M,{align:"left",label:"按钮组"},{default:f((e=>[g(n,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:l=>(async e=>{const l=await a(e);0===l.code&&z({type:"success",message:l.msg})})(e.row)},{default:f((()=>[b("签约申请")])),_:2},1032,["onClick"]),g(n,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:l=>(async e=>{const l=await t(e);0===l.code&&z({type:"success",message:l.msg})})(e.row)},{default:f((()=>[b("签约状态")])),_:2},1032,["onClick"]),g(n,{type:"primary",link:"",icon:"delete",size:"small",onClick:l=>{return a=e.row,void k.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{ee(a)}));var a}},{default:f((()=>[b("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),m("div",j,[g($,{layout:"total, sizes, prev, pager, next, jumper","current-page":F.value,"page-size":E.value,"page-sizes":[10,30,50,100],total:R.value,onCurrentChange:P,onSizeChange:N},null,8,["current-page","page-size","total"])])]),g(re,{modelValue:le.value,"onUpdate:modelValue":l[9]||(l[9]=e=>le.value=e),"before-close":te,title:"弹窗操作"},{footer:f((()=>[m("div",U,[g(n,{size:"small",onClick:te},{default:f((()=>[b("取 消")])),_:1}),g(n,{size:"small",type:"primary",onClick:se},{default:f((()=>[b("确 定")])),_:1})])])),default:f((()=>[g(d,{model:D.value,"label-position":"right",ref_key:"elFormRef",ref:B,rules:O,"label-width":"100px"},{default:f((()=>[g(o,{label:"到账方式:",prop:"busOpenType"},{default:f((()=>[g(ue,{modelValue:J.value,"onUpdate:modelValue":l[5]||(l[5]=e=>J.value=e)},{default:f((()=>[g(oe,{label:"扫码工作日到账",size:"large"}),g(oe,{label:"扫码实时到账",size:"large"}),g(oe,{label:"刷卡工作日到账",size:"large"}),g(oe,{label:"刷卡实时到账",size:"large"}),g(oe,{label:"D1到账 ",size:"large"})])),_:1},8,["modelValue"])])),_:1}),g(o,{label:"合同类型:",prop:"contractType"},{default:f((()=>[g(ne,{modelValue:D.value.contractType,"onUpdate:modelValue":l[6]||(l[6]=e=>D.value.contractType=e),placeholder:"请选择",style:{width:"100%"},clearable:!0},{default:f((()=>[(p(!0),v(C,null,w(q.value,((e,l)=>(p(),_(ie,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),g(o,{label:"签约通知:",prop:"isSendConMsg"},{default:f((()=>[g(ne,{modelValue:D.value.isSendConMsg,"onUpdate:modelValue":l[7]||(l[7]=e=>D.value.isSendConMsg=e),placeholder:"请选择",style:{width:"100%"},clearable:!0},{default:f((()=>[(p(!0),v(C,null,w(A.value,((e,l)=>(p(),_(ie,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),g(o,{label:"客户号:",prop:"custId"},{default:f((()=>[g(de,{modelValue:D.value.custId,"onUpdate:modelValue":l[8]||(l[8]=e=>D.value.custId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{q as default};