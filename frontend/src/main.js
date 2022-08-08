import "./style.css";
import { Cmd } from "../wailsjs/go/main/App";

const buttons = document.querySelectorAll("button");

/**
 * @param {MouseEvent} e
 */
function handleClick(e) {
  e.PreventDefault;
  const cmds = e.target.dataset.cmd;
  Cmd(cmds);
}
buttons.forEach((btn) => {
  btn.addEventListener("click", handleClick, true);
});
