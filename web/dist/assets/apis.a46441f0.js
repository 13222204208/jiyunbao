/*! 
 Build based on gin-vue-admin 
 Time : 1677739392000 */
import{e as a}from"./api.cc9aa125.js";import{s as e,_ as t,r as s,J as o,b as l,o as r,c as n,d as c,e as d,w as i,h as p,i as u}from"./index.533d8bbb.js";const h={class:"clearfix sticky-button"},f={class:"tree-content"},y={name:"Apis"},m=t(Object.assign(y,{props:{row:{default:function(){return{}},type:Object}},setup(t,{expose:y}){const m=t,v=s({children:"children",label:"description"}),b=s(""),k=s([]),I=s([]),w=s("");(async()=>{const t=(await a()).data.apis;k.value=j(t);const s=await(o={authorityId:m.row.authorityId},e({url:"/casbin/getPolicyPathByAuthorityId",method:"post",data:o}));var o;w.value=m.row.authorityId,I.value=[],s.data.paths&&s.data.paths.forEach((a=>{I.value.push("p:"+a.path+"m:"+a.method)}))})();const g=s(!1),x=()=>{g.value=!0},j=a=>{const e={};a&&a.forEach((a=>{a.onlyId="p:"+a.path+"m:"+a.method,Object.prototype.hasOwnProperty.call(e,a.apiGroup)?e[a.apiGroup].push(a):Object.assign(e,{[a.apiGroup]:[a]})}));const t=[];for(const s in e){const a={ID:s,description:s+"组",children:e[s]};t.push(a)}return t},O=s(null),C=async()=>{const a=O.value.getCheckedNodes(!0);var t=[];a&&a.forEach((a=>{var e={path:a.path,method:a.method};t.push(e)}));var s;0===(await(s={authorityId:w.value,casbinInfos:t},e({url:"/casbin/updateCasbin",method:"post",data:s}))).code&&u({type:"success",message:"api设置成功"})};y({needConfirm:g,enterAndNext:()=>{C()}});const _=(a,e)=>!a||-1!==e.description.indexOf(a);return o(b,(a=>{O.value.filter(a)})),(a,e)=>{const t=l("el-input"),s=l("el-button"),o=l("el-tree");return r(),n("div",null,[c("div",h,[d(t,{modelValue:b.value,"onUpdate:modelValue":e[0]||(e[0]=a=>b.value=a),class:"fitler",placeholder:"筛选"},null,8,["modelValue"]),d(s,{class:"fl-right",size:"small",type:"primary",onClick:C},{default:i((()=>[p("确 定")])),_:1})]),c("div",f,[d(o,{ref_key:"apiTree",ref:O,data:k.value,"default-checked-keys":I.value,props:v.value,"default-expand-all":"","highlight-current":"","node-key":"onlyId","show-checkbox":"","filter-node-method":_,onCheck:x},null,8,["data","default-checked-keys","props"])])])}}}),[["__scopeId","data-v-c620bda6"]]);export{m as default};