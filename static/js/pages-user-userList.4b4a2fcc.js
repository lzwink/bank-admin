(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-user-userList"],{"2b5e":function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={data:function(){return{userlList:[]}},onLoad:function(){var t=this;uni.request({url:"http://localhost:8080/GetUsersByGroupId",method:"GET",data:{groupId:getApp().globalData.groupId},success:function(e){0==e.data.code&&(t.userlList=e.data.data)},fail:function(){},complete:function(){}})},methods:{selectUser:function(t){getApp().globalData.uId=t,uni.navigateTo({url:"../user/userDetail"})}}};e.default=n},3751:function(t,e,a){e=t.exports=a("2350")(!1),e.push([t.i,".content-opp[data-v-ffd1708a]{\n\t/* 定义flex容器 */display:-webkit-box;display:-webkit-flex;display:-ms-flexbox;display:flex;\n\t/*设置容器内部容器的排列方向*/-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;-ms-flex-direction:column;flex-direction:column;-webkit-box-flex:1;-webkit-flex:1;-ms-flex:1;flex:1}.ul[data-v-ffd1708a]{font-size:%?30?%;color:#8f8f94;margin-top:%?50?%}.ul>uni-view[data-v-ffd1708a]{line-height:%?50?%;padding-left:%?10?%}.b-border[data-v-ffd1708a]{margin-top:%?20?%;margin-left:%?20?%;margin-right:%?20?%}.backcolor[data-v-ffd1708a]{background-color:#efeff4}.container[data-v-ffd1708a]{\n\t/* 定义flex容器 */display:-webkit-box;display:-webkit-flex;display:-ms-flexbox;display:flex;\n\t/*设置容器内部容器的排列方向*/-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;-ms-flex-direction:row;flex-direction:row;margin-top:%?20?%;margin-bottom:%?20?%}",""])},6977:function(t,e,a){var n=a("3751");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var i=a("4f06").default;i("7a523a50",n,!0,{sourceMap:!1,shadowMode:!1})},"782d":function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"content-opp backcolor"},t._l(t.userlList,function(e){return a("v-uni-view",[a("v-uni-button",{staticClass:"b-border",attrs:{type:"default"},on:{click:function(a){a=t.$handleEvent(a),t.selectUser(e["Id"])}}},[t._v(t._s(e.UserName)+"  "+t._s(e.RealName))])],1)}),1)},i=[];a.d(e,"a",function(){return n}),a.d(e,"b",function(){return i})},8768:function(t,e,a){"use strict";a.r(e);var n=a("782d"),i=a("c536");for(var o in i)"default"!==o&&function(t){a.d(e,t,function(){return i[t]})}(o);a("a999");var r=a("2877"),l=Object(r["a"])(i["default"],n["a"],n["b"],!1,null,"ffd1708a",null);e["default"]=l.exports},a999:function(t,e,a){"use strict";var n=a("6977"),i=a.n(n);i.a},c536:function(t,e,a){"use strict";a.r(e);var n=a("2b5e"),i=a.n(n);for(var o in n)"default"!==o&&function(t){a.d(e,t,function(){return n[t]})}(o);e["default"]=i.a}}]);