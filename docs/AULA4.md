# 4. Next.js e criação de Componentes com Material UI

**Next.js Serve Side Rendering (SSR)**

Criando o projeto base com REACT.

```bash
$ npx create-next-app --typescript
✔ What is your project named? … next
✔ Would you like to use ESLint? … No / Yes
✔ Would you like to use Tailwind CSS? … No / Yes
✔ Would you like to use `src/` directory? … No / Yes
✔ Would you like to use App Router? (recommended) … No / Yes
✔ Would you like to customize the default import alias (@/*)? … No / Yes
Creating a new Next.js app in /home/sidnei/developer/cartola-fc/3-next/next.

Using npm.

Initializing project with template: default 


Installing dependencies:
- react
- react-dom
- next

Installing devDependencies:
- typescript
- @types/node
- @types/react
- @types/react-dom


added 28 packages, and audited 29 packages in 5s

3 packages are looking for funding
  run `npm fund` for details

found 0 vulnerabilities
Success! Created next at /home/sidnei/developer/cartola-fc/3-next/next
```

**Adicionando o Material UI**

```bash
$ yarn add @mui/material @emotion/react @emotion/styled @emotion/cache @emotion/server @next/font @mui/icons-material
```

- Criar a pasta src dentro de next e mover pages para dentro src

> File: src/pages/_app.tsx

```tsx
import Head from "next/head";
import { AppProps } from "next/app";
import { ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import { CacheProvider, EmotionCache } from "@emotion/react";
import theme from "../util/theme";
import createEmotionCache from "../util/createEmotionCache";
import { Navbar } from "../components/Navbar";

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache();

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache;
}

export default function MyApp(props: MyAppProps) {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props;
  return (
    <CacheProvider value={emotionCache}>
      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <Navbar />
        <Component {...pageProps} />
      </ThemeProvider>
    </CacheProvider>
  );
}
```

> File: src/pages/_document.tsx

```tsx
import Document, { Html, Head, Main, NextScript } from "next/document";
import createEmotionServer from "@emotion/server/create-instance";
// import theme, { roboto } from "../util/theme";
import createEmotionCache from "../util/createEmotionCache";

export default class MyDocument extends Document {
  render() {
    return (
      // <Html lang="en" className={roboto.className}>
      <Html lang="en">
        <Head>
          {/* PWA primary color */}
          {/* <meta name="theme-color" content={theme.palette.primary.main} /> */}
          <link rel="shortcut icon" href="/favicon.ico" />
          <meta name="emotion-insertion-point" content="" />
          <meta name="viewport" content="initial-scale=1, width=device-width" />
          {(this.props as any).emotionStyleTags}
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

// `getInitialProps` belongs to `_document` (instead of `_app`),
// it's compatible with static-site generation (SSG).
MyDocument.getInitialProps = async (ctx) => {
  // Resolution order
  //
  // On the server:
  // 1. app.getInitialProps
  // 2. page.getInitialProps
  // 3. document.getInitialProps
  // 4. app.render
  // 5. page.render
  // 6. document.render
  //
  // On the server with error:
  // 1. document.getInitialProps
  // 2. app.render
  // 3. page.render
  // 4. document.render
  //
  // On the client
  // 1. app.getInitialProps
  // 2. page.getInitialProps
  // 3. app.render
  // 4. page.render

  const originalRenderPage = ctx.renderPage;

  // You can consider sharing the same Emotion cache between all the SSR requests to speed up performance.
  // However, be aware that it can have global side effects.
  const cache = createEmotionCache();
  const { extractCriticalToChunks } = createEmotionServer(cache);

  ctx.renderPage = () =>
    originalRenderPage({
      enhanceApp: (App: any) =>
        function EnhanceApp(props) {
          return <App emotionCache={cache} {...props} />;
        },
    });

  const initialProps = await Document.getInitialProps(ctx);
  // This is important. It prevents Emotion to render invalid HTML.
  // See https://github.com/mui/material-ui/issues/26561#issuecomment-855286153
  const emotionStyles = extractCriticalToChunks(initialProps.html);
  const emotionStyleTags = emotionStyles.styles.map((style) => (
    <style
      data-emotion={`${style.key} ${style.ids.join(" ")}`}
      key={style.key}
      // eslint-disable-next-line react/no-danger
      dangerouslySetInnerHTML={{ __html: style.css }}
    />
  ));

  return {
    ...initialProps,
    emotionStyleTags,
  };
};
```
**Criando Util**

> File: src/util/createEmotionCache.ts

```tsx
import createCache from '@emotion/cache';

const isBrowser = typeof document !== 'undefined';

// On the client side, Create a meta tag at the top of the <head> and set it as insertionPoint.
// This assures that MUI styles are loaded first.
// It allows developers to easily override MUI styles with other styling solutions, like CSS modules.
export default function createEmotionCache() {
  let insertionPoint;

  if (isBrowser) {
    const emotionInsertionPoint = document.querySelector<HTMLMetaElement>(
      'meta[name="emotion-insertion-point"]',
    );
    insertionPoint = emotionInsertionPoint ?? undefined;
  }

  return createCache({ key: 'mui-style', insertionPoint });
}
```

> File: src/util/models.ts

```tsx
export type Player = {
  id: string;
  name: string;
  price: number;
};

export const PlayersMap: { [key: string]: string } = {
  "Cristiano Ronaldo": "/img/players/Cristiano Ronaldo.png",
  "De Bruyne": "/img/players/De Bruyne.png",
  "Harry Kane": "/img/players/Harry Kane.png",
  Lewandowski: "/img/players/Lewandowski.png",
  Maguirre: "/img/players/Maguirre.png",
  Messi: "/img/players/Messi.png",
  Neymar: "/img/players/Neymar.png",
  Richarlison: "/img/players/Richarlison.png",
  "Vinicius Junior": "/img/players/Vinicius Junior.png",
};
```

> File: src/util/theme.ts

```tsx
import { createTheme } from "@mui/material";
import { Roboto } from "@next/font/google";

export const roboto = Roboto({
  weight: ["300", "400", "500", "700"],
  subsets: ["latin"],
  display: "swap",
  fallback: ["Helvetica", "Arial", "sans-serif"],
});

const theme = createTheme({
  palette: {
    primary: {
      main: "#B8C858",
    },
    secondary: {
      main: "#0E4987",
    },
    divider: '#1B73A7',
    background: {
      default: "#09345e",
      paper: "#0E4987",
    },
    mode: "dark",
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        "html, body, body>div": {
          padding: 0,
          margin: 0,
          width: "100%",
          height: "100%",
        },
        body: {
          background: "url(/img/background.png) no-repeat center center fixed",
          WebkitBackgroundSize: "cover",
          OBackgroundSize: "cover",
          BackgroundSize: "cover",
          MozBackgroundSize: "cover",
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: ({ theme, ownerState }) => ({
          ...(ownerState.variant === "outlined" && {
            border: `3px solid ${theme.palette.divider}`,
          }),
        }),
      },
    },
  },
});

export default theme;
```

**Criando Componente**

> File: src/components/Navbar.tsx

```tsx
import { AppBar, Avatar, Box, Button, Chip, Toolbar } from "@mui/material";
import Image from "next/image";
import Link, { LinkProps } from "next/link";
import { useRouter } from "next/router";
import { PropsWithChildren } from "react";
//import { useHttp } from "../hooks/useHttp";
//import { fetcherStats, httpStats } from "../util/http";

export type NavbarItemProps = LinkProps & { showUnderline: boolean };

export const NavbarItem = (props: PropsWithChildren<NavbarItemProps>) => {
  const { showUnderline, ...linkProps } = props;

  return (
    //@ts-expect-error
    <Button
      component={Link}
      sx={{
        color: "white",
        display: "inline-block",
        textAlign: "center",
        "&::after": (theme) => ({
          content: '""',
          borderBottom: showUnderline
            ? `4px solid ${theme.palette.primary.main}`
            : `4px solid transparent`,
          width: "100%",
          display: "block",
        })
      }}
      {...props}
    />
  );
};

export const Navbar = () => {

  const router = useRouter();

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static" sx={{ background: "none", boxShadow: "none" }}>
        <Toolbar>
          <Image
            src="/img/logo.png"
            width={315}
            height={58}
            alt="logo"
            priority={true}
          />

          <Box sx={{ flexGrow: 1, ml: (theme) => theme.spacing(4) }}>
            <NavbarItem href="/" showUnderline={router.pathname === "/"}>
              Home
            </NavbarItem>
            <NavbarItem
              href="/players"
              showUnderline={router.pathname === "/players"}
            >
              Escalação
            </NavbarItem>
            <NavbarItem
              href="/matches"
              showUnderline={["/matches", "/matches/[id]"].includes(
                router.pathname
              )}
            >
              Jogo
            </NavbarItem>
          </Box>

          <Chip
            //label={data ? data.balance : 0}
            label = {"300"}
            avatar={<Avatar>C$</Avatar>}
            color="secondary"
          />


        </Toolbar>
      </AppBar>
    </Box>
  );
};
```

> File: src/components/TeamLogo.tsx

```tsx
import { Box, BoxProps, Typography } from "@mui/material";
import Image from "next/image";
import { Label } from "./Label";

export type TeamLogoProps = BoxProps;

export const TeamLogo = (props: TeamLogoProps) => {
  return (
    <Box
        {...props}
        sx={{ display: "flex", flexDirection: "column", alignItems: "center", ...props.sx }}
    >
      <Image
        src="/img/my-team-logo.svg"
        width={84.5}
        height={88.5}
        alt="Meu time FC"
        priority={true}
      />
      <Label>Meu time FC</Label>

    </Box>
  );
};
```

> File: src/components/Section.tsx

```tsx
import { Paper, PaperProps } from "@mui/material";

export type SectionProps = PaperProps;

export const Section = (props: SectionProps) => {
  return (
    <Paper
      {...props}
      variant="outlined"
      sx={{
        padding: (theme) => theme.spacing(2),
        ...props.sx,
      }}
    />
  );
};
```

> File: src/components/Label.tsx

```tsx
import { Typography, TypographyProps } from "@mui/material";

export type LabelProps = TypographyProps;

export const Label = (props: TypographyProps) => {
  return <Typography variant="h6" component="span" {...props} />;
};
```

> File: src/components/Page.tsx

```tsx
import { Container } from "@mui/system";
import { PropsWithChildren } from "react";

export const Page = (props: PropsWithChildren) => {
  return (
    <Container sx={{ paddingTop: (theme) => theme.spacing(3) }}>
      {props.children}
    </Container>
  );
};
```

**Criando Paginas**

> File:  src/pages/players.tsx

```tsx
import React, { useCallback, useMemo, useState } from "react";
import { Box, Autocomplete, Grid, ListItem, ListItemAvatar, Avatar, ListItemText, Divider, TextField, List, IconButton, Button } from '@mui/material'
import type { NextPage } from 'next'
import Image from "next/image";
import { Page } from '../components/Page'
import { TeamLogo } from '../components/TeamLogo'
import { Section } from "../components/Section";
import { Label } from "../components/Label";
import { Player, PlayersMap } from "../util/models";
import DeleteIcon from "@mui/icons-material/Delete";
import PersonIcon from "@mui/icons-material/Person2";

const players = [
  {
    id: 1,
    name: "Messi",
    price: 35,
  },
  {
    id: 2,
    name: "Cristiano Ronaldo",
    price: 35,
  },
  {
    id: 3,
    name: "Neymar",
    price: 25,
  },
  {
    id: 4,
    name: "De Bruyne",
    price: 25,
  },
  {
    id: 5,
    name: "Vinicius Junior",
    price: 25,
  },
  {
    id: 6,
    name: "Lewandowski",
    price: 15,
  },
  {
    id: 7,
    name: "Maguirre",
    price: 15,
  },
  {
    id: 8,
    name: "Richarlison",
    price: 15,
  },
  {
    id: 9,
    name: "Harry Kane",
    price: 15,
  },
];

const fakePlayer = {
  id: "",
  name: "Escolha um jogador",
  price: 0,
};

const makeFakePlayer = (key: number) => ({
  ...fakePlayer,
  name: `${fakePlayer.name} ${key + 1}`,
});

const totalPlayers = 4;
const balance = 300;

const fakePlayers: Player[] = new Array(totalPlayers)
  .fill(0)
  .map((_, key) => makeFakePlayer(key));

const ListPlayerPage: NextPage = () => {

  const [playersSelected, setPlayersSelected] = useState(fakePlayers);

  const countPlayersUsed = useMemo(
    () => playersSelected.filter((player) => player.id !== "").length,
    [playersSelected]
  );

  const budgetRemaining = useMemo(
    () => balance - playersSelected.reduce((acc, player) => acc + player.price, 0),
    [playersSelected]
  );

  const addPlayer = useCallback((player: Player) => {
    setPlayersSelected((prev) => {
      // Verifica se o player já está na lista
      const hasFound = prev.find((p) => p.id === player.id);
      if (hasFound) return prev;

      // Verifica se existem jogadores disponiveis na lista
      const firstIndexFakerPlayer = prev.findIndex((p) => p.id === "");
      if (firstIndexFakerPlayer === -1) return prev;

      // Adicionando jogadores
      const newPlayers = [...prev];
      newPlayers[firstIndexFakerPlayer] = player;
      return newPlayers;
    });
  }, []);

  const removePlayer = useCallback((index: number) => {
    setPlayersSelected((prev) => {
      const newPlayers = prev.map((p, key) => {
        if (key === index) {
          return makeFakePlayer(key);
        }
        return p;
      });
      return newPlayers;
    });
  }, []);

  return (
    <Page>
      <Grid container
        sx={{
          display: "flex",
          justifyContent: "center",
          gap: (theme) => theme.spacing(4),
        }}
      >
        <Grid item xs={12}>
          <Section>
            <TeamLogo
              sx={{
                position: 'absolute',
                flexDirection: 'row',
                ml: (theme) => theme.spacing(-5.5),
                mt: (theme) => theme.spacing(-3.5),
                gap: (theme) => theme.spacing(1),
              }}/>
            <Box
              sx={{
                display: 'flex',
                justifyContent: 'flex-end',
                gap: (theme) => theme.spacing(2),
              }}>
              <Label>Você ainda tem</Label>
              <Label>C$ {budgetRemaining}</Label>
            </Box>
          </Section>
        </Grid>
        <Grid item xs={12}>
          <Section>
            <Grid container>
              <Grid item xs={6}>
                <Autocomplete
                  sx={{ width: 400 }}
                  isOptionEqualToValue={(option, value) => {
                    console.log(option);
                    return option.name
                      .toLowerCase()
                      .includes(value.name.toLowerCase());
                  }}
                  getOptionLabel={(option) => option.name}
                  options={players}
                  onChange={(_event, newValue) => {
                    if (!newValue) {
                      return;
                    }
                    addPlayer(newValue);
                  }}
                  renderOption={(props, option) => (
                    <React.Fragment key={option.name}>
                      <ListItem {...props}>
                        <ListItemAvatar>
                          <Avatar>
                            <Image
                              src={PlayersMap[option.name]}
                              width={40}
                              height={40}
                              alt=""
                            />
                          </Avatar>
                        </ListItemAvatar>
                        <ListItemText
                          primary={`${option.name}`}
                          secondary={`C$ ${option.price}`}
                        />
                      </ListItem>
                      <Divider variant="inset" component="li" />
                    </React.Fragment>
                  )}
                  renderInput={(params) => (
                    <TextField
                      {...params}
                      label="Pesquise um jogador"
                      InputProps={{
                        ...params.InputProps,
                        sx: {
                          backgroundColor: (theme) => theme.palette.background.default,
                        },
                      }}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
              <Label>Meu time</Label>
                <List>
                  {playersSelected.map((player, key) => (
                    <React.Fragment key={key}>
                      <ListItem
                        secondaryAction={
                          <IconButton
                            edge="end"
                            disabled={player.id === ""}
                            onClick={() => removePlayer(key)}
                          >
                            <DeleteIcon />
                          </IconButton>
                        }
                      >
                        <ListItemAvatar>
                          <Avatar>
                            {player.id === "" ? (
                              <PersonIcon />
                            ) : (
                              <Image
                                src={PlayersMap[player.name]}
                                width={40}
                                height={40}
                                alt=""
                              />
                            )}
                          </Avatar>
                        </ListItemAvatar>
                        <ListItemText
                          primary={player.name}
                          secondary={`C$ ${player.price}`}
                        />
                      </ListItem>
                      <Divider variant="inset" component="li" />
                    </React.Fragment>
                  ))}
                </List>
              </Grid>
            </Grid>
          </Section>
        </Grid>
        <Grid item>
          <Button
            variant="contained"
            size="large"
            disabled={countPlayersUsed < totalPlayers || budgetRemaining < 0}
            onClick={() => console.log("saveMyPlayers()")}
          >
            Salvar
          </Button>
        </Grid>
      </Grid>
    </Page>
  )
}

export default ListPlayerPage
```

> File: src/pages/index.tsx

```tsx
import GroupsIcon from "@mui/icons-material/Groups";
import Link from "next/link";
import { Button, Divider, Grid, Grid2, styled } from '@mui/material'
import type { NextPage } from 'next'
import { Page } from '../components/Page'
import { TeamLogo } from "../components/TeamLogo";
import { Section } from "../components/Section";
import { Label } from "../components/Label";

const BudgetContainer = styled(Section)(({ theme }) => ({
  width: "800px",
  height: "300px",
  marginTop: theme.spacing(8),
  display: "flex",
  alignItems: "center",
}));

const HomePage: NextPage = () => {
  return (
    <Page>
      <Grid container
          sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: (theme) => theme.spacing(3),
        }}>
        <Grid item
        >
          <TeamLogo
            sx={{ position: "absolute", left: 0, right: 0, m: "auto" }}
          />
          <BudgetContainer>
            <Grid container>
              <Grid
                item
                xs={5}
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Label>Última pontuação</Label>
                <Label>99.04</Label>
              </Grid>
              <Grid item xs={2} sx={{display: "flex", justifyContent: "center" }}>
                <Divider orientation="vertical" sx={{ height: "auto" }}/>
              </Grid>
              <Grid
                item
                xs={5}
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Label>Patrimônio</Label>
                <Label>300</Label>
              </Grid>
            </Grid>
          </BudgetContainer>
        </Grid>
        <Grid item>
          <Button component={Link}
            href="/players"
            variant="contained"
            startIcon={<GroupsIcon />}>
            Escalar Jogadores
          </Button>

        </Grid>
      </Grid>
    </Page>
  )
}

export default HomePage
```

> File: src/pages/matches/[id].tsx

```tsx
import { Button } from '@mui/material'
import type { NextPage } from 'next'

const ShowMatchPage: NextPage = () => {
  return (
    <div>
      <Button variant='contained'>Teste</Button>
    </div>
  )
}

export default ShowMatchPage
```
> File: src/pages/matches/index.tsx

```tsx
import { Button } from '@mui/material'
import type { NextPage } from 'next'

const ListMatchesPages: NextPage = () => {
  return (
    <div>
      <Button variant='contained'>Teste</Button>
    </div>
  )
}

export default ListMatchesPages
```

