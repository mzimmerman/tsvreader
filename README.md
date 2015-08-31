# tsvreader
Command line tool to input multi-line TSV/CSV files but only export certain fields into a new TSV/CSV 

# example
cat file.tsv | tsvreader 4 9 1 > file-new.tsv

This will export columns 4, 9, and 1 into file-new.tsv with the new order specified
