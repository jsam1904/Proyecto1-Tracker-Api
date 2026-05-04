CREATE TABLE IF NOT EXISTS teams (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    city        VARCHAR(100) NOT NULL,
    stadium     VARCHAR(100) NOT NULL,
    founded     INTEGER NOT NULL,
    image_url   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

INSERT INTO teams (name, city, stadium, founded, image_url) VALUES
('Bayern München',       'München',      'Allianz Arena',                1900, 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/FC_Bayern_M%C3%BCnchen_logo_%282017%29.svg/800px-FC_Bayern_M%C3%BCnchen_logo_%282017%29.svg.png'),
('Borussia Dortmund',    'Dortmund',     'Signal Iduna Park',            1909, 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Borussia_Dortmund_logo.svg/800px-Borussia_Dortmund_logo.svg.png'),
('Bayer Leverkusen',     'Leverkusen',   'BayArena',                     1904, 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1f/Bayer_04_Leverkusen_Logo.svg/800px-Bayer_04_Leverkusen_Logo.svg.png'),
('RB Leipzig',           'Leipzig',      'Red Bull Arena',               2009, 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/RB_Leipzig_2014_logo.svg/800px-RB_Leipzig_2014_logo.svg.png'),
('Eintracht Frankfurt',  'Frankfurt',    'Deutsche Bank Park',           1899, 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/Eintracht_Frankfurt_Logo.svg/800px-Eintracht_Frankfurt_Logo.svg.png'),
('VfB Stuttgart',        'Stuttgart',    'MHPArena',                     1893, 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a4/VfB_Stuttgart_1893_Logo.svg/800px-VfB_Stuttgart_1893_Logo.svg.png'),
('SC Freiburg',          'Freiburg',     'Europa-Park Stadion',          1904, 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f1/SC-Freiburg_Logo-new.svg/800px-SC-Freiburg_Logo-new.svg.png'),
('Union Berlin',         'Berlin',       'Stadion An der Alten Försterei',1906, 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/44/1._FC_Union_Berlin_Logo.svg/800px-1._FC_Union_Berlin_Logo.svg.png'),
('Werder Bremen',        'Bremen',       'wohninvest WESERSTADION',      1899, 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/be/SV-Werder-Bremen-Logo.svg/800px-SV-Werder-Bremen-Logo.svg.png'),
('Borussia Mönchengladbach', 'Mönchengladbach', 'Borussia-Park',        1900, 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/81/Borussia_M%C3%B6nchengladbach_logo.svg/800px-Borussia_M%C3%B6nchengladbach_logo.svg.png'),
('FC Augsburg',          'Augsburg',     'WWK Arena',                    1907, 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/FC_Augsburg_logo.svg/800px-FC_Augsburg_logo.svg.png'),
('TSG Hoffenheim',       'Sinsheim',     'PreZero Arena',                1899, 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/64/TSG_Logo-Standard_4c.svg/800px-TSG_Logo-Standard_4c.svg.png'),
('VfL Wolfsburg',        'Wolfsburg',    'Volkswagen Arena',             1945, 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f3/Logo-VfL-Wolfsburg.svg/800px-Logo-VfL-Wolfsburg.svg.png'),
('VfL Bochum',           'Bochum',       'Vonovia Ruhrstadion',          1848, 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/VfL_Bochum_logo.svg/800px-VfL_Bochum_logo.svg.png'),
('1. FC Köln',           'Köln',         'RheinEnergieStadion',          1948, 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/FC_Cologne_logo.svg/800px-FC_Cologne_logo.svg.png'),
('Hertha BSC',           'Berlin',       'Olympiastadion',               1892, 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/86/Hertha_BSC_Logo_2012.svg/800px-Hertha_BSC_Logo_2012.svg.png'),
('Mainz 05',             'Mainz',        'MEWA Arena',                   1905, 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Logo_Mainz_05.svg/800px-Logo_Mainz_05.svg.png'),
('Darmstadt 98',         'Darmstadt',    'Merck-Stadion am Böllenfalltor',1898, 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/57/SV_Darmstadt_98_Logo_2021.svg/800px-SV_Darmstadt_98_Logo_2021.svg.png')
ON CONFLICT DO NOTHING;