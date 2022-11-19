(()=>{function f(t,n){let e;n?e=n.querySelectorAll(t):e=document.querySelectorAll(t);let o=[];return e.forEach(r=>{o.push(r)}),o}function g(t,n){let e=f(t,n);switch(e.length){case 0:return;case 1:return e[0];default:console.warn(`found [${e.length}] elements with selector [${t}], wanted zero or one`)}}function c(t,n){let e=g(t,n);return e||console.warn(`no element found for selector [${t}]`),e}function E(t,n){return typeof t=="string"&&(t=c(t)),t.innerHTML=n,t}function l(t,n,e="block"){return typeof t=="string"&&(t=c(t)),t.style.display=n?e:"none",t}function v(t,n){let e=document.createElement(t);for(let o in n)if(o&&n.hasOwnProperty(o)){let r=n[o];o==="dangerouslySetInnerHTML"?E(e,r.__html):r===!0?e.setAttribute(o,o):r!==!1&&r!==null&&r!==void 0&&e.setAttribute(o,r.toString())}for(let o=2;o<arguments.length;o++){let r=arguments[o];if(Array.isArray(r))r.forEach(s=>{if(r==null)throw`child array for tag [${t}] is ${r}
${e.outerHTML}`;if(s==null)throw`child for tag [${t}] is ${s}
${e.outerHTML}`;typeof s=="string"&&(s=document.createTextNode(s)),e.appendChild(s)});else{if(r==null)throw`child for tag [${t}] is ${r}
${e.outerHTML}`;r.nodeType||(r=document.createTextNode(r.toString())),e.appendChild(r)}}return e}function L(){for(let t of Array.from(document.querySelectorAll(".menu-container .final")))t.scrollIntoView({block:"nearest"})}var k="mode-light",h="mode-dark";function M(){for(let t of Array.from(document.getElementsByClassName("mode-input"))){let n=t;n.onclick=function(){switch(n.value){case"":document.body.classList.remove(k),document.body.classList.remove(h);break;case"light":document.body.classList.add(k),document.body.classList.remove(h);break;case"dark":document.body.classList.remove(k),document.body.classList.add(h);break}}}}function b(){let t=document.getElementById("flash-container");if(t===null)return;let n=t.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let e of n){let o=e;o.style.opacity="0",setTimeout(()=>o.remove(),500)}},3e3)}function x(){for(let t of Array.from(document.getElementsByClassName("link-confirm"))){let n=t;n.onclick=function(){let e=n.dataset.message;return e&&e.length===0&&(e="Are you sure?"),confirm(e)}}}function H(){document.addEventListener("keydown",t=>{t.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function y(t,n){return`<svg class="icon" style="width: ${n}px; height: ${n}px;"><use xlink:href="#svg-${t}"></use></svg>`}function w(){for(let t of f(".tag-editor")){let n=c("input.result",t),e=c(".tags",t),o=n.value.split(",").map(s=>s.trim()).filter(s=>s!=="");l(n,!1),e.innerHTML="";for(let s of o)e.appendChild(S(s,t));g(".add-item",t)?.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=y("plus",22),r.onclick=function(){X(e,t)},t.insertBefore(r,c(".clear",t))}}function W(t,n){return t.parentElement!==n.parentElement?null:t===n?0:t.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var p;function S(t,n){let e=document.createElement("div");e.className="item",e.draggable=!0,e.ondragstart=function(i){i.dataTransfer.setDragImage(document.createElement("div"),0,0),e.classList.add("dragging"),p=e},e.ondragover=function(i){let d=W(e,p);if(!d)return;let u=d===-1?e:e.nextSibling;p.parentElement.insertBefore(p,u),T(n)},e.ondrop=function(i){i.preventDefault()},e.ondragend=function(i){e.classList.remove("dragging"),i.preventDefault()};let o=document.createElement("div");o.innerText=t,o.className="value",o.onclick=function(){I(e)},e.appendChild(o);let r=document.createElement("input");r.className="editor",e.appendChild(r);let s=document.createElement("div");return s.innerHTML=y("times",13),s.className="close",s.onclick=function(){B(e)},e.appendChild(s),e}function B(t){let n=t.parentElement.parentElement;t.remove(),T(n)}function X(t,n){let e=S("",n);t.appendChild(e),I(e)}function I(t){let n=c(".value",t),e=c(".editor",t);e.value=n.innerText;let o=function(){if(e.value===""){t.remove();return}n.innerText=e.value,l(n,!0),l(e,!1);let r=t.parentElement.parentElement;T(r)};e.onblur=o,e.onkeydown=function(r){if(r.code==="Enter")return r.preventDefault(),o(),!1},l(n,!1),l(e,!0),e.focus()}function T(t){let n=[],e=f(".item .value",t);for(let r of e)n.push(r.innerText);let o=c("input.result",t);o.value=n.join(", ")}var A="--selected";function R(t){let n=t.parentElement.parentElement.querySelector("input");if(!n)throw"no associated input found";n.value="\u2205"}function N(){window.rituals.setSiblingToNull=R;let t={},n={};for(let e of Array.from(document.querySelectorAll("form.editor"))){let o=e,r=()=>{t={},n={};for(let s of o.elements){let i=s;if(i.name.length>0)if(i.name.endsWith(A))n[i.name]=i;else{(i.type!=="radio"||i.checked)&&(t[i.name]=i.value);let d=()=>{let u=n[i.name+A];u&&(u.checked=t[i.name]!==i.value)};i.onchange=d,i.onkeyup=d}}};o.onreset=r,r()}}var U=[];function C(){let t=document.querySelectorAll(".color-var");if(t.length>0)for(let n of Array.from(t)){let e=n,o=e.dataset.var,r=e.dataset.mode;U.push(o),!(!o||o.length===0)&&(e.oninput=function(){_(r,o,e.value)})}}function _(t,n,e){let o=document.querySelector("#mockup-"+t);if(!o){console.error("can't find mockup for mode ["+t+"]");return}switch(n){case"color-foreground":a(o,".mock-main",e);break;case"color-background":m(o,".mock-main",e);break;case"color-foreground-muted":a(o,".mock-main .mock-muted",e);break;case"color-background-muted":m(o,".mock-main .mock-muted",e);break;case"color-link-foreground":a(o,".mock-main .mock-link",e);break;case"color-link-visited-foreground":a(o,".mock-main .mock-link-visited",e);break;case"color-nav-foreground":a(o,".mock-nav",e),a(o,".mock-nav .mock-link",e);break;case"color-nav-background":m(o,".mock-nav",e);break;case"color-menu-foreground":a(o,".mock-menu",e),a(o,".mock-menu .mock-link",e);break;case"color-menu-background":m(o,".mock-menu",e);break;case"color-menu-selected-foreground":a(o,".mock-menu .mock-link-selected",e);break;case"color-menu-selected-background":m(o,".mock-menu .mock-link-selected",e);break;default:console.error("invalid key ["+n+"]")}}function $(t,n,e){let o=t.querySelectorAll(n);if(o.length==0)throw"empty query selector ["+n+"]";o.forEach(r=>{e(r)})}function m(t,n,e){$(t,n,o=>o.style.backgroundColor=e)}function a(t,n,e){$(t,n,o=>o.style.color=e)}function O(){window.rituals.Socket=q}var q=class{constructor(n,e,o,r,s){this.debug=n,this.open=e,this.recv=o,this.err=r,this.url=D(s),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){this.connectTime=Date.now(),this.sock=new WebSocket(D(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open()},this.sock.onmessage=e=>{let o=JSON.parse(e.data);n.debug&&console.debug("in",o),n.recv(o)},this.sock.onerror=e=>()=>{n.err("socket",e.type)},this.sock.onclose=()=>{n.connected=!1;let e=n.connectTime?Date.now()-n.connectTime:0;0<e&&e<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+e+"ms]"),setTimeout(()=>{n.connect()},n.pauseSeconds*500))}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw"not initialized";if(this.connected){let e=JSON.stringify(n,null,2);this.sock.send(e)}else this.pendingMessages.push(n)}};function D(t){if(t||(t="/connect"),t.indexOf("ws")==0)return t;let n=document.location,e="ws";return n.protocol==="https:"&&(e="wss"),t.indexOf("/")!=0&&(t="/"+t),e+`://${n.host}${t}`}function j(){console.log("[socket]: open")}function F(t){let n=document.getElementById("socket-list");if(n){let e=document.createElement("pre");e.innerText=JSON.stringify(t,null,2),n.append(e)}}function G(t){console.log("[socket error]: "+t)}function P(t,n,e,o){new window.rituals.Socket(!0,j,F,G,"/"+t+"/"+n.id+"/connect"),console.log("loaded ["+t+"] workspace with ["+e?.length+"] members and ["+o?.length+"] permissions")}function J(){window.rituals.initWorkspace=P}function z(){window.rituals={},window.JSX=v,L(),M(),b(),x(),H(),w(),N(),C(),O(),J()}document.addEventListener("DOMContentLoaded",z);})();
//# sourceMappingURL=client.js.map
