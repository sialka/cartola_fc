INSERT INTO players (id, name, price) VALUES ('1', 'Cristiano Ronaldo', 10.00);
INSERT INTO players (id, name, price) VALUES ('2', 'De Bruyne', 10.00);
INSERT INTO players (id, name, price) VALUES ('3', 'Harry Kane', 10.00);
INSERT INTO players (id, name, price) VALUES ('4', 'Lewandowski', 10.00);
INSERT INTO players (id, name, price) VALUES ('5', 'Maguirre', 10.00);
INSERT INTO players (id, name, price) VALUES ('6', 'Messi', 10.00);
INSERT INTO players (id, name, price) VALUES ('7', 'Neymar', 10.00);
INSERT INTO players (id, name, price) VALUES ('8', 'Richarlison', 10.00);
INSERT INTO players (id, name, price) VALUES ('9', 'Vinicius Junior', 10.00);

INSERT INTO teams (id, name) VALUES ('1', 'Argentina');
INSERT INTO teams (id, name) VALUES ('2', 'Alemanha');
INSERT INTO teams (id, name) VALUES ('3', 'Brasil');
INSERT INTO teams (id, name) VALUES ('4', 'Bélgica');
INSERT INTO teams (id, name) VALUES ('5', 'Portugal');
INSERT INTO teams (id, name) VALUES ('6', 'Polônia');
INSERT INTO teams (id, name) VALUES ('7', 'Inglaterra');

-- /** fazer implementação com verificação da relação de teams com players **/
INSERT INTO team_players (team_id, player_id) VALUES ('1','1');
INSERT INTO team_players (team_id, player_id) VALUES ('2','2');
INSERT INTO team_players (team_id, player_id) VALUES ('3','3');
INSERT INTO team_players (team_id, player_id) VALUES ('3','4');
INSERT INTO team_players (team_id, player_id) VALUES ('4','5');
INSERT INTO team_players (team_id, player_id) VALUES ('5','6');
INSERT INTO team_players (team_id, player_id) VALUES ('6','7');
INSERT INTO team_players (team_id, player_id) VALUES ('7','8');
INSERT INTO team_players (team_id, player_id) VALUES ('7','9');

INSERT INTO my_team (id, name, score) VALUES ('1', 'Meu Time FC', 100);

INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '1');
INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '2');
INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '3');
INSERT INTO my_team_players (my_team_id, player_id) VALUES ('1', '4');

INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES ('1', '2024-08-02 00:00:00', '1', 'Argentina', '2', 'Alemanha', '1-0');
INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES ('2', '2024-08-02 00:00:00', '3', 'Brasil', '4', 'Bélgica', '1-0');

-- -- insert match actions
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('1', '1', '1', '1', 'goal', 1, 20);
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('2', '1', '1', '1', 'yellow card', 1, 5);
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('3', '2', '3', '3', 'assist', 1, 10);
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES ('4', '2', '3', '4', 'goal', 1, 10);