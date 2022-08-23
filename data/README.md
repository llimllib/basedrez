Opening DB from [lichess' opening book](https://github.com/lichess-org/chess-openings/)

1. Downloaded the repo as a zip
2. Went to the dist folder
3. concatenated the TSVs with `cat *.tsv > all.tsv`
4. loaded them into a sqlite database with: `sqlite3 openings.db -cmd '.mode tabs' '.import all.tsv openings' '.schema'`
5. added an index with `CREATE INDEX idx_epd ON openings(epd);`
