import React from 'react'
import styles from './index.module.scss'

type IProps = {
  value: string
}

const Kbd = ({ value }: IProps) => {
  return <div className={styles['kbd-wrapper']}>{value}</div>
}

export default Kbd
