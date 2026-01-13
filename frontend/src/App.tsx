import React from 'react'
import styles from './app.module.scss'
import { Typography } from '@mui/joy'
import { Kbd } from '@/components'
import { Footer } from '@/components'

function App() {
  return (
    <React.Fragment>
      <div className={styles['app-wrapper']}>
        <Typography level="h4">概览</Typography>
        <div className={styles['overview-wrapper']}>
          <div className={styles['overview-item']}>
            <div className={styles['overview-chunk']}>
              <span className={styles['overview-label']}>⌨️键盘敲击</span>
              <span className={styles['overview-value']}>1</span>
            </div>
          </div>
          <div className={styles['overview-item']}>
            <div className={styles['overview-chunk']}>
              <span className={styles['overview-label']}>🖱左键点击</span>
              <span className={styles['overview-value']}>1</span>
            </div>
            <div className={styles['overview-chunk']}>
              <span className={styles['overview-label']}>🖱右键点击</span>
              <span className={styles['overview-value']}>1</span>
            </div>
          </div>
          <div className={styles['overview-item']}>
            <div className={styles['overview-chunk']}>
              <span className={styles['overview-label']}>↔️鼠标移动</span>
              <span className={styles['overview-value']}>1</span>
            </div>
            <div className={styles['overview-chunk']}>
              <span className={styles['overview-label']}>↕️滚动距离</span>
              <span className={styles['overview-value']}>1</span>
            </div>
          </div>
        </div>

        <Typography level="h4">键位统计</Typography>

        <div className={styles['kdb-statistics-wrapper']}>
          {[
            'Ctrl',
            'Shift',
            'Alt',
            'A',
            'B',
            'C',
            'D',
            'E',
            'F',
            'G',
            'H',
            'I',
            'J',
            'K',
            'L',
            'M',
            'N',
            'O',
            'P',
            'Q',
            'R',
            'S',
            'T',
            'U',
            'V',
            'W',
            'X',
            'Y',
            'Z',
          ].map((item, index) => (
            <div
              className={styles['kdb-item']}
              key={item}
            >
              <Kbd value={item} />
              <span
                style={
                  (index + 1) % 3 === 0
                    ? { paddingRight: 0 }
                    : { paddingRight: 12 }
                }
              >
                12
              </span>
            </div>
          ))}
        </div>
      </div>
      <Footer />
    </React.Fragment>
  )
}

export default App
