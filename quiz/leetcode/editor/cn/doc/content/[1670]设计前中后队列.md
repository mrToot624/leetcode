<p>请你设计一个队列，支持在前，中，后三个位置的 <code>push</code>&nbsp;和 <code>pop</code>&nbsp;操作。</p>

<p>请你完成&nbsp;<code>FrontMiddleBack</code>&nbsp;类：</p>

<ul> 
 <li><code>FrontMiddleBack()</code>&nbsp;初始化队列。</li> 
 <li><code>void pushFront(int val)</code> 将&nbsp;<code>val</code>&nbsp;添加到队列的 <strong>最前面</strong>&nbsp;。</li> 
 <li><code>void pushMiddle(int val)</code> 将&nbsp;<code>val</code>&nbsp;添加到队列的 <strong>正中间</strong>&nbsp;。</li> 
 <li><code>void pushBack(int val)</code>&nbsp;将&nbsp;<code>val</code>&nbsp;添加到队里的 <strong>最后面</strong>&nbsp;。</li> 
 <li><code>int popFront()</code>&nbsp;将 <strong>最前面</strong> 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 <code>-1</code>&nbsp;。</li> 
 <li><code>int popMiddle()</code> 将 <b>正中间</b>&nbsp;的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 <code>-1</code>&nbsp;。</li> 
 <li><code>int popBack()</code> 将 <strong>最后面</strong> 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 <code>-1</code>&nbsp;。</li> 
</ul>

<p>请注意当有&nbsp;<strong>两个</strong>&nbsp;中间位置的时候，选择靠前面的位置进行操作。比方说：</p>

<ul> 
 <li>将 <code>6</code>&nbsp;添加到&nbsp;<code>[1, 2, 3, 4, 5]</code>&nbsp;的中间位置，结果数组为&nbsp;<code>[1, 2, <strong>6</strong>, 3, 4, 5]</code>&nbsp;。</li> 
 <li>从&nbsp;<code>[1, 2, <strong>3</strong>, 4, 5, 6]</code>&nbsp;的中间位置弹出元素，返回&nbsp;<code>3</code>&nbsp;，数组变为&nbsp;<code>[1, 2, 4, 5, 6]</code>&nbsp;。</li> 
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>
["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
[[], [1], [2], [3], [4], [], [], [], [], []]
<strong>输出：</strong>
[null, null, null, null, null, 1, 3, 4, 2, -1]

<strong>解释：</strong>
FrontMiddleBackQueue q = new FrontMiddleBackQueue();
q.pushFront(1);   // [<strong>1</strong>]
q.pushBack(2);    // [1, <strong>2</strong>]
q.pushMiddle(3);  // [1, <strong>3</strong>, 2]
q.pushMiddle(4);  // [1, <strong>4</strong>, 3, 2]
q.popFront();     // 返回 1 -&gt; [4, 3, 2]
q.popMiddle();    // 返回 3 -&gt; [4, 2]
q.popMiddle();    // 返回 4 -&gt; [2]
q.popBack();      // 返回 2 -&gt; []
q.popFront();     // 返回 -1 -&gt; [] （队列为空）
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= val &lt;= 10<sup>9</sup></code></li> 
 <li>最多调用&nbsp;<code>1000</code>&nbsp;次&nbsp;<code>pushFront</code>，&nbsp;<code>pushMiddle</code>，&nbsp;<code>pushBack</code>，&nbsp;<code>popFront</code>，&nbsp;<code>popMiddle</code>&nbsp;和&nbsp;<code>popBack</code> 。</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>设计 | 队列 | 数组 | 链表 | 数据流</details><br>

<div>👍 106, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/issues' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：为满足广大读者的需求，网站上架 [速成目录](https://labuladong.online/algo/intro/quick-learning-plan/)，如有需要可以看下，谢谢大家的支持~**

<details><summary><strong>labuladong 思路</strong></summary>

<!-- vip -->
<!-- i_62b43720e4b07bd2d7b1b6dd -->

本题思路为 labuladong 网站会员专属，请 [点击这里](https://labuladong.online/algo/intro/site-vip/) 购买会员并「按照各个插件的解锁方法手动刷新数据」。

若之前已经购买会员并成功解锁插件，现在却突然出现这个问题，是因为添加了新的题解数据。请尝试重新手动刷新插件数据。进入 [会员购买页](https://labuladong.online/algo/intro/site-vip/) 向下翻即可查看各个插件刷新数据的方法。

若依然无法解决问题，可以在按照 [bug 反馈页面](https://labuladong.online/algo/intro/bug-report/) 的提示像我反馈问题，如是 bug 我会立即修复。</details>
</div>

