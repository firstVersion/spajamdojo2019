\c spajamdojo2019

INSERT INTO users (user_name) VALUES('ばん');
INSERT INTO users (user_name) VALUES('みよし');
INSERT INTO users (user_name) VALUES('しゅう');
INSERT INTO users (user_name) VALUES('くどう');
INSERT INTO users (user_name) VALUES('やの');

INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 1, 1, 1, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 2, 1, 1, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 3, 1, 2, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 4, 1, 2, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 5, 1, 2, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 2, 2, 1, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 2, 3, 1, 0);
INSERT INTO scores ( user_id, term, team_id, distance ) VALUES( 2, 4, 1, 0);


INSERT INTO bets ( user_id, amount, term, team_id ) VALUES( 1, 1000, 1, 1);
INSERT INTO bets ( user_id, amount, term, team_id ) VALUES( 2, 1000, 1, 1);
INSERT INTO bets ( user_id, amount, term, team_id ) VALUES( 3, 1000, 2, 1);
INSERT INTO bets ( user_id, amount, term, team_id ) VALUES( 4, 1000, 2, 1);
INSERT INTO bets ( user_id, amount, term, team_id ) VALUES( 5, 1000, 2, 1);

INSERT INTO terms ( term_id, t_from, t_to ) VALUES( 1, '2019.04/14', '2019.04/20');
INSERT INTO terms ( term_id, t_from, t_to ) VALUES( 2, '2019.04/21', '2019.04/27');
INSERT INTO terms ( term_id, t_from, t_to ) VALUES( 3, '2019.04/28', '2019.05/04');
INSERT INTO terms ( term_id, t_from, t_to ) VALUES( 4, '2019.05/05', '2019.05/11');
