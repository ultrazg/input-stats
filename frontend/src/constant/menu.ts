import { MenuItem } from '@/types'

export const BASE_MENU: MenuItem[] = [
  {
    type: 'separator',
    title: '',
    tooltip: '',
    event: '',
    disabled: false,
  },
  {
    type: 'item',
    title: '退出',
    tooltip: '退出 Input Stats',
    event: 'onMenuItemClick:ExitApp',
    disabled: false,
  },
]
