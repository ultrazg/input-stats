import React from 'react'
import styles from './index.module.scss'
import { Button } from '@mui/joy'

export default function Footer() {
  return (
    <div className={styles['footer-wrapper']}>
      <Button
        variant="soft"
        size="sm"
        color="neutral"
      >
        Footer
      </Button>
    </div>
  )
}
