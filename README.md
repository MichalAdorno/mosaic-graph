# Mosaic Graph Generator

The program generates a mosaic graph. (This kind of graphs is commonly used in economics)

The usage of program parameters is described in the `main.go` file (lines: 14-19).
It takes a CSV file with the following format as an input:
```
textiles,25.0
clothing,25.0
footwear,50.0
```
and returns a graph as an output:

![Alt text](./output/out_data6.png "Mosaic Graph")

(Note: I skipped printing category names on the graph because I was mainly interested in the algorithm to generate the graph.) 
