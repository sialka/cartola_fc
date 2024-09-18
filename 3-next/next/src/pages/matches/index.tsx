import { MatchResult } from '@/src/components/MatchResult'
import { Page } from '@/src/components/Page'
import { Box, Button } from '@mui/material'
import type { NextPage } from 'next'

const ListMatchesPages: NextPage = () => {
  return (
    <Page>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          flexDirection: "column",
          gap: (theme) => theme.spacing(3),
        }}
      >
        <MatchResult match={{ team_a: "Brasil", team_b: "Argentina"}}/>
        <MatchResult match={{ team_a: "França", team_b: "Bélgica"}}/>
        <MatchResult match={{ team_a: "Portugal", team_b: "Inglaterra"}}/>
      </Box>
    </Page>
  )
}

export default ListMatchesPages