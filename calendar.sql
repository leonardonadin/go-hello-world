CREATE TABLE plants(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    planting_type TEXT,
    start_planting_month TEXT,
    end_planting_month TEXT,
    region TEXT,
    days_to_harvest INTEGER
);

INSERT INTO plants(name, planting_type, start_planting_month, end_planting_month, region, days_to_harvest)
VALUES(
    'Agrião',
    'mudas',
    'janeiro',
    'dezembro',
    'sul',
    60
), (
    'Alface',
    'mudas',
    'fevereiro',
    'agosto',
    'todas',
    80
);

