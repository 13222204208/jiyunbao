/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{s as e,i as a,_ as l,r as t,j as o,b as s,o as n,c as i,d,e as p,w as c,h as r,A as u,t as m,W as f}from"./index.42b7f490.js";const x=(e,l)=>{if(void 0!==e.data){if("application/json"===e.data.type){const l=new FileReader;l.onload=function(){const e=JSON.parse(l.result).msg;a({showClose:!0,message:e,type:"error"})},l.readAsText(new Blob([e.data]))}}else{var t=window.URL.createObjectURL(new Blob([e])),o=document.createElement("a");o.style.display="none",o.href=t,o.download=l;var s=new MouseEvent("click");o.dispatchEvent(s)}},w=()=>e({url:"/excel/loadExcel",method:"get"}),b={class:"upload"},h={class:"gva-table-box"},v={class:"gva-btn-list"},g=l(Object.assign({name:"Excel"},{setup(a){const l=t("/api"),g=t(1),y=t(0),E=t(999),_=t([]),k=async(e=(()=>{}))=>{const a=await e({page:g.value,pageSize:E.value});0===a.code&&(_.value=a.data.list,y.value=a.data.total,g.value=a.data.page,E.value=a.data.pageSize)};k(f);const j=o(),z=a=>{a&&"string"==typeof a||(a="ExcelExport.xlsx"),((a,l)=>{e({url:"/excel/exportExcel",method:"post",data:{fileName:l,infoList:a},responseType:"blob"}).then((e=>{x(e,l)}))})(_.value,a)},I=()=>{k(w)},T=()=>{var a;e({url:"/excel/downloadTemplate",method:"get",params:{fileName:a="ExcelTemplate.xlsx"},responseType:"blob"}).then((e=>{x(e,a)}))};return(e,a)=>{const t=s("el-button"),o=s("el-upload"),f=s("el-table-column"),x=s("el-table");return n(),i("div",b,[d("div",h,[d("div",v,[p(o,{class:"excel-btn",action:`${l.value}/excel/importExcel`,headers:{"x-token":u(j).token},"on-success":I,"show-file-list":!1},{default:c((()=>[p(t,{size:"small",type:"primary",icon:"upload"},{default:c((()=>[r("导入")])),_:1})])),_:1},8,["action","headers"]),p(t,{class:"excel-btn",size:"small",type:"primary",icon:"download",onClick:a[0]||(a[0]=e=>z("ExcelExport.xlsx"))},{default:c((()=>[r("导出")])),_:1}),p(t,{class:"excel-btn",size:"small",type:"success",icon:"download",onClick:a[1]||(a[1]=e=>T())},{default:c((()=>[r("下载模板")])),_:1})]),p(x,{data:_.value,"row-key":"ID"},{default:c((()=>[p(f,{align:"left",label:"ID","min-width":"100",prop:"ID"}),p(f,{align:"left","show-overflow-tooltip":"",label:"路由Name","min-width":"160",prop:"name"}),p(f,{align:"left","show-overflow-tooltip":"",label:"路由Path","min-width":"160",prop:"path"}),p(f,{align:"left",label:"是否隐藏","min-width":"100",prop:"hidden"},{default:c((e=>[d("span",null,m(e.row.hidden?"隐藏":"显示"),1)])),_:1}),p(f,{align:"left",label:"父节点","min-width":"90",prop:"parentId"}),p(f,{align:"left",label:"排序","min-width":"70",prop:"sort"}),p(f,{align:"left",label:"文件路径","min-width":"360",prop:"component"})])),_:1},8,["data"])])])}}}),[["__scopeId","data-v-d29f2fdc"]]);export{g as default};
