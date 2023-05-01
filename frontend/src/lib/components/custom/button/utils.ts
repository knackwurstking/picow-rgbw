export function createRippleAnimation(
  ev: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement }
) {
  const button = ev.currentTarget;
  const rect = button.getBoundingClientRect();

  const cursorX = ev.clientX - rect.x;
  const cursorY = ev.clientY - rect.y;

  const circle = document.createElement("span");
  const diameter = Math.max(rect.width, rect.height);
  const radius = diameter / 2;

  circle.style.width = circle.style.height = `${diameter}px`;
  circle.style.left = `${cursorX - radius}px`;
  circle.style.top = `${cursorY - radius}px`;
  circle.classList.add("ripple");

  const ripple = button.getElementsByClassName("ripple")[0];

  if (ripple) {
    ripple.remove();
  }

  button.appendChild(circle);
}
