/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{g as e,d as l,f as a,a as t,c as d,u as i}from"./appCategory.2eaec9b8.js";import{f as o}from"./format.215352f7.js";import{r as u,a as s,b as r,o as n,c as p,d as c,e as m,w as v,h as f,t as g,A as b,C as y,i as h}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const C={class:"gva-search-box"},_={class:"gva-table-box"},V={class:"gva-btn-list"},w=c("p",null,"确定要删除吗？",-1),k={style:{"text-align":"right","margin-top":"8px"}},I={class:"gva-pagination"},z={class:"dialog-footer"},x={name:"AppCategory"},D=Object.assign(x,{setup(x){const D=u({title:"",codeId:"",pid:0}),j=s({title:[{required:!0,message:"",trigger:["input","blur"]}],codeId:[{required:!0,message:"",trigger:["input","blur"]}]}),U=u(),A=u(1),S=u(0),T=u(10),q=u([]),B=u({}),F=()=>{B.value={}},M=()=>{A.value=1,T.value=10,E()},O=e=>{T.value=e,E()},R=e=>{A.value=e,E()},E=async()=>{const l=await e({page:A.value,pageSize:T.value,...B.value});0===l.code&&(q.value=l.data.list,S.value=l.data.total,A.value=l.data.page,T.value=l.data.pageSize)};E();(async()=>{})();const G=u([]),H=e=>{G.value=e},J=u(!1),K=async()=>{const e=[];if(0===G.value.length)return void h({type:"warning",message:"请选择要删除的数据"});G.value&&G.value.map((l=>{e.push(l.ID)}));0===(await l({ids:e})).code&&(h({type:"success",message:"删除成功"}),q.value.length===e.length&&A.value>1&&A.value--,J.value=!1,E())},L=u(""),N=async e=>{0===(await t({ID:e.ID})).code&&(h({type:"success",message:"删除成功"}),1===q.value.length&&A.value>1&&A.value--,E())},P=u(!1),Q=()=>{L.value="create",P.value=!0},W=()=>{P.value=!1,D.value={title:"",codeId:"",pid:0}},X=async()=>{var e;null==(e=U.value)||e.validate((async e=>{if(!e)return;let l;switch(L.value){case"create":default:l=await d(D.value);break;case"update":l=await i(D.value)}0===l.code&&(h({type:"success",message:"创建/更改成功"}),W(),E())}))};return(e,l)=>{const t=r("el-date-picker"),d=r("el-form-item"),i=r("el-input"),u=r("el-button"),s=r("el-form"),h=r("el-popover"),x=r("el-table-column"),E=r("el-table"),Y=r("el-pagination"),Z=r("el-dialog");return n(),p("div",null,[c("div",C,[m(s,{inline:!0,model:B.value,class:"demo-form-inline"},{default:v((()=>[m(d,{label:"创建时间"},{default:v((()=>[m(t,{modelValue:B.value.startCreatedAt,"onUpdate:modelValue":l[0]||(l[0]=e=>B.value.startCreatedAt=e),type:"datetime",placeholder:"开始时间"},null,8,["modelValue"]),f(" — "),m(t,{modelValue:B.value.endCreatedAt,"onUpdate:modelValue":l[1]||(l[1]=e=>B.value.endCreatedAt=e),type:"datetime",placeholder:"结束时间"},null,8,["modelValue"])])),_:1}),m(d,{label:"品类名称"},{default:v((()=>[m(i,{modelValue:B.value.title,"onUpdate:modelValue":l[2]||(l[2]=e=>B.value.title=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),m(d,{label:"类型ID"},{default:v((()=>[m(i,{modelValue:B.value.codeId,"onUpdate:modelValue":l[3]||(l[3]=e=>B.value.codeId=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),m(d,null,{default:v((()=>[m(u,{size:"small",type:"primary",icon:"search",onClick:M},{default:v((()=>[f("查询")])),_:1}),m(u,{size:"small",icon:"refresh",onClick:F},{default:v((()=>[f("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),c("div",_,[c("div",V,[m(u,{size:"small",type:"primary",icon:"plus",onClick:Q},{default:v((()=>[f("新增")])),_:1}),m(h,{visible:J.value,"onUpdate:visible":l[6]||(l[6]=e=>J.value=e),placement:"top",width:"160"},{reference:v((()=>[m(u,{icon:"delete",size:"small",style:{"margin-left":"10px"},disabled:!G.value.length,onClick:l[5]||(l[5]=e=>J.value=!0)},{default:v((()=>[f("删除")])),_:1},8,["disabled"])])),default:v((()=>[w,c("div",k,[m(u,{size:"small",type:"primary",link:"",onClick:l[4]||(l[4]=e=>J.value=!1)},{default:v((()=>[f("取消")])),_:1}),m(u,{size:"small",type:"primary",onClick:K},{default:v((()=>[f("确定")])),_:1})])])),_:1},8,["visible"])]),m(E,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:q.value,"row-key":"ID",onSelectionChange:H},{default:v((()=>[m(x,{type:"selection",width:"55"}),m(x,{align:"left",label:"日期",width:"180"},{default:v((e=>[f(g(b(o)(e.row.CreatedAt)),1)])),_:1}),m(x,{align:"left",label:"品类名称",prop:"title",width:"120"}),m(x,{align:"left",label:"类型ID",prop:"codeId",width:"120"}),m(x,{align:"left",label:"父id",prop:"pid",width:"120"}),m(x,{align:"left",label:"按钮组"},{default:v((e=>[m(u,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:l=>(async e=>{const l=await a({ID:e.ID});L.value="update",0===l.code&&(D.value=l.data.reappCategory,P.value=!0)})(e.row)},{default:v((()=>[f("变更")])),_:2},1032,["onClick"]),m(u,{type:"primary",link:"",icon:"delete",size:"small",onClick:l=>{return a=e.row,void y.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{N(a)}));var a}},{default:v((()=>[f("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),c("div",I,[m(Y,{layout:"total, sizes, prev, pager, next, jumper","current-page":A.value,"page-size":T.value,"page-sizes":[10,30,50,100],total:S.value,onCurrentChange:R,onSizeChange:O},null,8,["current-page","page-size","total"])])]),m(Z,{modelValue:P.value,"onUpdate:modelValue":l[10]||(l[10]=e=>P.value=e),"before-close":W,title:"弹窗操作"},{footer:v((()=>[c("div",z,[m(u,{size:"small",onClick:W},{default:v((()=>[f("取 消")])),_:1}),m(u,{size:"small",type:"primary",onClick:X},{default:v((()=>[f("确 定")])),_:1})])])),default:v((()=>[m(s,{model:D.value,"label-position":"right",ref_key:"elFormRef",ref:U,rules:j,"label-width":"180px"},{default:v((()=>[m(d,{label:"品类名称:",prop:"title"},{default:v((()=>[m(i,{modelValue:D.value.title,"onUpdate:modelValue":l[7]||(l[7]=e=>D.value.title=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(d,{label:"类型ID:",prop:"codeId"},{default:v((()=>[m(i,{modelValue:D.value.codeId,"onUpdate:modelValue":l[8]||(l[8]=e=>D.value.codeId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(d,{label:"父id:",prop:"pid"},{default:v((()=>[m(i,{modelValue:D.value.pid,"onUpdate:modelValue":l[9]||(l[9]=e=>D.value.pid=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{D as default};
