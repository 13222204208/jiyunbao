/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{f as a}from"./sysDictionary.e268ed4c.js";import{D as t,r as e}from"./index.42b7f490.js";const i=t("dictionary",(()=>{const t=e({}),i=a=>{t.value={...t.value,...a}};return{dictionaryMap:t,setDictionaryMap:i,getDictionary:async e=>{if(t.value[e]&&t.value[e].length)return t.value[e];{const s=await a({type:e});if(0===s.code){const a={},r=[];return s.data.resysDictionary.sysDictionaryDetails&&s.data.resysDictionary.sysDictionaryDetails.forEach((a=>{r.push({label:a.label,value:a.value})})),a[s.data.resysDictionary.type]=r,i(a),t.value[e]}}}}}));export{i as u};
