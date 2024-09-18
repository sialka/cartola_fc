import { MatchResult } from '@/src/components/MatchResult'
import Image from "next/image";
import { Page } from '@/src/components/Page'
import { Section } from '@/src/components/Section'
import { Box, Button, Chip, styled, Table, TableBody, TableCell, TableHead, TableRow, Typography } from '@mui/material'
import type { NextPage } from 'next'
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import router, { useRouter } from 'next/router';
import { Match } from "../../util/models";
import { green } from "@mui/material/colors";

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  "&>td": {
    border: 0,
  },
  "&:nth-of-type(odd)": {
    backgroundColor: theme.palette.divider,
    border: 0,
  },
}));

const StyledTableHead = styled(TableHead)({
  th: {
    border: 0,
  },
});


const HeadCellContent = styled(Box)({
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
});

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  textAlign: "center",
}));

type HeadImageProps = {
  src: string;
  alt: string;
};

const HeadImage = (props: HeadImageProps) => (
  <Image src={props.src} alt={props.alt} width={32} height={32} />
);

function formatAction(playerName: string, action: string) {
  switch (action) {
    case "goal":
      return `${playerName} fez um gol`;
    case "assist":
      return `${playerName} deu uma assistência`;
    case "yellow card":
      return `${playerName} levou um cartão amarelo`;
    case "red card":
      return `${playerName} levou um cartão vermelho`;
    default:
      return `${playerName} fez alguma coisa`;
  }
}

// Teste
const match: Match = {
  id: '1',
  team_a: "Brasil",
  team_b: "Argentina",
  match_date: "12/12/2024 16:00",
  result: "1-0",
  actions: [
    {
      action: "goal",
      minutes: 10,
      player_name: "Neymar",
      score: 5
    },
    {
      action: "assist",
      minutes: 14,
      player_name: "Neymar",
      score: 4
    },
  ]
}


const ShowMatchPage: NextPage = () => {

  const router = useRouter();
  const { id: matchId } = router.query;

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
        <MatchResult match={match}/>

        <Section
            sx={{
              marginTop: "-30px",
              zIndex: -10,
              width: 750,
              position: "relative",
            }}
          >
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <StyledTableHead>
                <TableRow>
                  <TableCell>
                    <HeadCellContent>
                      <HeadImage src="/img/time.svg" alt="" /> Tempo de jogo
                    </HeadCellContent>
                  </TableCell>
                  <TableCell>
                    <HeadCellContent>
                      <HeadImage src="/img/player.svg" alt="" /> Jogador
                    </HeadCellContent>
                  </TableCell>
                  <TableCell>
                    <HeadCellContent>
                      <HeadImage src="/img/score.svg" alt="" /> Pontuação
                    </HeadCellContent>
                  </TableCell>
                </TableRow>
              </StyledTableHead>
              <TableBody>
                { match!.actions.map((action, key) => (
                  <StyledTableRow key= { key }>
                    <StyledTableCell>{ action.minutes }&#39;</StyledTableCell>
                    <StyledTableCell>
                      { formatAction(action.player_name, action.action) }
                    </StyledTableCell>
                    <StyledTableCell
                      sx={{
                        color: (theme) =>
                          action.score > 0
                            ? green[500]
                            : theme.palette.primary.main,
                      }}
                    >
                      <Typography>{ action.score }  pts</Typography>
                    </StyledTableCell>
                  </StyledTableRow>
                 ))}
              </TableBody>
            </Table>
            <Chip
              label={
                <Box>
                  <Typography component="span">Total do jogo: </Typography>
                  <Typography
                    component="span"
                    sx={{
                      fontWeight: "bold",
                      color: (theme) =>
                        1 > 0 ? green[500] : theme.palette.primary.main,
                    }}
                  >
                    -- pts
                  </Typography>
                </Box>
              }
              sx={{
                bottom: -15,
                position: "absolute",
                right: 15,
                backgroundColor: (theme) => theme.palette.background.default,
              }}
            />
          </Section>
          <Button
            variant="contained"
            size="large"
            startIcon={<ArrowBackIcon />}
            onClick={() => router.back()}
          >
            Voltar
          </Button>
      </Box>
    </Page>
  )
}

export default ShowMatchPage

function useHttp<T>(arg0: string | null, fetcherStats: any, arg2: { refreshInterval: number; }): { data: any; } {
  throw new Error('Function not implemented.');
}
