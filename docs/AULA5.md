# 5. Next.js: Server Side Rendering e Integração entre microsserviços

**Util**

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

export type Action = {
  player_name: string;
  minutes: number;
  action: "goal" | "yellow card" | "red card" | "assist";
  score: number;
};

export type Match = {
  id: string;
  match_date: string;
  team_a: string; //Brasil
  team_b: string; //Argentina
  result: string; //'1-0'
  //score: number;
  actions: Action[];
};

export const TeamsImagesMap: { [key: string]: string } = {
  Alemanha: "/img/flags/Alemanha.png",
  Argentina: "/img/flags/Argentina.png",
  Bélgica: "/img/flags/Belgica.png",
  Brasil: "/img/flags/Brasil.png",
  França: "/img/flags/Franca.png",
  Inglaterra: "/img/flags/Inglaterra.png",
  Polônia: "/img/flags/Polonia.png",
  Portugal: "/img/flags/Portugal.png",
};
```

**Componente**

> File: src/components/MatchResult.tsx

```tsx
import { Box, styled, Typography } from "@mui/material";
import Image from "next/image";
import { Match, TeamsImagesMap } from "../util/models";
// import {format, parseISO} from 'date-fns';

type FlagProps = {
  src: string;
  alt: string;
};

const Flag = (props: FlagProps) => {
  return <Image src={props.src} alt={props.alt} width={121} height={76} style={{
    marginLeft: '-5px',
    marginRight: '-5px'
  }} />;
};

const ResultContainer = styled(Box)(({ theme }) => ({
  display: "flex",
  width: "400px",
  backgroundColor: theme.palette.background.default,
  alignItems: "center",
  padding: 0,
  border: 'none !important',
  boxShadow: 'none'
}));


const ResultItem = styled(Box)(({ theme }) => ({
  height: "55px",
  display: "flex",
  alignItems: "center",
}));

type MatchResultProps = {
  match: Match;
};

export const MatchResult = (props: MatchResultProps) => {
  const { match } = props;
  return (
    <Box display="flex">
      <Flag src={TeamsImagesMap[match.team_a]} alt={match.team_a} />
      <ResultContainer>

        <ResultItem width={"150px"} justifyContent="flex-end">
          <Typography variant="h6">{match.team_a}</Typography>
        </ResultItem>

        <ResultItem width={"100px"} justifyContent="center" position="relative">
          <Box sx={{ position: "absolute", top: 0, fontSize: "0.70rem" }}>
            { /* format(parseISO(match.match_date), 'dd/MM/yyyy HH:mm')*/ }
            12/12/2022 00:00
          </Box>
          <Typography variant="h6" sx={{ fontWeight: "900" }}>
            { /* match.result.split('-').join(' - ') */ }
            1 - 0
          </Typography>
        </ResultItem>

        <ResultItem width={"150px"} justifyContent="flex-start">
          <Typography variant="h6">{match.team_b}</Typography>
        </ResultItem>

      </ResultContainer>
      <Flag src={TeamsImagesMap[match.team_b]} alt={match.team_a} />
    </Box>
  );
};
```

**Page**

> File: src/pages/matches/index.tsx

```tsx
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
```

> File: src/pages/matches/[id].tsx

```tsx
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

```

**Integração**

django

$ pipenv shell

- delete o db.sqlite3
- delete app/migrations
- Update no models.py

**Ajustes**

> File: app/receivers.py

```py
from .models import Player, Team, Match, Action
from django.db.models.signals import post_save, pre_save
from django.dispatch import receiver

@receiver(post_save, sender=Player)
def publish_player_created(sender, instance: Player, created: bool, **kwargs):
  if created:
    print("Player created")


@receiver(post_save, sender=Team)
def publish_team_created(sender, instance: Team, created: bool, **kwargs):
  if created:
    print("Team created")

@receiver(post_save, sender=Match)
def publish_match_created(sender, instance: Match, created: bool, **kwargs):
  if created:
    print("Match created")

@receiver(pre_save, sender=Match)
def get_old_match(sender, instance: Match, **kwargs):
  try:
    instance._pre_save_instance = Match.objects.get(pk=instance.pk)
  except Match.DoesNotExist:
    instance._pre_save_instance = None

@receiver(post_save, sender=Match)
def publish_new_match_result(sender, instance: Match, created:bool, **kwargs):
  if not created and instance._pre_save_instance and (instance._pre_save_instance.team_a_goal != instance.team_a_goal or instance._pre_save_instance.team_b_goal != instance.team_b_goal):
    print("Match result published")

@receiver(post_save, sender=Action)
def publish_action_added(sender, instance: Action, created:bool, **kwargs):
  if created:
    print("Action created")
```

> File: app/models.py

```py
from django.db import models
import uuid

class UUIDModel(models.Model):
  uuid = models.UUIDField(default=uuid.uuid4, editable=False, unique=True)

  class Meta:
    abstract = True

# Create your models here.
class Player(UUIDModel):
  name = models.CharField(max_length=50, verbose_name='Nome')
  initial_price = models.FloatField(verbose_name='Preço inicial')

  def __str__(self):
    return self.name

  class Meta:
    verbose_name = 'Jogador'
    verbose_name_plural = 'Jogadores'


# Este Model é apenas decorativo
class Team(UUIDModel):
  name = models.CharField(max_length=50, verbose_name='Nome')

  def __str__(self):
    return self.name

  class Meta:
    verbose_name = 'Time'
    verbose_name_plural = 'Times'

class MyTeam(UUIDModel):
  players = models.ManyToManyField(Player, verbose_name='Jogadores')

  def __str__(self):
    return [ player.name for player in self.players.all() ].__str__()

  class Meta:
    verbose_name = 'Meu time'
    verbose_name_plural = 'Meus times'

class Match(UUIDModel):
  match_date = models.DateTimeField(verbose_name='Data do jogo')
  team_a = models.ForeignKey(
    Team,
    on_delete=models.PROTECT,
    related_name='team_a_matches',
    verbose_name='Time A'
  )
  team_a_goal = models.IntegerField(default=0, verbose_name='Gols do Time A')

  team_b = models.ForeignKey(
    Team,
    on_delete=models.PROTECT,
    related_name='team_b_matches',
    verbose_name='Time B'
  )
  team_b_goal = models.IntegerField(default=0, verbose_name='Gols do Time B')

  class Meta:
    verbose_name = 'Jogo'
    verbose_name_plural = 'Jogos'

  def __str__(self):
    return f"{self.team_a.name} x {self.team_b.name}"

class Action(UUIDModel):
  player = models.ForeignKey(Player, on_delete=models.PROTECT, verbose_name='Jogador')
  team = models.ForeignKey(Team, on_delete=models.PROTECT, verbose_name='Time')
  minutes = models.IntegerField(verbose_name='Minutos')
  match = models.ForeignKey(Match, on_delete=models.PROTECT, verbose_name='Jogo')

  class Actions(models.TextChoices):
    GOAL = 'goal', 'Gol'
    ASSIT = 'assist', 'Assistência'
    YELLOW_CARD = 'yellow card', 'Cartão amarelo'
    READ_CARD = 'red card', 'Cartão vermelho'

  action = models.CharField(max_length=50, choices=Actions.choices, verbose_name='Ação')

  def __str__(self):
    return f"{self.player} - {self.action}"

  class Meta:
    verbose_name = 'Ação do jogo'
    verbose_name_plural = 'Ações do jogo'
```

**Refazendo as migrações**

```bash
$ python manage.py migrate
$ python manage.py makemigrations app
```

**Adicionando fixtures**

> File: app/fixtures/initial_data.json

```json
[
  {
      "model": "auth.user",
      "pk": 1,
      "fields": {
          "password": "pbkdf2_sha256$390000$Wkm5wqZ5xIKSMmrALdAmKY$/hEivykHXfjSs73YJGQvM7Kq5/EQMIgB+Pfrj9RE1Sc=",
          "last_login": "2022-11-28T07:46:43.073Z",
          "is_superuser": true,
          "username": "admin@user.com",
          "first_name": "",
          "last_name": "",
          "email": "admin@user.com",
          "is_staff": true,
          "is_active": true,
          "date_joined": "2022-11-28T07:46:12.228Z",
          "groups": [],
          "user_permissions": []
      }
  },
  {
      "model": "app.player",
      "pk": 1,
      "fields": {
          "name": "Cristiano Ronaldo",
          "uuid": "4876d14f-d998-4abf-96ef-89fd53185464",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 2,
      "fields": {
          "name": "De Bruyne",
          "uuid": "0b8f08d8-d871-4a42-b395-17d698f477db",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 3,
      "fields": {
          "name": "Harry Kane",
          "uuid": "0c9ba4fb-4609-464d-9845-421ca1e1e3bd",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 4,
      "fields": {
          "name": "Lewandowski",
          "uuid": "67fbf409-d94f-4858-8423-8043576cda05",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 5,
      "fields": {
          "name": "Maguirre",
          "uuid": "c7830b65-cf79-49b7-a878-82250fec1d94",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 6,
      "fields": {
          "name": "Messi",
          "uuid": "64fb9c2f-a45b-4f96-9d8b-b127878ca6f3",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 7,
      "fields": {
          "name": "Neymar",
          "uuid": "0f463bea-1dbd-4765-b080-9f5f170b6ded",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 8,
      "fields": {
          "name": "Richarlison",
          "uuid": "5ce233a8-5cd8-4a85-8156-9ac255cf909e",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.player",
      "pk": 9,
      "fields": {
          "name": "Vinicius Junior",
          "uuid": "c707bfa9-074e-4636-8772-633e4b56248d",
          "initial_price": 10.0
      }
  },
  {
      "model": "app.team",
      "pk": 1,
      "fields": {
          "name": "Argentina",
          "uuid": "76ef612d-acdf-48cd-af1c-9313d16ab642"
      }
  },
  {
      "model": "app.team",
      "pk": 2,
      "fields": {
          "name": "Alemanha",
          "uuid": "b30cf8e9-d02e-408e-9494-c4ae95b85b49"
      }
  },
  {
      "model": "app.team",
      "pk": 3,
      "fields": {
          "name": "Brasil",
          "uuid": "4da19d91-c700-4865-8333-80efc789a70f"
      }
  },
  {
      "model": "app.team",
      "pk": 4,
      "fields": {
          "name": "Bélgica",
          "uuid": "52f02f8d-bf04-4bd2-8b3d-28ec97e562a8"
      }
  },
  {
      "model": "app.team",
      "pk": 5,
      "fields": {
          "name": "Portugal",
          "uuid": "b503dc80-1f8c-4c8e-8218-a37049676815"
      }
  },
  {
      "model": "app.team",
      "pk": 6,
      "fields": {
          "name": "Polônia",
          "uuid": "322dc42e-33a0-42cb-8c50-5d1e5fc0a308"
      }
  },
  {
      "model": "app.team",
      "pk": 7,
      "fields": {
          "name": "Inglaterra",
          "uuid": "ace3abc1-003d-4baf-8d93-51135c437c7a"
      }
  },
  {
      "model": "app.myteam",
      "pk": 1,
      "fields": {
          "uuid": "22087246-01bc-46ad-a9d9-a99a6d734167"
      }
  },
  {
      "model": "app.match",
      "pk": 1,
      "fields": {
          "uuid": "97ed7b97-ea2f-4bab-80e6-e14800e4bfcc",
          "match_date": "2022-11-28T07:51:48Z",
          "team_a": 1,
          "team_a_goal": 0,
          "team_b": 2,
          "team_b_goal": 0
      }
  },
  {
      "model": "app.match",
      "pk": 2,
      "fields": {
          "uuid": "a9d93ef1-7f1a-489f-89a3-d0321735c326",
          "match_date": "2022-11-28T07:52:12Z",
          "team_a": 3,
          "team_a_goal": 0,
          "team_b": 4,
          "team_b_goal": 0
      }
  },
  {
      "model": "app.action",
      "pk": 1,
      "fields": {
          "uuid": "ea32249c-4474-4e56-b6b6-747573c43552",
          "player": 6,
          "team": 1,
          "minutes": 20,
          "match": 1,
          "action": "goal"
      }
  },
  {
      "model": "app.action",
      "pk": 2,
      "fields": {
          "uuid": "24420fc7-9e3f-4b9a-af50-025f4b857c9e",
          "player": 8,
          "team": 3,
          "minutes": 5,
          "match": 2,
          "action": "yellow card"
      }
  },
  {
      "model": "app.action",
      "pk": 3,
      "fields": {
          "uuid": "b2c3e094-0c73-4a18-a1aa-b91b42124627",
          "player": 9,
          "team": 3,
          "minutes": 10,
          "match": 2,
          "action": "assist"
      }
  },
  {
      "model": "app.action",
      "pk": 4,
      "fields": {
          "uuid": "a315fee9-55de-4b5b-aed7-9ff7dee268dc",
          "player": 7,
          "team": 3,
          "minutes": 10,
          "match": 2,
          "action": "goal"
      }
  }
]
```

**LoadData**

Dando carga nas tabelas através de um JSOIN

```bash
$ python manage.py loaddata initial_data
```

**Instalando API REST**

```bash
$ pipenv install djangorestframework
$ pipenv install autopep8
```

> File: app/api.py

```python
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.decorators import api_view
from .use_cases import UpdatePlayersInMyTeam


#PUT
@api_view(['PUT'])
def update_players_in_my_team(request: Request, my_team_uuid):
    UpdatePlayersInMyTeam().execute(my_team_uuid, request.data.get('players_uuid'))
    return Response(None, 204)
```

> File: app/use_cases.py

```python
from typing import List
from app.models import MyTeam, Player

class UpdatePlayersInMyTeam:

    def execute(self, my_team_uuid, players_uuid: List[str]):
        my_team = MyTeam.objects.get(uuid=my_team_uuid)
        players = Player.objects.filter(uuid__in=players_uuid)
        my_team.players.set(players)
```

> File: app/urls.py

```python
from django.urls import path
from .api import update_players_in_my_team

urlpatterns = [
    path('my-teams/<uuid:my_team_uuid>/players', update_players_in_my_team)
]
```

> File: cartola_fc/urls.py

```python
"""
URL configuration for cartola_fc project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/5.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('app.urls'))
]

```

> File: api.http

```python
PUT http://localhost:8000/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/players
Content-Type: application/json

{
    "players_uuid": ["c707bfa9-074e-4636-8772-633e4b56248d"]
}
```

**Integrando serviços**

$ pipenv install django-cors-headers

> File: cartola_fc/settings.py

```python
"""
Django settings for cartola_fc project.

Generated by 'django-admin startproject' using Django 5.1.

For more information on this file, see
https://docs.djangoproject.com/en/5.1/topics/settings/

For the full list of settings and their values, see
https://docs.djangoproject.com/en/5.1/ref/settings/
"""

from pathlib import Path
import os
import environ
import dj_database_url

# Build paths inside the project like this: BASE_DIR / 'subdir'.
BASE_DIR = Path(__file__).resolve().parent.parent

environ.Env.read_env(os.path.join(BASE_DIR, '.env'))

env = environ.Env(
  # set casting, default value
  DEBUG=(bool, False),
  ALLOWED_HOSTS=(list,['*']),
)

# Quick-start development settings - unsuitable for production
# See https://docs.djangoproject.com/en/5.1/howto/deployment/checklist/

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = env('SECRET_KEY')

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = env('DEBUG')

ALLOWED_HOSTS = env('ALLOWED_HOSTS')


# Application definition

INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'corsheaders',
    'app',
]

MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'corsheaders.middleware.CorsMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
]

ROOT_URLCONF = 'cartola_fc.urls'

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [],
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]

WSGI_APPLICATION = 'cartola_fc.wsgi.application'


# Database
# https://docs.djangoproject.com/en/5.1/ref/settings/#databases


DATABASE_URL = 'sqlite:///' + os.path.join(BASE_DIR, 'db.sqlite3')
DATABASES = {'default': dj_database_url.config(default=DATABASE_URL)}

# Password validation
# https://docs.djangoproject.com/en/5.1/ref/settings/#auth-password-validators

AUTH_PASSWORD_VALIDATORS = [
    {
        'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator',
    },
]


# Internationalization
# https://docs.djangoproject.com/en/5.1/topics/i18n/

LANGUAGE_CODE = 'en-us'

TIME_ZONE = 'UTC'

USE_I18N = True

USE_TZ = True


# Static files (CSS, JavaScript, Images)
# https://docs.djangoproject.com/en/5.1/howto/static-files/

STATIC_URL = 'static/'

# Default primary key field type
# https://docs.djangoproject.com/en/5.1/ref/settings/#default-auto-field

DEFAULT_AUTO_FIELD = 'django.db.models.BigAutoField'

CORS_ALLOW_ALL_ORIGINS = True
```

> File: docker-composer

```yaml
version: '3'

services:
  app:
    build: .
    volumes:
      - .:/home/python/app
    ports:
      - 8000:8000
    extra_hosts:
    - "host.docker.internal:172.17.0.1"
```

```bash
$ pipenv install confluent-kafka
```

> File: 

```python

```

## Next.js

npm install axios

> File: src/util/http.ts

```ts
import axios from "axios";

export const httpAdmin = axios.create({
  baseURL: "http://localhost:8000"
});
```

> File: src/pages/players.tsx

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
import { httpAdmin } from "../util/http";

const players = [
  {
    id: "64fb9c2f-a45b-4f96-9d8b-b127878ca6f3",
    name: "Messi",
    price: 35,
  },
  {
    id: "4876d14f-d998-4abf-96ef-89fd53185464",
    name: "Cristiano Ronaldo",
    price: 35,
  },
  {
    id: "0f463bea-1dbd-4765-b080-9f5f170b6ded",
    name: "Neymar",
    price: 25,
  },
  {
    id: "0b8f08d8-d871-4a42-b395-17d698f477db",
    name: "De Bruyne",
    price: 25,
  },
  {
    id: "c707bfa9-074e-4636-8772-633e4b56248d",
    name: "Vinicius Junior",
    price: 25,
  },
  {
    id: "67fbf409-d94f-4858-8423-8043576cda05",
    name: "Lewandowski",
    price: 15,
  },
  {
    id: "c7830b65-cf79-49b7-a878-82250fec1d94",
    name: "Maguirre",
    price: 15,
  },
  {
    id: "5ce233a8-5cd8-4a85-8156-9ac255cf909e",
    name: "Richarlison",
    price: 15,
  },
  {
    id: "0c9ba4fb-4609-464d-9845-421ca1e1e3bd",
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


  const saveMyPlayers = useCallback(async () => {
    await httpAdmin.put(
      "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/players",
      {
        players_uuid: playersSelected.map((player) => player.id),
      }
    );
  }, [playersSelected]);

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
            onClick={() => saveMyPlayers()}
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
