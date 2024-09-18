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
