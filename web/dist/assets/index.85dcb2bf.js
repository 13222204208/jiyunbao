/*! 
 Build based on gin-vue-admin 
 Time : 1675935707000 */
import{j as a,r as e,b as s,o as l,c as o,e as t,w as d,z as n,d as i,h as r,i as u}from"./index.3b117381.js";const c=i("div",{class:"el-upload__text"},[r(" 拖拽或"),i("em",null,"点击上传")],-1),p=i("div",{class:"el-upload__tip"}," 请把安装包的zip拖拽至此处上传 ",-1),f={__name:"index",setup(i){const r=a(),f=e("/api"),m=a=>{if(0===a.code){let e="";a.data&&a.data.forEach(((a,s)=>{e+=`${s+1}.${a.msg}\n`})),alert(e)}else u.error(a.msg)};return(a,e)=>{const i=s("upload-filled"),u=s("el-icon"),_=s("el-upload");return l(),o("div",null,[t(_,{class:"upload-demo",drag:"",action:`${f.value}/autoCode/installPlugin`,headers:{"x-token":n(r).token},"show-file-list":!1,"on-success":m,"on-error":m,name:"plug"},{tip:d((()=>[p])),default:d((()=>[t(u,{class:"el-icon--upload"},{default:d((()=>[t(i)])),_:1}),c])),_:1},8,["action","headers"])])}}};export{f as default};