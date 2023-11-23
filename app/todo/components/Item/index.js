"use client";

import { useState } from "react";

// 列表项
export default function Item() {
  // 记录鼠标移入
  const { mouse, setMouse } = useState(false);
  // 接收参数，返回函数用于事件绑定
  handleMouse = (flag) => () => setMouse(flag);

  return (
    <li
      style={{ backgroundColor: mouse ? "#ddd" : "white" }}
      onMouseEnter={handleMouse(true)}
      onMouseLeave={handleMouse(false)}
    >
      <button style={{ display: mouse ? "block" : "none" }}></button>
    </li>
  );
}
