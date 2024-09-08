# 3. Desenvolvendo Microsserviço a jato com Python e Django

#### Preparar Ambiente de Desenvolvimento

**Visual Studio**

- Instalar extension: dev container

**Configurar**

ctrl+shit+p

> dev container: rebuild and reopen in container add configuration to workspace

** Ajustar configurações  .devcontainer**

> File: .devcontainer/devcontainer.json

```
// For format details, see https://aka.ms/devcontainer.json. For config options, see the
// README at: https://github.com/devcontainers/templates/tree/main/src/docker-existing-docker-compose
{
	"name": "Cartola Full Cycle - Django",

	// Update the 'dockerComposeFile' list if you have more compose files or use different names.
	// The .devcontainer/docker-compose.yml file contains any overrides you need/want to make.
	"dockerComposeFile": [
		"../docker-composer.yaml",
		"docker-compose.yml"
	],

	// The 'service' property is the name of the service for the container that VS Code should
	// use. Update this value and .devcontainer/docker-compose.yml to the real service name.
	"service": "app",

	// The optional 'workspaceFolder' property is the path VS Code should open by default when
	// connected. This is typically a file mount in .devcontainer/docker-compose.yml
	"workspaceFolder": "/home/python/app",

	"customizations": {
		"vscode": {
			"extensions": [
				"ms-python.python",
				"sourcery.sourcery"
			]
		},
		"settings": {
			"editor.tabSize": 2
		}
	},
	// Features to add to the dev container. More info: https://containers.dev/features.
	// "features": {},

	// Use 'forwardPorts' to make a list of ports inside the container available locally.
	// "forwardPorts": [],

	// Uncomment the next line if you want start specific services in your Docker Compose config.
	// "runServices": [],

	// Uncomment the next line if you want to keep your containers running after VS Code shuts down.
	// "shutdownAction": "none",

	// Uncomment the next line to run commands after the container is created.
	// "postCreateCommand": "cat /etc/os-release",

	// Configure tool-specific properties.
	// "customizations": {},

	// Uncomment to connect as an existing user other than the container default. More info: https://aka.ms/dev-containers-non-root.
	// "remoteUser": "devcontainer"
}
```

> File:  .devcontainer/docker-compose.yml

```
version: '3'
services:
  # Update this to the name of the service you want to work with in your docker-compose.yml file
  app:
    # Uncomment if you want to override the service's Dockerfile to one in the .devcontainer
    # folder. Note that the path of the Dockerfile and context is relative to the *primary*
    # docker-compose.yml file (the first in the devcontainer.json "dockerComposeFile"
    # array). The sample below assumes your primary file is in the root of your project.
    #
    # build:
    #   context: .
    #   dockerfile: .devcontainer/Dockerfile

    volumes:
      # Update this to wherever you want VS Code to mount the folder of your project
      - .:/home/python/app:cached

    # Uncomment the next four lines if you will use a ptrace-based debugger like C++, Go, and Rust.
    # cap_add:
    #   - SYS_PTRACE
    # security_opt:
    #   - seccomp:unconfined

    # Overrides default command so things don't shut down after the process ends.
    # command: sleep infinity
```

> File: .vscode/settings.json

```
{
  "python.defaultInterpreterPath": "/home/python/app/.venv/bin/python",
  "python.terminal.activateEnvironment": true,
}
```

**Fazer o   Rebuild**

```bash
$ docker-compose up -d --build  
```

#### Instalando Django

Abrir o Terminal de dentro do vscode depois de subir o container.

```bash
$ pipenv install django
```

**Ajustar arquivo Pipfile e .gitignore**

> File: Pipfile

```
[[source]]
url = "https://pypi.org/simple"
verify_ssl = true
name = "pypi"

[packages]
django = "*"
django-environ = "*"
dj-database-url = "*"

[dev-packages]

[requires]
python_version = "3.10"
```

**Preparando Ambiente**

```bash
$ pipenv shell
$ django-admin
$ django-admin startproject cartola_fc
```
**Update**
-  mover manger.py para raiz
-  mover conteúdo do cartola_fc/cartola_fc para cartola_fc/ 

**Migrações**

```bash
python manage.py migrate
python manage.py runserver 0.0.0.0:8000
```

**Criando usuario admin**

```bash
$ python manage.py createsuperuser
(app) root@88e0285cfc79:/home/python/app# python manage.py createsuperuser
Username (leave blank to use 'root'): admin@user.com
Email address: admin@user.com
Password: 12345678
```

**Criando Módulos - app**

```bash
$ django-admin startapp app
```

Dentro da pasta app temos:
- admin
- apps
- models
- tests
- views

**Adicionando o app no cartola_fc**

> File: cartola_fc/settings.py

```python
...
INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'app',
]
...
```

**Modelando**

> File: app/models.py

```python
from django.db import models

# Create your models here.
class Player(models.Model):
  name = models.CharField(max_length=50)
  initial_price = models.FloatField()

  def __str__(self):
    return self.name

class Team(models.Model):
  name = models.CharField(max_length=50)

  def __str__(self):
    return self.name

class MyTeam(models.Model):
  players = models.ManyToManyField(Player)

  def __str__(self):
    return [ player.name for player in self.players.all() ].__str__()

class Match(models.Model):
  match_date = models.DateTimeField()
  team_a = models.ForeignKey(Team, on_delete=models.PROTECT, related_name='team_a_matches')
  team_a_goal = models.IntegerField(default=0)
  team_b = models.ForeignKey(Team, on_delete=models.PROTECT, related_name='team_b_matches')
  team_b_goal = models.IntegerField(default=0)

  def __str__(self):
    return f"{self.team_a.name} x {self.team_b.name}"

class Action(models.Model):
  player = models.ForeignKey(Player, on_delete=models.PROTECT)
  team = models.ForeignKey(Team, on_delete=models.PROTECT)
  minutes = models.IntegerField()
  match = models.ForeignKey(Match, on_delete=models.PROTECT)

  class Actions(models.TextChoices):
    GOAL = 'goal', 'Goal'
    ASSIT = 'assist', 'Assist'
    YELLOW_CARD = 'yellow card', 'Yellow Card'
    READ_CARD = 'red card', 'Red Card'

  action = models.CharField(max_length=50, choices=Actions.choices)

  def __str__(self):
    return f"{self.minutes}' - {self.player.name} - {self.action}"
```

**Gerando e migrando Player a plugando no Admin **

```bash
$ python manage.py makemigrations
$ python manage.py migrate
```

> File: app/admin.py

```python
from django.contrib import admin
from .models import Player, Team, MyTeam, Match, Action

# Register your models here.

class ActionInline(admin.TabularInline):
  model = Action

class MatchAdmin(admin.ModelAdmin):
  inlines = [ActionInline]

admin.site.register(Player)
admin.site.register(Team)
admin.site.register(MyTeam)
admin.site.register(Match, MatchAdmin)
admin.site.register(Action)

# signals == events
# receivers == listeners
```

**Criando estrutura para captar os Eventos**

> File: app/receivers.py

```python
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
  if not created and instance._pre_save_instance.team_a_goal != instance.team_a_goal or instance._pre_save_instance.team_b_goal != instance.team_b_goal:
    print("Match result published")

@receiver(post_save, sender=Action)
def publish_action_added(sender, instance: Action, created:bool, **kwargs):
  if created:
    print("Action created")
```

**Bootstrap do Django**

Adicionando receiver no app

> File: app/apps.py

```python
from django.apps import AppConfig


class AppConfig(AppConfig):
    default_auto_field = 'django.db.models.BigAutoField'
    name = 'app'

    def ready(self) -> None:
        import app.receivers
```

```bash
$ pipenv install django-environ
$ pipenv install dj-database-url
```

> File: .env

```
DEBUG=true
SECRET_KEY='django-insecure-kob=zl-ar$9!)9m4#xdo%&t4=81157=_#nxq%u@jd7czp=qrzr'
ALLOWED_HOSTS=*
DATABASES=sqlite:///db.sqlite3
```

> File: cartola_fc/settings.py

```
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
    'app',
]

MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
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
```
