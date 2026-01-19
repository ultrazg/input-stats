export type StatsSnapshot = {
  modifierStats: Record<string, number>
  keyStats: Record<string, number>
  comboStats: Record<string, number>
  mouseLeftClick: number
  mouseRightClick: number
  mouseMovePixels: number
  mouseWheel: number
}
