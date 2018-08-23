package views

const blog = `<!doctype html>
<html>
<head><meta charset="utf-8"><style>#right-panel {
    background-color: #fff;
}

#right-panel .cover-top {
    background: linear-gradient(to bottom, #fff 50%, transparent);
}

#cover-bottom #cover-bottom-background-right {
    background: #fff;
}

@font-face {
  font-family: octicons-link;
  src: url(data:font/woff;charset=utf-8;base64,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) format('woff');
}

#container {
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  color: #373737;
  font-family: "Roboto", "Noto Sans", "Ubuntu", "Helvetica Neue", Helvetica, "Segoe UI", Arial, sans-serif, "Noto Sans CJK SC", "Source Han Sans SC", "Microsoft Yahei";
  font-size: 14px;
  line-height: 2;
  word-wrap: break-word;
  background-color: #fff;
}

#container a {
  background-color: transparent;
  -webkit-text-decoration-skip: objects;
}

#container a:active,
#container a:hover {
  outline-width: 0;
}

#container strong {
  font-weight: inherit;
}

#container strong {
  font-weight: bolder;
}

#container h1 {
  font-size: 2em;
  margin: 0.67em 0;
}

#container img {
  border-style: none;
}

#container svg:not(:root) {
  overflow: hidden;
}

#container code,
#container kbd,
#container pre {
  font-family: "Roboto Mono", "Ubuntu Mono", "Menlo", "Consolas", monospace;
  font-size: 1em;
}

#container hr {
  box-sizing: content-box;
  height: 0;
  overflow: visible;
}

#container input {
  font: inherit;
  margin: 0;
}

#container input {
  overflow: visible;
}

#container button:-moz-focusring,
#container [type="button"]:-moz-focusring,
#container [type="reset"]:-moz-focusring,
#container [type="submit"]:-moz-focusring {
  outline: 1px dotted ButtonText;
}

#container [type="checkbox"] {
  box-sizing: border-box;
  padding: 0;
}

#container table {
  border-spacing: 0;
  border-collapse: collapse;
}

#container td,
#container th {
  padding: 0;
}

#container * {
  box-sizing: border-box;
}

#container input {
  font: 13px/1.4 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
}

#container a {
  color: #4078c0;
  text-decoration: none;
}

#container a:hover,
#container a:active {
  text-decoration: underline;
}

#container hr {
  height: 0;
  margin: 15px 0;
  overflow: hidden;
  background: transparent;
  border: 0;
  border-bottom: 1px solid #ddd;
}

#container hr::before {
  display: table;
  content: "";
}

#container hr::after {
  display: table;
  clear: both;
  content: "";
}

#container h1,
#container h2,
#container h3,
#container h4,
#container h5,
#container h6 {
  margin-top: 0;
  margin-bottom: 0;
  line-height: 1.5;
}

#container h1 {
  font-size: 30px;
}

#container h2 {
  font-size: 21px;
}

#container h3 {
  font-size: 16px;
}

#container h4 {
  font-size: 14px;
}

#container h5 {
  font-size: 12px;
}

#container h6 {
  font-size: 11px;
}

#container p {
  margin-top: 0;
  margin-bottom: 10px;
}

#container blockquote {
  margin: 0;
}

#container ul,
#container ol {
  padding-left: 0;
  margin-top: 0;
  margin-bottom: 0;
}

#container ol ol,
#container ul ol {
  list-style-type: lower-roman;
}

#container ul ul ol,
#container ul ol ol,
#container ol ul ol,
#container ol ol ol {
  list-style-type: lower-alpha;
}

#container dd {
  margin-left: 0;
}

#container code {
  font-family: "Roboto Mono", "Ubuntu Mono", "Menlo", "Consolas", monospace;
  font-size: 12px;
}

#container pre {
  margin-top: 0;
  margin-bottom: 0;
  font: 12px "Roboto Mono", "Ubuntu Mono", "Menlo", "Consolas", monospace;
}

#container .pl-0 {
  padding-left: 0 !important;
}

#container .pl-1 {
  padding-left: 3px !important;
}

#container .pl-2 {
  padding-left: 6px !important;
}

#container .pl-3 {
  padding-left: 12px !important;
}

#container .pl-4 {
  padding-left: 24px !important;
}

#container .pl-5 {
  padding-left: 36px !important;
}

#container .pl-6 {
  padding-left: 48px !important;
}

#container .form-select::-ms-expand {
  opacity: 0;
}

#container:before {
  display: table;
  content: "";
}

#container:after {
  display: table;
  clear: both;
  content: "";
}

#container>*:first-child {
  margin-top: 0 !important;
}

#container>*:last-child {
  margin-bottom: 0 !important;
}

#container a:not([href]) {
  color: inherit;
  text-decoration: none;
}

#container .anchor {
  display: inline-block;
  padding-right: 2px;
  margin-left: -18px;
}

#container .anchor:focus {
  outline: none;
}

#container h1,
#container h2,
#container h3,
#container h4,
#container h5,
#container h6 {
  margin-top: 1em;
  margin-bottom: 16px;
  font-weight: bold;
  line-height: 1.4;
}

#container h1 .octicon-link,
#container h2 .octicon-link,
#container h3 .octicon-link,
#container h4 .octicon-link,
#container h5 .octicon-link,
#container h6 .octicon-link {
  color: #000;
  vertical-align: middle;
  visibility: hidden;
}

#container h1:hover .anchor,
#container h2:hover .anchor,
#container h3:hover .anchor,
#container h4:hover .anchor,
#container h5:hover .anchor,
#container h6:hover .anchor {
  text-decoration: none;
}

#container h1:hover .anchor .octicon-link,
#container h2:hover .anchor .octicon-link,
#container h3:hover .anchor .octicon-link,
#container h4:hover .anchor .octicon-link,
#container h5:hover .anchor .octicon-link,
#container h6:hover .anchor .octicon-link {
  visibility: visible;
}

#container h1 {
  padding-bottom: 0.3em;
  font-size: 2.25em;
  line-height: 1.2;
  border-bottom: 1px solid #eee;
}

#container h1 .anchor {
  line-height: 1;
}

#container h2 {
  padding-bottom: 0.3em;
  font-size: 1.75em;
  line-height: 1.225;
  border-bottom: 1px solid #eee;
}

#container h2 .anchor {
  line-height: 1;
}

#container h3 {
  font-size: 1.5em;
  line-height: 1.43;
}

#container h3 .anchor {
  line-height: 1.2;
}

#container h4 {
  font-size: 1.25em;
}

#container h4 .anchor {
  line-height: 1.2;
}

#container h5 {
  font-size: 1em;
}

#container h5 .anchor {
  line-height: 1.1;
}

#container h6 {
  font-size: 1em;
  color: #777;
}

#container h6 .anchor {
  line-height: 1.1;
}

#container p,
#container blockquote,
#container ul,
#container ol,
#container dl,
#container table,
#container pre {
  margin-top: 0;
  margin-bottom: 16px;
}

#container hr {
  height: 4px;
  padding: 0;
  margin: 16px 0;
  background-color: #e7e7e7;
  border: 0 none;
}

#container ul,
#container ol {
  padding-left: 2em;
}

#container ul ul,
#container ul ol,
#container ol ol,
#container ol ul {
  margin-top: 0;
  margin-bottom: 0;
}

#container li>p {
  margin-top: 16px;
}

#container dl {
  padding: 0;
}

#container dl dt {
  padding: 0;
  margin-top: 16px;
  font-size: 1em;
  font-style: italic;
  font-weight: bold;
}

#container dl dd {
  padding: 0 16px;
  margin-bottom: 16px;
}

#container blockquote {
  padding: 0 15px;
  color: #777;
  border-left: 4px solid #ddd;
}

#container blockquote>:first-child {
  margin-top: 0;
}

#container blockquote>:last-child {
  margin-bottom: 0;
}

#container table {
  display: block;
  width: 100%;
  overflow: auto;
  word-break: normal;
  word-break: keep-all;
}

#container table th {
  font-weight: bold;
}

#container table th,
#container table td {
  padding: 6px 13px;
  border: 1px solid #ddd;
}

#container table tr {
  background-color: #fff;
  border-top: 1px solid #ccc;
}

#container table tr:nth-child(2n) {
  background-color: #f8f8f8;
}

#container img {
  max-width: 100%;
  box-sizing: content-box;
  background-color: #fff;
}

#container code {
  padding: 0;
  padding-top: 0.2em;
  padding-bottom: 0.2em;
  margin: 0;
  font-size: 85%;
  background-color: rgba(0,0,0,0.04);
  border-radius: 3px;
}

#container code:before,
#container code:after {
  letter-spacing: -0.2em;
  content: "\00a0";
}

#container pre>code {
  padding: 0;
  margin: 0;
  font-size: 100%;
  word-break: normal;
  white-space: pre;
  background: transparent;
  border: 0;
}

#container .highlight {
  margin-bottom: 16px;
}

#container .highlight pre,
#container pre {
  padding: 16px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
  background-color: #f7f7f7;
  border-radius: 3px;
}

#container .highlight pre {
  margin-bottom: 0;
  word-break: normal;
}

#container pre {
  word-wrap: normal;
}

#container pre code {
  display: inline;
  max-width: initial;
  padding: 0;
  margin: 0;
  overflow: initial;
  line-height: inherit;
  word-wrap: normal;
  background-color: transparent;
  border: 0;
}

#container pre code:before,
#container pre code:after {
  content: normal;
}

#container kbd {
  display: inline-block;
  padding: 3px 5px;
  font-size: 11px;
  line-height: 10px;
  color: #555;
  vertical-align: middle;
  background-color: #fcfcfc;
  border: solid 1px #ccc;
  border-bottom-color: #bbb;
  border-radius: 3px;
  box-shadow: inset 0 -1px 0 #bbb;
}

#container .pl-c {
  color: #969896;
}

#container .pl-c1,
#container .pl-s .pl-v {
  color: #0086b3;
}

#container .pl-e,
#container .pl-en {
  color: #795da3;
}

#container .pl-s .pl-s1,
#container .pl-smi {
  color: #333;
}

#container .pl-ent {
  color: #63a35c;
}

#container .pl-k {
  color: #a71d5d;
}

#container .pl-pds,
#container .pl-s,
#container .pl-s .pl-pse .pl-s1,
#container .pl-sr,
#container .pl-sr .pl-cce,
#container .pl-sr .pl-sra,
#container .pl-sr .pl-sre {
  color: #183691;
}

#container .pl-v {
  color: #ed6a43;
}

#container .pl-id {
  color: #b52a1d;
}

#container .pl-ii {
  background-color: #b52a1d;
  color: #f8f8f8;
}

#container .pl-sr .pl-cce {
  color: #63a35c;
  font-weight: bold;
}

#container .pl-ml {
  color: #693a17;
}

#container .pl-mh,
#container .pl-mh .pl-en,
#container .pl-ms {
  color: #1d3e81;
  font-weight: bold;
}

#container .pl-mq {
  color: #008080;
}

#container .pl-mi {
  color: #333;
  font-style: italic;
}

#container .pl-mb {
  color: #333;
  font-weight: bold;
}

#container .pl-md {
  background-color: #ffecec;
  color: #bd2c00;
}

#container .pl-mi1 {
  background-color: #eaffea;
  color: #55a532;
}

#container .pl-mdr {
  color: #795da3;
  font-weight: bold;
}

#container .pl-mo {
  color: #1d3e81;
}

#container kbd {
  display: inline-block;
  padding: 3px 5px;
  font: 11px "Roboto Mono", "Ubuntu Mono", "Menlo", "Consolas", monospace;
  line-height: 10px;
  color: #555;
  vertical-align: middle;
  background-color: #fcfcfc;
  border: solid 1px #ccc;
  border-bottom-color: #bbb;
  border-radius: 3px;
  box-shadow: inset 0 -1px 0 #bbb;
}

#container .full-commit .btn-outline:not(:disabled):hover {
  color: #4078c0;
  border: 1px solid #4078c0;
}

#container :checked+.radio-label {
  position: relative;
  z-index: 1;
  border-color: #4078c0;
}

#container .octicon {
  display: inline-block;
  vertical-align: text-top;
  fill: currentColor;
}

#container .task-list-item {
  list-style-type: none;
}

#container .task-list-item+.task-list-item {
  margin-top: 3px;
}

#container .task-list-item input {
  margin: 0 0.2em 0.25em -1.6em;
  vertical-align: middle;
}

#container hr {
  border-bottom-color: #eee;
}
</style></head><body id="container" class="export export-html"><span i="1"><h1 id="-day-2">短学期 Day 2</h1>
</span><span i="2"><h2 id="-">今天回顾了昨天的内容</h2>
</span><span i="4"><p>pwd 绝对路径</p>
</span><span i="6"><p>cd 切换路径</p>
</span><span i="8"><p>ls 显示文件及文件夹
   -a 显示隐藏
   -l 显示详细消息</p>
</span><span i="12"><h2 id="-">上午</h2>
</span><span i="13"><ol>
<li><span i="13"><code>touch</code> 创建文件</span></li>
<li><span i="14">编辑文件 <code>gedit</code>（nonono）</span></li>
<li><span i="15"><code>cat</code> 显示内容（终端内）</span></li>
<li><span i="16"><code>rm</code> 删除</span></li>
<li><span i="17"><code>mkdir</code> 新建文件夹</span></li>
<li><span i="18"><code>rmdir</code> 删除空文件夹 ---<code>rm -rf</code> 删除非空文件夹</span></li>
<li><span i="19"><code>cp</code> 拷贝</span></li>
<li><span i="20"><code>mv</code> 移动（<strong><em>重命名也用mv</em></strong>）</span></li>
<li><span i="21"><code>vi</code> 编辑器(三种模式)， </span><span i="22"><ul>
<li><span i="22">默认命令行模式， <code>vi</code> 命令行复制 <code>yy</code>（<code>nyy</code>表示复制n行）； 粘贴<code>p</code>（默认在光标下一行）； <code>dd</code> 删除（<code>ndd</code>表示删除n行）； <code>u</code>撤销上一步操作</span></li>
<li><span i="23"><code>a, i, o</code>进入插入模式， </span></li>
<li><span i="24">冒号进入底行模式，输入命令，<code>：wq</code> 保存并退出， <code>：w</code>保存， <code>：q</code>退出， <code>：q!</code> 强制退出</span></li>
</ul>
</span></li>
<li><span i="25">gcc步骤： 预处理 -&gt; 编译 -&gt; 汇编 -&gt; 链接 </span><span i="26"><ul>
<li><span i="26"><code>gcc filemane</code> 默认生成可执行二进制文件</span></li>
<li><span i="27"><code>gcc filename -o newname</code> 生成规定名字的可执行文件 </span></li>
</ul>
</span></li>
</ol>
</span><span i="30"><h2 id="-">下午</h2>
</span><span i="31"><ol>
<li><span i="31"><strong>常量</strong>：不变的量 ； 标识常量（不单独开辟空间，与代码放在一起）：  <code>#define N 10</code></span></li>
<li><span i="32"><strong>变量</strong>：一个内存空间的别名 </span><span i="33"><ul>
<li><span i="33"><strong>声明</strong>： <strong><em>&lt;存储类型&gt;  &lt;数据类型&gt; 变量名</em></strong>：</span></li>
<li><span i="34">局部变量： 作用于某个函数， 存放在栈区。直接使用未初始化的编译器会赋随机值</span></li>
<li><span i="35">全局变量： 作用于整个代码，所有函数公用。直接使用未初始化的编译器会赋零值</span></li>
<li><span i="36">32 个关键字 </span></li>
<li><span i="37">变量名遵从标识符命名规则： </span><span i="38"><ol>
<li><span i="38">不能以数字开头</span></li>
<li><span i="39">只能以字母和下划线开头</span></li>
<li><span i="40">不能与关键字重名</span></li>
</ol>
</span></li>
<li><span i="41"><strong>数据类型</strong>：</span><span i="42"><ol>
<li><span i="42">字符型 char： <strong>占1个字节</strong></span></li>
<li><span i="43">整形：</span><span i="44"><ul>
<li><span i="44">短整形（short）: 占2个字节</span></li>
<li><span i="45">整形（int）： 占4个字节</span></li>
<li><span i="46">长整形： </span><span i="47"><ul>
<li><span i="47">（long）占4个字节（32 bit system） </span></li>
<li><span i="48">（long long）占8个字节（64 bit system） </span></li>
</ul>
</span></li>
</ul>
</span></li>
<li><span i="49">浮点型 </span><span i="50"><ul>
<li><span i="50">单精度（float）： 占4字节  6~7位</span></li>
<li><span i="51">双精度（double）： 占8字节  15～16位</span></li>
</ul>
</span></li>
<li><span i="52">void： </span><span i="53"><ul>
<li><span i="53">修饰指针的时候表示任意类型指针</span></li>
<li><span i="54">修饰函数的时候表示函数没有返回值</span></li>
</ul>
</span></li>
</ol>
</span></li>
<li><span i="55"><strong>存储类型</strong>：</span><span i="56"><ol>
<li><span i="56">auto： 自动存储类型，可以被缺省</span></li>
<li><span i="57">regiser： 寄存器存储类型，特点：速度快、价格高、数目少、不容易申请，申请失败时自动转为auto</span></li>
<li><span i="58">static： </span><span i="59"><ul>
<li><span i="59">修饰局部变量时：生命周期延长，作用域不变</span></li>
<li><span i="60">修饰全局变量时：生命周期不变，限制了作用域（仅当前文件可以使用）</span></li>
<li><span i="61">修饰函数时： 生命周期不变，限制了作用域（仅当前文件可以使用）</span></li>
</ul>
</span></li>
<li><span i="62">extern：引用外部变量</span></li>
</ol>
</span></li>
</ul>
</span></li>
<li><span i="63">运算符：（日志只写逻辑运算符和位运算符和前加加后加加）</span><span i="64"><ul>
<li><span i="64">算术运算： + - * / ++ --</span><span i="65"><ol>
<li><span i="65">前++： 先让变量自加，再使用</span></li>
<li><span i="66">后++: 先使用变量，再自加</span></li>
</ol>
</span></li>
<li><span i="67">赋值： = 把右边的内容赋值给左边；还有<code>+=， -=，\*=， /=</code></span></li>
<li><span i="68">关系运算符： <code>==，&lt;, &gt;, &lt;=, &gt;=</code></span></li>
<li><span i="69">逻辑运算符：<code>&amp;&amp;， ||， ！</code></span></li>
<li><span i="70">位运算： <code>&amp;(按位与)  |（按位或）  ～（按位取反）  ^（异或：相同为0，不同为1）  &gt;&gt;（右位移）  &lt;&lt;（左位移）</code></span></li>
<li><span i="71">sizeof()： 求字节大小 sizeof(int) ==&gt; 4</span></li>
</ul>
</span></li>
<li><span i="72">输入输出</span><span i="73"><ul>
<li><span i="73">C语言中没有专门输入输出的语句，都是调用stdio.h里面的函数。</span><span i="74"><ol>
<li><span i="74">格式化输入输出：</span><span i="75"><ul>
<li><span i="75"><code>scanf("格式控制串"， 地址...);</code></span></li>
<li><span i="76"><code>printf("格式控制串"， 变量， 表达式...);</code></span></li>
<li><span i="77">格式控制串：除了格式控制符和转义字符，其它原样输出</span><span i="78"><ol>
<li><span i="78">%d：十进制整数</span></li>
<li><span i="79">%c：单个字符</span></li>
<li><span i="80">%s：字符串 (日志只写前三个)</span></li>
<li><span i="81">%f：单精度浮点数</span></li>
<li><span i="82">%lf：双精度浮点数</span></li>
<li><span i="83">%u：无符号整数</span></li>
<li><span i="84">%x：十六进制</span></li>
<li><span i="85">%p：地址</span></li>
<li><span i="86">%.nf：保留小数点后n位</span></li>
<li><span i="87">%ns：n为域宽，字符长度超过域宽就原样输出；未超过域宽前面用空格补齐</span></li>
<li><span i="88">\n：换行</span></li>
<li><span i="89">\t：制表符</span></li>
</ol>
</span></li>
</ul>
</span></li>
<li><span i="90">字符输入输出:</span><span i="91"><ul>
<li><span i="91"><code>getchar();</code></span></li>
<li><span i="92"><code>putchar();</code></span></li>
</ul>
</span></li>
<li><span i="93">字符串输入输出：</span><span i="94"><ul>
<li><span i="94"><code>gets();</code></span></li>
<li><span i="95"><code>puts();</code></span></li>
</ul>
</span></li>
</ol>
</span></li>
</ul>
</span></li>
</ol>
</span></body>
</html>`

const test = `
alhd\t 
\n
\0
`

func GetBlog() string {
	return blog
}
