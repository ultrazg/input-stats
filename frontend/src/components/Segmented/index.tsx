import React, { useState } from 'react'
import { Box, Radio, RadioGroup, Typography } from '@mui/joy'

type IProps = {
  label?: string
  value: number
  options: { label: string; value: number }[]
  onChange: (value: number) => void
}

const Segmented: React.FC<IProps> = ({ label, value, options, onChange }) => {
  const [v, setV] = useState<Number>(value)

  return (
    <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
      {label && <Typography sx={{ fontSize: 'sm' }}>{label}：</Typography>}

      <RadioGroup
        orientation="horizontal"
        name="segmented"
        value={v}
        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
          setV(Number(event.target.value))
          onChange(Number(event.target.value))
        }}
        sx={{
          minHeight: 12,
          padding: '2px',
          borderRadius: '4px',
          bgcolor: 'neutral.softBg',
          '--RadioGroup-gap': '4px',
          '--Radio-actionRadius': '4px',
        }}
      >
        {options.map((item) => (
          <Radio
            key={item.label}
            color="neutral"
            value={item.value}
            disableIcon
            label={item.label}
            variant="soft"
            sx={{ px: 1, alignItems: 'center' }}
            slotProps={{
              action: ({ checked }) => ({
                sx: {
                  ...(checked && {
                    bgcolor: 'background.surface',
                    boxShadow: 'sm',
                    '&:hover': {
                      bgcolor: 'background.surface',
                    },
                  }),
                },
              }),
            }}
          />
        ))}
      </RadioGroup>
    </Box>
  )
}

export default Segmented
