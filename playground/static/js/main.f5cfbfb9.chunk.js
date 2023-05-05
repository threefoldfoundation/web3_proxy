(this["webpackJsonp@open-rpc/playground"]=this["webpackJsonp@open-rpc/playground"]||[]).push([[1],{1350:function(e,t,n){},1382:function(e,t){},1390:function(e,t){},1392:function(e,t){},1417:function(e,t){},1534:function(e,t,n){"use strict";n.r(t);var a,r=n(0),o=n.n(r),c=n(64),i=n.n(c),s=(n(667),n(652)),l=n(125),u=n.n(l),p=n(181),m=n(110),d=n(41),f=n(16),h=Object(f.a)((function(e){return{list:{color:e.palette.secondary.main,height:"100%",listStyle:"none",overflow:"scroll"},wrapper:{maxHeight:"200px"}}}))((function(e){var t=e.classes,n=e.markers;return n&&0!==n.length?o.a.createElement("div",{className:t.wrapper},o.a.createElement("ul",{className:t.list},n&&n.map((function(e){return o.a.createElement("li",{key:e.message},e.startLineNumber,":",e.startColumn," - ",e.message)})))):null})),b=n(79),g=n(624),j=(n(1350),n(296)),O=n(297),v=n(299),y=n(298),E=n(1571),x=n(1580),k=n(539),w=n(522),S=n(1584),C=n(360),M=n(1581),T=n(1546),B=n(636),D=n.n(B),I=n(639),P=n.n(I),N=n(638),R=n.n(N),z=n(637),K=n.n(z),U=n(1576),L=Object(f.a)((function(e){return{title:{marginLeft:e.spacing(2)},appBar:{}}}))((function(e){var t=e.uiSchema,n=e.searchBarUrl,a=e.onChangeUrl;return o.a.createElement(U.a,{value:n,placeholder:t&&t.appBar["ui:inputPlaceholder"],style:{width:"100%"},onChange:function(e){a(e.target.value)}})})),A=n(523),V=n(625),W=n(1577),q=n(355),J=n.n(q),F=function(e){var t=e.examples,n=e.onChange,a=r.useState(null),o=Object(d.a)(a,2),c=o[0],i=o[1];return t?r.createElement(r.Fragment,null,r.createElement(M.a,{title:"Example OpenRPC Documents"},r.createElement(A.a,{onClick:function(e){i(e.currentTarget)},variant:"outlined",endIcon:r.createElement(J.a,null),style:{height:"38px",fontSize:"11px",marginLeft:"10px"}},"examples")),r.createElement(V.a,{id:"simple-menu",anchorEl:c,keepMounted:!0,open:Boolean(c),onClose:function(){i(null)}},r.createElement("div",{style:{maxWidth:"525px"}},t.map((function(e){return r.createElement(W.a,{onClick:function(t){return function(e){i(null),n&&n(e)}(e)}},r.createElement(k.a,{container:!0,spacing:0},r.createElement(k.a,{item:!0,xs:12},r.createElement(w.a,{variant:"subtitle1"},e.name)),r.createElement(k.a,{item:!0,xs:12},r.createElement(w.a,{variant:"caption",style:{fontSize:"9px"}},e.url))))}))))):null},H=n(1575),Y=n(1573),_=n(635),G=n.n(_),Q=function(e){var t=e.selectedTransport,n=e.transports,a=e.onChange,c=e.style,i=e.onAddTransport,s=Object(r.useState)(null),l=Object(d.a)(s,2),u=l[0],p=l[1],m=Object(r.useState)(!1),f=Object(d.a)(m,2),h=f[0],b=f[1],g=Object(r.useState)(),j=Object(d.a)(g,2),O=j[0],v=j[1],y=Object(r.useState)(),E=Object(d.a)(y,2),x=E[0],S=E[1],C=Object(r.useState)(),M=Object(d.a)(C,2),T=M[0],B=M[1],D=Object(r.useState)(null),I=Object(d.a)(D,2),P=I[0],N=I[1];return o.a.createElement("div",{style:c},o.a.createElement(H.a,{onClose:function(){return b(!1)},"aria-labelledby":"simple-dialog-title",open:h},o.a.createElement(Y.a,{maxWidth:"sm"},o.a.createElement(k.a,{container:!0,justify:"space-between",alignItems:"center",style:{padding:"30px",paddingTop:"10px",paddingBottom:"10px",marginTop:"10px"}},o.a.createElement(w.a,{variant:"h6"},"Custom Transport Plugin"),o.a.createElement(w.a,{variant:"caption",gutterBottom:!0},'Transport plugins are created by implementing the "connect", "sendData", and "close" methods over JSON-RPC.'),o.a.createElement(k.a,{container:!0,direction:"column",spacing:1},o.a.createElement(k.a,{item:!0},o.a.createElement(U.a,{placeholder:"Plugin Name",onChange:function(e){S(e.target.value)},style:{display:"block",background:"rgba(0,0,0,0.1)",borderRadius:"4px",padding:"0px 10px",marginRight:"5px"}})),o.a.createElement(k.a,{item:!0},o.a.createElement(U.a,{placeholder:"Plugin URI",onChange:function(e){B(e.target.value)},style:{display:"block",background:"rgba(0,0,0,0.1)",borderRadius:"4px",padding:"0px 10px",marginRight:"5px"}})),o.a.createElement(k.a,{item:!0},o.a.createElement(A.a,{variant:"outlined",onClick:function(e){N(e.currentTarget)}},O?O.name:"Select A Transport"))),o.a.createElement(V.a,{id:"transport-menu",anchorEl:P,keepMounted:!0,open:Boolean(P),onClose:function(){N(null)}},n&&n.filter((function(e){return"plugin"!==e.type})).map((function(e,t){return o.a.createElement(W.a,{onClick:function(){return function(e){N(null),v(e)}(e)}},e.name)}))),o.a.createElement(A.a,{style:{marginTop:"10px",marginBottom:"10px"},onClick:function(){(N(null),O&&x&&T)&&(i&&i({type:"plugin",transport:O,name:x,uri:T}),b(!1))},disabled:!x||!T||!O,variant:"contained"},"Add Transport")))),o.a.createElement(A.a,{style:{marginRight:"10px",marginLeft:"5px"},variant:"outlined",onClick:function(e){p(e.currentTarget)},endIcon:o.a.createElement(J.a,null)},t&&t.name),o.a.createElement(V.a,{id:"transport-menu",anchorEl:u,keepMounted:!0,open:Boolean(u),onClose:function(){p(null)}},n&&n.map((function(e,t){return o.a.createElement(W.a,{onClick:function(){return function(e){p(null),a&&a(e)}(e)}},e.name)})),o.a.createElement(W.a,{onClick:function(){return b(!0)}},o.a.createElement(G.a,{style:{marginRight:"5px"}}),o.a.createElement(w.a,{variant:"caption"},"Add Transport"))))},$=function(e){Object(v.a)(n,e);var t=Object(y.a)(n);function n(){return Object(j.a)(this,n),t.apply(this,arguments)}return Object(O.a)(n,[{key:"render",value:function(){var e=this.props,t=e.uiSchema,n=e.classes,a=e.onSplitViewChange,r=e.onDarkModeChange,c=e.examples,i=e.onExampleDocumentsDropdownChange;return o.a.createElement(E.a,{position:"fixed",color:"default",elevation:0,className:n.appBar},o.a.createElement(x.a,null,o.a.createElement(k.a,{alignItems:"center",container:!0},o.a.createElement(k.a,{item:!0,xs:6,sm:6,md:2,direction:"row",container:!0},this.props.uiSchema&&this.props.uiSchema.appBar&&this.props.uiSchema.appBar["ui:logoUrl"]&&o.a.createElement(k.a,null,o.a.createElement("img",{alt:"playground-title",height:"30",src:this.props.uiSchema.appBar["ui:logoUrl"]})),o.a.createElement(k.a,{style:{overflow:"hidden"}},o.a.createElement(w.a,{className:n.title,variant:"h6"},t&&t.appBar["ui:title"]))),o.a.createElement(S.a,{smDown:!0},o.a.createElement(k.a,{item:!0,container:!0,justify:"center",alignItems:"center",sm:8},o.a.createElement(k.a,{item:!0},o.a.createElement(Q,{transports:this.props.transportList,onAddTransport:this.props.onTransportAdd,selectedTransport:this.props.selectedTransport,onChange:this.props.onTransportChange,style:{marginLeft:"10px"}})),o.a.createElement(k.a,{item:!0,sm:6},this.props.uiSchema&&this.props.uiSchema.appBar&&this.props.uiSchema.appBar["ui:input"]&&o.a.createElement(C.a,{style:{background:"rgba(0, 0, 0, 0.1)",padding:"0px 10px 0px 10px",width:"100%"},elevation:0},o.a.createElement(L,{searchBarUrl:this.props.searchBarUrl,onChangeUrl:this.props.onChangeUrl,uiSchema:t}))),this.props.uiSchema&&this.props.uiSchema.appBar&&this.props.uiSchema.appBar["ui:examplesDropdown"]&&o.a.createElement(F,{examples:c,onChange:i}))),o.a.createElement(k.a,{item:!0,xs:6,sm:6,md:2,container:!0,justify:"flex-end",alignItems:"center"},t&&t.appBar["ui:splitView"]?o.a.createElement(M.a,{title:"Full Screen"},o.a.createElement(T.a,{onClick:function(){a&&a(!1)}},o.a.createElement(D.a,null))):o.a.createElement(M.a,{title:"Edit"},o.a.createElement(T.a,{onClick:function(){a&&a(!0)}},o.a.createElement(K.a,null))),o.a.createElement(M.a,{title:"Toggle Dark Theme"},o.a.createElement(T.a,null,t&&t.appBar["ui:darkMode"]?o.a.createElement(R.a,{onClick:function(){return r(!1)}}):o.a.createElement(P.a,{onClick:function(){return r(!0)}})))))))}}]),n}(r.Component),X=Object(f.a)((function(e){return{title:{marginLeft:e.spacing(2)},appBar:{}}}))($),Z=n(653),ee=n(1556),te=n(1578),ne=n(446),ae=n.n(ne),re=n(146),oe=n.n(re),ce=n(640),ie=n.n(ce),se=n(645),le=n.n(se),ue=n(642),pe=n.n(ue),me=n(447),de=n.n(me),fe=n(641),he=n.n(fe),be=n(643),ge=n.n(be),je=n(644),Oe=n.n(je),ve={success:ie.a,warning:he.a,error:pe.a,info:de.a,debug:de.a};!function(e){e.error="error",e.warn="warning",e.info="info",e.success="success",e.debug="debug"}(a||(a={}));var ye=function(e){Object(v.a)(n,e);var t=Object(y.a)(n);function n(){return Object(j.a)(this,n),t.apply(this,arguments)}return Object(O.a)(n,[{key:"render",value:function(){var e=this.props,t=e.classes,n=e.className,a=e.message,r=e.onClose,c=e.variant,i=Object(Z.a)(e,["classes","className","message","onClose","variant"]),s=ve[c];return o.a.createElement(ee.a,Object.assign({className:ae()(t[c],n),"aria-describedby":"client-snackbar",message:o.a.createElement("span",{id:"client-snackbar",className:t.message},o.a.createElement(s,{className:ae()(t.icon,t.iconVariant)}),a),action:[o.a.createElement(T.a,{key:"close","aria-label":"Close",color:"inherit",className:t.close,onClick:r},o.a.createElement(le.a,{className:t.icon}))]},i))}}]),n}(o.a.Component),Ee=Object(f.a)((function(e){return{success:{color:ge.a[600]},error:{backgroundColor:e.palette.error.dark},info:{color:e.palette.primary.dark},debug:{color:e.palette.secondary.dark},warning:{color:Oe.a[700]},icon:{fontSize:20},iconVariant:{opacity:.9,marginRight:e.spacing(2)},message:{display:"flex",alignItems:"center"},close:{padding:e.spacing(1)},margin:{margin:e.spacing(2)}}}))(ye),xe=function(e){Object(v.a)(n,e);var t=Object(y.a)(n);function n(){return Object(j.a)(this,n),t.apply(this,arguments)}return Object(O.a)(n,[{key:"render",value:function(){var e=this.props,t=e.classes,n=e.notification,a=e.close;return Object(re.isEmpty)(n)?null:o.a.createElement(te.a,{open:!0,autoHideDuration:1e4,anchorOrigin:{vertical:"bottom",horizontal:"left"}},o.a.createElement(Ee,{onClose:a,variant:n.type,message:o.a.createElement("span",null,n.message),className:t.margin}))}}]),n}(r.Component),ke=Object(f.a)((function(e){return{title:{marginLeft:e.spacing(2)},close:{padding:e.spacing(1)},margin:{margin:e.spacing(2)}}}))(xe),we=n(1570),Se=n(328),Ce=n(358),Me=n.n(Ce),Te=Object(Se.a)({props:{MuiAppBar:{position:"sticky"},MuiCard:{elevation:0}},overrides:{MuiPaper:{root:{zIndex:1,opacity:1}},MuiToolbar:{root:{background:"transparent !important"}},MuiAppBar:{root:{backgroundColor:"transparent !important"},colorDefault:{background:"transparent !important"},colorPrimary:{background:"transparent !important"}}},palette:{background:{default:"#fff"}}}),Be=Object(Se.a)({props:{MuiAppBar:{position:"sticky"}},palette:{type:"dark",background:{default:Me.a[900],paper:Me.a[900]}},overrides:{MuiPaper:{root:{zIndex:1,opacity:1}},MuiAppBar:{root:{background:"transparent !important"},colorPrimary:{background:"transparent !important"}},MuiToolbar:{root:{background:"transparent !important"}},MuiTable:{root:{background:"transparent !important"}},MuiTypography:{root:{color:Me.a[400]}}}}),De=n(213),Ie=n.n(De),Pe=n(651),Ne=n.n(Pe),Re=n(1574),ze=n(1582),Ke=n(1579),Ue=n(444),Le=function(e){if(!1===e.split&&e.onlyRenderSplit)return o.a.createElement("div",{key:2,style:e.splitLeft?e.leftStyle:e.rightStyle},e.splitLeft?e.left:e.right);var t=e.direction||"vertical",n=e.split?"horizontal"===t?.35*window.innerHeight:window.innerWidth/2:"horizontal"===t?window.innerHeight:window.innerWidth;return o.a.createElement(Ue.default,{split:t,style:e.style,className:"playground-splitview",minSize:100,maxSize:0,defaultSize:n,size:n,onChange:function(t){e.onChange&&e.onChange(t)}},o.a.createElement("div",{style:e.leftStyle?Object(m.a)(Object(m.a)({},e.leftStyle),{display:"flex",flexDirection:"column",height:"100%"}):{display:"flex",flexDirection:"column",height:"100%"},key:1},e.left),o.a.createElement("div",{key:2,style:e.rightStyle},e.right))},Ae=n(646),Ve=n.n(Ae),We=function(e){var t=Object(r.useState)(),n=Object(d.a)(t,2),a=n[0],o=n[1],c=function(e){var t;try{t=JSON.parse(e)}catch(n){}t&&Ve.a.dereference(t).then((function(t){o(t),oe.a.defer((function(){return window.localStorage.setItem("schema",e)}))}))};return Object(r.useEffect)((function(){e&&c(e)}),[]),[a,c]};var qe=function(){var e=Object(r.useState)((function(){return window.localStorage.getItem("schema")})),t=Object(d.a)(e,2),n=t[0],a=t[1];return[n,function(e){oe.a.defer((function(){return window.localStorage.setItem("schema",e)})),a(e)}]},Je=n(169),Fe=function(e){var t=Object(r.useState)(e),n=Object(d.a)(t,2);return[n[0],n[1]]},He=n(647),Ye=function(){var e=Object(r.useState)(He.parse(window.location.search,{ignoreQueryPrefix:!0,depth:100,decoder:function(e){return/^(\d+|\d*\.\d+)$/.test(e)?parseFloat(e):"false"!==e&&("true"===e||decodeURIComponent(e))}}));return[Object(d.a)(e,1)[0]]},_e=Object(Je.createStore)((function(){return Ye()})),Ge=Object(Je.createStore)((function(){var e=_e(),t=Object(d.a)(e,1)[0];return Fe(t.schemaUrl||t.url)})),Qe=Object(Je.createStore)((function(){var e=Object(r.useState)(),t=Object(d.a)(e,2);return[t[0],t[1]]})),$e=function(e){var t,n,a=Ge(),r=Object(d.a)(a,1)[0],c=Qe(),i=Object(d.a)(c,2)[1],s=e.openrpcMethodObject,l=[];if(s&&s.examples&&s.examples[0]){t=s.examples[0];var u=s.paramStructure||"either";n="by-name"===u?t.params.reduce((function(e,t){return e[t.name]=t.value,e}),{}):t.params.map((function(e){return e.value}))}return s&&"by-name"===s.paramStructure&&(l={}),o.a.createElement(M.a,{title:"Open in Inspector"},o.a.createElement(A.a,{variant:"contained",onClick:function(){return i({url:r,openrpcMethodObject:e.openrpcMethodObject,request:{jsonrpc:"2.0",method:s.name,params:n||l,id:0}})}},o.a.createElement("span",{role:"img","aria-label":"try-it-inspector"},"\ud83d\udd75\ufe0f\u200d\u2642\ufe0f"),"\ufe0f\ufe0f Try It Now"))},Xe=n(285),Ze=function(e){var t=Object(r.useState)(e),n=Object(d.a)(t,2),a=n[0],o=n[1];return[a,function(e){var t=e.section,n=e.key,r=e.value;o(Object(m.a)(Object(m.a)({},a),{},Object(Xe.a)({},t,Object(m.a)(Object(m.a)({},a.appBar),{},Object(Xe.a)({},n,r)))))}]},et=Object(Je.createStore)((function(){var e,t,n=_e(),a=Object(d.a)(n,1)[0];return Ze((e={appBar:{"ui:input":!0,"ui:inputPlaceholder":"Enter OpenRPC Document Url or rpc.discover Endpoint","ui:logoUrl":"https://github.com/open-rpc/design/raw/master/icons/open-rpc-logo-noText/open-rpc-logo-noText%20(PNG)/128x128.png","ui:splitView":!0,"ui:darkMode":!1,"ui:title":"Playground","ui:examplesDropdown":!0},methods:{"ui:defaultExpanded":!1,"ui:methodPlugins":!0},params:{"ui:defaultExpanded":!1}},t=a.uiSchema,e&&t?{appBar:Object(m.a)(Object(m.a)({},e.appBar),t.appBar||{}),methods:Object(m.a)(Object(m.a)({},e.methods),t.methods||{}),params:Object(m.a)(Object(m.a)({},e.params),t.params||{})}:e||t))})),tt=[{name:"api-with-examples",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/api-with-examples-openrpc.json"},{name:"link-example",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/link-example-openrpc.json"},{name:"params-by-name-petstore",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/params-by-name-petstore-openrpc.json"},{name:"petstore-expanded",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/petstore-expanded-openrpc.json"},{name:"petstore",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/petstore-openrpc.json"},{name:"simple-math",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/simple-math-openrpc.json"},{name:"empty-document",url:"https://raw.githubusercontent.com/open-rpc/examples/master/service-descriptions/empty-openrpc.json"}],nt=n(234),at=n.n(nt),rt=n(327),ot=n.n(rt),ct=n(279),it=n(356),st=n.n(it),lt=function(e){var t,n=Object(r.useRef)(),a=ot()();Object(r.useEffect)((function(){void 0!==n&&void 0!==n.current&&n.current.layout()}),[a]),Object(r.useEffect)((function(){return function(){t&&t.dispose()}}),[]);return o.a.createElement(at.a,{height:"100%",editorOptions:{useShadows:!1,minimap:{enabled:!1},scrollBeyondLastLine:!1,lineNumbers:"on"},value:e.value,editorDidMount:function(a,r){n.current=r;var o=b.Uri.parse("inmemory://openrpc-playground.json");t=b.editor.createModel(e.value||"","json",o),r.setModel(t),Object(ct.addDiagnostics)(o.toString(),st.a,b),e.editorDidMount&&e.editorDidMount(oe.a,r),e.onMarkerChange&&r.onDidChangeModelDecorations(oe.a.debounce((function(){if(e.onMarkerChange){var t=b.editor.getModelMarkers({resource:o});e.onMarkerChange(t)}}),300))},language:"json",onChange:function(t,n){e.onChange&&e.onChange(n)}})},ut=function(e){var t=Object(r.useState)(st.a),n=Object(d.a)(t,2),a=n[0],o=n[1];return Object(r.useEffect)((function(){e&&e.addAction({id:"replace-meta-schema",label:"Replace Meta Schema",keybindings:[b.KeyMod.chord(b.KeyMod.CtrlCmd|b.KeyCode.KEY_K,b.KeyMod.CtrlCmd|b.KeyCode.KEY_R)],contextMenuGroupId:"navigation",contextMenuOrder:1.5,run:function(){var e=window.prompt("Paste schema to replace current meta schema","{}");e&&o(JSON.parse(e))}})}),[e]),[a]},pt=n(595).initVimMode,mt=function(e){var t=Object(r.useState)(),n=Object(d.a)(t,2),a=n[0],o=n[1];return Object(r.useEffect)((function(){if(e)return e.addAction({id:"vim-mode",label:"Vim Mode",keybindings:[b.KeyMod.chord(b.KeyMod.CtrlCmd|b.KeyCode.KEY_K,b.KeyMod.CtrlCmd|b.KeyCode.KEY_V)],contextMenuGroupId:"navigation",contextMenuOrder:1.5,run:function(){a&&a.dispose(),o(pt(e))}}),function(){a&&a.dispose()}}),[e]),[e,a]},dt=n(648),ft=n.n(dt),ht=n(259),bt=n(189),gt=[{type:"http",name:"HTTP",schema:{type:"object",properties:{headers:{patternProperties:{"":{type:"string"}}},credentials:{type:"string",enum:["omit","same-origin","include"]}},examples:[{headers:{}}]}},{type:"websocket",name:"WebSocket"},{type:"postmessagewindow",name:"PostMessageWindow"},{type:"postmessageiframe",name:"PostMessageIframe"}],jt=function(){var e=Object(p.a)(u.a.mark((function e(t,n,a,r){var o,c,i,s;return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if("websocket"!==(null===(c=n.find((function(e){return e.type===a.type})))||void 0===c?void 0:c.type)){e.next=5;break}o=new ht.WebSocketTransport(t),e.next=23;break;case 5:if("http"!==(null===c||void 0===c?void 0:c.type)){e.next=9;break}o=new ht.HTTPTransport(t,r),e.next=23;break;case 9:if("postmessageiframe"!==(null===c||void 0===c?void 0:c.type)){e.next=13;break}o=new ht.PostMessageIframeTransport(t),e.next=23;break;case 13:if("postmessagewindow"!==(null===c||void 0===c?void 0:c.type)){e.next=17;break}o=new ht.PostMessageWindowTransport(t),e.next=23;break;case 17:if("plugin"!==(null===c||void 0===c?void 0:c.type)){e.next=23;break}return e.next=20,jt(c.uri,n,c.transport);case 20:i=e.sent,s=Object.assign({},bt.Transport,{connect:function(){return Object(p.a)(u.a.mark((function e(){var n;return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,i.connect();case 2:return e.next=4,i.sendData({internalID:0,request:{jsonrpc:"2.0",method:"connect",params:[t],id:0}});case 4:return n=e.sent,e.abrupt("return",n);case 6:case"end":return e.stop()}}),e)})))()},sendData:function(e){return i.sendData({internalID:0,request:{jsonrpc:"2.0",method:"sendData",params:[e.request],id:0}})},close:function(){return i.sendData({internalID:0,request:{jsonrpc:"2.0",method:"close",params:[],id:0}})}}),o=new s;case 23:return e.abrupt("return",o);case 24:case"end":return e.stop()}}),e)})));return function(t,n,a,r){return e.apply(this,arguments)}}(),Ot=function(e,t,n,a){var o=Object(r.useState)(),c=Object(d.a)(o,2),i=c[0],s=c[1],l=Object(r.useState)(!1),m=Object(d.a)(l,2),f=m[0],h=m[1],b=Object(r.useState)(n),g=Object(d.a)(b,2),j=g[0],O=g[1],v=Object(r.useState)(),y=Object(d.a)(v,2),E=y[0],x=y[1];return Object(r.useEffect)((function(){""!==t&&void 0!==t?j&&function(){var n=Object(p.a)(u.a.mark((function n(){var r;return u.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return n.next=2,jt(t,e,j,a);case 2:return r=n.sent,n.abrupt("return",r.connect().then((function(){h(!0),s(r)})));case 4:case"end":return n.stop()}}),n)})));return function(){return n.apply(this,arguments)}}()().catch((function(e){h(!1),s(void 0),x(e)})):s(void 0)}),[j,t,e,a]),[i,j,function(){var e=Object(p.a)(u.a.mark((function e(t){return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:h(!1),O(t);case 2:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),E,f]},vt=function(){var e=Object(p.a)(u.a.mark((function e(t){var n;return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch(t);case 2:if(404!==(n=e.sent).status){e.next=5;break}throw new Error("404: Not Found");case 5:return e.abrupt("return",n.text());case 6:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),yt=n(650),Et=function(){var e=qe(),t=Object(d.a)(e,2),n=t[0],c=t[1],i=Object(r.useState)([]),l=Object(d.a)(i,2),f=l[0],j=l[1],O=Ge(),v=Object(d.a)(O,2),y=v[0],E=v[1],x=Object(yt.a)(y,1e3),k=Object(d.a)(x,1)[0],S=Object(r.useState)(),C=Object(d.a)(S,2),B=C[0],D=C[1],I=Object(r.useState)(),P=Object(d.a)(I,2),N=P[0],R=P[1],z=Object(r.useState)(),K=Object(d.a)(z,2),U=K[0],L=K[1],A=et(),V=Object(d.a)(A,2),W=V[0],q=V[1],J=Object(r.useState)(),F=Object(d.a)(J,2),H=F[0],_=F[1],G=Object(r.useState)(!1),Q=Object(d.a)(G,2),$=Q[0],Z=Q[1],ee=We(n?JSON.parse(n):null),te=Object(d.a)(ee,2),ne=te[0],ae=te[1],re=_e(),oe=Object(d.a)(re,1)[0],ce=function(e){H&&setTimeout((function(){H.layout()}),0),Z(e)},ie=Qe(),se=Object(d.a)(ie,1)[0];ut(H),mt(H);Object(r.useEffect)((function(){b.editor.setTheme(W.appBar["ui:darkMode"]?"vs-dark":"vs")}),[]),Object(r.useEffect)((function(){var e=tt.find((function(e){return"petstore"===e.name}));n||y||!e||E(e.url)}),[n]),Object(r.useEffect)((function(){me(Object(m.a)(Object(m.a)({},pe),{},{theme:W.appBar["ui:darkMode"]?"summerfruit":"summerfruit:inverted"}))}),[W.appBar["ui:darkMode"]]),Object(r.useEffect)((function(){B&&H&&H.setValue(B),B&&ae(B)}),[B]),Object(r.useEffect)((function(){N&&L({type:a.error,message:N})}),[N]),Object(r.useEffect)((function(){ae(n||"")}),[n]);var le=Object(r.useState)({theme:"summerfruit:inverted",collapseStringsAfterLength:25,displayDataTypes:!1,displayObjectSize:!1,indentWidth:2,name:!1}),ue=Object(d.a)(le,2),pe=ue[0],me=ue[1],de=Object(r.useState)(gt),fe=Object(d.a)(de,2),he=fe[0],be=fe[1],ge=W.appBar["ui:darkMode"]?Be:Te,je=Ot(he,k,oe.transport&&he.find((function(e){return e.type===oe.transport}))||he[0]),Oe=Object(d.a)(je,3),ve=Oe[0],ye=Oe[1],Ee=Oe[2],xe=function(){var e=Object(p.a)(u.a.mark((function e(){var t,n,a;return u.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!k||!k.includes(".json")){e.next=6;break}return e.next=3,vt(k);case 3:return t=e.sent,c(t),e.abrupt("return",D(t));case 6:return e.prev=6,e.next=9,null===ve||void 0===ve?void 0:ve.sendData({internalID:999999,request:{jsonrpc:"2.0",params:[],id:999999,method:"rpc.discover"}});case 9:n=e.sent,(a=JSON.stringify(n,null,2))&&(c(a),D(a)),e.next=17;break;case 14:e.prev=14,e.t0=e.catch(6),R(e.t0.message);case 17:case"end":return e.stop()}}),e,null,[[6,14]])})));return function(){return e.apply(this,arguments)}}();return Object(r.useEffect)((function(){k&&ve&&xe()}),[k,ve]),Object(r.useEffect)((function(){se&&ce(!0)}),[se]),o.a.createElement(we.a,{theme:ge},o.a.createElement(Re.a,null),o.a.createElement(X,{searchBarUrl:y,uiSchema:W,examples:tt,onExampleDocumentsDropdownChange:function(e){return E(e.url)},selectedTransport:ye,transportList:he,onTransportChange:function(e){return Ee(e)},onTransportAdd:function(e){be((function(t){return[].concat(Object(s.a)(t),[e])}))},onSplitViewChange:function(e){q({value:e,key:"ui:splitView",section:"appBar"})},onDarkModeChange:function(e){b.editor.setTheme(e?"vs-dark":"vs"),q({value:e,key:"ui:darkMode",section:"appBar"})},onChangeUrl:E}),o.a.createElement(Le,{direction:"horizontal",splitLeft:!0,split:$,leftStyle:{width:"100%",height:"100%"},right:o.a.createElement(o.a.Fragment,null,o.a.createElement(ft.a,{hideToggleTheme:!0,url:k&&k.includes(".json")?null:k,transport:"plugin"!==ye.type?ye.type:void 0,request:se&&se.request,openrpcDocument:ne})),onChange:function(){return H&&H.layout()},left:o.a.createElement(Le,{onlyRenderSplit:!0,split:W.appBar["ui:splitView"],leftStyle:{paddingTop:"58px",height:"94%",width:"100%"},rightStyle:{height:"94%",width:"100%",overflowY:"auto",marginTop:"58px",paddingBottom:"58px"},onChange:function(){return H&&H.layout()},left:o.a.createElement(o.a.Fragment,null,o.a.createElement(h,{markers:f}),o.a.createElement(lt,{editorDidMount:function(e,t){_(t)},onMarkerChange:function(e){j(e)},onChange:function(e){ae(e)},value:n||""})),right:o.a.createElement(o.a.Fragment,null,o.a.createElement(Y.a,null,o.a.createElement(g.Documentation,{schema:ne,uiSchema:W,reactJsonOptions:pe,methodPlugins:W.methods["ui:methodPlugins"]?[$e]:void 0})),o.a.createElement(ze.a,{variant:"scrollable",indicatorColor:"primary",value:0,style:{position:"absolute",bottom:"0",right:"25px",zIndex:1,marginBottom:"0px"}},o.a.createElement(Ke.a,{onClick:function(){return ce(!$)},style:{background:ge.palette.background.default,width:"165px",paddingRight:"30px",border:"1px solid ".concat(ge.palette.text.hint)},label:o.a.createElement("div",null,o.a.createElement(w.a,{variant:"body1"},o.a.createElement("span",{role:"img","aria-label":"inspector"},"\ud83d\udd75\ufe0f\u200d\u2642\ufe0f"),"\ufe0f Inspector"),o.a.createElement(M.a,{title:"Toggle Inspector"},o.a.createElement(T.a,{style:{position:"absolute",right:"5px",top:"20%"},size:"small"},$?o.a.createElement(Ie.a,null):o.a.createElement(Ne.a,null))))})))})}),o.a.createElement(ke,{close:function(){return L({})},notification:U}))};i.a.render(o.a.createElement(Je.ReusableProvider,null,o.a.createElement(Et,null)),document.getElementById("root"))},662:function(e,t,n){e.exports=n(1534)},667:function(e,t,n){}},[[662,2,3]]]);
//# sourceMappingURL=main.f5cfbfb9.chunk.js.map